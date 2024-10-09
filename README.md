# Kafka Producer & Consumer with Go

This project demonstrates a Kafka producer and consumer implemented in Go, using the [Confluent Kafka Go client](https://github.com/confluentinc/confluent-kafka-go).

## Requirements

- **Go 1.23** or higher
- A **Kafka cluster** with High Availability (HA) using the **KRaft setup**.
  
  Kafka brokers should be accessible with a multi-node KRaft configuration for redundancy. Ensure that the Kafka cluster is properly set up and running with multiple brokers.

## Environment Variables

You will need to configure the following environment variables in your `docker-compose.yml` or in your environment before running the producer and consumer services:

- `KAFKA_BROKERS`: Comma-separated list of Kafka broker addresses (e.g., `192.168.8.71:9092,192.168.8.72:9092,192.168.8.73:9092`).
- `KAFKA_CLIENT_ID`: The client ID for the Kafka producer.
- `KAFKA_GROUP_ID`: The consumer group ID for the Kafka consumer.
- `KAFKA_TOPIC`: The Kafka topic name to which the producer will send messages and from which the consumer will read messages.

Example configuration:

```yaml
services:
  producer:
    environment:
      KAFKA_BROKERS: "192.168.8.71:9092,192.168.8.72:9092,192.168.8.73:9092"
      KAFKA_CLIENT_ID: "myProducer"
      KAFKA_TOPIC: "myTopic"

  consumer:
    environment:
      KAFKA_BROKERS: "192.168.8.71:9092,192.168.8.72:9092,192.168.8.73:9092"
      KAFKA_GROUP_ID: "myConsumerGroup"
      KAFKA_TOPIC: "myTopic"
```

## Build and Run

You can build and run the services using Docker and Docker Compose.

### 1. Clone the Repository

```bash
git clone https://github.com/Rishats/kafka_worker.git
cd kafka-producer-consumer-go
```

### 2. Build and Start Services

Use Docker Compose to build the images and start the services:

```bash
docker-compose up --build
```

### 3. Verify

- **Producer Logs**: Check if messages are being produced to the Kafka topic.
- **Consumer Logs**: Check if messages are being consumed from the Kafka topic.

```bash
docker-compose logs -f producer
docker-compose logs -f consumer
```

## Kafka Cluster with KRaft Setup

Ensure that you have a Kafka cluster with KRaft mode enabled. KRaft mode removes the dependency on ZooKeeper, providing a simplified, high-availability architecture for Kafka brokers.

You can learn more about KRaft setup in the [Kafka documentation](https://kafka.apache.org/documentation/#kraft).

## Additional Information

- This project is built using a multi-stage Dockerfile to optimize the build process and keep the final Docker images lightweight.
- Make sure the Kafka brokers and topic configurations match your environment.


