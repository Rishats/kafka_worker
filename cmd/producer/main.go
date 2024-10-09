// cmd/producer/main.go
package main

import (
	"fmt"
	kafkainternal "kafka_worker/internal/kafka"
	"os"
	"time"
)

func main() {
	// Get Kafka configuration from environment variables
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		fmt.Println("KAFKA_BROKERS environment variable is not set")
		os.Exit(1)
	}

	clientID := os.Getenv("KAFKA_CLIENT_ID")
	if clientID == "" {
		fmt.Println("KAFKA_CLIENT_ID environment variable is not set")
		os.Exit(1)
	}

	// Initialize Kafka producer
	producer, err := kafkainternal.NewKafkaProducer(brokers, clientID)
	if err != nil {
		fmt.Printf("Error creating producer: %v\n", err)
		os.Exit(1)
	}

	// Produce a message
	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		fmt.Println("KAFKA_TOPIC environment variable is not set")
		os.Exit(1)
	}

	message := "Hello From Go!"

	// every second produce a message and add in to message the current time in human readable format

	for {
		message = fmt.Sprintf("Hello From Go! %s", time.Now().Format("2006-01-02 15:04:05"))
		err = producer.ProduceMessage(topic, message)
		if err != nil {
			fmt.Printf("Error producing message: %v\n", err)
			os.Exit(1)
		}
		time.Sleep(time.Second)

		// Handle events asynchronously
		producer.HandleEvents()
	}
}
