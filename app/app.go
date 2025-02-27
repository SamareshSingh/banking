package app

import (
	"log"
	"net/http"

	"github.com/SamareshSingh/banking/domain"
	"github.com/SamareshSingh/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	// Create new multiplexer
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// Wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Defining routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// Starting server
	log.Fatal(http.ListenAndServe(":9000", router))
}
