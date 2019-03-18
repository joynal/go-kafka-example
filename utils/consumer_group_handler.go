package utils

import (
	"encoding/json"
	"fmt"
	"go-kafka-example/models"
	"log"

	"github.com/Shopify/sarama"
)

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// construct struct from byte
		var user models.User
		err := json.Unmarshal(msg.Value, &user)
		if err != nil {
			log.Fatal(err)
		}

		// Do your things here, I just printed it
		fmt.Println(user.Email, user.FirstName, user.LastName)

		sess.MarkMessage(msg, "")
	}

	return nil
}
