package web

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

type WebCargo struct {
	Logger log.Logger
}

func (cargo *WebCargo) Load() {
	engine := gin.Default()

	engine.GET("/", webMiddleware)
	go engine.Run(":8080")
}

func webMiddleware(context *gin.Context) {
	context.JSON(200, "Hey men!")
	context.Next()
}
