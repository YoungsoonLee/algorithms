package main

import "fmt"

// Graph struct
type Graph struct {
	vertices []*Vertex
}

// Vertex struct
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add vertex
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Println("Vertex already exists")
		return
	}

	g.vertices = append(g.vertices, &Vertex{key: k})
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		// error
		fmt.Println("ERROR: not exists vertex")
		return
	} else if contains(fromVertex.adjacent, to) {
		fmt.Printf("ERROR: Existing edge(%v -> %v)\n", from, to)
		return
	}

	// add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}

	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.AddEdge(6, 2)

	test.Print()
}
