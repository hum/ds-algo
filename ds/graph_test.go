package ds_test

import (
	"strconv"
	"testing"

	"github.com/hum/ds-algo/ds"
)

func TestGraphVertexesAndEdges(t *testing.T) {
	graph := ds.Graph{}
	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")

	graph.AddEdge("3", "1")
	graph.AddEdge("3", "4")
	graph.AddEdge("4", "2")
	graph.AddEdge("4", "5")
	graph.AddEdge("1", "2")
	graph.AddEdge("1", "0")
	graph.AddEdge("0", "2")
	graph.AddEdge("6", "5")

	tests := []struct {
		wrong string
	}{
		{"50"},
		{"100"},
		{"?"},
		{"-150"},
		{"10"},
		{"abc"},
		{"77"},
	}

	for i, tt := range tests {
		if err := graph.AddEdge(tt.wrong, strconv.Itoa(i)); err == nil {
			t.Fatalf("Graph should have errored out, because node 1 - [%s] does not exist. Expected=error, got a valid response.", tt.wrong)
		}
	}

	for i := 0; i < 6; i++ {
		if err := graph.AddVertex(strconv.Itoa(i)); err == nil {
			t.Fatalf("Node already exists, but graph did not error out. Expected=error, got a valid response.")
		}
	}
}
