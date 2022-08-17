package kafka_handler

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

const (
	topic         = "message-log"
	brokerAddress = "localhost:9092"
	partition     = 0
)

var writer *kafka.Conn

func Configure() (w *kafka.Writer, err error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//conn.SetWriteDeadline(time.Now().Add(time.Hour))

	writer = conn
	return w, nil
}
