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

	for i := 0; i < 10; i++ {
		data, _ := json.Marshal(models.Subscription{
			APIKey: os.Getenv("KLAVIYO_API_KEY"),
			Email:  fake.EmailAddress(),
			Properties: models.Property{
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
			},
			ConfirmOptin: false,
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

}
