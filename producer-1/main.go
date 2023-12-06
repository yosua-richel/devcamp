package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func main() {
	// Configuration producer
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal("Error creating producer:", err)
	}

	// Send message nsq to the topic
	msg := `{"name": "jane doe","email": "janedoe@example.com"}`
	err = producer.Publish("email_notif_topic", []byte(msg))
	if err != nil {
		log.Fatal("Error publishing message:", err)
	}

	// Wait the signal to end gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Stop the producer
	producer.Stop()
}
