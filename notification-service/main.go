package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	defer consumer.Close()

	partitionConsumer, _ := consumer.ConsumePartition("todo-events", 0, sarama.OffsetNewest)
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
		// Add your notification logic here
	}
}
