package database

import (
	"github.com/lucashtc/contact_api/app/handler"
)

func (d *db) GetAll() []handler.Contact {
	var c []handler.Contact
	return c
}

func (d *db) GetContact(id int) handler.Contact {
	var c handler.Contact
	return c
}

func (d *db) CreateContact(c handler.Contact) error {
	return nil
}

func (d *db) EditContact(c handler.Contact) error {
	return nil
}

func (d *db) DeleteContact(id int) error {
	return nil
}
