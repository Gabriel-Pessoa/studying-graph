package graph

import "github.com/Gabriel-Pessoa/studying-graph/data-structures/dictionary"

type Graph interface {
}

type graph struct {
	isDirected bool
	vertices   []interface{} // string or number
	adjList    []dictionary.Dictionary
}

func NewGraph(isDirected bool)
