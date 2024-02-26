package handler

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"sync" // Import sync package for synchronization primitives

	"github.com/segmentio/kafka-go"
)

var (
	ReceivedMessages []Person
	mu               sync.Mutex // Mutex for synchronizing access to ReceivedMessages
)

type Person struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func ConsumeMessage() {
	log.Println("masuk ke consume message")
	brokerAddress := "localhost:9092"
	topic := "test-topic"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "example-group",
	})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case <-signals:
			log.Println("Received interrupt signal, shutting down...")
			return
		default:

			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("Error reading message: %s", err)
			}

			receiveMessage := Person{}

			err = json.Unmarshal(message.Value, &receiveMessage)
			if err != nil {

				continue
			}

			// Lock before modifying ReceivedMessages
			mu.Lock()
			ReceivedMessages = append(ReceivedMessages, receiveMessage)

			log.Println("Received message:", receiveMessage)
			mu.Unlock()
		}
	}

	reader.Close()
}
