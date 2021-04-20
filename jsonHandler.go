package main

import (
	"encoding/json"
	"github.com/mofax/iso8583"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func rintisReqHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New Request To Rintis")

	// Get msg from body
	body, _ := ioutil.ReadAll(r.Body)
	var request rintisRequest

	// Unmarshal body to struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	// Run convertion
	reqISO, head := structRintisToIso(request)

	// Put header and message on variable
	data := resConsume{
		Head:    head,
		Content: reqISO,
	}
	// Send variable to channel to produce to kafka
	channelArrChan <- data

	// Get data from arr consume (tempStorage) with same head
	isoResponse := getValueFromArray(head)

	if (resConsume{} == isoResponse) {
		responseFormatter(w, "failed, esception timeout", 504)
		return
	}

	isoStructResponse := parseResponse(isoResponse.Content)

	structResponse := makeRinRespFromIso(isoStructResponse)

	// Make response to request
	responseFormatter(w, structResponse, 200)
}

func makeRinRespFromIso(parsedIso iso8583.IsoStruct) (response rintisResponse) {
	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Pan = strings.Trim(emap[2], " ")
	response.ProcessingCode = strings.Trim(emap[3], " ")
	response.TotalAmount, _ = strconv.Atoi(emap[4])
	response.TransmissionDateTime = strings.Trim(emap[7], " ")
	response.Stan = strings.Trim(emap[11], " ")
	response.LocalTransactionTime = strings.Trim(emap[12], " ")
	response.LocalTransactionDate = strings.Trim(emap[13], " ")
	response.SettlementDate = strings.Trim(emap[15], " ")
	response.CaptureDate = strings.Trim(emap[17], " ")
	response.AcquirerID = strings.Trim(emap[32], " ")
	response.Track2Data = strings.Trim(emap[35], " ")
	response.Refnum = strings.Trim(emap[37], " ")
	response.AuthIdResponse = strings.Trim(emap[38], " ")
	response.ResponseCode = strings.Trim(emap[39], " ")
	response.TerminalID = strings.Trim(emap[41], " ")
	response.AdditionalResponseData = strings.Trim(emap[44], " ")
	response.Currency = strings.Trim(emap[49], " ")
	response.TerminalData = strings.Trim(emap[60], " ")
	response.ReceivingInstitutionIDCode = strings.Trim(emap[100], " ")
	response.AccountTo = strings.Trim(emap[103], " ")
	response.TokenData = strings.Trim(emap[126], " ")
	log.Println("Convert success")
	log.Printf("Topup Check Response (JSON): %+v\n", response)
	return response

}
