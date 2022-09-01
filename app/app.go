package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Printf("Starting new server")

	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/getcustomers", getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/createcustomer", createCustomer).Methods(http.MethodPost)

	http.ListenAndServe("localhost:8080", router)
}
