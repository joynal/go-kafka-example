package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"go-kafka-example/models"
	"go-kafka-example/utils"

	"github.com/Shopify/sarama"
	"github.com/icrowley/fake"
)

func main() {
	utils.LoadConfigs()

	brokerList := strings.Split(os.Getenv("KAFKA_SERVER_URL"), ",")
	topic := os.Getenv("PROVIDER_TOPIC")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max, _ = strconv.Atoi(os.Getenv("PRODUCER_RETRY_MAX"))
	config.Producer.Return.Successes, _ = strconv.ParseBool(os.Getenv("PRODUCER_RETRY_RETURN_SUCCESSES"))

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	data, _ := json.Marshal(models.User{
		Email:     fake.EmailAddress(),
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
	})

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(data)),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("send error ------------->", err)
	}

	log.Printf("sent at -------------> %s/partition - %d/offset - %d\n", topic, partition, offset)
}
