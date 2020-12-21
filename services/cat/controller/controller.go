package controller

import (
	cats "MeowGoWithDB/services/cat/service"

	"github.com/gorilla/mux"
)

//Controller func
func Controller(r *mux.Router) {
	r.StrictSlash(true)

	r.HandleFunc("", cats.GetAll).Methods("GET")
	r.HandleFunc("/all", cats.GetAllWithDeleted).Methods("GET")
	r.HandleFunc("/{id}", cats.GetByID).Methods("GET")
	r.HandleFunc("", cats.Create).Methods("POST")
	r.HandleFunc("/{id}", cats.Update).Methods("PUT")
	r.HandleFunc("/{id}", cats.Delete).Methods("DELETE")
}
