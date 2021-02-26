package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
		return
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

//GoRoutine Refactor Handler
func goTest(w http.ResponseWriter, r *http.Request) {
	log.Println("GoRoutine Endpoint requested")

	//get body json
	body, _ := ioutil.ReadAll(r.Body)

	var request GoRoutineReq
	var response GoRoutineRes

	//unmarshal json request with GoRoutineReq struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("JSON Unmarshal Error: %s\n", err.Error())
		response.Response = "JSON Unmarshal Error"
		responseFormatter(w, response, 500)
	}

	//Convert request message to ISO
	requestIso := "This is ISO message. The message is " + request.Data

	// send request to channelChan
	channelChan <- requestIso

	// get response
	response.Response = <-consumerChan
	log.Println("Goroutine Passed!")
	responseFormatter(w, response, 200)

}

//ChipSakti
//PPOB Inquiry
func ppobInquiry(w http.ResponseWriter, r *http.Request) {

	log.Println("PPOB Inquiry requested")

	// get body json
	body, _ := ioutil.ReadAll(r.Body)

	var response Response
	var reqBody PPOBInquiryRequest
	//var resBody PPOBInquiryResponse

	// unmarshal json with request struct
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	// convert json to iso
	reqISO := convIsoPPOBInquiry(reqBody)

	err = doProducer(broker, "chipsakti-channel", reqISO)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		log.Println(err)
		responseFormatter(w, response, 500)
	} else {

		msg, err := consumeResponse(broker, group, []string{"chipsakti-biller"})
		if err != nil {
			errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(w, response, 500)
		} else {
			// Parse response string to ISO8583 data
			header := msg[0:4]
			data := msg[4:]

			isoStruct := iso8583.NewISOStruct("spec1987.yml", true)

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

			rc := isoParsed.Elements.GetElements()[39]

			if rc == "00" {
				Resp := convJsonPPOBInquiry(isoParsed)

				desc := "PPOB Inquiry Success"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			} else {
				Resp := convJsonUnsuccessfulChipsakti(isoParsed)

				desc := "PPOB Inquiry Unsuccessful"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			}

		}

	}
}

//PPOBPayment
func ppobPayment(w http.ResponseWriter, r *http.Request) {

	log.Println("PPOB Payment requested")

	// get body json
	body, _ := ioutil.ReadAll(r.Body)

	var response Response
	var reqBody PPOBPaymentRequest
	//var resBody PPOBInquiryResponse

	// unmarshal json with request struct
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	// convert json to iso
	reqISO := convIsoPPOBPayment(reqBody)

	err = doProducer(broker, "chipsakti-channel", reqISO)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		log.Println(err)
		responseFormatter(w, response, 500)
	} else {

		msg, err := consumeResponse(broker, group, []string{"chipsakti-biller"})
		if err != nil {
			errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(w, response, 500)
		} else {
			// Parse response string to ISO8583 data
			header := msg[0:4]
			data := msg[4:]

			isoStruct := iso8583.NewISOStruct("spec1987.yml", true)

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

			rc := isoParsed.Elements.GetElements()[39]

			if rc == "00" {
				Resp := convJsonPPOBPayment(isoParsed)

				desc := "PPOB Payment Success"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			} else {
				Resp := convJsonUnsuccessfulChipsakti(isoParsed)

				desc := "PPOB Payment Unsuccessful"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			}

		}

	}
}

//PPOBStatus
func ppobStatus(w http.ResponseWriter, r *http.Request) {

	log.Println("PPOB Status requested")

	// get body json
	body, _ := ioutil.ReadAll(r.Body)

	var response Response
	var reqBody PPOBStatusRequest
	//var resBody PPOBInquiryResponse

	// unmarshal json with request struct
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	// convert json to iso
	reqISO := convIsoPPOBStatus(reqBody)

	err = doProducer(broker, "chipsakti-channel", reqISO)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		log.Println(err)
		responseFormatter(w, response, 500)
	} else {

		msg, err := consumeResponse(broker, group, []string{"chipsakti-biller"})
		if err != nil {
			errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(w, response, 500)
		} else {
			// Parse response string to ISO8583 data
			header := msg[0:4]
			data := msg[4:]

			isoStruct := iso8583.NewISOStruct("spec1987.yml", true)

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

			rc := isoParsed.Elements.GetElements()[39]

			if rc == "00" {
				Resp := convJsonPPOBStatus(isoParsed)

				desc := "PPOB Status Success"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			} else {
				Resp := convJsonUnsuccessfulChipsakti(isoParsed)

				desc := "PPOB Status Unsuccessful"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			}

		}

	}
}

//TopupBuy
func topupBuy(w http.ResponseWriter, r *http.Request) {

	log.Println("Topup Buy requested")

	// get body json
	body, _ := ioutil.ReadAll(r.Body)

	var response Response
	var reqBody TopupBuyRequest
	//var resBody PPOBInquiryResponse

	// unmarshal json with request struct
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	// convert json to iso
	reqISO := convIsoTopupBuy(reqBody)

	err = doProducer(broker, "chipsakti-channel", reqISO)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		log.Println(err)
		responseFormatter(w, response, 500)
	} else {

		msg, err := consumeResponse(broker, group, []string{"chipsakti-biller"})
		if err != nil {
			errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(w, response, 500)
		} else {
			// Parse response string to ISO8583 data
			header := msg[0:4]
			data := msg[4:]

			isoStruct := iso8583.NewISOStruct("spec1987.yml", true)

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

			rc := isoParsed.Elements.GetElements()[39]

			if rc == "00" {
				Resp := convJsonTopupBuy(isoParsed)

				desc := "Topup Buy Success"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			} else {
				Resp := convJsonUnsuccessfulChipsakti(isoParsed)

				desc := "Topup Buy Unsuccessful"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			}
		}

	}
}

//TopupCheck
func topupCheck(w http.ResponseWriter, r *http.Request) {

	log.Println("Topup Check requested")

	// get body json
	body, _ := ioutil.ReadAll(r.Body)

	var response Response
	var reqBody TopupCheckRequest
	//var resBody PPOBInquiryResponse

	// unmarshal json with request struct
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	// convert json to iso
	reqISO := convIsoTopupCheck(reqBody)

	err = doProducer(broker, "chipsakti-channel", reqISO)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		log.Println(err)
		responseFormatter(w, response, 500)
	} else {

		msg, err := consumeResponse(broker, group, []string{"chipsakti-biller"})
		if err != nil {
			errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(w, response, 500)
		} else {
			// Parse response string to ISO8583 data
			header := msg[0:4]
			data := msg[4:]

			isoStruct := iso8583.NewISOStruct("spec1987.yml", true)

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

			rc := isoParsed.Elements.GetElements()[39]

			if rc == "00" {
				Resp := convJsonTopupCheck(isoParsed)

				desc := "Topup Check Success"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			} else {
				Resp := convJsonUnsuccessfulChipsakti(isoParsed)

				desc := "Topup Check Unsuccessful"
				log.Println(desc)

				responseFormatter(w, Resp, 200)
			}

		}

	}
}
