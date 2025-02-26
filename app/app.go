package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// Create new multiplexer
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	// Defining routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	// Starting server
	log.Fatal(http.ListenAndServe(":9000", router))
}
