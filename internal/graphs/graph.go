package graphs

import (
	"errors"
	"log"
	"sync"

	"github.com/yourbasic/graph"
)

type GraphImpl struct {
	mapOfGraphs map[int]*graph.Mutable
	graphId     int
	sync.Mutex
}

type GraphDetails struct {
	Reprentation string
	NoOfNode     int
}

func InitGraph() *GraphImpl {
	return &GraphImpl{
		mapOfGraphs: make(map[int]*graph.Mutable, 0),
		graphId:     0,
	}
}

func (g *GraphImpl) CreateMap(numberOfVertices int) int {
	g.graphId += 1
	newGraph := graph.New(numberOfVertices)
	g.mapOfGraphs[g.graphId] = newGraph
	return g.graphId
}

func (g *GraphImpl) AddEdgeToGraph(graphId, source, destination int) error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	mutableGraph, ok := g.mapOfGraphs[graphId]
	if !ok {
		log.Printf("graph with id %d not found\n", graphId)
		return errors.New("graph not found")
	}
	mutableGraph.AddBothCost(source, destination, 1)
	g.mapOfGraphs[graphId] = mutableGraph
	log.Println(mutableGraph.String())
	return nil
}

func (g *GraphImpl) FetchGraph(graphId int) *graph.Mutable {
	var mutableGraph *graph.Mutable
	var ok bool
	if mutableGraph, ok = g.mapOfGraphs[graphId]; !ok {
		return nil
	}

	return mutableGraph
}

func (g *GraphImpl) FetchGraphDetails(graphId int) *GraphDetails {
	var mutableGraph *graph.Mutable
	var ok bool
	if mutableGraph, ok = g.mapOfGraphs[graphId]; !ok {
		return nil
	}

	return &GraphDetails{
		Reprentation: mutableGraph.String(),
		NoOfNode:     mutableGraph.Order(),
	}
}

func (g *GraphImpl) GetShortestPath(graphId, source, destination int) ([]int, int64) {
	if mutableGraph, ok := g.mapOfGraphs[graphId]; ok {
		return graph.ShortestPath(mutableGraph, source, destination)
	}
	return nil, 0
}

func (g *GraphImpl) DeleteGraph(graphId int) error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	if _, ok := g.mapOfGraphs[graphId]; !ok {
		return errors.New("graph not found")
	}
	delete(g.mapOfGraphs, graphId)
	return nil
}
