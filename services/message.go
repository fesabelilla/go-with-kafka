package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/dtos"
	kafka_handler "project/kafka-handler"
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

	if err = kafka_handler.Produce(context, newMessage.Body); err != nil {
		log.Fatal("failed to write messages:", err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"Message": "Message send Successfully",
	})
}
