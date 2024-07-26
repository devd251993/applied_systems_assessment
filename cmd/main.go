package main

import (
	"log"
	"net/http"

	"github.com/devd251993/applied_systems_assessment/internal/server/handlers"
	"github.com/devd251993/applied_systems_assessment/internal/server/routes"
)

func main() {
	rounter := routes.InitRoutes()
	handlers.InitHandlers()
	log.Println("Starting server on :8080 port")
	if err := http.ListenAndServe(":8080", rounter); err != nil {
		panic(err)
	}
}
