# Kafka Communication Testing with Go
## This project demonstrates how to implement and test communication with Apache Kafka using the Go programming language. It includes a Kafka producer for sending messages to a specified topic and a consumer for reading messages from the same topic. The project leverages the Sarama library for Kafka integration and provides Docker Compose configurations to set up the required Kafka and Zookeeper services locally. This setup enables seamless testing of Kafka's publish-subscribe messaging pattern.

## To execute this project use: 

### docker-compose up -d
### Terminal 1 - go run consumer/consumer.go 
### Terminal 2 - go run producer/producer.go

