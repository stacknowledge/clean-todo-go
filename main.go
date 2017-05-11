package main

import (
	"net/http"

	"log"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/stacknowledge/go-clean-todo/src/domain"
	database "github.com/stacknowledge/go-clean-todo/src/infrastructure/database"
	"github.com/stacknowledge/go-clean-todo/src/interfaces"
	"github.com/stacknowledge/go-clean-todo/src/usecases"
)

func main() {
	dbHandler, err := database.NewSqliteHandler("./database.sqlite")
	router := mux.NewRouter()

	if err != nil {
		log.Fatal(err)
	}

	handlers := make(map[string]database.DbHandler)
	handlers["UserRepository"] = dbHandler

	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDbUserRepository(handlers)

	router.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		_, error := userInteractor.Register(domain.User{
			Name:     "Peters",
			Email:    "peters@badass.com",
			Password: "petersboysband",
		})

		if error != nil {
			panic(error)
		}
	})

	router.HandleFunc("/user/{id}", func(res http.ResponseWriter, req *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(req)["id"])
		panic(userInteractor.Profile(uint(id)))
	})

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
