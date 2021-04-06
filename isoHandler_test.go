package main

import (
	"github.com/mofax/iso8583"
	"testing"
)

// Test field lebih banyak
// Test koneksi

func TestGetIso(t *testing.T) {

	request := map[int]string{
		3:  "380001",
		48: "2021                     USER01          WOM             2                        KIOS01                   2018-05-15 15:10:052020",
	}
	mti := "0200"
	isoStruct := getIso(request, mti)

	expected := "0200a00000000001000000000000000000003800011302021                     USER01          WOM             2                        KIOS01                   2018-05-15 15:10:052020"
	result, _ := isoStruct.ToString()

	if result != expected {
		t.Errorf("getIso() failed, \nexpected\t: %v, \ngot\t\t\t: %v", expected, result)
	} else {
		t.Log("getIso() success")
	}
}

func TestGetIsoPPOBInquiry(t *testing.T) {

	var jsonRequest PPOBInquiryRequest

	jsonRequest.TransactionID = "2021"
	jsonRequest.PartnerID = "USER01"
	jsonRequest.ProductCode = "WOM"
	jsonRequest.CustomerNo = "2"
	jsonRequest.MerchantCode = "KIOS01"
	jsonRequest.RequestTime = "2018-05-15 15:10:05"
	jsonRequest.Periode = "2020"

	isoRequest := getIsoPPOBInquiry(jsonRequest)

	expected := "01750200a00000000001000000000000000000003800011302021                     USER01          WOM             2                        KIOS01                   2018-05-15 15:10:052020"
	result := isoRequest

	if result != expected {
		t.Errorf("getIsoPPOBInquiry() failed, \nexpected\t: %v, \ngot\t\t\t: %v", expected, result)
	} else {
		t.Log("getIsoPPOBInquiry() success")
	}
}

func TestGetIsoPPOBPayment(t *testing.T) {

	var jsonRequest PPOBPaymentRequest

	jsonRequest.TransactionID = "2015"
	jsonRequest.PartnerID = "USER01"
	jsonRequest.ProductCode = "WOM"
	jsonRequest.CustomerNo = "2"
	jsonRequest.MerchantCode = "KIOS01"
	jsonRequest.RequestTime = "2018-05-15 15:10:05"
	jsonRequest.ReffID = "12345"
	jsonRequest.Amount = 873300

	isoRequest := getIsoPPOBPayment(jsonRequest)

	expected := "01950200b000000008010000000000000000000081000100000087330012345       1262015                     USER01          WOM             2                        KIOS01                   2018-05-15 15:10:05"
	result := isoRequest

	if result != expected {
		t.Errorf("getIsoPPOBPayment() failed, \nexpected\t: %v, \ngot\t\t\t: %v", expected, result)
	} else {
		t.Log("getIsoPPOBPayment() success")
	}
}

func TestGetIsoPPOBStatus(t *testing.T) {

	var jsonRequest PPOBStatusRequest

	jsonRequest.TransactionID = "2021"
	jsonRequest.PartnerID = "USER01"
	jsonRequest.ProductCode = "WOM"
	jsonRequest.CustomerNo = "2"
	jsonRequest.MerchantCode = "KIOS01"
	jsonRequest.RequestTime = "2018-05-15 15:10:05"
	jsonRequest.ReffID = "12345"
	jsonRequest.Amount = 10000

	isoRequest := getIsoPPOBStatus(jsonRequest)

	expected := "01950200b000000008010000000000000000000038000200000001000012345       1262021                     USER01          WOM             2                        KIOS01                   2018-05-15 15:10:05"
	result := isoRequest

	if result != expected {
		t.Errorf("getIsoPPOBStatus() failed, \nexpected\t: %v, \ngot\t\t\t: %v", expected, result)
	} else {
		t.Log("getIsoPPOBStatus() success")
	}
}

func TestGetIsoTopupBuy(t *testing.T) {

	var jsonRequest TopupBuyRequest

	jsonRequest.TransactionID = "2021"
	jsonRequest.PartnerID = "USER01"
	jsonRequest.ProductCode = "WOM"
	jsonRequest.CustomerNo = "1"
	jsonRequest.MerchantCode = "KIOS01"
	jsonRequest.RequestTime = "2018-05-15 15:10:05"

	isoRequest := getIsoTopupBuy(jsonRequest)

	expected := "01710200a00000000001000000000000000000008100021262021                     USER01          WOM             1                        KIOS01                   2018-05-15 15:10:05"
	result := isoRequest

	if result != expected {
		t.Errorf("getIsoTopupBuy() failed, \nexpected\t: %v, \ngot\t\t\t: %v", expected, result)
	} else {
		t.Log("getIsoTopupBuy() success")
	}
}

func TestGetIsoTopupCheck(t *testing.T) {

	var jsonRequest TopupCheckRequest

	jsonRequest.TransactionID = "2021"
	jsonRequest.PartnerID = "USER01"
	jsonRequest.ProductCode = "WOM"
	jsonRequest.CustomerNo = "1"
	jsonRequest.MerchantCode = "KIOS01"
	jsonRequest.RequestTime = "2018-05-15 15:10:05"

	isoRequest := getIsoTopupCheck(jsonRequest)

	expected := "01710200a00000000001000000000000000000003800031262021                     USER01          WOM             1                        KIOS01                   2018-05-15 15:10:05"
	result := isoRequest

	if result != expected {
		t.Errorf("getIsoTopupCheck() failed, \nexpected\t: %v, \ngot\t\t\t: %v", expected, result)
	} else {
		t.Log("getIsoTopupCheck() success")
	}
}

func TestParseResponse(t *testing.T) {
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	parsed := "0210a0000000020100000000000000000100380001002200192021-03-18 08:03:30023Tagihan tidak ditemukan"
	isoExpected, err := isoStruct.Parse(parsed)
	if err != nil {
		t.Errorf("Error parsing iso message. Error: %v", err)
	}
	expected, _ := isoExpected.ToString()
	header := "0095"
	isoResult := parseResponse(header + parsed)
	result, _ := isoResult.ToString()

	if result != expected {
		t.Errorf("parseResponse() failed, \nexpected\t: %v, \ngot\t\t\t: %v", expected, result)
	} else {
		t.Log("parseResponse() success")
	}
}
