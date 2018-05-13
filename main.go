package main

import (
	"github.com/stacknowledge/clean-todo-go/src/delivery"
)

//var userId uint

func main() {
	application := new(delivery.Application)
	application.Ship()
}
