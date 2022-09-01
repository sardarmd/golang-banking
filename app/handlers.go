package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name"`
	City string `json:"city"`
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hi Good Morning")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{{"Sardar", "Bangalore"}, {"Sadia", "Kanpur"}}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post method received")
}
