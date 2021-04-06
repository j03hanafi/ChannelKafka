package main

import (
	"encoding/json"
	"github.com/mofax/iso8583"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Handler to PPOB Inquiry request
func ppobInquiry(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	ipReq := getIP(r)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Received new PPOB Inquiry Request\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBInquiryRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New PPOB Inquiry request from %v (Customer_No: %v)", ipReq, request.CustomerNo)

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBInquiry(request)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request converted from JSON to ISO\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request sent to Producer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	//fmt.Println("PPOB Inquiry sleep for 5 seconds to process response from channel")
	//time.Sleep(5 * time.Second)

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response received from consumer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	log.Printf("PPOB Inquiry Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response parsed from ISO to JSON\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonPPOBInquiry(isoResponse)
		desc := "PPOB Inquiry Success"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, response.Nopel)
		log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response Sent\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

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

	start := time.Now()
	ipReq := getIP(r)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Received new PPOB Payment Request\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBPaymentRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New PPOB Payment request from %v (Customer_No: %v)", ipReq, request.CustomerNo)

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBPayment(request)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request converted from JSON to ISO\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request sent to Producer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	//fmt.Println("PPOB Payment sleep for 2 seconds to process response from channel")
	//time.Sleep(2 * time.Second)

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response received from consumer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	log.Printf("PPOB Payment Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response parsed from ISO to JSON\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonPPOBPayment(isoResponse)
		desc := "PPOB Payment Success"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, response.Nopel)
		log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response Sent\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

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

	start := time.Now()
	ipReq := getIP(r)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Received new PPOB Status Request\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request PPOBStatusRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New PPOB Status request from %v (Customer_No: %v)", ipReq, request.CustomerNo)

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoPPOBStatus(request)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request converted from JSON to ISO\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request sent to Producer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response received from consumer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	log.Printf("PPOB Status Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response parsed from ISO to JSON\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonPPOBStatus(isoResponse)
		desc := "PPOB Status Success"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, response.Nopel)
		log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response Sent\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

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

	start := time.Now()
	ipReq := getIP(r)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Received new Topup Buy Request\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request TopupBuyRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New Topup Buy request from %v (Customer_No: %v)", ipReq, request.CustomerNo)

	// Convert request JSON to ISO message and create request file
	isoRequest := getIsoTopupBuy(request)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request converted from JSON to ISO\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Send ISO8583 request to channelChan
	channelChan <- isoRequest
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request sent to Producer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response received from consumer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	log.Printf("Topup Buy Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response parsed from ISO to JSON\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonTopupBuy(isoResponse)
		desc := "Topup Buy Success"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)
		log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response Sent\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

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

	start := time.Now()
	ipReq := getIP(r)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Received new Topup Check Request\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	var request TopupCheckRequest

	// Unmarshal JSON with request struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	log.Printf("New Topup Check request from %v (Customer_No: %v)", ipReq, request.CustomerNo)

	// Convert request JSON to ISO message and create request file
	reqISO := getIsoTopupCheck(request)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request converted from JSON to ISO\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Send ISO8583 request to channelChan
	channelChan <- reqISO
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Request sent to Producer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get response from consumerChan
	msg := <-consumerChan
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response received from consumer\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	log.Printf("Topup Check Response (ISO8583): %s\n", msg)

	// Parse response string to ISO8583 data and create response file
	isoResponse := parseResponse(msg)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response parsed from ISO to JSON\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())
	rc := isoResponse.Elements.GetElements()[39]

	// Check for Response Code (RC)
	if rc == "00" {
		response := getJsonTopupCheck(isoResponse)
		desc := "Topup Check Success"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)
		log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Response Sent\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

		responseFormatter(w, response, 200)
	} else {
		response := getJsonUnsuccessfulChipsakti(isoResponse)
		desc := "Topup Check Unsuccessful"
		log.Printf("%v to %v (Customer_No: %v)\n", desc, ipReq, request.CustomerNo)

		responseFormatter(w, response, 200)
	}
}

