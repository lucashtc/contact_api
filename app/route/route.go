package route

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lucashtc/contact_api/app/handler"
)

// Routers ...
func Routers() {
	m := mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/", handler.Index)
	m.HandleFunc("/de", handler.DeleteContact)

	srv := &http.Server{
		Handler:      m,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Erro ao iniciar server  => ", err)
	}
}
