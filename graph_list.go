package main

import (
	"errors"
	"fmt"
)

type Edge struct {
	connection *Vertex
	weight     int
}

type Vertex struct {
	key      interface{}
	adjacent *LinkedList
}

type GraphList struct {
	vertices *LinkedList
}

func (g *GraphList) AddVertex(value interface{}) error {
	if g.vertices.containsVertex(value) {
		return errors.New("Vertex not added because it is already an existing key")
	}
	g.vertices.Insert(1, &Vertex{key: value, adjacent: &LinkedList{}})
	return nil
}

func (g *GraphList) Print() {
	g.vertices.Print()
}

func (g *GraphList) AddEdge(from, to interface{}, weight int) error {
	fromVertex := g.vertices.selectVertex(from)
	toVertex := g.vertices.selectVertex(to)

	if fromVertex == nil || toVertex == nil {
		return errors.New("introduced 'from' or 'to' vertex does not exist")
	}

	if fromVertex.adjacent.containsVertex(to) {
		return errors.New("connection between vertices in this direction has been established already")
	}

	fromVertex.adjacent.Insert(1, &Edge{connection: toVertex, weight: weight})
	return nil
}

func (l *LinkedList) selectVertex(v interface{}) *Vertex {
	p := l.head

	for p != nil {
		vertex := p.value.(*Vertex)
		if vertex.key == v {
			return vertex
		}
		p = p.next
	}
	return nil
}

func (l *LinkedList) containsVertex(v interface{}) bool {
	p := l.head
	for p != nil {
		vertex := p.value.(*Vertex)
		if vertex.key == v {
			return true
		}
		p = p.next
	}
	return false
}

func (l *LinkedList) selectEdge(v interface{}) *Edge {
	p := l.head

	for p != nil {
		edge := p.value.(*Edge)
		if edge.connection.key == v {
			return edge
		}
		p = p.next
	}
	return nil
}

func (l *LinkedList) containsEdge(v interface{}) bool {
	p := l.head
	for p != nil {
		edge := p.value.(*Edge)
		if edge.connection.key == v {
			return true
		}
		p = p.next
	}
	return false
}

func InitGraphList() *GraphList {
	graph := &GraphList{}
	graph.vertices = &LinkedList{}
	return graph
}

func main() {
	graph := &GraphList{}
	graph.vertices = &LinkedList{}
	graph.AddVertex("Primer")
	graph.AddVertex("Second")
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(0)
	graph.AddVertex("Last")

	graph.AddVertex("Last")
	graph.AddVertex("Last")
	graph.Print()

	graph.AddEdge("Last", "Primer", 34)
	graph.Print()
	err := graph.AddEdge("Last", "Primer", 10)
	if err != nil {
		fmt.Println(err)
	}
	last := graph.vertices.selectVertex("Last")
	last.adjacent.Print()
}
