package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/gotutorial/phonebook-api/services"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", Login).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}