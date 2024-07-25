package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/devd251993/applied_systems_assessment/internal/graphs"
)

func GetGraphDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	graphId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid graph id"))
		return
	}
	graphDetails := graphs.FetchGraphDetails(graphId)
	if graphDetails == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("graph not found"))
		return
	}

	details, err := json.Marshal(graphDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("got get graph request with id : %s\n", id)

	w.WriteHeader(http.StatusOK)
	w.Write(details)
}

func CreateGraphHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got create graph request")
	vertices := r.URL.Query().Get("vertices")
	noOfVertices, err := strconv.Atoi(vertices)
	if err != nil {
		log.Println("invalid number of vertices got")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please enter valid number of vertices"))
		return
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
		return
	}

	destination, err := strconv.Atoi(dest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid destination node"))
		return
	}

	graphID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid graphId"))
		return
	}

	err = graphs.AddEdgeToGraph(graphID, source, destination)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("graph not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetShortestPathHandler(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	id := r.URL.Query().Get("graphId")

	source, err := strconv.Atoi(start)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid start node"))
		return
	}

	destination, err := strconv.Atoi(end)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid end node"))
		return
	}
	graphID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid graph id"))
		return
	}

	path, distance := graphs.GetShortestPath(graphID, source, destination)
	if path == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("unable to found shortest path"))
		return
	}
	type Output struct {
		Path     []int
		Distance int64
	}

	output := Output{
		Path:     path,
		Distance: distance,
	}

	response_data, err := json.Marshal(&output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to write data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response_data)
}

func DeleteGraphHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	log.Printf("got delete graph request with id : %s\n", id)
	w.WriteHeader(http.StatusOK)
}
