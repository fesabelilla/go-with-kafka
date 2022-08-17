package kafka_handler

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func Produce(ctx context.Context, msg string) {
	_, err := writer.WriteMessages(
		kafka.Message{Value: []byte(msg)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}
