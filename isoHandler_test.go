package main

import "testing"

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

// TestStructRintisToIso
func TestStructRintisToIso(t *testing.T) {
	jsonRequest := rintisRequest{
		Pan:                  "",
		ProcessingCode:       "",
		TotalAmount:          0,
		TransmissionDateTime: "",
		Stan:                 "",
		LocalTransactionTime: "",
		LocalTransactionDate: "",
		CaptureDate:          "",
		AcquirerID:           "",
		Track2Data:           "",
		Refnum:               "",
		TerminalID:           "",
		CardAcceptorData:     "",
		AdditionalData:       "",
		Currency:             "",
		PIN:                  "",
		TerminalData:         "",
		AccountTo:            "",
		TokenData:            "",
	}

	reqISO, head := structRintisToIso(jsonRequest)

	reqIsoExpected := ""
	headExpected := ""

	if reqISO != reqIsoExpected {
		t.Errorf("return different reqISO from structRintisToIso(), \nexpected\t: %v, \ngot\t\t\t: %v", reqIsoExpected, reqISO)
	} else {
		t.Log("return reqISO structRintisToIso() success")
	}

	if head != headExpected {
		t.Errorf("return different head from structRintisToIso(), \nexpected\t: %v, \ngot\t\t\t: %v", headExpected, head)
	} else {
		t.Log("return head structRintisToIso() success")
	}

}

func TestParseResponse(t *testing.T) {
	expected := ""

	isoStructResponse := parseResponse(expected)
	result, _ := isoStructResponse.ToString()

	if result != expected {
		t.Errorf("parseResponse() failed, \nexpected\t: %v, \ngot\t\t\t: %v", expected, result)
	} else {
		t.Log("parseResponse() success")
	}
}
