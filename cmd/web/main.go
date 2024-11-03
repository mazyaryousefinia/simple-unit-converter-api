package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/length", convertLength)
	mux.HandleFunc("/temp", convertTemperature)
	mux.HandleFunc("/weight", convertWeight)
	log.Print("Starting server on :5000")

	err := http.ListenAndServe(":5000", mux)

	log.Fatal(err)

}
