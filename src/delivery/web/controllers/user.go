package controllers

import (
	"net/http"

	"github.com/stacknowledge/clean-todo-go/src/domain"
	"github.com/stacknowledge/clean-todo-go/src/usecases"

	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

func GetLoginPage(context *gin.Context) {
	successMessages, _ := context.Get("success")
	errorMessages, _ := context.Get("errors")

	context.HTML(http.StatusOK, "login.tmpl", gin.H{
		"success": successMessages,
		"errors":  errorMessages,
	})
}

func GetRegisterPage(context *gin.Context) {
	successMessages, _ := context.Get("success")
	errorMessages, _ := context.Get("errors")

	log.Println(errorMessages)

	context.HTML(http.StatusOK, "register.tmpl", gin.H{
		"success": successMessages,
		"errors":  errorMessages,
	})
}

func PostLogin(context *gin.Context, interactor *usecases.UserInteractor) {
	name := context.PostForm("name")
	email := context.PostForm("email")
	password := context.PostForm("password")

	if name == "" || email == "" || password == "" {
		context.Set("errors", "Invalid register parameters.")
		context.Redirect(302, "/register")
	}

	userID, _ := interactor.Register(&domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	})

	log.Println(userID)

	if userID > 0 {
		context.Set("success", "User created with success.")
		context.Redirect(302, "/")
	}

	context.Set("errors", "Email already registered.")
	context.Redirect(302, "/register")
}
