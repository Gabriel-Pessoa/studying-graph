package dijkstra

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Gabriel-Pessoa/studying-graph/data-structures/stack"
)

const INF = math.MaxInt64

// arguments: adjacency matrix, origin vertex.
// returns: distance, route. Ex: route[x] have distance[x]
func Dijkstra(adjMatrix [][]int, src int) (route []string, dist []int) {
	length := len(adjMatrix)

	dist = make([]int, length)
	visited := make([]bool, length)
	predecessors := make(map[int]int)

	for i := 0; i < length; i++ {
		dist[i] = INF
	}

	dist[src] = 0

	for i := 0; i < length-1; i++ {
		u := minDistance(dist, visited)

		visited[u] = true

		for v := 0; v < length; v++ {
			if !visited[v] && adjMatrix[u][v] != 0 && dist[u] != INF && dist[u]+adjMatrix[u][v] < dist[v] {
				dist[v] = dist[u] + adjMatrix[u][v]
				predecessors[v] = u
			}
		}
	}

	route = road(predecessors, length, src)

	return
}

func minDistance(dist []int, visited []bool) int {
	min := INF
	minIndex := -1
	for v := 0; v < len(dist); v++ {
		if !visited[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex
}

func road(predecessors map[int]int, size, src int) []string {
	route := make([]string, size)
	route[src] = "0"

	for i := 1; i < size; i++ {
		toVertex := i
		path := stack.NewStack()

		for v := toVertex; v != src; v = predecessors[v] {
			path.Push(v)
		}

		path.Push(src)

		elem, err := path.Pop()
		if err != nil {
			panic("fail to build path")
		}

		s := strconv.Itoa(elem)

		for !path.IsEmpty() {
			elem, err = path.Pop()
			if err != nil {
				panic("fail to build path")
			}

			s += fmt.Sprintf(" - %v", elem)
		}

		route[i] = s
	}

	return route
}
