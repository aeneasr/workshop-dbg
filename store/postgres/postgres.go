package postgres

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/ory/workshop-dbg/store"
)

const contactTable = "dbg_contacts"

type PostgresStore struct {
	DB *sqlx.DB
}

var schemata = []string{
	fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id       	text NOT NULL PRIMARY KEY,
	name		text NULL,
	department	text NULL,
	company		text NULL
)
`, contactTable),
}

func (s *PostgresStore) CreateSchemas() error {
	for k, schema := range schemata {
		if _, err := s.DB.Exec(schema); err != nil {
			log.Warnf("Error creating schema %d with error %s: %s", k, err, schema)
			return err
		}
	}
	return nil
}

func (s *PostgresStore) FetchContacts() (store.Contacts, error) {
	var cs []*store.Contact
	csi := store.Contacts{}
	if err := s.DB.Select(&cs, fmt.Sprintf("SELECT * FROM %s", contactTable)); err != nil {
		return csi, err
	}

	for _, c := range cs {
		csi[c.ID] = c
	}
	return csi, nil
}

func (s *PostgresStore) GetContact(id string) (*store.Contact, error) {
	var c store.Contact
	if err := s.DB.Get(&c, fmt.Sprintf("SELECT * FROM %s WHERE id = $1", contactTable), id); err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *PostgresStore) DeleteContact(id string) error {
	if _, err := s.DB.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = $1", contactTable), id); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) UpdateContact(c *store.Contact) error {
	if _, err := s.DB.NamedExec(
		fmt.Sprintf("UPDATE %s SET name = :name, department = :department, company = :company WHERE id = :id", contactTable),
		&c,
	); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateContact(c *store.Contact) error {
	if _, err := s.DB.NamedExec(
		fmt.Sprintf(
			`INSERT INTO %s (id, name, department, company) VALUES (:id, :name, :department, :company)`,
			contactTable,
		),
		*c,
	); err != nil {
		return err
	}
	return nil
}
