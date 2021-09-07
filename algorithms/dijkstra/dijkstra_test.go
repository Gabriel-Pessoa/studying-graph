package dijkstra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	expectedRoute := []string{"0", "0 - 1", "0 - 2", "0 - 1 - 3", "0 - 1 - 4", "0 - 1 - 4 - 5"}
	expectedDist := []int{0, 2, 4, 6, 4, 6}
	adjMatrix := [][]int{
		{0, 2, 4, 0, 0, 0},
		{2, 0, 2, 4, 2, 0},
		{4, 2, 0, 0, 3, 0},
		{0, 4, 0, 0, 3, 2},
		{0, 2, 3, 3, 0, 2},
		{0, 0, 0, 2, 2, 0},
	}

	route, dist := Dijkstra(adjMatrix, 0)

	assert.Equal(t, len(route), len(dist))
	assert.Equal(t, expectedRoute, route)
	assert.Equal(t, expectedDist, dist)
}
