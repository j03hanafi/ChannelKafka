package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Handler to PPOB Inquiry request
func ppobInquiry(w http.ResponseWriter, r *http.Request) {

	ipReq := getIP(r)

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBInquiryRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New PPOB Inquiry request from %v (Customer_No: %v) @%v", ipReq, request.CustomerNo, time.Now().Format("01-02 15:04:05.000000"))

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBInquiry(request)

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest
	//fmt.Println("PPOB Inquiry sleep for 5 seconds to process response from channel")
	//time.Sleep(5 * time.Second)

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
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, response.Nopel)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "PPOB Inquiry Unsuccessful"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)

		responseFormatter(w, response, 200)
	}
}

// Handler to PPOB Payment request
func ppobPayment(w http.ResponseWriter, r *http.Request) {

	ipReq := getIP(r)

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBPaymentRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New PPOB Payment request from %v (Customer_No: %v) @%v", ipReq, request.CustomerNo, time.Now().Format("01-02 15:04:05.000000"))

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBPayment(request)

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest
	//fmt.Println("PPOB Payment sleep for 2 seconds to process response from channel")
	//time.Sleep(2 * time.Second)

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
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, response.Nopel)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "PPOB Payment Unsuccessful"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)

		responseFormatter(w, response, 200)
	}
}

// Handler to PPOB Status request
func ppobStatus(w http.ResponseWriter, r *http.Request) {

	ipReq := getIP(r)

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBStatusRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New PPOB Status request from %v (Customer_No: %v) @%v", ipReq, request.CustomerNo, time.Now().Format("01-02 15:04:05.000000"))

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBStatus(request)

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("PPOB Status Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonPPOBStatus(isoResponse)
		desc := "PPOB Status Success"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, response.Nopel)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "PPOB Status Unsuccessful"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)

		responseFormatter(w, response, 200)
	}
}

// Handler to Topup Buy request
func topupBuy(w http.ResponseWriter, r *http.Request) {

	ipReq := getIP(r)

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request TopupBuyRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New Topup Buy request from %v (Customer_No: %v) @%v", ipReq, request.CustomerNo, time.Now().Format("01-02 15:04:05.000000"))

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoTopupBuy(request)

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("Topup Buy Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonTopupBuy(isoResponse)
		desc := "Topup Buy Success"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "Topup Buy Unsuccessful"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)

		responseFormatter(w, response, 200)
	}
}

// Handler to Topup Check request
func topupCheck(w http.ResponseWriter, r *http.Request) {

	ipReq := getIP(r)

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request TopupCheckRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New Topup Check request from %v (Customer_No: %v) @%v", ipReq, request.CustomerNo, time.Now().Format("01-02 15:04:05.000000"))

	// Convert request JSON to ISO message and create request file
	reqISO := getIsoTopupCheck(request)

	// Send ISO8583 request to channelChan
	channelChan <- reqISO

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("Topup Check Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonTopupCheck(isoResponse)
		desc := "Topup Check Success"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "Topup Check Unsuccessful"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)

		responseFormatter(w, response, 200)
	}
}
