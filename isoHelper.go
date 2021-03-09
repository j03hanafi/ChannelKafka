package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/mofax/iso8583"
	"github.com/rivo/uniseg"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
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

// Return ISO message for PPOB Inquiry JSON request
func getIsoPPOBInquiry(jsonRequest PPOBInquiryRequest) (isoRequest string) {

	log.Println("Converting PPOB Inquiry JSON Request to ISO8583")
	log.Printf("PPOB Inquiry Request (JSON): %v\n", jsonRequest)

	transactionID := jsonRequest.TransactionID
	partnerID := jsonRequest.PartnerID
	productCode := jsonRequest.ProductCode
	customerNo := jsonRequest.CustomerNo
	merchantCode := jsonRequest.MerchantCode
	requestTime := jsonRequest.RequestTime
	periode := jsonRequest.Periode

	// Checking data length, match it to Spec Document by adding pad
	if len(jsonRequest.TransactionID) < 25 {
		transactionID = rightPad(jsonRequest.TransactionID, 25, " ")
	}
	if len(jsonRequest.PartnerID) < 16 {
		partnerID = rightPad(jsonRequest.PartnerID, 16, " ")
	}
	if len(jsonRequest.ProductCode) < 16 {
		productCode = rightPad(jsonRequest.ProductCode, 16, " ")
	}
	if len(jsonRequest.CustomerNo) < 25 {
		customerNo = rightPad(jsonRequest.CustomerNo, 25, " ")
	}
	if len(jsonRequest.MerchantCode) < 25 {
		merchantCode = rightPad(jsonRequest.MerchantCode, 25, " ")
	}
	if len(jsonRequest.RequestTime) < 19 {
		requestTime = rightPad(jsonRequest.RequestTime, 19, " ")
	}

	// Assign data to map, Adding PAN for PPOB Inquiry Request, and add MTI
	request := map[int]string{
		3:  "380001",
		48: transactionID + partnerID + productCode + customerNo + merchantCode + requestTime + periode,
	}
	mti := "0200"

	// Converting request map to isoStruct
	isoStruct := getIso(request, mti)

	isoMessage, _ := isoStruct.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoRequest = isoHeader + isoMessage

	// Create file from request
	filename := "Request_from_PPOBInquiry@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/request/"+filename, isoRequest)
	log.Println("Request file: ", file)

	log.Println("Convert success")
	log.Printf("PPOB Inquiry Request (ISO8583): %s\n", isoRequest)
	return isoRequest
}

// Return ISO message for PPOB Payment JSON request
func getIsoPPOBPayment(jsonRequest PPOBPaymentRequest) (isoRequest string) {

	log.Println("Converting PPOB Payment JSON Request to ISO8583")
	log.Printf("PPOB Payment Request (JSON): %v\n", jsonRequest)

	transactionID := jsonRequest.TransactionID
	partnerID := jsonRequest.PartnerID
	productCode := jsonRequest.ProductCode
	customerNo := jsonRequest.CustomerNo
	merchantCode := jsonRequest.MerchantCode
	requestTime := jsonRequest.RequestTime
	reffId := jsonRequest.ReffID
	amount := jsonRequest.Amount

	// Checking data length, match it to Spec Document by adding pad
	if len(jsonRequest.TransactionID) < 25 {
		transactionID = rightPad(jsonRequest.TransactionID, 25, " ")
	}
	if len(jsonRequest.PartnerID) < 16 {
		partnerID = rightPad(jsonRequest.PartnerID, 16, " ")
	}
	if len(jsonRequest.ProductCode) < 16 {
		productCode = rightPad(jsonRequest.ProductCode, 16, " ")
	}
	if len(jsonRequest.CustomerNo) < 25 {
		customerNo = rightPad(jsonRequest.CustomerNo, 25, " ")
	}
	if len(jsonRequest.MerchantCode) < 25 {
		merchantCode = rightPad(jsonRequest.MerchantCode, 25, " ")
	}
	if len(jsonRequest.RequestTime) < 19 {
		requestTime = rightPad(jsonRequest.RequestTime, 19, " ")
	}
	if len(jsonRequest.ReffID) < 12 {
		reffId = rightPad(jsonRequest.ReffID, 12, " ")
	}

	// Assign data to map, Adding PAN for PPOB Payment Request, and add MTI
	request := map[int]string{
		3:  "810001",
		4:  strconv.Itoa(amount),
		37: reffId,
		48: transactionID + partnerID + productCode + customerNo + merchantCode + requestTime,
	}
	mti := "0200"

	// Converting request map to isoStruct
	isoStruct := getIso(request, mti)

	isoMessage, _ := isoStruct.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoRequest = isoHeader + isoMessage

	// Create file from request
	filename := "Request_from_PPOBPayment@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/request/"+filename, isoRequest)
	log.Println("Request file: ", file)

	log.Println("Convert success")
	log.Printf("PPOB Payment Request (ISO8583): %s\n", isoRequest)
	return isoRequest
}

