package graph

import "sync"

// Graph base struct
type Graph struct {
	sync.RWMutex
	Nodes map[string]struct{}
	Edges map[string]map[string]struct{}
}

// NewGraph Create graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]struct{}),
		Edges: make(map[string]map[string]struct{}),
	}
}

// AddNode Add node id to graph, return true if added (string's are unique).
func (g *Graph) AddNode(id string) bool {
	g.Lock()
	defer g.Unlock()
	if _, ok := g.Nodes[id]; ok {
		return false
	}
	g.Nodes[id] = struct{}{}
	return true
}

// AddEdge Add an edge from u to v.
func (g *Graph) AddEdge(u, v string) {
	g.Lock()
	defer g.Unlock()
	if _, ok := g.Nodes[u]; !ok {
		g.Unlock()
		g.AddNode(u)
		g.Lock()
	}
	if _, ok := g.Nodes[v]; !ok {
		g.Unlock()
		g.AddNode(v)
		g.Lock()
	}

	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = make(map[string]struct{})
	}
	g.Edges[u][v] = struct{}{}
}
