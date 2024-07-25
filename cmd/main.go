package main

import (
	"net/http"

	"github.com/devd251993/applied_systems_assessment/internal/server/routes"
)

func main() {
	rounter := routes.InitRoutes()

	if err := http.ListenAndServe(":8080", rounter); err != nil {
		panic(err)
	}
}
