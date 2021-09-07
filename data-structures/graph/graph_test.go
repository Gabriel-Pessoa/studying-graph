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

func TestAddVertexGraph(t *testing.T) {
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
			name:       "insert invalid: with nil",
			vertices:   []i{"A", nil, "C", nil, "D"},
			wantError:  []error{nil, errAddVertex, nil, errAddVertex, nil},
			wantResult: []i{"A", "C", "D"},
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

func TestAddVertexFromAddEdgeGraph(t *testing.T) {
	errAddVertex := errors.New("fail to add vertex in the graph")

	adjacencies := []*valuePair{
		{Vertex: "A", Edge: "B"},
		{Vertex: "A", Edge: "C"},
		{Vertex: "A", Edge: "D"},
		{Vertex: "C", Edge: "D"},
		{Vertex: "", Edge: nil},
		{Vertex: "   ", Edge: ""},
		{Vertex: nil, Edge: "   "},
		{Vertex: nil, Edge: "E"},
		{Vertex: "F", Edge: "   "},
		{Vertex: nil, Edge: "F"},
		{Vertex: "E", Edge: nil},
	}
	wantError := []error{nil, nil, nil, nil, errAddVertex, errAddVertex, errAddVertex, errAddVertex, errAddVertex, errAddVertex, errAddVertex}
	wantResult := []i{"A", "B", "C", "D", "F", "E"}

	graphTest := NewGraph(false)

	for i, valuePair := range adjacencies {
		err := graphTest.AddEdge(valuePair.Vertex, valuePair.Edge)
		assert.Equal(t, wantError[i], err)
	}

	assert.Equal(t, wantResult, graphTest.GetVertices())
}

func TestAddEdgeGraph(t *testing.T) {
	var tests = []struct {
		name            string
		isDirectedGraph bool
		adjacencies     []*valuePair
		wantResult      map[i]*[]dictionary.I
	}{
		{
			name:            "insert with success: graph is not directed",
			isDirectedGraph: false,
			adjacencies: []*valuePair{
				{Vertex: "A", Edge: "B"},
				{Vertex: "A", Edge: "C"},
				{Vertex: "A", Edge: "D"},
				{Vertex: "C", Edge: "D"},
			},
			wantResult: map[i]*[]dictionary.I{
				"A": {"B", "C", "D"},
				"B": {"A"},
				"C": {"A", "D"},
				"D": {"A", "C"},
			},
		},
		{
			name:            "insert with success: graph is directed",
			isDirectedGraph: true,
			adjacencies: []*valuePair{
				{Vertex: "A", Edge: "B"},
				{Vertex: "A", Edge: "C"},
				{Vertex: "A", Edge: "D"},
				{Vertex: "C", Edge: "D"},
			},
			wantResult: map[i]*[]dictionary.I{
				"A": {"B", "C", "D"},
				"B": {},
				"C": {"D"},
				"D": {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graphTest := makeGraphWithVertices(tt.isDirectedGraph)

			for _, valuePair := range tt.adjacencies {
				err := graphTest.AddEdge(valuePair.Vertex, valuePair.Edge)
				assert.Equal(t, nil, err)
			}

			keys := make([]i, 0, len(tt.wantResult))
			for k := range tt.wantResult {
				keys = append(keys, k)
			}

			for _, k := range keys {
				adjList, err := graphTest.GetAdjList(k)
				assert.Equal(t, nil, err)
				assert.Equal(t, *tt.wantResult[k], adjList)
			}
		})
	}
}

func TestToStringGraph(t *testing.T) {
	expectedGraphToString := "A -> B C D \nB -> A E F \nC -> A D G \nD -> A C E G H \nE -> D B I \nF -> B \nG -> C D \nH -> D \nI -> E \n"

	graphTest := NewGraph(false)
	myVertices := [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	for _, v := range myVertices {
		graphTest.AddVertex(v)
	}
	graphTest.AddEdge("A", "B")
	graphTest.AddEdge("A", "C")
	graphTest.AddEdge("A", "D")
	graphTest.AddEdge("C", "D")
	graphTest.AddEdge("D", "E")
	graphTest.AddEdge("C", "G")
	graphTest.AddEdge("D", "G")
	graphTest.AddEdge("D", "H")
	graphTest.AddEdge("B", "E")
	graphTest.AddEdge("B", "F")
	graphTest.AddEdge("E", "I")

	assert.Equal(t, expectedGraphToString, graphTest.ToString())
}

func makeGraphWithVertices(isDirected bool) Graph {
	graph := NewGraph(isDirected)
	vertices := [4]string{"A", "B", "C", "D"}

	for _, v := range vertices {
		graph.AddVertex(v)
	}

	return graph
}
