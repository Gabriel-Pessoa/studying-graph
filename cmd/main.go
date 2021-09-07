package main

import (
	"fmt"

	"github.com/Gabriel-Pessoa/studying-graph/algorithms/dijkstra"
	"github.com/Gabriel-Pessoa/studying-graph/data-structures/graph"
)

func main() {
	fmt.Printf("Start\n\n")

	// *********** Graph data structure  ****************
	graph := graph.NewGraph(false)
	myVertices := [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	for _, v := range myVertices {
		graph.AddVertex(v)
	}
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("A", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "E")
	graph.AddEdge("C", "G")
	graph.AddEdge("D", "G")
	graph.AddEdge("D", "H")
	graph.AddEdge("B", "E")
	graph.AddEdge("B", "F")
	graph.AddEdge("E", "I")

	fmt.Println(graph.ToString())

	// *********** Dijkstra's algorithm ****************
	adjMatrix := [][]int{
		{0, 2, 4, 0, 0, 0},
		{2, 0, 2, 4, 2, 0},
		{4, 2, 0, 0, 3, 0},
		{0, 4, 0, 0, 3, 2},
		{0, 2, 3, 3, 0, 2},
		{0, 0, 0, 2, 2, 0},
	}

	route, dist := dijkstra.Dijkstra(adjMatrix, 0)

	for i := 0; i < len(route) && i < len(dist); i++ {
		fmt.Printf("route: %s, distance: %v \n", route[i], dist[i])
	}

	fmt.Printf("\nEnd\n")
}
