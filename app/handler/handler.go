package handler

import (
	"encoding/json"
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

	var c []app.Contact

	//Json deve ser um array
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		fmt.Printf("Error ao decodificar json %v => ", err)
	}

	qtd, err := d.CreateContact(c)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Fprintf(w, "resultado: '%v'", qtd)

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
