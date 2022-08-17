package main

import (
	kafka_handler "project/kafka-handler"
	"project/routers"
)

func main() {

	go kafka_handler.Configure()

	// create a new context
	// ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking

	//go kafka_handler.Produce(ctx)
	go kafka_handler.Consumer()

	routers.MyRouters()
}
