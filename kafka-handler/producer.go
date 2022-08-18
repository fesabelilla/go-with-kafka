package kafka_handler

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context, msg string) (err error) {
	_, err = writer.WriteMessages(
		kafka.Message{Value: []byte(msg)},
	)
	return err
}