// Return ISO message for PPOB Status JSON request
func getIsoPPOBStatus(jsonRequest PPOBStatusRequest) (isoRequest string) {

	log.Println("Converting PPOB Status JSON Request to ISO8583")
	log.Printf("PPOB Status Request (JSON): %v\n", jsonRequest)

	transactionID := jsonRequest.TransactionID
	partnerID := jsonRequest.PartnerID
	productCode := jsonRequest.ProductCode
	customerNo := jsonRequest.CustomerNo
	merchantCode := jsonRequest.MerchantCode
	requestTime := jsonRequest.RequestTime
	reffId := jsonRequest.ReffID
	amount := jsonRequest.Amount

	// Checking data length, match it to Spec Document by adding pad
	if len(jsonRequest.TransactionID) < 25 {
		transactionID = rightPad(jsonRequest.TransactionID, 25, " ")
	}
	if len(jsonRequest.PartnerID) < 16 {
		partnerID = rightPad(jsonRequest.PartnerID, 16, " ")
	}
	if len(jsonRequest.ProductCode) < 16 {
		productCode = rightPad(jsonRequest.ProductCode, 16, " ")
	}
	if len(jsonRequest.CustomerNo) < 25 {
		customerNo = rightPad(jsonRequest.CustomerNo, 25, " ")
	}
	if len(jsonRequest.MerchantCode) < 25 {
		merchantCode = rightPad(jsonRequest.MerchantCode, 25, " ")
	}
	if len(jsonRequest.RequestTime) < 19 {
		requestTime = rightPad(jsonRequest.RequestTime, 19, " ")
	}
	if len(jsonRequest.ReffID) < 12 {
		reffId = rightPad(jsonRequest.ReffID, 12, " ")
	}

	// Assign data to map, Adding PAN for PPOB Status Request, and add MTI
	request := map[int]string{
		3:  "380002",
		4:  strconv.Itoa(amount),
		37: reffId,
		48: transactionID + partnerID + productCode + customerNo + merchantCode + requestTime,
	}
	mti := "0200"

	// Converting request map to isoStruct
	isoStruct := getIso(request, mti)

	isoMessage, _ := isoStruct.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoRequest = isoHeader + isoMessage

	// Create file from request
	filename := "Request_from_PPOBStatus@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/request/"+filename, isoRequest)
	log.Println("Request file: ", file)

	log.Println("Convert success")
	log.Printf("PPOB Status Request (ISO8583): %s\n", isoRequest)
	return isoRequest
}

