package main

import (
	"context"
	"github.com/MszSabab/go-kafka/pkg/kafka_producer"
	"github.com/namsral/flag"
	"runtime"
	"strings"
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
		go kafka_producer.ProduceMessage(ctx, topic, kafkaBrokerUrl)
	}
	runtime.Goexit()
}
