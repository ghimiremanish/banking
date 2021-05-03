package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name"`
	City string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func getAllCustomers(w http.ResponseWriter, r *http.Request){
	/*
		list of our customers
	*/
	customers := []Customer {
		{"Ram","Ktm","44600"},
		{"shyam","pkr","44602"},
		{"hari","Ktm","44601"},
	}

	w.Header().Add("Content-Type","application/json")
	/*
		encoding all customer into json format
	*/
	json.NewEncoder(w).Encode(customers)
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello world")
}