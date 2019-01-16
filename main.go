package main

// The import section defines libraries that we are going to use in our program.
import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ory-am/common/env"
	"github.com/rs/cors"
)

// In a 12 factor app, we must obey the environment variables.
var envHost = env.Getenv("HOST", "")
var envPort = env.Getenv("PORT", "5678")

// Contact defines the structure of a contact which including name, department and company.
type Contact struct {
	ID string `json:"id"`

	// Name is the contact's full name.
	Name string `json:"name"`

	// Department is the contact's department in a company.
	Department string `json:"department"`

	// Company is the name of the company the contact works for.
	Company string `json:"company"`
}

// Contacts is a list of contact structs.
type Contacts []Contact

// myContacts is an exemplary
var MyContacts = Contacts{
	{
		ID:         "john-bravo",
		Name:       "John Bravo",
		Department: "IT",
		Company:    "ACME Inc",
	},
	{
		ID:         "cathrine-mueller",
		Name:       "Cathrine Müller",
		Department: "HR",
		Company:    "Grove AG",
	},
	{
		ID:         "john-bravo",
		Name:       "Maximilian Schmidt",
		Department: "PR",
		Company:    "Titanpad AG",
	},
}

// The main routine is going the "entry" point.
func main() {
	// Create a new router.////
	router := mux.NewRouter()

	// Requests to "/" are by method listContacts.
	router.HandleFunc("/", ListContacts(&MyContacts)).Methods("GET")
	router.HandleFunc("/", AddContact(&MyContacts)).Methods("POST")

	// Print some information.
	fmt.Printf("Listening on %s\n", "http://localhost:5678")

	// Cross origin resource requests
	c := cors.New(cors.Options{AllowedOrigins: []string{"*"}})

	// Start up the server and check for errors.
	listenOn := fmt.Sprintf("%s:%s", envHost, envPort)
	err := http.ListenAndServe(listenOn, c.Handler(router))
	if err != nil {
		log.Fatalf("Could not set up server because %s", err)
	}
}

// listContacts takes a contact list and outputs it.
func ListContacts(contacts *Contacts) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		output, err := json.MarshalIndent(contacts, "", "\t")

		// If an error occurs, handle it.
		if err != nil {
			http.Error(rw, fmt.Sprintf("Could not read vcards because %s", err), http.StatusInternalServerError)
			return
		}

		rw.Write(output)
	}
}

// addContact will add a contact to the list
func AddContact(contacts *Contacts) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		// newContact poses as a placeholder for the contact that the request is going to add.
		var newContact Contact

		// We decode the information and "inject" it to newContact.
		err := json.NewDecoder(r.Body).Decode(&newContact)

		// If an error occurs while decoding, handle it.
		if err != nil {
			http.Error(rw, fmt.Sprintf("Could not read vcards because %s", err), http.StatusInternalServerError)
			return
		}

		// Save newContact to the list of available contacts.
		*contacts = append(*contacts, newContact)

		// Per specification, RESTful may return an empty response when a POST request was successful
		rw.WriteHeader(http.StatusNoContent)
	}
}
