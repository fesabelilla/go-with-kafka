package routers

import (
	"github.com/gin-gonic/gin"
	"project/config"
	messageService "project/services"
)

func MyRouters() *gin.Engine {
	router := gin.Default()
	router.POST("/message", messageService.PublishMessage)

	router.Run(config.GetEnv().RouterPort)
	return router
}
