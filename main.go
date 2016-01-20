package main

// The import section defines libraries that we are going to use in our program.
import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/ory-am/common/pkg"
	"github.com/ory-am/common/env"
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
type Contacts map[string]Contact

// MyContacts is an exemplary list of contacts.
var MyContacts = Contacts{
	"john-bravo": Contact{
		Name:       "John Bravo",
		Department: "IT",
		Company:    "ACME Inc",
	},
	"cathrine-mueller": Contact{
		Name:       "Cathrine Müller",
		Department: "HR",
		Company:    "Grove AG",
	},
	"maximilian-schmidt": Contact{
		Name:       "Maximilian Schmidt",
		Department: "PR",
		Company:    "Titanpad AG",
	},
}

// The main routine is going the "entry" point.
func main() {
	// Create a new router.
	router := mux.NewRouter()

	// RESTful defines operations
	// * GET for fetching data
	// * POST for inserting data
	// * PUT for updating existing data
	// * DELETE for deleting data
	router.HandleFunc("/contacts/{id}", UpdateContact(MyContacts)).Methods("PUT")
	router.HandleFunc("/contacts/{id}", DeleteContact(MyContacts)).Methods("DELETE")
	router.HandleFunc("/contacts", ListContacts(MyContacts)).Methods("GET")
	router.HandleFunc("/contacts", AddContact(MyContacts)).Methods("POST")

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

// ListContacts takes a contact list and outputs it.
func ListContacts(contacts Contacts) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		// Write contact list to output
		pkg.WriteIndentJSON(rw, contacts)
	}
}

// AddContact will add a contact to the list
func AddContact(contacts Contacts) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		// We parse the request's information into contactToBeAdded
		contactToBeAdded, err := ReadContactData(rw, r)

		// Abort handling the request if an error occurs.
		if err != nil {
			return
		}

		// Save newContact to the list of contacts.
		contacts[contactToBeAdded.ID] = contactToBeAdded

		// Output our newly created contact
		pkg.WriteIndentJSON(rw, contactToBeAdded)
	}
}

// DeleteContact will delete a contact from the list
func DeleteContact(contacts Contacts) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Fetch the ID of the contact that is going to be deleted
		contactToBeDeleted := mux.Vars(r)["id"]

		// Check if the contact exists and return an error if not
		if _, found := contacts[contactToBeDeleted]; !found {
			http.Error(rw, "I do not know any contact by that ID.", http.StatusNotFound)
			return
		}

		// Delete the contact from the list
		delete(contacts, contactToBeDeleted)

		// Per specification, RESTful may return an empty response when a DELETE request was successful
		rw.WriteHeader(http.StatusNoContent)
	}
}

// UpdateContact will update a contact on the list
func UpdateContact(contacts Contacts) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Fetch the ID of the contact that is going to be updated
		contactToBeUpdated := mux.Vars(r)["id"]

		// Check if the contact exists
		if _, found := contacts[contactToBeUpdated]; !found {
			http.Error(rw, "I don't know any contact by that ID.", http.StatusNotFound)
			return
		}

		// We parse the request's information into newContactData.
		newContactData, err := ReadContactData(rw, r)

		// Abort handling the request if an error occurs.
		if err != nil {
			return
		}

		// Update the data in the contact list.
		delete(contacts, contactToBeUpdated)
		contacts[newContactData.ID] = newContactData

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