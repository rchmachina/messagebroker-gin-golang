package main

import (
	"context"
	"math/rand"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/google/uuid"
	"github.com/rchmachina/playingBrokerMessage/producer/helper"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka broker address
	brokerAddress := "localhost:9092"
	// Kafka topic to produce messages to
	topic := "test-topic"

	// Create a new Kafka producer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	// Trap SIGINT to gracefully shutdown the producer
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	var mu sync.Mutex
	// Number of worker goroutines
	maxWorkers := 10
	// Channel to communicate with worker goroutines
	messageChan := make(chan struct{}, maxWorkers)
	// Wait group to wait for all worker goroutines to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		log.Println("worker ke ", i)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-interrupt:
					log.Println("Received interrupt signal, shutting down worker...")
					return
				case <-messageChan:
					mu.Lock()
					uuid := uuid.New()
	
    				randomNumber := rand.Intn(100)
					person := struct {
						Message string `json:"message"`
						ID      string `json:"id"`
					}{
						Message: helper.RandomWord(randomNumber),
						ID:      uuid.String(),
					}
					mu.Unlock()

					messageMarshal, err := json.Marshal(person)
					if err != nil {
						panic(err)
					}
					err = writer.WriteMessages(context.Background(), kafka.Message{
						Value: messageMarshal,
					})
					if err != nil {
						log.Fatalf("Error producing message: %s", err)
					}
					log.Println("success sending message", person)
					
				}
			}
		}()
	}

	// Produce messages
	for i := 0; i < 9999; i++ {
		// Send a message to one of the worker goroutines
		messageChan <- struct{}{}
	}

	// Wait for all worker goroutines to finish
	//wg.Wait()

	// Close the writer
	writer.Close()
}
