package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var arr []resConsume

type resConsume struct {
	stan  string
	msgin string
}

// Struct for kafkaConfig.json
type Config struct {
	Broker         string   `json:"broker"`
	ProducerTopics []string `json:"producer_topics"`
	ConsumerTopics []string `json:"consumer_topics"`
	Group          string   `json:"group"`
}

// Return config for setting up Kafka Producer and Consumer
func configKafka() (broker string, producerTopics []string, consumerTopics []string, group string) {
	log.Printf("Get config for current request")

	file, _ := os.Open("./kafkaConfig.json")
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(b, &config)

	return config.Broker, config.ProducerTopics, config.ConsumerTopics, config.Group
}
