package main

import (
	"errors"
	"fmt"
	"math"
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

	if fromVertex.adjacent.containsEdge(to) {
		return errors.New("connection between vertices in this direction has been established already")
	}

	fromVertex.adjacent.Insert(1, &Edge{connection: toVertex, weight: weight})
	return nil
}

func (g *GraphList) DJP() (map[string]*Edge, error) {
	// Initialize cost table that has Edge with connection and cheapest weight
	// Map key is Vertex id and value is the cheapest Edge to that Vertex
	costTable := make(map[string]*Edge)
	notVisited := &Queue{}
	p := g.vertices.head
	for p != nil {
		vertex := p.value.(*Vertex)
		notVisited.Enqueue(vertex)
		costTable[vertex.key.(string)] = &Edge{connection: nil, weight: math.MaxInt32}
		p = p.next
	}

	// Cost of initial node is 0
	root := g.vertices.head.value.(*Vertex)

	costTable[root.key.(string)].weight = 0
	costTable[root.key.(string)].connection = root

	for !notVisited.IsEmpty() {
		v, err := notVisited.Dequeue()
		if err != nil {
			return costTable, errors.New("something went wrong with the queue")
		}
		vertex := v.(*Vertex)
		vertexCost := costTable[vertex.key.(string)].weight

		edges := vertex.adjacent
		k := edges.head
		for k != nil {
			edge := k.value.(*Edge)
			currentCost := costTable[edge.connection.key.(string)].weight

			if currentCost > vertexCost+edge.weight {
				costTable[edge.connection.key.(string)].weight = vertexCost + edge.weight
				costTable[edge.connection.key.(string)].connection = vertex
			}
			k = k.next
		}
	}

	return costTable, nil
}

func (l *LinkedList) findShortestPath() *Edge {
	p := l.head
	min := &Edge{}
	minWeight := math.MaxInt32

	for p != nil {
		edge := p.value.(*Edge)
		p = p.next
		if edge.weight < minWeight {
			min = edge
		}
	}

	return min
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

	// graph.AddVertex("E")
	graph.AddVertex("D")
	graph.AddVertex("C")
	graph.AddVertex("B")
	graph.AddVertex("A")

	// graph.AddEdge("A", "B", 1)
	// graph.AddEdge("A", "D", 4)
	// graph.AddEdge("B", "E", 5)
	// graph.AddEdge("B", "C", 3)
	// graph.AddEdge("B", "D", 2)
	// graph.AddEdge("D", "C", 5)
	graph.AddEdge("A", "B", 2)
	graph.AddEdge("B", "A", 2)
	graph.AddEdge("A", "D", 1)
	graph.AddEdge("D", "A", 1)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("D", "B", 2)
	graph.AddEdge("C", "D", 3)
	graph.AddEdge("D", "C", 3)

	// graph.Print()

	cost, err := graph.DJP()
	if err != nil {
		fmt.Println("error during DJP")
	}
	for k, c := range cost {
		fmt.Println(k, c)
	}
}
