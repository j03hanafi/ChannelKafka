package main

import (
	"github.com/mofax/iso8583"
	"reflect"
	"testing"
)

func TestGetJsonUnsuccessfulChipsakti(t *testing.T) {
	var result UnsuccessfulChipsakti
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	parsedIso := "0210a0000000020100000000000000000100380001002200192021-03-18 08:03:30023Tagihan tidak ditemukan"
	iso, err := isoStruct.Parse(parsedIso)
	if err != nil {
		t.Errorf("Error parsing iso message. Error: %v", err)
	}
	result = getJsonUnsuccessfulChipsakti(iso)
	expected := UnsuccessfulChipsakti{
		Rc:      "20",
		Msg:     "Tagihan tidak ditemukan",
		Restime: "2021-03-18 08:03:30",
	}

	if result != expected {
		t.Errorf("getJsonUnsuccessfulChipsakti() failed. \nExpected\t: %v. Got\t: %v", expected, result)
	} else {
		t.Log("getJsonUnsuccessfulChipsakti() success")
	}
}

func TestGetJsonPPOBInquiry(t *testing.T) {
	var result PPOBInquiryResponse
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	parsedIso := "0210bc0000000a21000400000000000001c038000100000000150000000000330000000000480012345       00200HANAFI                                  0192021-03-18 08:03:230042020007approve003WOM0012"
	iso, err := isoStruct.Parse(parsedIso)
	if err != nil {
		t.Errorf("Error parsing iso message. Error: %v", err)
	}
	result = getJsonPPOBInquiry(iso)
	expected := PPOBInquiryResponse{
		Rc:           "00",
		Msg:          "approve",
		Produk:       "WOM",
		Nopel:        "2",
		Nama:         "HANAFI",
		Tagihan:      1500,
		Admin:        3300,
		TotalTagihan: 4800,
		Reffid:       "12345",
		Data:         "2020",
		Restime:      "2021-03-18 08:03:23",
	}

	if result != expected {
		t.Errorf("getJsonPPOBInquiry() failed. \nExpected\t: %v. Got\t: %v", expected, result)
	} else {
		t.Log("getJsonPPOBInquiry() success")
	}
}

func TestGetsonPPOBPayment(t *testing.T) {
	var result PPOBPaymentResponse
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	parsedIso := "0210bc0000000a21000400000000000001e081000100000087000000000000330000000087330012345       00200HANAFI                                  0192021-03-18 08:03:35191pembayaranWOM,,ID PEL :2,NAMA :HANAFI,REF : 5/4-3-2-1,ANGSURAN KE: 5,TAGIHAN : Rp 870000,BIAYA ADMIN : Rp 3300,TTL TAGIHAN : Rp 873300,,STRUK INI ADALAH BUKTI PEMBAYARAN YANG SAH,TERIMA KASIH007approve003WOM001200554321"
	iso, err := isoStruct.Parse(parsedIso)
	if err != nil {
		t.Errorf("Error parsing iso message. Error: %v", err)
	}
	result = getJsonPPOBPayment(iso)
	expected := PPOBPaymentResponse{
		Rc:           "00",
		Msg:          "approve",
		Produk:       "WOM",
		Nopel:        "2",
		Nama:         "HANAFI",
		Tagihan:      870000,
		Admin:        3300,
		TotalTagihan: 873300,
		Reffid:       "12345",
		TglLunas:     "2021-03-18 08:03:35",
		Struk: []string{
			"pembayaranWOM",
			"",
			"ID PEL :2",
			"NAMA :HANAFI",
			"REF : 5/4-3-2-1",
			"ANGSURAN KE: 5",
			"TAGIHAN : Rp 870000",
			"BIAYA ADMIN : Rp 3300",
			"TTL TAGIHAN : Rp 873300",
			"",
			"STRUK INI ADALAH BUKTI PEMBAYARAN YANG SAH",
			"TERIMA KASIH",
		},
		ReffNo:  "54321",
		Restime: "",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getJsonPPOBPayment() failed. \nExpected\t: %v. Got\t: %v", expected, result)
	} else {
		t.Log("getJsonPPOBPayment() success")
	}
}

