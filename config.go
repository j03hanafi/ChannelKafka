package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

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
