package graph

import (
	"errors"
	"fmt"

	"github.com/Gabriel-Pessoa/studying-graph/data-structures/dictionary"
	"github.com/Gabriel-Pessoa/studying-graph/utils"
)

type i interface{}

type Graph interface {
	AddVertex(v i) error
	AddEdge(v, w i) error
	GetVertices() []i
	GetAdjList(v i) ([]dictionary.I, error)
	ToString() string
}

type graph struct {
	IsDirected bool
	Vertices   []i // slice of string or number
	AdjList    dictionary.Dictionary
}

func NewGraph(isDirected bool) Graph {
	return &graph{
		IsDirected: isDirected, // default false
		Vertices:   []i{},
		AdjList:    dictionary.NewDictionary(),
	}
}

func (g *graph) AddVertex(v i) error {
	if !g.includes(v) && !utils.IsEmpty(v) {
		g.Vertices = append(g.Vertices, v)
		g.AdjList.Set(v, []dictionary.I{})

		return nil
	}

	return errors.New("fail to add vertex in the graph")
}

func (g *graph) AddEdge(v, w i) error {
	if !g.AdjList.HasKey(v) {
		g.AddVertex(v)
	}

	if !g.AdjList.HasKey(w) {
		g.AddVertex(w)
	}

	adjListV, err := g.AdjList.Get(v)
	if err != nil {
		return err
	}

	adjListV = append(adjListV, w)
	g.AdjList.Set(v, adjListV)

	if !g.IsDirected {
		adjListW, err := g.AdjList.Get(w)
		if err != nil {
			return err
		}

		adjListW = append(adjListW, v)
		g.AdjList.Set(w, adjListW)
	}

	return nil
}

func (g graph) GetVertices() []i {
	return g.Vertices
}

func (g graph) GetAdjList(v i) ([]dictionary.I, error) {
	return g.AdjList.Get(v)
}

func (g graph) ToString() string {
	var s string

	for i := 0; i < len(g.Vertices); i++ {
		s += fmt.Sprintf("%v -> ", g.Vertices[i])

		neighbors, err := g.AdjList.Get(g.Vertices[i])
		if err != nil {
			panic("fail to get adjacencies")
		}

		for j := 0; j < len(neighbors); j++ {
			s += fmt.Sprintf("%v ", neighbors[j])
		}

		s += "\n"
	}

	return s
}

func (g graph) includes(searchElement i) bool {
	for _, value := range g.Vertices {
		if value == searchElement {
			return true
		}
	}
	return false
}
