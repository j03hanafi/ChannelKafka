package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Handle all JSON Client request
func sendJSON(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var response Response
	var reqBody Transaction

	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Printf("json error with msg : %s", err.Error())
		return
	}

	resISO, _ := convJSON(reqBody)

	err = doProducer(broker, topic1, resISO)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		log.Println(err)
		responseFormatter(w, response, 500)
	} else {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "channel",
			"auto.offset.reset": "earliest",
		})

		if err != nil {
			fmt.Println(err.Error())
		}

		c.SubscribeTopics([]string{topic2}, nil)

		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

		c.Close()
		strRes := fromISO(string(msg.Value))

		responseFormatter(w, strRes, 200)
	}
}
