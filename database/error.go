package database

import (
	"log"
)

// Check error ...
func Check(e error) {
	if e != nil {
		log.Fatal("\nError => ", e)
	}
}
