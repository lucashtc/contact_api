package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Db ...
type Db struct {
	db *sql.DB
}

// Conn connection database
func (d *Db) Conn() {
	db, err := sql.Open("mysql", "root:@/contact_api")
	fmt.Println(db.Ping())
	if err != nil {
		log.Fatal(err)
	}
	d.db = db
}
