package main

import "github.com/gorilla/mux"

// return HTTP handler
func server() *mux.Router {

	// create new handler instance
	router := mux.NewRouter()

	// Endpoints, Handler function, and HTTP request Method
	router.HandleFunc("/ppob/inquiry", ppobInquiry).Methods("POST")
	router.HandleFunc("/ppob/payment", ppobPayment).Methods("POST")
	router.HandleFunc("/ppob/status", ppobStatus).Methods("POST")
	router.HandleFunc("/topup/buy", topupBuy).Methods("POST")
	router.HandleFunc("/topup/check", topupCheck).Methods("POST")
	router.HandleFunc("/epay/rintis", rintis).Methods("POST")

	return router
}
