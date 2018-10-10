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
func Create(w http.ResponseWriter, r *http.Request) {
	c := []app.Contact{}
	c[1].Person.Name = "Alguem"
	c[2].Person.Name = "Eu mesmo"

	qtd, err := d.CreateContact(c)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	fmt.Fprintf(w, "resultado: '%d'", qtd)
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

	qtd, err := d.DeleteContact(4)
	if err != nil {
		fmt.Fprintf(w, "{status: %v}", err)
	}

	if qtd > 0 {
		fmt.Fprintf(w, "{status:'Deletado => %v'}", qtd)
	} else {
		fmt.Fprintf(w, "{status: '%s'}", notDeleted)
	}

}
