package api

import (
	"github.com/stacknowledge/clean-todo-go/src/infrastructure"
	"gopkg.in/gin-gonic/gin.v1"
)

type ApiCargo struct {
	StorageRepositories map[string]infrastructure.DbHandler
}

func (cargo *ApiCargo) Load() {
	engine := gin.Default()
	api := engine.Group("/api/v1/")

	api.Use()
	{
		api.GET("/", apiMiddleware)
	}

	go engine.Run(":8081")
}

func apiMiddleware(context *gin.Context) {
	context.JSON(200, "HEY YOU!")
	context.Next()
}
