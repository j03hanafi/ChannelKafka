package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consumeMsgFromKafka(broker string, topics []string, group string) {

	// Setting up Consumer (Kafka) config
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": "latest",
	})

	if err != nil {
		panic(err)
	}

	// Subscribe to topics
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Message consumed on %s: %s\n", msg.TopicPartition, string(msg.Value))

			//put value in variable
			header := string(msg.Headers[0].Value)
			content := string(msg.Value)

			//func for put value in array
			putValueToArray(header, content)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
