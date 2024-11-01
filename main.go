package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
	visited  bool
}

// add vertex
func (g *Graph) AddVertex(key int) {
	if contains(g.vertices, key) {
		fmt.Printf("the key %v already exist", key)
	} else {
		g.vertices = append(g.vertices, &Vertex{key: key})
	}
}

func (g *Graph) print() {
	for _, v := range g.vertices {

		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
}

// add Edge to the graph
func (g *Graph) addEdge(from, to int) {
	// get Vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check Error if the vertices exist
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v -> %v); missing vertex", from, to)
		fmt.Println(err.Error())
		return
	}

	// add Edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// check if vertices contains
func contains(g []*Vertex, k int) bool {
	for _, v := range g {
		if v.key == k {
			return true
		}
	}
	return false
}

func (g *Graph) bfs(startKey int) {
	startVertex := g.getVertex(startKey)
	if startVertex == nil {
		fmt.Printf("Vertex %v not found in the graph.\n", startKey)
		return
	}
	for _, v := range g.vertices {
		v.visited = false
	}
	queue := []*Vertex{startVertex}
	startVertex.visited = true
	fmt.Printf("BFS starting from vertex %v:\n", startKey)
	for len(queue) > 0 {
		// dequeue the first vertex
		current := queue[0]
		queue = queue[1:]
		fmt.Printf("%v ", current.key)
		// Enqueue all adjacent, unvisited vertices
		for _, neighbor := range current.adjacent {
			if !neighbor.visited {
				neighbor.visited = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}
	test.AddVertex(2)
	test.AddVertex(55)
	test.AddVertex(22)
	test.AddVertex(45)
	test.addEdge(2, 55)  // Valid edge
	test.addEdge(55, 22) // Invalid edge, vertex 99 doesn't exist
	test.addEdge(55, 45) // Valid edge
	test.bfs(55)

	test.print()
}
