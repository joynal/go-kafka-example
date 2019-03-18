package utils

import (
	"log"

	"github.com/Shopify/sarama"
)

func TrackGroupErrors(group sarama.ConsumerGroup) {
	for err := range group.Errors() {
		log.Fatal(err)
	}
}
