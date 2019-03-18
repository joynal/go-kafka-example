package main

import (
	"go-kafka-example/utils"
	"os"
)

func main() {
	// give your group name and custom group handler
	utils.LoadConfigs()
	utils.GetNewConsumerGroup("confirmation-worker", os.Getenv("PROVIDER_TOPIC"), utils.ConsumerGroupHandler{})
}
