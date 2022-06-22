package generate_message

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage(ctx context.Context, topic string, kafkaBrokerUrl string) {
	i := 0
	w := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBrokerUrl),
		Topic:    topic,
		Balancer: &kafka.Hash{},
	}

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second)
	}
}
