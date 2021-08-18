package graph

import (
	"errors"
	"testing"

	"github.com/Gabriel-Pessoa/studying-graph/data-structures/dictionary"
	"github.com/stretchr/testify/assert"
)

type valuePair struct {
	Vertex, Edge i
}

func TestAddVertex(t *testing.T) {
	errAddVertex := errors.New("fail to add vertex in the graph")

	var tests = []struct {
		name       string
		vertices   []i
		wantError  []error
		wantResult []i
	}{
		{
			name:       "insert invalid: duplicate vertex string",
			vertices:   []i{"A", "B", "C", "A"},
			wantError:  []error{nil, nil, nil, errAddVertex},
			wantResult: []i{"A", "B", "C"},
		},
		{
			name:       "insert invalid: duplicate vertex int",
			vertices:   []i{1, 2, 2, 3},
			wantError:  []error{nil, nil, errAddVertex, nil},
			wantResult: []i{1, 2, 3},
		},
		{
			name:       "insert invalid: duplicate vertex float",
			vertices:   []i{1.1, 2.2, 2.3, 2.2, 3.2},
			wantError:  []error{nil, nil, nil, errAddVertex, nil},
			wantResult: []i{1.1, 2.2, 2.3, 3.2},
		},
		{
			name:       "insert invalid: empty vertex",
			vertices:   []i{"A", "", "C", ""},
			wantError:  []error{nil, errAddVertex, nil, errAddVertex},
			wantResult: []i{"A", "C"},
		},
		{
			name:       "insert invalid: white space",
			vertices:   []i{"A", "B", "C", "   ", "D"},
			wantError:  []error{nil, nil, nil, errAddVertex, nil},
			wantResult: []i{"A", "B", "C", "D"},
		},
		{
			name:       "insert with success: vertex as string",
			vertices:   []i{"A", "B", "C", "D"},
			wantError:  []error{nil, nil, nil, nil},
			wantResult: []i{"A", "B", "C", "D"},
		},
		{
			name:       "insert with success: vertex as number",
			vertices:   []i{1, 2, 3, 4},
			wantError:  []error{nil, nil, nil, nil},
			wantResult: []i{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graphTest := NewGraph(false)

			for i := range tt.vertices {
				err := graphTest.AddVertex(tt.vertices[i])
				assert.Equal(t, tt.wantError[i], err)
			}

			assert.Equal(t, tt.wantResult, graphTest.GetVertices())
		})
	}
}

func TestAddEdge(t *testing.T) {
	//errAddEdge := errors.New("fail to get the element from dictionary")

	var tests = []struct {
		name              string
		isDirectedGraph   bool
		adjacencies       []*valuePair
		wantErrorsAddEdge []error
		wantErrorsGetAdj  []error
		wantResult        map[i]*[]dictionary.I
	}{
		{
			name:            "insert with success: vertex and edge as string",
			isDirectedGraph: false,
			adjacencies: []*valuePair{
				{Vertex: "A", Edge: "B"},
				{Vertex: "A", Edge: "C"},
				{Vertex: "A", Edge: "D"},
				{Vertex: "C", Edge: "D"},
			},
			wantErrorsAddEdge: []error{nil, nil, nil, nil},
			wantErrorsGetAdj:  []error{nil, nil, nil, nil},
			wantResult: map[i]*[]dictionary.I{
				"A": {"B", "C", "D"},
				"B": {"A"},
				"C": {"A", "D"},
				"D": {"A", "C"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graphTest := makeGraphWithVertices(tt.isDirectedGraph)

			for i, valuePair := range tt.adjacencies {
				err := graphTest.AddEdge(valuePair.Vertex, valuePair.Edge)

				assert.Equal(t, tt.wantErrorsAddEdge[i], err)
			}

			keys := make([]i, 0, len(tt.wantResult))
			for k := range tt.wantResult {
				keys = append(keys, k)
			}

			for i, k := range keys {
				adjList, err := graphTest.GetAdjList(k)

				assert.Equal(t, tt.wantErrorsGetAdj[i], err)
				assert.Equal(t, *tt.wantResult[k], adjList)
			}
		})
	}
}

func makeGraphWithVertices(isDirected bool) Graph {
	graph := NewGraph(isDirected)
	vertices := [4]string{"A", "B", "C", "D"}

	for _, v := range vertices {
		graph.AddVertex(v)
	}

	return graph
}
