package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/length", convertLength)

	log.Print("Starting server on :5000")

	err := http.ListenAndServe(":5000", mux)

	log.Fatal(err)

}
