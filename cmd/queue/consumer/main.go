package main

import (
	"os"
	"sync"

	"go-kafka-example/utils"

	"github.com/Shopify/sarama"
)

const (
	maxConcurrency = 100
)

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	utils.LoadConfigs()

	var wg sync.WaitGroup
	maxChan := make(chan bool, maxConcurrency)

	for msg := range claim.Messages() {
		maxChan <- true
		go utils.AddSubscription(msg, sess, os.Getenv("KLAVIYO_URL"), maxChan, &wg)
	}
	wg.Wait()

	return nil
}

func main() {
	utils.LoadConfigs()
	utils.GetNewConsumerGroup("subscription-worker", os.Getenv("PROVIDER_TOPIC"), consumerGroupHandler{})
}
