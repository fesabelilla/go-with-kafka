package kafka_handler

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"project/config"
)

var writer *kafka.Conn

func Configure() (w *kafka.Writer, err error) {
	connection, err := kafka.DialLeader(context.Background(), "tcp", config.GetEnv().BrokerAddress, config.GetEnv().Topic, config.GetEnv().Partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	writer = connection
	return w, nil
}
