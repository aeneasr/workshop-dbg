package main

// The import section defines libraries that we are going to use in our program.
import (
	"fmt"
	"log"
	"net/http"
	"math/rand"

	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ory-am/common/env"
	"github.com/ory-am/common/pkg"
	"github.com/pborman/uuid"
	"github.com/rs/cors"
	"math"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	. "github.com/ory-am/workshop-dbg/store"
	"github.com/ory-am/workshop-dbg/store/memory"
	"github.com/ory-am/workshop-dbg/store/postgres"
)

// In a 12 factor app, we must obey the environment variables.
var envHost = env.Getenv("HOST", "")
var envPort = env.Getenv("PORT", "5678")
var databaseURL = env.Getenv("DATABASE_URL", "")
var thisID = uuid.New()

// MyContacts is an exemplary list of contacts.
var MyContacts = Contacts{
	// Each contact hs identified by its ID which is prepended with "my-id":
	// We are doing this because it is easier to manage and simpler to read.
	"john-bravo": &Contact{
		Name:       "Andreas Preuss",
		Department: "IT",
		Company:    "ACME Inc",
	},
	"cathrine-mueller": &Contact{
		Name:       "Cathrine Müller",
		Department: "HR",
		Company:    "Grove AG",
	},
	"maximilian-schmidt": &Contact{
		Name:       "Maximilian Schmidt",
		Department: "PR",
		Company:    "Titanpad AG",
	},
	"uwe-charly": &Contact{
		Name:       "Uwe Charly",
		Department: "FAC",
		Company:    "KPMG",
	},
	"Thomas-Aidan": &Contact{
		Name:       "Thomas Aidan",
		Department: "INO",
		Company:    "OuterSpace",
	},
	"frank-sec": &Contact{
		Name:       "Frank Secure",
		Department: "Unknown",
		Company:    "Secret",
	},
	"juergen-elsner": &Contact{
		Name:       "Jürgen Elsner",
		Department: "DaCS",
		Company:    "DBG",
	},
}

var memoryStore = &memory.InMemoryStore{Contacts: MyContacts}

// The main routine is going the "entry" point.
func main() {
	// Create a new router.
	router := mux.NewRouter()

	// RESTful defines operations
	// * GET for fetching data
	// * POST for inserting data
	// * PUT for updating existing data
	// * DELETE for deleting data
	router.HandleFunc("/memory/contacts", ListContacts(memoryStore)).Methods("GET")
	router.HandleFunc("/memory/contacts", AddContact(memoryStore)).Methods("POST")
	router.HandleFunc("/memory/contacts/{id}", UpdateContact(memoryStore)).Methods("PUT")
	router.HandleFunc("/memory/contacts/{id}", DeleteContact(memoryStore)).Methods("DELETE")

	// Connect to database store
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		log.Printf("Could not connect to database because %s", err)
	} else {
		databaseStore := &postgres.PostgresStore{DB: db}
		if err := databaseStore.CreateSchemas(); err != nil {
			log.Printf("Could not set up relations %s", err)
		} else {
			router.HandleFunc("/database/contacts", ListContacts(databaseStore)).Methods("GET")
			router.HandleFunc("/database/contacts", AddContact(databaseStore)).Methods("POST")
			router.HandleFunc("/database/contacts/{id}", UpdateContact(databaseStore)).Methods("PUT")
			router.HandleFunc("/database/contacts/{id}", DeleteContact(databaseStore)).Methods("DELETE")
		}
	}

	// The info endpoint is for showing demonstration purposes only and is not subject to any task.
	router.HandleFunc("/info", InfoHandler).Methods("GET")
	router.HandleFunc("/pi", ComputePi).Methods("GET")
	router.HandleFunc("/allocate", Allocate).Methods("GET")

	// Print where to point the browser at.
	fmt.Printf("Listening on %s\n", "http://localhost:5678")

	// Cross origin resource requests
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"}},
	)

	// Start up the server and check for errors.
	listenOn := fmt.Sprintf("%s:%s", envHost, envPort)
	if err := http.ListenAndServe(listenOn, c.Handler(router)); err != nil {
		log.Fatalf("Could not set up server because %s", err)
	}
}

// ListContacts takes a contact list and outputs it.
func ListContacts(store ContactStorer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		// Write contact list to output
		contacts, err := store.FetchContacts()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		pkg.WriteIndentJSON(rw, contacts)
	}
}

// AddContact will add a contact to the list
func AddContact(contacts ContactStorer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		// We parse the request's information into contactToBeAdded
		contactToBeAdded, err := ReadContactData(rw, r)

		// Abort handling the request if an error occurs.
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		// Save newContact to the list of contacts.
		if err = contacts.CreateContact(&contactToBeAdded); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		// Output our newly created contact
		pkg.WriteIndentJSON(rw, contactToBeAdded)
	}
}

// DeleteContact will delete a contact from the list
func DeleteContact(contacts ContactStorer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Fetch the ID of the contact that is going to be deleted
		contactToBeDeleted := mux.Vars(r)["id"]

		// Delete the contact from the list
		if err := contacts.DeleteContact(contactToBeDeleted); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		// Per specification, RESTful may return an empty response when a DELETE request was successful
		rw.WriteHeader(http.StatusNoContent)
	}
}

// UpdateContact will update a contact on the list
func UpdateContact(store ContactStorer) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		// We parse the request's information into newContactData.
		newContactData, err := ReadContactData(rw, r)

		// Abort handling the request if an error occurs.
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the data in the contact list.
		if err := store.UpdateContact(&newContactData); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the new data
		pkg.WriteIndentJSON(rw, newContactData)
	}
}

// ReadContactData is a helper function for parsing a HTTP request body. It returns a contact on success and an
// error if something went wrong.
func ReadContactData(rw http.ResponseWriter, r *http.Request) (contact Contact, err error) {
	err = json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Could not read input data because %s", err), http.StatusBadRequest)
		return contact, err
	}

	return contact, nil
}

func Allocate(rw http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		n = 0
	}
	m := make([][]byte, n + 1)

	for i := 0; i < n; i++ {
		z := make([]byte, n + 1)
		_, _ = rand.Read(z)
		m[i] = z
	}

	pkg.WriteIndentJSON(rw, struct {
		Result string `json:"result"`
		N int `json:"n"`
	}{
		Result: "Processed!",
		N: n,
	})
}

func ComputePi(rw http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		n = 0
	}

	pkg.WriteIndentJSON(rw, struct {
		Pi string `json:"pi"`
		N  int    `json:"n"`
	}{
		Pi: strconv.FormatFloat(pi(n), 'E', -1, 64),
		N:  n,
	})
}

func InfoHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(thisID))
}

// pi launches n goroutines to compute an
// approximation of pi.
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
