package utils

import (
	"encoding/json"
	"go-kafka-example/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/go-querystring/query"
)

var httpClient *http.Client

const (
	maxIdleConnections int = 20
	requestTimeout     int = 10
)

func init() {
	httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
	}

	return client
}

// AddSubscription will create subscription at klaviyo
func AddSubscription(msg *sarama.ConsumerMessage, sess sarama.ConsumerGroupSession, klaviyoURL string, maxChan chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer func(maxChan chan bool) { <-maxChan }(maxChan)

	// construct struct from byte
	var subscription models.Subscription
	err := json.Unmarshal(msg.Value, &subscription)
	if err != nil {
		log.Fatal(err)
	}

	// url encode data
	data, _ := query.Values(subscription)

	// prepare http request
	req, err := http.NewRequest("POST", klaviyoURL, strings.NewReader(data.Encode()))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Connection", "close")

	res, err := httpClient.Do(req)

	if err != nil {
		log.Println("client error ------------>", err)
		return
	}

	defer res.Body.Close()

	log.Println("response Status ------------>", res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	log.Println("response Body ------------>", string(body))

	// commit kafka message
	sess.MarkMessage(msg, "")
}
