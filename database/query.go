package database

import (
	"fmt"

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
func (d *Db) CreateContact(c []app.Contact) (int64, error) {
	d.Conn()
	d.VerifiqueConnection()
	var qtdCreated int64
	qtdCreated = 0

	stmt, err := d.db.Prepare("INSERT INTO user(nome) VALUE(?)")
	if err != nil {
		return 0, fmt.Errorf("Falha ao prepara %v => ", err)
	}
	for _, v := range c {

		res, err := stmt.Exec(v.Person.Name)
		if err != nil {
			return 0, fmt.Errorf("Falha ao inserir %v => ", err)
		}
		qtd, _ := res.RowsAffected()
		qtdCreated += qtd
	}
	return qtdCreated, nil
}

// EditContact ...
func (d *Db) EditContact(c app.Contact) error {
	return nil
}

// DeleteContact ....
func (d *Db) DeleteContact(id int) (int64, error) {
	d.Conn()
	defer d.db.Close()

	d.VerifiqueConnection()

	stmt, err := d.db.Prepare("DELETE FROM user WHERE id=?")
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

// VerifiqueConnection ...
func (d Db) VerifiqueConnection() {
	if err := d.db.Ping(); err != nil {
		fmt.Printf("Error => %v", err)
	}
}
