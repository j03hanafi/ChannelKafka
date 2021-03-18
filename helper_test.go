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
