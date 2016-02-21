package store

type ContactStorer interface {
	FetchContacts() (Contacts, error)
	GetContact(id string) (*Contact, error)
	DeleteContact(id string) error
	CreateContact(*Contact) error
	UpdateContact(*Contact) error
}

// Contacts is a list of contacts.
type Contacts map[string]*Contact

// Contact defines the structure of a contact which including name, department and company.
type Contact struct {
	// The unique identifier of this contact.
	// omitempty hides this field when exporting to json. Because it is common for json
	// to be lowercase, we additionally define `json:"id"` to tell the "exporter" that this
	// field should be called id, not ID.
	ID string `json:"id,omitempty" db:"id"`

	// Name is the contact's full name.
	Name string `json:"name" db:"name"`

	// Department is the contact's department in a company.
	Department string `json:"department" db:"department"`

	// Company is the name of the company the contact works for.
	Company string `json:"company" db:"company"`

	// Here is room for improvements like adding new fields
}
