package main

import "github.com/gorilla/mux"

// return HTTP handler
func pathHandler() *mux.Router {

	// create new handler instance
	router := mux.NewRouter()

	// Endpoints, Handler function, and HTTP request Method
	router.HandleFunc("/epay/rintis", rintisReqHandler).Methods("POST")
	router.HandleFunc("/json/adapter_iso20022", adapterReqHandler).Methods("POST")

	return router
}
