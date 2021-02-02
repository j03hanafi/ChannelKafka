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
)

// Any helper to process ISO data
// converter, formatter, etc
func convIsoPPOBInquiry(data PPOBInquiryRequest) string {

	log.Println("New Request from JSON")
	log.Printf("original: %v\n", data)

	transaction_id := data.TransactionID
	if len(data.TransactionID) < 25 {
		transaction_id = rightPad(data.TransactionID, 25, " ")
	}
	partner_id := data.PartnerID
	if len(data.PartnerID) < 16 {
		partner_id = rightPad(data.PartnerID, 16, " ")
	}
	product_code := data.ProductCode
	if len(data.ProductCode) < 16 {
		product_code = rightPad(data.ProductCode, 16, " ")
	}
	customer_no := data.CustomerNo
	if len(data.CustomerNo) < 25 {
		customer_no = rightPad(data.CustomerNo, 25, " ")
	}
	merchant_code := data.MerchantCode
	if len(data.MerchantCode) < 25 {
		merchant_code = rightPad(data.MerchantCode, 25, " ")
	}
	request_time := data.RequestTime
	if len(data.RequestTime) < 19 {
		request_time = rightPad(data.RequestTime, 19, " ")
	}

	signature := data.Signature
	periode := data.Periode

	val := map[int]string{
		48: transaction_id + partner_id + product_code + customer_no + merchant_code + request_time + periode,
		62: signature,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range val {

		fieldSpec := spec.fields[field]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(int64(field), data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	one.AddMTI("0200")
	one.AddField(3, "380001")
	isoMessage, _ := one.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoReq := isoHeader + isoMessage

	log.Printf("result: %s\n", isoReq)

	return isoReq
}

func convIsoPPOBPayment(data PPOBPaymentRequest) string {

	log.Println("New Request from JSON")
	log.Printf("original: %v\n", data)

	transaction_id := data.TransactionID
	if len(data.TransactionID) < 25 {
		transaction_id = rightPad(data.TransactionID, 25, " ")
	}
	partner_id := data.PartnerID
	if len(data.PartnerID) < 16 {
		partner_id = rightPad(data.PartnerID, 16, " ")
	}
	product_code := data.ProductCode
	if len(data.ProductCode) < 16 {
		product_code = rightPad(data.ProductCode, 16, " ")
	}
	customer_no := data.CustomerNo
	if len(data.CustomerNo) < 25 {
		customer_no = rightPad(data.CustomerNo, 25, " ")
	}
	merchant_code := data.MerchantCode
	if len(data.MerchantCode) < 25 {
		merchant_code = rightPad(data.MerchantCode, 25, " ")
	}
	request_time := data.RequestTime
	if len(data.RequestTime) < 19 {
		request_time = rightPad(data.RequestTime, 19, " ")
	}
	reffId := data.ReffID
	if len(data.ReffID) < 12 {
		reffId = rightPad(data.ReffID, 12, " ")
	}

	signature := data.Signature
	amount := data.Amount

	val := map[int]string{
		4:  strconv.Itoa(amount),
		37: reffId,
		48: transaction_id + partner_id + product_code + customer_no + merchant_code + request_time,
		62: signature,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range val {

		fieldSpec := spec.fields[field]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(int64(field), data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	one.AddMTI("0200")
	one.AddField(3, "810001")
	isoMessage, _ := one.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoReq := isoHeader + isoMessage

	log.Printf("result: %s\n", isoReq)

	return isoReq
}

func convIsoPPOBStatus(data PPOBStatusRequest) string {

	log.Println("New Request from JSON")
	log.Printf("original: %v\n", data)

	transaction_id := data.TransactionID
	if len(data.TransactionID) < 25 {
		transaction_id = rightPad(data.TransactionID, 25, " ")
	}
	partner_id := data.PartnerID
	if len(data.PartnerID) < 16 {
		partner_id = rightPad(data.PartnerID, 16, " ")
	}
	product_code := data.ProductCode
	if len(data.ProductCode) < 16 {
		product_code = rightPad(data.ProductCode, 16, " ")
	}
	customer_no := data.CustomerNo
	if len(data.CustomerNo) < 25 {
		customer_no = rightPad(data.CustomerNo, 25, " ")
	}
	merchant_code := data.MerchantCode
	if len(data.MerchantCode) < 25 {
		merchant_code = rightPad(data.MerchantCode, 25, " ")
	}
	request_time := data.RequestTime
	if len(data.RequestTime) < 19 {
		request_time = rightPad(data.RequestTime, 19, " ")
	}
	reffId := data.ReffID
	if len(data.ReffID) < 12 {
		reffId = rightPad(data.ReffID, 12, " ")
	}

	signature := data.Signature
	amount := data.Amount

	val := map[int]string{
		4:  strconv.Itoa(amount),
		37: reffId,
		48: transaction_id + partner_id + product_code + customer_no + merchant_code + request_time,
		62: signature,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range val {

		fieldSpec := spec.fields[field]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(int64(field), data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	one.AddMTI("0200")
	one.AddField(3, "380002")
	isoMessage, _ := one.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoReq := isoHeader + isoMessage

	log.Printf("result: %s\n", isoReq)

	return isoReq
}

func convIsoTopupBuy(data TopupBuyRequest) string {

	log.Println("New Request from JSON")
	log.Printf("original: %v\n", data)

	transaction_id := data.TransactionID
	if len(data.TransactionID) < 25 {
		transaction_id = rightPad(data.TransactionID, 25, " ")
	}
	partner_id := data.PartnerID
	if len(data.PartnerID) < 16 {
		partner_id = rightPad(data.PartnerID, 16, " ")
	}
	product_code := data.ProductCode
	if len(data.ProductCode) < 16 {
		product_code = rightPad(data.ProductCode, 16, " ")
	}
	customer_no := data.CustomerNo
	if len(data.CustomerNo) < 25 {
		customer_no = rightPad(data.CustomerNo, 25, " ")
	}
	merchant_code := data.MerchantCode
	if len(data.MerchantCode) < 25 {
		merchant_code = rightPad(data.MerchantCode, 25, " ")
	}
	request_time := data.RequestTime
	if len(data.RequestTime) < 19 {
		request_time = rightPad(data.RequestTime, 19, " ")
	}

	signature := data.Signature

	val := map[int]string{
		48: transaction_id + partner_id + product_code + customer_no + merchant_code + request_time,
		62: signature,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range val {

		fieldSpec := spec.fields[field]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(int64(field), data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	one.AddMTI("0200")
	one.AddField(3, "810002")
	isoMessage, _ := one.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoReq := isoHeader + isoMessage

	log.Printf("result: %s\n", isoReq)

	return isoReq
}

func convIsoTopupCheck(data TopupCheckRequest) string {

	log.Println("New Request from JSON")
	log.Printf("original: %v\n", data)

	transaction_id := data.TransactionID
	if len(data.TransactionID) < 25 {
		transaction_id = rightPad(data.TransactionID, 25, " ")
	}
	partner_id := data.PartnerID
	if len(data.PartnerID) < 16 {
		partner_id = rightPad(data.PartnerID, 16, " ")
	}
	product_code := data.ProductCode
	if len(data.ProductCode) < 16 {
		product_code = rightPad(data.ProductCode, 16, " ")
	}
	customer_no := data.CustomerNo
	if len(data.CustomerNo) < 25 {
		customer_no = rightPad(data.CustomerNo, 25, " ")
	}
	merchant_code := data.MerchantCode
	if len(data.MerchantCode) < 25 {
		merchant_code = rightPad(data.MerchantCode, 25, " ")
	}
	request_time := data.RequestTime
	if len(data.RequestTime) < 19 {
		request_time = rightPad(data.RequestTime, 19, " ")
	}

	signature := data.Signature

	val := map[int]string{
		48: transaction_id + partner_id + product_code + customer_no + merchant_code + request_time,
		62: signature,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range val {

		fieldSpec := spec.fields[field]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(int64(field), data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	one.AddMTI("0200")
	one.AddField(3, "380003")
	isoMessage, _ := one.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))
	isoReq := isoHeader + isoMessage

	log.Printf("result: %s\n", isoReq)

	return isoReq
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