// Return ISO message for Topup Buy JSON request
func getIsoTopupBuy(jsonRequest TopupBuyRequest) (isoRequest string) {

	log.Println("Converting Topup Buy JSON Request to ISO8583")
	log.Printf("Topup Buy Request (JSON): %v\n", jsonRequest)

	transactionID := jsonRequest.TransactionID
	partnerID := jsonRequest.PartnerID
	productCode := jsonRequest.ProductCode
	customerNo := jsonRequest.CustomerNo
	merchantCode := jsonRequest.MerchantCode
	requestTime := jsonRequest.RequestTime

	// Checking data length, match it to Spec Document by adding pad
	if len(jsonRequest.TransactionID) < 25 {
		transactionID = rightPad(jsonRequest.TransactionID, 25, " ")
	}
	if len(jsonRequest.PartnerID) < 16 {
		partnerID = rightPad(jsonRequest.PartnerID, 16, " ")
	}
	if len(jsonRequest.ProductCode) < 16 {
		productCode = rightPad(jsonRequest.ProductCode, 16, " ")
	}
	if len(jsonRequest.CustomerNo) < 25 {
		customerNo = rightPad(jsonRequest.CustomerNo, 25, " ")
	}
	if len(jsonRequest.MerchantCode) < 25 {
		merchantCode = rightPad(jsonRequest.MerchantCode, 25, " ")
	}
	if len(jsonRequest.RequestTime) < 19 {
		requestTime = rightPad(jsonRequest.RequestTime, 19, " ")
	}

	// Assign data to map, Adding PAN for Topup Buy Request, and add MTI
	request := map[int]string{
		3:  "810002",
		48: transactionID + partnerID + productCode + customerNo + merchantCode + requestTime,
	}
	mti := "0200"

	// Converting request map to isoStruct
	isoStruct := getIso(request, mti)

	isoMessage, _ := isoStruct.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoRequest = isoHeader + isoMessage

	// Create file from request
	filename := "Request_from_TopupBuy@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/request/"+filename, isoRequest)
	log.Println("Request file: ", file)

	log.Println("Convert success")
	log.Printf("Topup Buy Request (ISO8583): %s\n", isoRequest)
	return isoRequest
}

// Return ISO message for Topup Check JSON request
func getIsoTopupCheck(jsonRequest TopupCheckRequest) (isoRequest string) {

	log.Println("Converting Topup Check JSON Request to ISO8583")
	log.Printf("Topup Check Request (JSON): %v\n", jsonRequest)

	transactionID := jsonRequest.TransactionID
	partnerID := jsonRequest.PartnerID
	productCode := jsonRequest.ProductCode
	customerNo := jsonRequest.CustomerNo
	merchantCode := jsonRequest.MerchantCode
	requestTime := jsonRequest.RequestTime

	// Checking data length, match it to Spec Document by adding pad
	if len(jsonRequest.TransactionID) < 25 {
		transactionID = rightPad(jsonRequest.TransactionID, 25, " ")
	}
	if len(jsonRequest.PartnerID) < 16 {
		partnerID = rightPad(jsonRequest.PartnerID, 16, " ")
	}
	if len(jsonRequest.ProductCode) < 16 {
		productCode = rightPad(jsonRequest.ProductCode, 16, " ")
	}
	if len(jsonRequest.CustomerNo) < 25 {
		customerNo = rightPad(jsonRequest.CustomerNo, 25, " ")
	}
	if len(jsonRequest.MerchantCode) < 25 {
		merchantCode = rightPad(jsonRequest.MerchantCode, 25, " ")
	}
	if len(jsonRequest.RequestTime) < 19 {
		requestTime = rightPad(jsonRequest.RequestTime, 19, " ")
	}

	// Assign data to map, Adding PAN for Topup Check Request, and add MTI
	request := map[int]string{
		3:  "380003",
		48: transactionID + partnerID + productCode + customerNo + merchantCode + requestTime,
	}
	mti := "0200"

	// Converting request map to isoStruct
	isoStruct := getIso(request, mti)

	isoMessage, _ := isoStruct.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoRequest = isoHeader + isoMessage

	// Create file from request
	filename := "Request_from_TopupCheck@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/request/"+filename, isoRequest)
	log.Println("Request file: ", file)

	log.Println("Convert success")
	log.Printf("Topup Check Request (ISO8583): %s\n", isoRequest)
	return isoRequest
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
