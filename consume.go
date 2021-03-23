package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consumer(broker string, topics []string, group string) {
	log.Println("Consumer (Kafka) started!")

	// Setting up Consumer (Kafka) config
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": "earliest",
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

			// Send any consumed event to consumerChan
			consumerChan <- string(msg.Value)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func arrConsumer(broker string, topics []string, group string) {

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
	c.SubscribeTopics([]string{"kafkaBiller"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Message consumed on %s: %s\n", msg.TopicPartition, string(msg.Value))
			header := msg.Headers
			arr = append(arr, resConsume{
				Stan:  string(header[0].Value),
				Msgin: string(msg.Value)})
			fmt.Println(arr)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
