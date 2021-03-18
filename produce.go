package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func producer(wg *sync.WaitGroup, broker string, topics []string, message <-chan string) {
	fmt.Println("Producer started!")

	// Setting up Consumer (Kafka) config
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Run go routine for produce available event to Kafka
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Produce failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Produced message to %v. Message: %s (Header: %s)\n", ev.TopicPartition, ev.Value, ev.Headers)
				}
			}
		}
	}()

	// Setting up kafka message to get ready to be produced
	// Selected producer topic
	topic := topics[0]

	// message to be produced
	msg := <-message

	// header for the message
	header := map[string]string{
		"key":   "testHeader",
		"value": "headers value are binary",
	}
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
		Headers:        []kafka.Header{{Key: header["key"], Value: []byte(header["value"])}},
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(3 * 1000)
	log.Println("Producer closing!")

	// Done with worker
	wg.Done()
}

func arrProduce(msg string, head string) {
	fmt.Println("Producer started!")

	// Setting up Consumer (Kafka) config
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	topic := "channelKafka"
	if err != nil {
		panic(err)
	}
	defer p.Close()

	header := map[string]string{
		"key":   "uniqueKey",
		"value": head,
	}
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
		Headers:        []kafka.Header{{Key: header["key"], Value: []byte(header["value"])}},
	}, nil)
}

func checkResponse(head string) string {
	search := true
	var res string
	for search {
		for _, ele := range arr {
			if ele.stan == head {
				search = false
				res = ele.msgin
			}
		}
	}
	return res
}
