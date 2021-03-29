package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/mofax/iso8583"
)

// Return unsuccessful Chipsakti Response (rc != "00")
func getJsonUnsuccessfulChipsakti(parsedIso iso8583.IsoStruct) UnsuccessfulChipsakti {
	var response UnsuccessfulChipsakti

	log.Println("Converting Unsuccessful Chipsakti ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Rc = strings.Trim(emap[39], " ")
	response.Restime = emap[48]
	response.Msg = emap[120]

	log.Println("Convert success")
	log.Printf("Unsuccessful Chipsakti Response (JSON): %+v\n", response)
	return response
}

// Return JSON for PPOB Inquiry ISO8583 response
func getJsonPPOBInquiry(parsedIso iso8583.IsoStruct) PPOBInquiryResponse {
	var response PPOBInquiryResponse

	log.Println("Converting PPOB Inquiry ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Convert non-string field
	tagihan, _ := strconv.Atoi(emap[4])
	admin, _ := strconv.Atoi(emap[5])
	totalTagihan, _ := strconv.Atoi(emap[6])

	// Map ISO8583 format to JSON data
	response.Tagihan = tagihan
	response.Admin = admin
	response.TotalTagihan = totalTagihan
	response.Reffid = strings.Trim(emap[37], " ")
	response.Rc = strings.Trim(emap[39], " ")
	response.Nama = strings.Trim(emap[43], " ")
	response.Restime = emap[48]
	response.Data = emap[62]
	response.Msg = emap[120]
	response.Produk = strings.Trim(emap[121], " ")
	response.Nopel = strings.Trim(emap[122], " ")

	log.Println("Convert success")
	log.Printf("PPOB Inquiry Response (JSON): %+v\n", response)
	return response
}

// Return JSON for PPOB Payment ISO8583 response
func getJsonPPOBPayment(parsedIso iso8583.IsoStruct) PPOBPaymentResponse {
	var response PPOBPaymentResponse

	log.Println("Converting PPOB Payment ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Convert non-string field
	tagihan, _ := strconv.Atoi(emap[4])
	admin, _ := strconv.Atoi(emap[5])
	totalTagihan, _ := strconv.Atoi(emap[6])
	struk := strings.Split(emap[62], ",")

	// Map ISO8583 format to JSON data
	response.Tagihan = tagihan
	response.Admin = admin
	response.TotalTagihan = totalTagihan
	response.Reffid = strings.Trim(emap[37], " ")
	response.Rc = strings.Trim(emap[39], " ")
	response.Nama = strings.Trim(emap[43], " ")
	response.TglLunas = strings.Trim(emap[48], " ")
	response.Struk = struk
	response.Msg = strings.Trim(emap[120], " ")
	response.Produk = strings.Trim(emap[121], " ")
	response.Nopel = strings.Trim(emap[122], " ")
	response.ReffNo = strings.Trim(emap[123], " ")

	log.Println("Convert success")
	log.Printf("PPOB Payment Response (JSON): %+v\n", response)
	return response
}

// Return JSON for PPOB Status ISO8583 response
func getJsonPPOBStatus(parsedIso iso8583.IsoStruct) PPOBStatusResponse {
	var response PPOBStatusResponse

	log.Println("Converting PPOB Status ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Convert non-string field
	tagihan, _ := strconv.Atoi(emap[4])
	admin, _ := strconv.Atoi(emap[5])
	totalTagihan, _ := strconv.Atoi(emap[6])
	struk := strings.Split(emap[62], ",")

	// Map ISO8583 format to JSON data
	response.Tagihan = tagihan
	response.Admin = admin
	response.TotalTagihan = totalTagihan
	response.Reffid = strings.Trim(emap[37], " ")
	response.Rc = strings.Trim(emap[39], " ")
	response.Nama = strings.Trim(emap[43], " ")
	response.TglLunas = strings.Trim(emap[48], " ")
	response.Struk = struk
	response.Msg = strings.Trim(emap[120], " ")
	response.Produk = strings.Trim(emap[121], " ")
	response.Nopel = strings.Trim(emap[122], " ")
	response.ReffNo = strings.Trim(emap[123], " ")
	response.Status = strings.Trim(emap[124], " ")

	log.Println("Convert success")
	log.Printf("PPOB Status Response (JSON): %+v\n", response)
	return response
}

// Return JSON for Topup Buy ISO8583 response
func getJsonTopupBuy(parsedIso iso8583.IsoStruct) TopupBuyResponse {
	var response TopupBuyResponse

	log.Println("Converting Topup Buy ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Rc = strings.Trim(emap[39], " ")
	response.Restime = strings.Trim(emap[48], " ")
	response.Msg = strings.Trim(emap[120], " ")
	response.SN = strings.Trim(emap[121], " ")
	response.Price = strings.Trim(emap[122], " ")

	log.Println("Convert success")
	log.Printf("Topup Buy Response (JSON): %+v\n", response)
	return response
}

// Return JSON for Topup Check ISO8583 response
func getJsonTopupCheck(parsedIso iso8583.IsoStruct) TopupCheckResponse {
	var response TopupCheckResponse

	log.Println("Converting Topup Check ISO8583 to JSON Response")
	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Rc = strings.Trim(emap[39], " ")
	response.Restime = strings.Trim(emap[48], " ")
	response.Msg = strings.Trim(emap[120], " ")
	response.SN = strings.Trim(emap[121], " ")
	response.Price = strings.Trim(emap[122], " ")

	log.Println("Convert success")
	log.Printf("Topup Check Response (JSON): %+v\n", response)
	return response
}

func makeRinRespFromIso(parsedIso iso8583.IsoStruct) (response rintisResponse) {
	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Pan = strings.Trim(emap[2], " ")
	response.ProcessingCode = strings.Trim(emap[3], " ")
	response.TotalAmount, _ = strconv.Atoi(emap[4])
	response.TransmissionDateTime = strings.Trim(emap[7], " ")
	response.Stan = strings.Trim(emap[11], " ")
	response.LocalTransactionTime = strings.Trim(emap[12], " ")
	response.LocalTransactionDate = strings.Trim(emap[13], " ")
	response.SettlementDate = strings.Trim(emap[15], " ")
	response.CaptureDate = strings.Trim(emap[17], " ")
	response.AcquirerID = strings.Trim(emap[32], " ")
	response.Track2Data = strings.Trim(emap[35], " ")
	response.Refnum = strings.Trim(emap[37], " ")
	response.AuthIdResponse = strings.Trim(emap[38], " ")
	response.ResponseCode = strings.Trim(emap[39], " ")
	response.TerminalID = strings.Trim(emap[41], " ")
	response.AdditionalResponseData = strings.Trim(emap[44], " ")
	response.Currency = strings.Trim(emap[49], " ")
	response.TerminalData = strings.Trim(emap[60], " ")
	response.ReceivingInstitutionIDCode = strings.Trim(emap[100], " ")
	response.AccountTo = strings.Trim(emap[103], " ")
	response.TokenData = strings.Trim(emap[126], " ")
	log.Println("Convert success")
	log.Printf("Topup Check Response (JSON): %+v\n", response)
	return response

}
