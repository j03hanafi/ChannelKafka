package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Response formatter
func responseFormatter(w http.ResponseWriter, data interface{}, statusCode int) {
	j, err := json.Marshal(&data)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	fmt.Println(string(j))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Create file for request/response
func CreateFile(fileName string, content string) string {

	log.Println("Creating new file")

	if !strings.Contains(fileName, ".txt") {
		fileName += ".txt"
	}

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}

	log.Println("File created!")
	return fileName
}
