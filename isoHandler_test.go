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
