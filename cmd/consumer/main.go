// cmd/consumer/main.go
package main

import (
	"fmt"
	kafkainternal "kafka_worker/internal/kafka"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Get Kafka configuration from environment variables
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		fmt.Println("KAFKA_BROKERS environment variable is not set")
		os.Exit(1)
	}

	groupID := os.Getenv("KAFKA_GROUP_ID")
	if groupID == "" {
		fmt.Println("KAFKA_GROUP_ID environment variable is not set")
		os.Exit(1)
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		fmt.Println("KAFKA_TOPIC environment variable is not set")
		os.Exit(1)
	}

	// Initialize Kafka consumer
	consumer, err := kafkainternal.NewKafkaConsumer(brokers, groupID, []string{topic})
	if err != nil {
		fmt.Printf("Error creating consumer: %v\n", err)
		os.Exit(1)
	}

	// Handle interrupts for graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Run the consumer in a goroutine
	go func() {
		consumer.ConsumeMessages()
	}()

	// Wait for termination signal
	<-sigchan
	fmt.Println("Shutting down consumer...")

	// Close the consumer gracefully
	consumer.Close()
}
