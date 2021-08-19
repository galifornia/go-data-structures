package main

import (
	"testing"
)

func TestInitGraph(t *testing.T) {
	graph := InitGraphList()

	if graph == nil || graph.vertices == nil {
		t.Errorf("Graph has not been created properly")
	}

}

func TestAddVertex(t *testing.T) {
	graph := InitGraphList()
	graph.AddVertex("Primer")
	graph.AddVertex("Second")

	if !graph.vertices.containsVertex("Primer") {
		t.Errorf("vertices are not been inserted")
	}

	if graph.vertices.containsVertex("Outro") {
		t.Errorf("vertices contains Outro vertex which was not inserted")
	}

	prevLength := graph.vertices.length
	graph.AddVertex("Primer")
	if prevLength < graph.vertices.length {
		t.Errorf("vertices is able to contain duplicates")
	}
}

func TestAddEdge(t *testing.T) {
	graph := InitGraphList()
	graph.AddVertex("Primer")
	graph.AddVertex("Second")

	graph.AddEdge("Primer", "Second", 351)

	vertex := graph.vertices.selectVertex("Primer")
	if !vertex.adjacent.containsEdge("Second") {
		t.Errorf("edge is not being added")
	}

	secondV := vertex.adjacent.selectEdge("Second")
	if secondV.weight != 351 {
		t.Errorf("edge is not succesfully adding weight")
	}
}
