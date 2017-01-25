package memory

import (
	"errors"
	"github.com/ory/workshop-dbg/store"
)

type InMemoryStore struct {
	Contacts store.Contacts
}

func (s *InMemoryStore) FetchContacts() (store.Contacts, error) {
	return s.Contacts, nil
}

func (s *InMemoryStore) GetContact(id string) (*store.Contact, error) {
	if c, ok := s.Contacts[id]; !ok {
		return nil, errors.New("Not found")
	} else {
		return c, nil
	}
}

func (s *InMemoryStore) DeleteContact(id string) error {
	delete(s.Contacts, id)
	return nil
}

func (s *InMemoryStore) CreateContact(c *store.Contact) error {
	s.Contacts[c.ID] = c
	return nil
}

func (s *InMemoryStore) UpdateContact(c *store.Contact) error {
	s.Contacts[c.ID] = c
	return nil
}
