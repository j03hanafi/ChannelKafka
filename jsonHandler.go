package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"log"
	"net/http"
	"github.com/mofax/iso8583"

)

// Handle all JSON Client request
func sendJSON(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var response Response
	var reqBody Transaction

	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	
	resISO, err := convJSON(reqBody)

	if err != nil {
		resp := Response{}
		resp.ResponseCode = 500
		resp.ReasonCode = 0
		resp.ResponseDescription = err.Error()
		responseFormatter(w, resp, 500)
	}

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
			// Parse response string to ISO8583 data
			header := msg[0:4]
			data := msg[4:]

			isoStruct := iso8583.NewISOStruct("spec1987.yml", false)

			isoParsed, err := isoStruct.Parse(data)
			if err != nil {
				log.Printf("Error parsing iso message\nError: %v", err)
			}

			isoMsg, err := isoParsed.ToString()
			if err != nil {
				log.Printf("Iso Parsed failed convert to string.\nError: %v", err)
			}

			// create file from response
			event := header + isoMsg
			filename := "Response_from_" + isoParsed.Elements.GetElements()[3] + "@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
			file := CreateFile("storage/response/"+filename, event)
			log.Println("File created: ", file)

			strRes := fromISO(string(msg))

			responseFormatter(w, strRes, 200)
		}

	}
}
