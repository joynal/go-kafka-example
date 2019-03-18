package utils

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
)

func GetNewConsumerGroup(name string, topic string, handler sarama.ConsumerGroupHandler) {
	LoadConfigs()

	brokerList := strings.Split(os.Getenv("KAFKA_SERVER_URL"), ",")
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Return.Errors, _ = strconv.ParseBool(os.Getenv("CONSUMER_RETRY_RETURN_SUCCESSES"))
	group, err := sarama.NewConsumerGroup(brokerList, name, config)
	if err != nil {
		log.Fatal(err)
	}
	defer group.Close()

	// Track errors
	go TrackGroupErrors(group)

	ctx := context.Background()
	for {
		group.Consume(ctx, []string{topic}, handler)
	}
}
