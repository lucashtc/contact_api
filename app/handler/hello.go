package handler

import (
	"fmt"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'msg':'Welcome contact API'}")
}
