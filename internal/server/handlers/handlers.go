package handlers

import (
	"log"
	"net/http"
)

func GetGraphHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	log.Printf("got get graph request with id : %s\n", id)

	w.WriteHeader(http.StatusOK)
}

func CreateGraphHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got create graph request")
	w.WriteHeader(http.StatusOK)
}

func GetShortestPathHandler(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	log.Printf("got shortest path graph request with start %s and end %s\n", start, end)
	w.WriteHeader(http.StatusOK)
}

func DeleteGraphHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	log.Printf("got delete graph request with id : %s\n", id)
	w.WriteHeader(http.StatusOK)
}
