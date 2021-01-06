package main

import (
	"fmt"
	"github.com/mofax/iso8583"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Handle all ISO Client request

// Process ISO message in body request
func sendIso(writer http.ResponseWriter, request *http.Request) {

	var response Response
	var iso Iso8583

	// Read body request
	reqBody, _ := ioutil.ReadAll(request.Body)
	req := string(reqBody)
	log.Println("New Request")
	log.Printf("ISO Message: %v\n", req)

	// Produce event
	err := doProducer(broker, topic1, req)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		log.Println(err)
		responseFormatter(writer, response, 500)
	} else {
		// Read response
		msg, err := consumeResponse(broker, group, []string{topic2})
		if err != nil {
			errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(writer, response, 500)
		} else {

			// Return empty response
			if msg == "" {
				errDesc := "Got empty response"
				response.ResponseCode, response.ResponseDescription = 500, errDesc
				log.Println(errDesc)
				responseFormatter(writer, response, 500)
			} else {

				// Parse response string to ISO8583 data
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

				// create file from response
				event := header + iso.Message
				filename := "Response_to_" + isoParsed.Elements.GetElements()[3] + "@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
				file := CreateFile("storage/response/"+filename, event)
				log.Println("File created: ", file)

				desc := "Success"
				iso.ResponseStatus.ResponseCode, iso.ResponseStatus.ResponseDescription = 200, desc
				log.Println(desc)
				responseFormatter(writer, iso, 200)

			}

		}
	}

}

// Process file path in body request
func sendFile(writer http.ResponseWriter, request *http.Request) {

	var response Response
	var iso Iso8583

	// Read body request
	reqBody, _ := ioutil.ReadAll(request.Body)
	filename := string(reqBody)

	// Check if file in .txt extension
	if !strings.Contains(filename, ".txt") {
		filename += ".txt"
	}

	log.Println("New Request")

	// Check if file exist in storage
	if CheckExist(filename) {
		req := ReadFile(filename)

		log.Printf("ISO Message: %v\n", req)

		// Produce event
		err := doProducer(broker, topic1, req)

		if err != nil {
			errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(writer, response, 500)
		} else {
			// Read response
			msg, err := consumeResponse(broker, group, []string{topic2})
			if err != nil {
				errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
				response.ResponseCode, response.ResponseDescription = 500, errDesc
				log.Println(err)
				responseFormatter(writer, response, 500)
			} else {

				// Return empty response
				if msg == "" {
					errDesc := "Got empty response"
					response.ResponseCode, response.ResponseDescription = 500, errDesc
					log.Println(errDesc)
					responseFormatter(writer, response, 500)
				} else {

					// Parse response string to ISO8583 data
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

					// create file from response
					event := header + iso.Message
					filename := "Response_from_" + isoParsed.Elements.GetElements()[3] + "@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
					file := CreateFile("storage/response/"+filename, event)
					log.Println("File created: ", file)

					desc := "Success"
					iso.ResponseStatus.ResponseCode, iso.ResponseStatus.ResponseDescription = 200, desc
					log.Println(desc)
					responseFormatter(writer, iso, 200)

				}

			}
		}

	} else {

		// File not found
		errDesc := "File not found"
		response.ResponseCode, response.ResponseDescription = 404, errDesc
		log.Println(errDesc)
		log.Println("Process failed")
		responseFormatter(writer, response, 404)

	}

}

// Process uploaded file in body request
func uploadFile(writer http.ResponseWriter, request *http.Request) {

	var response Response
	var iso Iso8583

	log.Println("File Upload Endpoint Hit")

	// Check maximum file upload size
	request.ParseMultipartForm(10 << 20)

	// Get file from body
	uploadedFile, handler, err := request.FormFile("isoFile")
	if err != nil {
		log.Fatalf("Failed to retreive ISO File. Error: %v\n", err)
	}

	defer uploadedFile.Close()
	log.Printf("ISO File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// Read file content
	reqBody, _ := ioutil.ReadAll(uploadedFile)
	filename := "storage/request/" + handler.Filename + "@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))

	// Add .txt extension
	if !strings.Contains(filename, ".txt") {
		filename += ".txt"
	}

	// Create and write the content to new file in storage
	file := CreateFile(filename, string(reqBody))
	log.Printf("Uploaded file saved. %v\n", file)

	log.Println("New Request")

	// Check if file exist
	if CheckExist(filename) {
		// Read file content
		req := ReadFile(filename)

		log.Printf("ISO Message: %v\n", req)

		// Produce event
		err := doProducer(broker, topic1, req)

		if err != nil {
			errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			log.Println(err)
			responseFormatter(writer, response, 500)
		} else {
			// Read response
			msg, err := consumeResponse(broker, group, []string{topic2})
			if err != nil {
				errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
				response.ResponseCode, response.ResponseDescription = 500, errDesc
				log.Println(err)
				responseFormatter(writer, response, 500)
			} else {

				// Return empty response
				if msg == "" {
					errDesc := "Got empty response"
					response.ResponseCode, response.ResponseDescription = 500, errDesc
					log.Println(errDesc)
					responseFormatter(writer, response, 500)
				} else {

					// Parse response string to ISO8583 data
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

					// create file from response
					event := header + iso.Message
					filename := "Response_from_" + isoParsed.Elements.GetElements()[3] + "@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
					file := CreateFile("storage/response/"+filename, event)
					log.Println("File created: ", file)

					desc := "Success"
					iso.ResponseStatus.ResponseCode, iso.ResponseStatus.ResponseDescription = 200, desc
					log.Println(desc)
					responseFormatter(writer, iso, 200)

				}

			}
		}

	} else {

		// Return file not found
		errDesc := "File not found"
		response.ResponseCode, response.ResponseDescription = 404, errDesc
		log.Println(errDesc)
		log.Println("Process failed")
		responseFormatter(writer, response, 404)

	}

}
