package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/mofax/iso8583"
	"github.com/rivo/uniseg"
)

// Return ISO Message by converting data from map[int]string
func getIso(data map[int]string, mti string) (iso iso8583.IsoStruct) {
	log.Println("Converting to ISO8583...")

	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if isoStruct.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	// Compare request data length and spec data length, add padding if different
	for field, data := range data {

		fieldSpec := spec.fields[field]

		// Check length for field with Length Type "fixed"
		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(field, fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					// Add padding for numeric field
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					// Add padding for non-numeric field
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		// Add data to isoStruct
		isoStruct.AddField(int64(field), data)
	}

	// Add MTI to isoStruct
	isoStruct.AddMTI(mti)

	// Logging isoStruct field and value
	printSortedDE(isoStruct)
	log.Println("Convert Success")

	return isoStruct
}

func structRintisToIso(jsonRequest rintisRequest) (string, string) {
	log.Println("Converting rintis Request JSON Request to ISO8583")
	log.Printf("Topup Check Request (JSON): %v\n", jsonRequest)

	// Conv amount to int
	amount := strconv.Itoa(jsonRequest.TotalAmount)

	// Mapping struct
	request := map[int]string{
		2:   jsonRequest.Pan,
		3:   jsonRequest.ProcessingCode,
		4:   amount,
		7:   jsonRequest.TransmissionDateTime,
		11:  jsonRequest.Stan,
		12:  jsonRequest.LocalTransactionTime,
		13:  jsonRequest.LocalTransactionDate,
		17:  jsonRequest.CaptureDate,
		32:  jsonRequest.AcquirerID,
		35:  jsonRequest.Track2Data,
		37:  jsonRequest.Refnum,
		41:  jsonRequest.TerminalID,
		43:  jsonRequest.CardAcceptorData,
		48:  jsonRequest.AdditionalData,
		49:  jsonRequest.Currency,
		52:  jsonRequest.PIN,
		60:  jsonRequest.TerminalData,
		103: jsonRequest.AccountTo,
		126: jsonRequest.TokenData,
	}

	// Add MTI
	mti := "0200"

	// Converting request map to isoStruct
	isoStruct := getIso(request, mti)

	// Add len header
	isoMessage, _ := isoStruct.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoRequest := isoHeader + isoMessage

	// Create file from request
	filename := "Request_from_ePayRintis@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/request/"+filename, isoRequest)
	log.Println("Request file: ", file)

	log.Printf("Topup Check Request (ISO8583): %s\n", isoRequest)
	return isoRequest, jsonRequest.Stan

}

// Log sorted converted ISO Message
func printSortedDE(parsedMessage iso8583.IsoStruct) {
	dataElement := parsedMessage.Elements.GetElements()
	int64toSort := make([]int, 0, len(dataElement))
	for key := range dataElement {
		int64toSort = append(int64toSort, int(key))
	}
	sort.Ints(int64toSort)
	for _, key := range int64toSort {
		log.Printf("[%v] : %v\n", int64(key), dataElement[int64(key)])
	}
}

// Return ISO Struct for response and create response file
func parseResponse(response string) (isoResponse iso8583.IsoStruct) {
	log.Println("Parsing response...")

	header := response[0:4]
	data := response[4:]
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)

	isoResponse, err := isoStruct.Parse(data)
	if err != nil {
		log.Printf("Error parsing iso message. Error: %v", err)
	}

	isoMsg, err := isoResponse.ToString()
	if err != nil {
		log.Printf("Iso Parsed failed convert to string. Error: %v", err)
	}

	// Create file from response
	event := header + isoMsg
	filename := "Response_from_" + isoResponse.Elements.GetElements()[3] + "@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/response/"+filename, event)
	log.Println("Response file: ", file)

	return isoResponse
}

// Spec contains a structured description of an iso8583 spec
// properly defined by a spec file
type Spec struct {
	fields map[int]fieldDescription
}

// readFromFile reads a yaml specfile and loads
// and iso8583 spec from it
func (s *Spec) readFromFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	yaml.Unmarshal(content, &s.fields) // expecting content to be valid yaml
	return nil
}

// SpecFromFile returns a brand new empty spec
func specFromFile(filename string) (Spec, error) {
	s := Spec{}
	err := s.readFromFile(filename)
	if err != nil {
		return s, err
	}
	return s, nil
}

// Add pad on left of data,
// Used to format number by adding "0" in front of number data
func leftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return padding + s
}

// Add pad on right of data,
// Used to format string by adding " " at the end of string data
func rightPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return s + padding
}
