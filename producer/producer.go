package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Error on creat producer: %v", err)
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: "test_topic",
		Value: sarama.StringEncoder("Hello Kafka!"),
	}

	for i := 0; i < 10; i++ {
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Fatalf("Error on send message: %v", err)
		}

		log.Printf("Message send to %d, offset %d\n", partition, offset)
	}
}
