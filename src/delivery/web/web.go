package web

import (
	"github.com/stacknowledge/clean-todo-go/src/delivery/web/controllers"
	"github.com/stacknowledge/clean-todo-go/src/infrastructure"
	"github.com/stacknowledge/clean-todo-go/src/infrastructure/database"
	"github.com/stacknowledge/clean-todo-go/src/interfaces"
	"github.com/stacknowledge/clean-todo-go/src/usecases"

	"gopkg.in/gin-gonic/gin.v1"
)

type WebCargo struct {
	StorageRepositories *map[string]infrastructure.DbHandler
}

func (cargo *WebCargo) Load() {
	engine := gin.Default()
	dbHandler, _ := database.NewSqliteHandler("./database.sqlite")

	handlers := make(map[string]infrastructure.DbHandler)
	handlers["UserRepository"] = dbHandler
	handlers["TodoRepository"] = dbHandler

	cargo.StorageRepositories = &handlers

	engine.LoadHTMLGlob("./src/delivery/web/templates/*/*")

	engine.GET("/", webMiddleware)
	engine.GET("/login", controllers.GetLoginPage)
	engine.GET("/register", controllers.GetRegisterPage)

	engine.POST("/register", func(context *gin.Context) {
		userInteractor := new(usecases.UserInteractor)
		userInteractor.UserRepository = interfaces.NewDbUserRepository(*cargo.StorageRepositories)

		controllers.PostLogin(context, userInteractor)
	})

	go engine.Run(":8080")
}

func webMiddleware(context *gin.Context) {
	context.JSON(200, "Hey men!")
	context.Next()
}
