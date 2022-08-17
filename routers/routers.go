package routers

import (
	"github.com/gin-gonic/gin"
	messageService "project/src/message/services"
)

func MyRouters() *gin.Engine {
	router := gin.Default()
	router.POST("/message", messageService.PublishMessage)

	router.Run("localhost:9090")
	return router
}
