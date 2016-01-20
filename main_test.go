package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockedContactList = Contacts{
	"john-bravo": Contact{
		ID:         "john-bravo",
		Name:       "John Bravo",
		Department: "IT",
		Company:    "ACME Inc",
	},
	"cathrine-mueller": Contact{
		ID:         "cathrine-mueller",
		Name:       "Cathrine Müller",
		Department: "HR",
		Company:    "Grove AG",
	},
}

var mockContact = Contact{
	ID:         "eddie-markson",
	Name:       "Eddie Markson",
	Department: "Finance",
	Company:    "ACME Inc",
}

func TestListContacts(t *testing.T) {
	// Initialize everything (very similar to main() function).
	router := mux.NewRouter()
	router.HandleFunc("/contacts", ListContacts(mockedContactList)).Methods("GET")
	ts := httptest.NewServer(router)

	// This helper function makes an http request to ListContacts and validates its output.
	fetchAndTestContactList(t, ts, mockedContactList)
}

func TestAddContacts(t *testing.T) {
	// We create a copy of the store
	contactListForThisTest := copyContacts(mockedContactList)

	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/contacts", AddContact(contactListForThisTest)).Methods("POST")
	ts := httptest.NewServer(router)

	// Make the request
	_, _, errs := gorequest.New().Post(ts.URL + "/contacts").SendStruct(mockContact).End()
	require.Len(t, errs, 0)
	require.Equal(t, contactListForThisTest[mockContact.ID], mockContact)
}

func TestDeleteContacts(t *testing.T) {
	// We create a copy of the store
	contactListForThisTest := copyContacts(mockedContactList)

	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/contacts", AddContact(contactListForThisTest)).Methods("DELETE")
	ts := httptest.NewServer(router)

	// Make the request
	_, _, errs := gorequest.New().Delete(ts.URL + "/contacts/john-bravo").End()
	require.Len(t, errs, 0)

	_, found := contactListForThisTest[mockContact.ID]
	require.False(t, found)
}

func TestUpdateContacts(t *testing.T) {
	// We create a copy of the store
	contactListForThisTest := copyContacts(mockedContactList)

	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/contacts", AddContact(contactListForThisTest)).Methods("UPDATE")
	ts := httptest.NewServer(router)

	// Make the request
	_, _, errs := gorequest.New().Put(ts.URL + "/contacts/john-bravo").SendStruct(mockContact).End()
	require.Len(t, errs, 0)

	// The new contact should be inserted
	_, found := contactListForThisTest[mockContact.ID]
	require.True(t, found)

	// The old one should be removed
	_, found = contactListForThisTest["john-bravo"]
	require.False(t, found)
}

func fetchAndTestContactList(t *testing.T, ts *httptest.Server, compareWith Contacts) {
	// Request ListContacts
	resp, err := http.Get(ts.URL + "/contacts")

	// Verify that no errors occurred
	require.Nil(t, err)

	// Unmarshal the output
	var result Contacts
	err = json.NewDecoder(resp.Body).Decode(&result)

	// Make sure that no error occurred
	require.Nil(t, err)

	// Compare the outputs
	assert.Equal(t, compareWith, result)
}

func copyContacts(original Contacts) Contacts {
	result := Contacts{}
	for k, v := range original {
		result[k] = v
	}
	return result
}