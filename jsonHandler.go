package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/mofax/iso8583"
)

// Handle all JSON Client request
func sendJSON(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var response Response
	var reqBody Transaction
	var iso Iso8583

	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	resISO, _ := convJSON(reqBody)

	err = doProducer(broker, topic1, resISO)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		log.Println(err)
		responseFormatter(w, response, 500)
	} else {
		msg, err := consumeResponse(broker, group, []string{topic2})
		if err != nil {
			errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(w, response, 500)
		} else {
			if msg == "" {
				errDesc := "Got empty response"
				response.ResponseCode, response.ResponseDescription = 500, errDesc
				log.Println(errDesc)
				responseFormatter(w, response, 500)
			} else {
				header := msg[0:4]
				data := msg[4:]

				isoStruct := iso8583.NewISOStruct("spec1987.yml", false)

				isoParsed, err := isoStruct.Parse(data)
				if err != nil {
					log.Printf("Error parsing iso message\nError: %v", err)
				}

				iso.Header, _ = strconv.Atoi(header)
				iso.MTI = isoParsed.Mti.String()
				iso.Hex, _ = iso8583.BitMapArrayToHex(isoParsed.Bitmap)

				iso.Message, err = isoParsed.ToString()
				if err != nil {
					log.Printf("Iso Parsed failed convert to string.\nError: %v", err)
				}
				desc := "Success"
				iso.ResponseStatus.ResponseCode, iso.ResponseStatus.ResponseDescription = 200, desc
				log.Println(desc)
				responseFormatter(w, iso, 200)
			}
		}
	}
}
