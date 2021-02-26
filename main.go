package main

import (
	"log"
	"net/http"
	"os"
	"sync"
)

var channelChan = make(chan string)
var producerChan = make(chan string)
var consumerChan = make(chan string)

func main() {
	// Service log setup
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Found error in log ", err)
	}
	log.SetOutput(file)

	// Service setup
	router := server()
	go func() {
		log.Fatal(http.ListenAndServe(":6010", router))
	}()

	//create wait group
	var wg sync.WaitGroup

	//consumer goroutine
	go consumer()

	// producer goroutine
	for {
		select {
		case x := <-channelChan:
			log.Println("New request in channelChan is ready to produce")
			wg.Add(1)
			go producer(&wg, producerChan)
			producerChan <- x
			wg.Wait()
		default:
			continue
		}
	}

}
