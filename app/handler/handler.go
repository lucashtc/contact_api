package handler

import (
	"github.com/lucashtc/contact_api/app"
	"github.com/lucashtc/contact_api/database"
)

var db database.db

// Create contact ...
func Create(c app.Contact) error {
	return nil
}

// GetContact ...
func GetContact(id string) app.Contact {
	var c Contact
	return c
}

// GetAll ...
func GetAll() []app.Contact {
	var c []app.Contact

	return c
}

// EditContact ...
func EditContact(c app.Contact) error {
	return nil
}

// DeleteContact ...
func DeleteContact() error {
	return nil
}
