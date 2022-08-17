package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	kafka_handler "project/kafka-handler"
	"project/src/message/dtos"
)

func init() {
	fmt.Println("INIT func called")
}

func PublishMessage(context *gin.Context) {
	var newMessage dtos.MessageDto

	err := context.ShouldBindJSON(&newMessage)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(newMessage)
	kafka_handler.Produce(context, newMessage.Body)

	context.JSON(http.StatusOK, nil)
}
