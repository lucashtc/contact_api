package database

import (
	"log"
)

// Check error default ...
func Check(e error) {
	if e != nil {
		log.Fatal("\nError => ", e)
	}
}

// CheckError return error
func CheckError(e error) error {
	if e != nil {
		return e
	}
	return nil
}
