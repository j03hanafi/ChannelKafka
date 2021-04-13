package main

import (
	"reflect"
	"testing"
)

// TestMakeRinRespFromIso
func TestMakeRinRespFromIso(t *testing.T) {

	isoStructResponse := ""
	parsedIso := parseResponse(isoStructResponse)

	response := rintisResponse{
		Pan:                        "",
		ProcessingCode:             "",
		TotalAmount:                0,
		TransmissionDateTime:       "",
		Stan:                       "",
		LocalTransactionTime:       "",
		LocalTransactionDate:       "",
		SettlementDate:             "",
		CaptureDate:                "",
		AcquirerID:                 "",
		Track2Data:                 "",
		Refnum:                     "",
		AuthIdResponse:             "",
		ResponseCode:               "",
		TerminalID:                 "",
		AdditionalResponseData:     "",
		Currency:                   "",
		TerminalData:               "",
		ReceivingInstitutionIDCode: "",
		AccountTo:                  "",
		TokenData:                  "",
	}

	structResponse := makeRinRespFromIso(parsedIso)
	result := reflect.ValueOf(structResponse)
	expected := reflect.ValueOf(response)

	for i := 0; i < result.NumField(); i++ {
		if expected.Field(i).Interface() != result.Field(i).Interface() {
			t.Errorf("makeRinRespFromIso() failed, \nexpected (%v)\t: %v, \ngot (%v)\t\t\t: %v", expected.Field(i), expected.Field(i).Interface(), result.Field(i), result.Field(i).Interface())
		} else {
			t.Logf("makeRinRespFromIso() %v success", result.Field(i))
		}
	}

}
