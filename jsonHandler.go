package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler to PPOB Inquiry request
func ppobInquiry(w http.ResponseWriter, r *http.Request) {

	log.Println("New PPOB Inquiry request")

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBInquiryRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBInquiry(request)

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("PPOB Inquiry Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonPPOBInquiry(isoResponse)
		desc := "PPOB Inquiry Success"
		log.Println(desc)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "PPOB Inquiry Unsuccessful"
		log.Println(desc)

		responseFormatter(w, response, 200)
	}
}

// Handler to PPOB Payment request
func ppobPayment(w http.ResponseWriter, r *http.Request) {

	log.Println(" New PPOB Payment request")

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBPaymentRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBPayment(request)

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("PPOB Payment Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonPPOBPayment(isoResponse)
		desc := "PPOB Payment Success"
		log.Println(desc)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "PPOB Payment Unsuccessful"
		log.Println(desc)

		responseFormatter(w, response, 200)
	}
}

// Handler to PPOB Status request
func ppobStatus(w http.ResponseWriter, r *http.Request) {

	log.Println("New PPOB Status request")

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBStatusRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBStatus(request)

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest

	// Get response from consumerChan
	msg := <-consumerChan

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonPPOBStatus(isoResponse)
		desc := "PPOB Status Success"
		log.Println(desc)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "PPOB Status Unsuccessful"
		log.Println(desc)

		responseFormatter(w, response, 200)
	}
}

// Handler to Topup Buy request
func topupBuy(w http.ResponseWriter, r *http.Request) {

	log.Println(" New Topup Buy request")

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request TopupBuyRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoTopupBuy(request)

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest

	// Get response from consumerChan
	msg := <-consumerChan

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonTopupBuy(isoResponse)
		desc := "Topup Buy Success"
		log.Println(desc)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "Topup Buy Unsuccessful"
		log.Println(desc)

		responseFormatter(w, response, 200)
	}
}

// Handler to Topup Checl request
func topupCheck(w http.ResponseWriter, r *http.Request) {

	log.Println(" New Topup Check request")

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request TopupCheckRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	// Convert request JSON to ISO message and create request file
	reqISO := getIsoTopupCheck(request)

	// Send ISO8583 request to channelChan
	channelChan <- reqISO

	// Get response from consumerChan
	msg := <-consumerChan

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonTopupCheck(isoResponse)
		desc := "Topup Check Success"
		log.Println(desc)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "Topup Check Unsuccessful"
		log.Println(desc)

		responseFormatter(w, response, 200)
	}
}
