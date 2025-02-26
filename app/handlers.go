package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Data Transfer Object (DTO)
type Customer struct {
	Name    string `json:"full_name" "xml:name`
	City    string `json:"city" xml:"city`
	ZipCode string `json:"zip_code" xml:"zipcode`
}

func greet(w http.ResponseWriter, r *http.Request) {
	response := "Hello World!!"
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(response)))
	fmt.Fprint(w, response)
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Samaresh", City: "Austin", ZipCode: "78717"},
		{Name: "Amrita", City: "San Antanio", ZipCode: "89876"},
		{Name: "Prisha", City: "Dallas", ZipCode: "76758"},
		{Name: "Shreya", City: "Houston", ZipCode: "87098"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
