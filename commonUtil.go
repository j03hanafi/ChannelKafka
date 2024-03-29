package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	limit := 0
	for search {
		index := 0
		fmt.Println(len(tempStorage))
		for _, ele := range tempStorage {
			if ele.Head == head {
				search = false
				res = ele
				tempStorage = append(tempStorage[:index], tempStorage[index+1:]...)
			}
			index++
		}
		time.Sleep(100 * time.Millisecond)
		limit++
		if limit == 50 {
			break
		}

	}
	return res
}

func putValueToArray(data resConsume) {
	tempStorage = append(tempStorage, data)
}
