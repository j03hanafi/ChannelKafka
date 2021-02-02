package main

import "github.com/gorilla/mux"

func server() *mux.Router {
	router := mux.NewRouter()

	// Endpoints

	//JSON Client

	// ISO Client
	router.HandleFunc("/payment/channel/iso", sendIso).Methods("POST")
	router.HandleFunc("/payment/channel/file", sendFile).Methods("POST")
	router.HandleFunc("/payment/channel/upload", uploadFile).Methods("POST")
	router.HandleFunc("/payment/channel/json", sendJSON).Methods("POST")

	//Chipsakti
	router.HandleFunc("/ppob/inquiry", ppobInquiry).Methods("POST")
	router.HandleFunc("/ppob/payment", ppobPayment).Methods("POST")
	router.HandleFunc("/ppob/status", ppobStatus).Methods("POST")
	router.HandleFunc("/topup/buy", topupBuy).Methods("POST")
	router.HandleFunc("/topup/check", topupCheck).Methods("POST")

	return router
}
