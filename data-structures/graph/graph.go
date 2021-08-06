package graph

import (
	"errors"

	"github.com/Gabriel-Pessoa/studying-graph/data-structures/dictionary"
)

type i interface{}

type Graph interface {
	AddVertex(v i) error
	AddEdge(v, w i) error
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
	if !g.includes(v) {
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

	newAdjListV := adjListV
	newAdjListV = append(newAdjListV, w)
	g.AdjList.Set(v, newAdjListV)

	if !g.IsDirected {
		adjListW, err := g.AdjList.Get(w)
		if err != nil {
			return err
		}

		newAdjListW := adjListW
		newAdjListW = append(newAdjListW, v)
		g.AdjList.Set(w, newAdjListW)
	}

	return nil
}

func (g graph) includes(searchElement i) bool {
	for _, value := range g.Vertices {
		if value == searchElement {
			return true
		}
	}

	return false
}
