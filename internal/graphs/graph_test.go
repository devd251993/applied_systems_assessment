package graphs_test

import (
	"testing"

	"github.com/devd251993/applied_systems_assessment/internal/graphs"
)

func TestCreateMap(t *testing.T) {
	testCases := []struct {
		desc             string
		numberOfVertices int
	}{
		{
			desc:             "create a empty graph with 3 vertices",
			numberOfVertices: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			graphId := graphs.CreateMap(tC.numberOfVertices)
			graph := graphs.FetchGraph(graphId)
			if graph == nil || graph.Order() != tC.numberOfVertices {
				t.Errorf("error creating graph")
			}
		})
	}
}

func TestAddEdgeToGraph(t *testing.T) {
	testGraphId := graphs.CreateMap(3)
	type args struct {
		source  int
		dest    int
		graphId int
	}
	testCases := []struct {
		desc      string
		args      args
		wantError bool
	}{
		{
			desc: "happy path",
			args: args{
				source:  0,
				dest:    1,
				graphId: testGraphId,
			},
			wantError: false,
		},
		{
			desc: "graph is not present",
			args: args{
				source:  0,
				dest:    1,
				graphId: 10,
			},
			wantError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := graphs.AddEdgeToGraph(tC.args.graphId, tC.args.source, tC.args.dest)
			if tC.wantError != (err != nil) {
				t.Errorf("Wanted Error : %v but got Error: %v", tC.wantError, !tC.wantError)
			}
		})
	}
}
