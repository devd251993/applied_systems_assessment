package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devd251993/applied_systems_assessment/internal/server/handlers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/graph/create", handlers.CreateGraphHandler).Methods("POST")
	router.HandleFunc("/graph/get", handlers.GetGraphHandler).Queries("id", "{id}").Methods("GET")
	router.HandleFunc("/graph/shortest-path", handlers.GetShortestPathHandler).Queries("start", "{start}").Queries("end", "{end}").Methods("GET")
	router.HandleFunc("/graph/delete", handlers.DeleteGraphHandler).Queries("id", "{id}").Methods("DELETE")

	router.Use(Middleware)
	return router
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		w.Header().Set("Content-Type", "json")
		next.ServeHTTP(w, r)
	})
}
