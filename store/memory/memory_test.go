package memory

import (
	"github.com/ory/workshop-dbg/store"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	s := &InMemoryStore{
		Contacts: store.Contacts{},
	}

	c1 := &store.Contact{ID: uuid.New(), Name: "a", Department: "a1", Company: "a2"}
	c2 := &store.Contact{ID: uuid.New(), Name: "b", Department: "b1", Company: "b2"}
	c3 := &store.Contact{ID: c2.ID, Name: "ba", Department: "ba1", Company: "ba2"}
	r, err := s.GetContact(c1.ID)
	assert.NotNil(t, err)

	assert.Nil(t, s.DeleteContact(c1.ID))
	assert.Nil(t, s.CreateContact(c1))
	assert.Nil(t, s.CreateContact(c2))

	cs, err := s.FetchContacts()
	assert.Nil(t, err)
	assert.Len(t, cs, 2)

	r, err = s.GetContact(c1.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, c1, r)

	r, err = s.GetContact(c2.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, c2, r)

	assert.Nil(t, s.UpdateContact(c3))
	r, err = s.GetContact(c3.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, c3, r)

	r, err = s.GetContact(c2.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, c3, r)

	assert.Nil(t, s.DeleteContact(c1.ID))
	assert.Nil(t, s.DeleteContact(c1.ID))
	cs, err = s.FetchContacts()
	assert.Nil(t, err)
	assert.Len(t, cs, 1)

	assert.Nil(t, s.DeleteContact(c2.ID))
	cs, err = s.FetchContacts()
	assert.Nil(t, err)
	assert.Len(t, cs, 0)
}
