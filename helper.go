package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// Response formatter
func responseFormatter(w http.ResponseWriter, data interface{}, statusCode int) {
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

func getValueFromArray(head string) resConsume {
	search := true
	var res resConsume
	for search {
		index := 0
		for _, ele := range tempStorage {
			if ele.Head != head {
				tempStorage[index] = ele
				index++
			} else {
				search = false
				res = ele
			}
		}
	}
	return res
}

func putValueToArray(data resConsume) {
	tempStorage = append(tempStorage, data)
}
