# Stage 1: Builder
FROM golang:1.23 AS builder

# Install dependencies for building the Go application and Kafka
RUN apt-get update && apt-get install -y \
    git \
    build-essential \
    librdkafka-dev \
    pkg-config \
    && apt-get clean

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the producer binary (CGO_ENABLED=1 for librdkafka)
RUN CGO_ENABLED=1 GOOS=linux go build -o /producer ./cmd/producer

# Build the consumer binary
RUN CGO_ENABLED=1 GOOS=linux go build -o /consumer ./cmd/consumer

# Stage 2: Final image for producer
FROM debian:bookworm-slim AS producer

# Install required Kafka runtime dependencies
RUN apt-get update && apt-get install -y librdkafka1 && apt-get clean

# Copy the producer binary from the builder stage
COPY --from=builder /producer /app/producer

# Set the working directory
WORKDIR /app

# Expose the port if needed
EXPOSE 8080

# Run the producer app
CMD ["./producer"]

# Stage 3: Final image for consumer
FROM debian:bookworm-slim AS consumer

# Install required Kafka runtime dependencies
RUN apt-get update && apt-get install -y librdkafka1 && apt-get clean

# Copy the consumer binary from the builder stage
COPY --from=builder /consumer /app/consumer

# Set the working directory
WORKDIR /app

# Expose the port if needed
EXPOSE 8081

# Run the consumer app
CMD ["./consumer"]
