package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/devd251993/applied_systems_assessment/internal/graphs"
)

func GetGraphHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	log.Printf("got get graph request with id : %s\n", id)

	w.WriteHeader(http.StatusOK)
}

func CreateGraphHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got create graph request")
	vertices := r.URL.Query().Get("vertices")
	noOfVertices, err := strconv.Atoi(vertices)
	if err != nil {
		log.Println("invalid number of vertices got")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please enter valid number of vertices"))
	}
	graphId := graphs.CreateMap(noOfVertices)
	outputString := "Graph with id: " + strconv.Itoa(graphId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outputString))
}

func AddEdgeToGraph(w http.ResponseWriter, r *http.Request) {
	sour := r.URL.Query().Get("source")
	dest := r.URL.Query().Get("destination")
	id := r.URL.Query().Get("graphId")

	source, err := strconv.Atoi(sour)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid source node"))
	}

	destination, err := strconv.Atoi(dest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid destination node"))
	}

	graphID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid graphId"))
	}

	err = graphs.AddEdgeToGraph(graphID, source, destination)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("graph not found"))
	}

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
