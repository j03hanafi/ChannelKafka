package main

import "testing"

func TestCreateFile(t *testing.T) {
	createdFile := CreateFile("testFile", "testContent")
	expectedFileName := "testFile.txt"

	if createdFile != expectedFileName {
		t.Errorf("CreateFile() failed. Expected: %v. Got: %v", expectedFileName, createdFile)
	} else {
		t.Log("CreateFile() success")
	}
}

func TestGetValueFromArray(t *testing.T) {
	dummy := []resConsume{
		{
			Head:    "A",
			Content: "Aa",
		},
		{
			Head:    "B",
			Content: "Bb",
		},
		{
			Head:    "C",
			Content: "Cc",
		},
	}

	for _, data := range dummy {
		putValueToArray(data)
	}

	for index, data := range dummy {
		result := getValueFromArray(data.Head)

		if data.Head != result.Head {
			t.Errorf("return different head from getValueFromArray(), \nexpected\t: %v, \ngot\t\t\t: %v", data.Head, result.Head)
		} else {
			t.Logf("return head getValueFromArray() success (dummy %v)", index)
		}
	}
}

func TestPutValueToArray(t *testing.T) {

	dummy := []resConsume{
		{
			Head:    "A",
			Content: "Aa",
		},
		{
			Head:    "B",
			Content: "Bb",
		},
		{
			Head:    "C",
			Content: "Cc",
		},
	}

	for _, data := range dummy {
		putValueToArray(data)
	}

	for index, data := range tempStorage {
		if data.Head != dummy[index].Head {
			t.Errorf("return different head from putValueToArray(), \nexpected\t: %v, \ngot\t\t\t: %v", dummy[index].Head, data.Head)
		} else {
			t.Logf("return head putValueToArray() success (dummy %v)", index)
		}

		if data.Content != dummy[index].Content {
			t.Errorf("return different content from putValueToArray(), \nexpected\t: %v, \ngot\t\t\t: %v", dummy[index].Head, data.Head)
		} else {
			t.Logf("return content putValueToArray() success (dummy %v)", index)
		}
	}
}
