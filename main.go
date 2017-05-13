package main

import (
	"net/http"

	"log"

	"strconv"

	"fmt"

	"time"

	"github.com/gorilla/mux"
	"github.com/stacknowledge/go-clean-todo/src/domain"
	"github.com/stacknowledge/go-clean-todo/src/infrastructure"
	"github.com/stacknowledge/go-clean-todo/src/infrastructure/database"
	"github.com/stacknowledge/go-clean-todo/src/interfaces"
	"github.com/stacknowledge/go-clean-todo/src/usecases"
)

var userId uint

func main() {

	dbHandler, err := database.NewSqliteHandler("./database.sqlite")
	router := mux.NewRouter()

	if err != nil {
		log.Fatal(err)
	}

	handlers := make(map[string]infrastructure.DbHandler)
	handlers["UserRepository"] = dbHandler
	handlers["TodoRepository"] = dbHandler

	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDbUserRepository(handlers)

	todoInteractor := new(usecases.TodoInteractor)
	todoInteractor.TodoRepository = interfaces.NewDbTodoRepository(handlers)

	router.HandleFunc("/user/register", func(res http.ResponseWriter, req *http.Request) {
		userId, error := userInteractor.Register(domain.User{
			Name:     "Peters",
			Email:    "petersx@badass.com",
			Password: "petersboysband",
		})

		if error != nil {
			panic(error)
		}

		log.Println(fmt.Sprintf("User added ID {%d}", userId))
	})

	router.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(req)["id"])
		log.Println(fmt.Sprintf("User profile: {%v}", userInteractor.Profile(uint(id))))
	})

	router.HandleFunc("/todo/create", func(res http.ResponseWriter, req *http.Request) {
		todoId, err := todoInteractor.Store(&domain.Todo{
			Title:   "O todo que a gente faz whateevveerssss",
			Expires: time.Now().Format("02-01-2006 03:04:05"),
			User:    userInteractor.Profile(userId),
		})

		if err != nil {
			panic(err)
		}

		log.Println(fmt.Sprintf("Todo added ID {%d}", todoId))
	})

	router.HandleFunc("/todo/{id}", func(res http.ResponseWriter, req *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(req)["id"])
		log.Println(fmt.Sprintf("Todo: {%v}", todoInteractor.FindById(uint(id))))
	})

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
