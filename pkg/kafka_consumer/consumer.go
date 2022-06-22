package kafka_consumer

import (
	"context"
	"fmt"

	"github.com/namsral/flag"
	"github.com/segmentio/kafka-go"
)

func Readcon() {
	var (
		topicName    string
		consumerName string
		groupName    string
	)

	flag.StringVar(&topicName, "topic_name", "topic1", "topic_name")
	flag.StringVar(&consumerName, "consumer_name", "default_consumer", "consumer_name")
	flag.StringVar(&groupName, "group_name", "g1", "group_name")

	flag.Parse()

	fmt.Println(topicName, consumerName, groupName)

	readcon := kafka.ReaderConfig{
		Brokers:  []string{"localhost: 9092"},
		Topic:    topicName,
		GroupID:  groupName,
		MaxBytes: 10,
	}
	reader := kafka.NewReader(readcon)

	fmt.Println(readcon)

	for {
		fmt.Println("yes")
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("some error occured", err)
			continue
		}
		fmt.Println(consumerName + ": message from " + topicName + " under group" + groupName + " => " + string(m.Value))
	}
}
