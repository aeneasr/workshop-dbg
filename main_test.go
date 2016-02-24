package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/ory-am/workshop-dbg/store"
	"github.com/ory-am/workshop-dbg/store/memory"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockedContactList = Contacts{
	"john-bravo": &Contact{
		ID:         "john-bravo",
		Name:       "John Bravo",
		Department: "IT",
		Company:    "ACME Inc",
	},
	"cathrine-mueller": &Contact{
		ID:         "cathrine-mueller",
		Name:       "Cathrine Müller",
		Department: "HR",
		Company:    "Grove AG",
	},
}

var mockContact = &Contact{
	ID:         "eddie-markson",
	Name:       "Eddie Markson",
	Department: "Finance",
	Company:    "ACME Inc",
}

func TestListContacts(t *testing.T) {
	store := &memory.InMemoryStore{Contacts: mockedContactList}

	// Initialize everything (very similar to main() function).
	router := mux.NewRouter()
	router.HandleFunc("/contacts", ListContacts(store)).Methods("GET")
	ts := httptest.NewServer(router)

	// This helper function makes an http request to ListContacts and validates its output.
	fetchAndTestContactList(t, ts, mockedContactList)
}

func TestAddContacts(t *testing.T) {
	// We create a copy of the store
	contactListForThisTest := copyContacts(mockedContactList)
	store := &memory.InMemoryStore{Contacts: contactListForThisTest}

	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/contacts", AddContact(store)).Methods("POST")
	ts := httptest.NewServer(router)

	// Make the request
	resp, _, errs := gorequest.New().Post(ts.URL + "/contacts").SendStruct(mockContact).End()
	require.Len(t, errs, 0)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, contactListForThisTest[mockContact.ID], mockContact)
}

func TestDeleteContacts(t *testing.T) {
	// We create a copy of the store
	contactListForThisTest := copyContacts(mockedContactList)
	store := &memory.InMemoryStore{Contacts: contactListForThisTest}

	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/contacts/{id}", DeleteContact(store)).Methods("DELETE")
	ts := httptest.NewServer(router)

	// Make the request
	resp, _, errs := gorequest.New().Delete(ts.URL + "/contacts/john-bravo").End()
	require.Len(t, errs, 0)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)

	_, found := contactListForThisTest["john-bravo"]
	require.False(t, found)
}

func TestUpdateContacts(t *testing.T) {
	// We create a copy of the store
	contactListForThisTest := copyContacts(mockedContactList)
	store := &memory.InMemoryStore{Contacts: contactListForThisTest}

	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/contacts/{id}", UpdateContact(store)).Methods("PUT")
	ts := httptest.NewServer(router)

	// Make the request
	resp, _, errs := gorequest.New().Put(ts.URL + "/contacts/john-bravo").SendStruct(mockContact).End()
	require.Len(t, errs, 0)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// The new contact should be inserted
	_, found := contactListForThisTest[mockContact.ID]
	require.True(t, found)
}

func TestPis(t *testing.T) {
	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/pis", ComputePis).Methods("GET")
	ts := httptest.NewServer(router)

	// Make the request
	resp, body, errs := gorequest.New().Get(ts.URL + "/pis?n=2").End()
	require.Len(t, errs, 0)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	res := struct {
		Pi string `json:"pi"`
		N  int    `json:"n"`
	}{}
	require.Nil(t, json.Unmarshal([]byte(body), &res))
	assert.Equal(t, 2, res.N)
}

func TestPi(t *testing.T) {
	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/pi", ComputePi).Methods("GET")
	ts := httptest.NewServer(router)

	// Make the request
	resp, body, errs := gorequest.New().Get(ts.URL + "/pi?n=100").End()
	require.Len(t, errs, 0)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	res := struct {
		Pi string `json:"pi"`
		N  int    `json:"n"`
	}{}
	require.Nil(t, json.Unmarshal([]byte(body), &res))
	assert.Equal(t, 100, res.N)
}

func TestAllocate(t *testing.T) {
	// Initialize the HTTP routes, similar to main()
	router := mux.NewRouter()
	router.HandleFunc("/allocate", Allocate).Methods("GET")
	ts := httptest.NewServer(router)

	// Make the request
	resp, body, errs := gorequest.New().Get(ts.URL + "/allocate?n=100").End()
	require.Len(t, errs, 0)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	res := struct {
		Result string `json:"result"`
		N  int    `json:"n"`
	}{}
	require.Nil(t, json.Unmarshal([]byte(body), &res))
	assert.Equal(t, 100, res.N)
	assert.Equal(t, "Processed!", res.Result)
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
