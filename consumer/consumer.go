package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{"localhost:9092"} // Substitua pelo endereço do seu broker
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Error on creat consumer: %v", err)
	}
	defer consumer.Close()

	topic := "test_topic"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	log.Println("Getting messages...")
	for message := range partitionConsumer.Messages() {
		log.Printf("Message received: %s\n", string(message.Value))
	}
}
