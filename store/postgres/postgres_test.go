package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/workshop-dbg/store"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
)

var s *PostgresStore

func TestMain(m *testing.M) {
	var db *sqlx.DB
	var err error
	var c dockertest.ContainerID
	if c, err = dockertest.ConnectToPostgreSQL(15, time.Second, func(url string) bool {
		var err error
		db, err = sqlx.Open("postgres", url)
		if err != nil {
			return false
		}
		return db.Ping() == nil
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	s = &PostgresStore{DB: db}
	if err := s.CreateSchemas(); err != nil {
		log.Fatalf("Could not set up schemas: %v", err)
	}
	if err := s.CreateSchemas(); err != nil {
		log.Fatalf("Schema did fail on second time: %v", err)
	}

	result := m.Run()
	c.KillRemove()
	os.Exit(result)
}

func TestStore(t *testing.T) {
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
