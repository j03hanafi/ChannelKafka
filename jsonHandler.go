package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handle all JSON Client request
func sendJSON(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var reqBody Transaction

	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	resISO, _ := convJSON(reqBody)

	fmt.Print(resISO)
	responseFormatter(w, resISO, 200)
}
