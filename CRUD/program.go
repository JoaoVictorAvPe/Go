package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.GetUser).Methods(http.MethodGet)

	fmt.Println("Linstening to port 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}