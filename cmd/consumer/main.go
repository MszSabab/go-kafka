package main

import (
	"fmt"

	"github.com/MszSabab/go-kafka/pkg/kafka_consumer"
)

func main() {
	fmt.Println("kafka consumer started ====================>")
	kafka_consumer.Readcon()
}