// Return unsuccessful Chipsakti Response (rc != "00")
func getJsonUnsuccessfulChipsakti(parsedIso iso8583.IsoStruct) UnsuccessfulChipsakti {
	var response UnsuccessfulChipsakti

	log.Println("Converting Unsuccessful Chipsakti ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Rc = strings.Trim(emap[39], " ")
	response.Restime = emap[48]
	response.Msg = emap[120]

	log.Println("Convert success")
	log.Printf("Unsuccessful Chipsakti Response (JSON): %+v\n", response)
	return response
}

// Return JSON for PPOB Inquiry ISO8583 response
func getJsonPPOBInquiry(parsedIso iso8583.IsoStruct) PPOBInquiryResponse {
	var response PPOBInquiryResponse

	log.Println("Converting PPOB Inquiry ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Convert non-string field
	tagihan, _ := strconv.Atoi(emap[4])
	admin, _ := strconv.Atoi(emap[5])
	totalTagihan, _ := strconv.Atoi(emap[6])

	// Map ISO8583 format to JSON data
	response.Tagihan = tagihan
	response.Admin = admin
	response.TotalTagihan = totalTagihan
	response.Reffid = strings.Trim(emap[37], " ")
	response.Rc = strings.Trim(emap[39], " ")
	response.Nama = strings.Trim(emap[43], " ")
	response.Restime = emap[48]
	response.Data = emap[62]
	response.Msg = emap[120]
	response.Produk = strings.Trim(emap[121], " ")
	response.Nopel = strings.Trim(emap[122], " ")

	log.Println("Convert success")
	log.Printf("PPOB Inquiry Response (JSON): %+v\n", response)
	return response
}

// Return JSON for PPOB Payment ISO8583 response
func getJsonPPOBPayment(parsedIso iso8583.IsoStruct) PPOBPaymentResponse {
	var response PPOBPaymentResponse

	log.Println("Converting PPOB Payment ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Convert non-string field
	tagihan, _ := strconv.Atoi(emap[4])
	admin, _ := strconv.Atoi(emap[5])
	totalTagihan, _ := strconv.Atoi(emap[6])
	struk := strings.Split(emap[62], ",")

	// Map ISO8583 format to JSON data
	response.Tagihan = tagihan
	response.Admin = admin
	response.TotalTagihan = totalTagihan
	response.Reffid = strings.Trim(emap[37], " ")
	response.Rc = strings.Trim(emap[39], " ")
	response.Nama = strings.Trim(emap[43], " ")
	response.TglLunas = strings.Trim(emap[48], " ")
	response.Struk = struk
	response.Msg = strings.Trim(emap[120], " ")
	response.Produk = strings.Trim(emap[121], " ")
	response.Nopel = strings.Trim(emap[122], " ")
	response.ReffNo = strings.Trim(emap[123], " ")

	log.Println("Convert success")
	log.Printf("PPOB Payment Response (JSON): %+v\n", response)
	return response
}

// Return JSON for PPOB Status ISO8583 response
func getJsonPPOBStatus(parsedIso iso8583.IsoStruct) PPOBStatusResponse {
	var response PPOBStatusResponse

	log.Println("Converting PPOB Status ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Convert non-string field
	tagihan, _ := strconv.Atoi(emap[4])
	admin, _ := strconv.Atoi(emap[5])
	totalTagihan, _ := strconv.Atoi(emap[6])
	struk := strings.Split(emap[62], ",")

	// Map ISO8583 format to JSON data
	response.Tagihan = tagihan
	response.Admin = admin
	response.TotalTagihan = totalTagihan
	response.Reffid = strings.Trim(emap[37], " ")
	response.Rc = strings.Trim(emap[39], " ")
	response.Nama = strings.Trim(emap[43], " ")
	response.TglLunas = strings.Trim(emap[48], " ")
	response.Struk = struk
	response.Msg = strings.Trim(emap[120], " ")
	response.Produk = strings.Trim(emap[121], " ")
	response.Nopel = strings.Trim(emap[122], " ")
	response.ReffNo = strings.Trim(emap[123], " ")
	response.Status = strings.Trim(emap[124], " ")

	log.Println("Convert success")
	log.Printf("PPOB Status Response (JSON): %+v\n", response)
	return response
}

// Return JSON for Topup Buy ISO8583 response
func getJsonTopupBuy(parsedIso iso8583.IsoStruct) TopupBuyResponse {
	var response TopupBuyResponse

	log.Println("Converting Topup Buy ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Rc = strings.Trim(emap[39], " ")
	response.Restime = strings.Trim(emap[48], " ")
	response.Msg = strings.Trim(emap[120], " ")
	response.SN = strings.Trim(emap[121], " ")
	response.Price = strings.Trim(emap[122], " ")

	log.Println("Convert success")
	log.Printf("Topup Buy Response (JSON): %+v\n", response)
	return response
}

// Return JSON for Topup Check ISO8583 response
func getJsonTopupCheck(parsedIso iso8583.IsoStruct) TopupCheckResponse {
	var response TopupCheckResponse

	log.Println("Converting Topup Check ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Rc = strings.Trim(emap[39], " ")
	response.Restime = strings.Trim(emap[48], " ")
	response.Msg = strings.Trim(emap[120], " ")
	response.SN = strings.Trim(emap[121], " ")
	response.Price = strings.Trim(emap[122], " ")

	log.Println("Convert success")
	log.Printf("Topup Check Response (JSON): %+v\n", response)
	return response
}
