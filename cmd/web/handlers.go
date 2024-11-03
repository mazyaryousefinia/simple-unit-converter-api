package main

import (
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

func convertTemperature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fromUnit := r.FormValue("temperatureFromUnit")
	toUnit := r.FormValue("temperatureToUnit")
	value, err := strconv.ParseFloat(r.FormValue("temperatureInput"), 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	tempConverter := converter.TemperatureConvert{
		FromUnit: fromUnit,
		ToUnit:   toUnit,
		Value:    value,
	}

	result, err := tempConverter.Convert()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

func convertWeight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fromUnit := r.FormValue("weightFromUnit")
	toUnit := r.FormValue("weightToUnit")
	value, err := strconv.ParseFloat(r.FormValue("weightInput"), 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	weightConverter := converter.WeightConverter{
		FromUnit: fromUnit,
		ToUnit:   toUnit,
		Value:    value,
	}

	result, err := weightConverter.Convert()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

func convertLength(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fromUnit := r.FormValue("lengthFromUnit")
	toUnit := r.FormValue("lengthToUnit")
	value, err := strconv.ParseFloat(r.FormValue("lengthInput"), 64)
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