func TestGetJsonPPOBStatus(t *testing.T) {
	var result PPOBStatusResponse
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	parsedIso := "0210bc0000000a21000400000000000001f038000200000001000000000000000000000001000012345       00200HANAFI                                  0192021-03-18 08:03:38230<b>PT. MULTI ACCESS INDONESIA - CHIPSAKTI</b>,,LOKET : ZONATIK,TGL BAYAR : 02/07/2018 / 14:16:44,,STRUK PEMBAYARAN LANGGANANWOM,,IDPEL 2,NAMA : HANAFI,TTL TAGIHAN : Rp 10000,,STRUK INI ADALAH BUKTI PEMBAYARAN YANG SAH,TERIMA KASIH007approve003WOM00120040123019payment Successfull"
	iso, err := isoStruct.Parse(parsedIso)
	if err != nil {
		t.Errorf("Error parsing iso message. Error: %v", err)
	}
	result = getJsonPPOBStatus(iso)
	expected := PPOBStatusResponse{
		Rc:           "00",
		Msg:          "approve",
		Produk:       "WOM",
		Nopel:        "2",
		Nama:         "HANAFI",
		Tagihan:      10000,
		Admin:        0,
		TotalTagihan: 10000,
		Reffid:       "12345",
		TglLunas:     "2021-03-18 08:03:38",
		Struk: []string{
			"<b>PT. MULTI ACCESS INDONESIA - CHIPSAKTI</b>",
			"",
			"LOKET : ZONATIK",
			"TGL BAYAR : 02/07/2018 / 14:16:44",
			"",
			"STRUK PEMBAYARAN LANGGANANWOM",
			"",
			"IDPEL 2",
			"NAMA : HANAFI",
			"TTL TAGIHAN : Rp 10000",
			"",
			"STRUK INI ADALAH BUKTI PEMBAYARAN YANG SAH",
			"TERIMA KASIH",
		},
		ReffNo: "0123",
		Status: "payment Successfull",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getJsonPPOBStatus() failed. \nExpected\t: %v. Got\t: %v", expected, result)
	} else {
		t.Log("getJsonPPOBStatus() success")
	}
}

func TestGetJsonTopupBuy(t *testing.T) {
	var result TopupBuyResponse
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	parsedIso := "0210a00000000201000000000000000001c0810002002000192021-03-18 08:03:42036PembelianWOMberhasil. Harga Rp. 1000008123456780041000"
	iso, err := isoStruct.Parse(parsedIso)
	if err != nil {
		t.Errorf("Error parsing iso message. Error: %v", err)
	}
	result = getJsonTopupBuy(iso)
	expected := TopupBuyResponse{
		Rc:      "00",
		Msg:     "PembelianWOMberhasil. Harga Rp. 1000",
		Restime: "2021-03-18 08:03:42",
		SN:      "12345678",
		Price:   "1000",
	}

	if result != expected {
		t.Errorf("getJsonTopupBuy() failed. \nExpected\t: %v. Got\t: %v", expected, result)
	} else {
		t.Log("getJsonTopupBuy() success")
	}
}

func TestGetJsonTopupCheck(t *testing.T) {
	var result TopupCheckResponse
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	parsedIso := "0210a00000000201000000000000000001c0380003002000192021-03-18 08:03:45036PembelianWOMberhasil. Harga Rp. 1000008123456780041000"
	iso, err := isoStruct.Parse(parsedIso)
	if err != nil {
		t.Errorf("Error parsing iso message. Error: %v", err)
	}
	result = getJsonTopupCheck(iso)
	expected := TopupCheckResponse{
		Rc:      "00",
		Msg:     "PembelianWOMberhasil. Harga Rp. 1000",
		Restime: "2021-03-18 08:03:45",
		SN:      "12345678",
		Price:   "1000",
	}

	if result != expected {
		t.Errorf("getJsonTopupCheck() failed. \nExpected\t: %v. Got\t: %v", expected, result)
	} else {
		t.Log("getJsonTopupCheck() success")
	}
}
