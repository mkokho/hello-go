package main

import (
	"github.com/Shopify/sarama"
	"time"
	"fmt"
	"strings"
	"log"
)

func main() {
	c, err := setupConsumer()
	if err != nil {
		log.Fatalf("Failed to setup consumer: %v", err)
	}

	total := 0
	start := time.Now()

	for {
		select {
		case msg := <-c.Messages():
			if total == 0 {
				start = time.Now()
			}
			total += 1

			if total % 100 == 0 {
				elapsed := time.Since(start)
				log.Printf("offset: %d, elapsed: %v, total: %d, speed: %v",
					msg.Offset,
					elapsed.Seconds(),
					total,
					float64(total) / elapsed.Seconds(),
				)
			}
		}
	}
}

func setupConsumer() (pc sarama.PartitionConsumer, err error) {
	topic := "mc.red.email.sender.in.request"
	brokers := strings.Split("eu-mq-a1.movio.co:9092", ",")

	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V0_10_1_0
	saramaConfig.Consumer.MaxProcessingTime = 1 * time.Second
	saramaConfig.ClientID = fmt.Sprintf("go-email-sender-%v", 0)
	saramaConfig.Consumer.Return.Errors = true

	client, err := sarama.NewClient(brokers, saramaConfig)
	if err != nil {
		return
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return
	}


	pc, err = consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return
	}

	return pc, nil
}
