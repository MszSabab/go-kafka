package main

import (
	"context"
	"runtime"
	"strings"

	"github.com/MszSabab/go-kafka/pkg/generate_message"
	"github.com/namsral/flag"
)

func main() {
	var (
		kafkaBrokerUrl string
		kafkaTopics    string
	)

	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "localhost:9092", "Kafka brokers in comma separated value")
	flag.StringVar(&kafkaTopics, "kafka-topic", "defaultTopic", "Kafka topic to push")
	flag.Parse()

	ctx := context.Background()
	for _, topic := range strings.Split(kafkaTopics, ",") {
		go generate_message.ProduceMessage(ctx, topic, kafkaBrokerUrl)
	}
	runtime.Goexit()
}
