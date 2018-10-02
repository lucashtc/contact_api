package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type db struct {
	db *sql.DB
}

func (d *db) Conn() {
	db, err := sql.Open("mysql", "root:@/contact_api")
	if err != nil {
		log.Fatal(err)
	}
	d.db = db
}
