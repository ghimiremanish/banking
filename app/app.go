package main

import (
	"log"
	"net/http"
)

func start(){
	/*
		defining all our routes
	*/
	http.HandleFunc("/test", hello)
	http.HandleFunc("/customers", getAllCustomers)

	/*
		starting server
	*/
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}