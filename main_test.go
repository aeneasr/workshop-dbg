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
	router.HandleFunc("/", ListContacts(&mockedContactList)).Methods("GET")
	ts := httptest.NewServer(router)

	// This helper function makes an http request to ListContacts and validates its output.
	fetchAndTestContactList(t, ts, mockedContactList)
}

func TestAddContacts(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/", AddContact(&mockedContactList)).Methods("POST")
	router.HandleFunc("/", ListContacts(&mockedContactList)).Methods("GET")
	ts := httptest.NewServer(router)

	// We are expecting that our new contact is now in the list of contacts.
	expectContactList := append(mockedContactList, mockContact)

	_, body, errs := gorequest.New().Post(ts.URL).SendStruct(mockContact).End()
	require.Len(t, errs, 0)
	require.Empty(t, body)

	// Like in TestListContacts, this helper function makes an http request to ListContacts and validates its output.
	fetchAndTestContactList(t, ts, expectContactList)
}

func fetchAndTestContactList(t *testing.T, ts *httptest.Server, compareWith Contacts) {
	// Request ListContacts
	resp, err := http.Get(ts.URL)

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
