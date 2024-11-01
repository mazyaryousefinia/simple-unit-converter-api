package main

import (
	"fmt"
	"github.com/mazyaryousefinia/simple-unit-converter-app/internal/converter"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func convertLength(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fromUnit := r.FormValue("lengthFromUnit")
	toUnit := r.FormValue("lengthToUnit")
	value, err := strconv.ParseFloat(r.FormValue("lengthInput"), 64)
	fmt.Printf("from: %s, to: %s, v: %s", fromUnit, toUnit, value)
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return
	}

	lengthConverter := converter.LengthConverter{
		FromUnit: fromUnit,
		ToUnit:   toUnit,
		Value:    value,
	}

	result, err := lengthConverter.Convert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, map[string]any{
		"Result":   result,
		"FromUnit": fromUnit,
		"ToUnit":   toUnit,
		"Value":    value,
	})

}
