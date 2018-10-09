package database

import (
	"github.com/lucashtc/contact_api/app"
)

// GetAll ..
func (d *Db) GetAll() []app.Contact {
	var c []app.Contact
	return c
}

// GetContact ...
func (d *Db) GetContact(id int) app.Contact {
	var c app.Contact
	return c
}

// CreateContact ...
func (d *Db) CreateContact(c app.Contact) error {
	return nil
}

// EditContact ...
func (d *Db) EditContact(c app.Contact) error {
	return nil
}

// DeleteContact ....
func (d *Db) DeleteContact(id int) (int64, error) {
	d.Conn()
	defer d.db.Close()

	stmt, err := d.db.Prepare("DELETE user where id=?")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	qtd, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return qtd, nil
}
