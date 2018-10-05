package database

import "github.com/lucashtc/contact_api/app"

func (d *db) GetAll() []app.Contact {
	var c []app.Contact
	return c
}

func (d *db) GetContact(id int) app.Contact {
	var c app.Contact
	return c
}

func (d *db) CreateContact(c app.Contact) error {
	return nil
}

func (d *db) EditContact(c app.Contact) error {
	return nil
}

// DeleteContact ...
func (d *db) DeleteContact(id int) int64 {
	d.Conn()
	defer d.db.Close()

	stmt, err := d.db.Prepare("DELETE user where id=?")
	Check(err)

	res, err := stmt.Exec(id)
	Check(err)

	qtd, err := res.RowsAffected()
	Check(err)

	return qtd
}
