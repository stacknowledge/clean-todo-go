package delivery

import (
	"fmt"
	"log"

	"github.com/stacknowledge/clean-todo-go/src/delivery/api"
	"github.com/stacknowledge/clean-todo-go/src/delivery/web"
)

type Application struct {
	Cargo  map[string]ShippableCargo
	Logger log.Logger
}

func (app *Application) Load() *Application {
	app.Cargo = map[string]ShippableCargo{
		"web": new(web.WebCargo),
		"api": new(api.ApiCargo),
	}

	return app
}

func (app *Application) Ship() {
	log.Println("[Shipping]")

	for cargoID, cargo := range app.Load().Cargo {
		log.Println(fmt.Sprintf(" ~> Loading : [%v] cargo", cargoID))
		cargo.Load()
	}

	for {
	}
}
