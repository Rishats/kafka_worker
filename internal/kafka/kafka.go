// internal/kafka/kafka.go
package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
}

type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

// Producer initialization
func NewKafkaProducer(brokers string, clientID string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"client.id":         clientID,
		"acks":              "all"})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	return &KafkaProducer{Producer: p}, nil
}

// Produce messages
func (kp *KafkaProducer) ProduceMessage(topic string, value string) error {
	err := kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value)},
		nil) // delivery channel

	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	return nil
}

// Handle producer events
func (kp *KafkaProducer) HandleEvents() {
	go func() {
		for e := range kp.Producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()
}

// Consumer initialization
func NewKafkaConsumer(brokers string, groupID string, topics []string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to topics: %w", err)
	}

	return &KafkaConsumer{Consumer: c}, nil
}

// Consume messages
func (kc *KafkaConsumer) ConsumeMessages() {
	for {
		msg, err := kc.Consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		fmt.Printf("Consumed message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}
}

// Close consumer
func (kc *KafkaConsumer) Close() {
	if err := kc.Consumer.Close(); err != nil {
		fmt.Printf("Failed to close consumer: %v\n", err)
	}
}
