package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Struct for kafkaConfig.json
type Config struct {
	Broker         string   `json:"broker"`
	ProducerTopic  string   `json:"producer_topic"`
	ConsumerTopics []string `json:"consumer_topics"`
	Group          string   `json:"group"`
}

// Return config for setting up Kafka Producer and Consumer
func configKafka() (broker string, producerTopic string, consumerTopics []string, group string) {
	log.Printf("Get config for current request")

	file, _ := os.Open("./kafkaConfig.json")
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(b, &config)

	log.Printf("Kafka Config -> Broker: `%v`, Producer Topic: `%v`, Consumer Topics: `%v`, Group: `%v`",
		config.Broker, config.ProducerTopic, config.ConsumerTopics, config.Group)
	return config.Broker, config.ProducerTopic, config.ConsumerTopics, config.Group
}
