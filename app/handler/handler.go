package handler

import (
	"fmt"
	"net/http"

	"github.com/lucashtc/contact_api/app"
	"github.com/lucashtc/contact_api/database"
)

var d struct {
	database.Db
}

// Create contact ...
func Create(c app.Contact) error {
	return nil
}

// GetContact ...
func GetContact(id string) app.Contact {
	var c app.Contact
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
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	d.Conn()
	result := d.DeleteContact(1)

	fmt.Fprint(w, result)
}
