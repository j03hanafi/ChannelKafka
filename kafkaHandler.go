package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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

func produceMsgToKafka(broker string, topics []string, message <-chan resConsume) {
	//put channel data in variable
	data := <-message
	fmt.Println("Producer started!")

	// Setting up Consumer (Kafka) config
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	//set variable for topic
	topic := topics[0]
	if err != nil {
		panic(err)
	}
	defer p.Close()

	//set msg header
	header := map[string]string{
		"key":   "uniqueKey",
		"value": data.Head,
	}

	//produce msg
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data.Content),
		Headers:        []kafka.Header{{Key: header["key"], Value: []byte(header["value"])}},
	}, nil)

	p.Flush(3 * 1000)

	log.Println("Producer closing!")

	// Done with worker
}

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
			consumeResult := resConsume{
				Head:    string(msg.Headers[0].Value),
				Content: string(msg.Value),
			}
			//func for put value in array
			putValueToArray(consumeResult)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
