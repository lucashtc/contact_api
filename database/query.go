package database

import "github.com/lucashtc/contact_api/app"

func (d *Db) GetAll() []app.Contact {
	var c []app.Contact
	return c
}

func (d *Db) GetContact(id int) app.Contact {
	var c app.Contact
	return c
}

func (d *Db) CreateContact(c app.Contact) error {
	return nil
}

func (d *Db) EditContact(c app.Contact) error {
	return nil
}

// DeleteContact ...
func (d *Db) DeleteContact(id int) int64 {
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
