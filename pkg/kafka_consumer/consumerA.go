package kafka_consumer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Readcon() {
	readcon := kafka.ReaderConfig{
		Brokers:  []string{"localhost: 9092"},
		Topic:    "topic1",
		GroupID:  "g1",
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
		fmt.Println("message is:", string(m.Value))

	}
}
