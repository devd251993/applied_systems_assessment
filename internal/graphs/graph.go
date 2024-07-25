package graphs

import (
	"errors"
	"log"

	"github.com/yourbasic/graph"
)

var (
	MapOfGraphs map[int]*graph.Mutable
	graphId     int
)

func init() {
	MapOfGraphs = make(map[int]*graph.Mutable)
	graphId = 0
}

func CreateMap(numberOfVertices int) int {
	graphId += 1
	newGraph := graph.New(numberOfVertices)
	MapOfGraphs[graphId] = newGraph
	return graphId
}

func AddEdgeToGraph(graphId, source, destination int) error {
	mutableGraph, ok := MapOfGraphs[graphId]
	if !ok {
		log.Printf("graph with id %d not found\n", graphId)
		return errors.New("graph not found")
	}
	mutableGraph.AddBothCost(source, destination, 1)
	MapOfGraphs[graphId] = mutableGraph
	log.Println(mutableGraph.String())
	return nil
}

func FetchGraph(graphId int) *graph.Mutable {
	var mutableGraph *graph.Mutable
	var ok bool
	if mutableGraph, ok = MapOfGraphs[graphId]; !ok {
		return nil
	}

	return mutableGraph
}
