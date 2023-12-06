package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

type (
	// Handler for NSQ message
	MessageHandler struct{}

	MessagePayload struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

// HandleMessage implementation from interface `nsq.Handler`
func (h *MessageHandler) HandleMessage(message *nsq.Message) error {
	body := message.Body

	// print the message
	log.Printf("Received message: %s \n", body)

	// var data MessagePayload
	// err := json.Unmarshal(body, &data)
	// if err != nil {
	// 	message.Finish()
	// 	return nil
	// }

	// log.Printf("Received message struct: %+v \n", data)

	return nil
}

func main() {
	// Configuration producer
	config := nsq.NewConfig()

	// Init the consumer
	consumer, err := nsq.NewConsumer("email_notif_topic", "email_notif_3_channel", config)
	if err != nil {
		log.Fatal("Error creating consumer:", err)
	}

	// Set handler untuk consumer
	consumer.AddHandler(&MessageHandler{})

	// Connect to NSQD
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal("Error connecting to NSQ:", err)
	}

	// Wait the signal to end gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Stop the consumer
	consumer.Stop()
}
