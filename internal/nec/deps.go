package nec

import (
	"fmt"
	"path/filepath"

	"github.com/atakanozceviz/nec/pkg/csproj"
	"github.com/atakanozceviz/nec/pkg/graph"
)

var g = graph.NewGraph()

func addToGraph(csp *csproj.Csproj) (*graph.Graph, error) {
	g.AddNode(csp.Name)
	for _, ig := range csp.ItemGroups {
		for _, pr := range ig.ProjectReferences {
			cspFilePath := filepath.Join(filepath.Dir(csp.FilePath), pr.Include)
			dep, err := csproj.Parse(cspFilePath)
			if err != nil {
				return nil, fmt.Errorf("cannot parse csproj file referenced in %s: %v", csp.FilePath, err)
			}

			if csp.FilePath == "" || dep.FilePath == "" {
				continue
			}
			g.AddEdge(csp.Name, dep.Name)
			_, _ = addToGraph(dep)
		}
	}
	return g, nil
}

var deps = make(map[string]struct{})

func DepsOf(id string, g *graph.Graph) map[string]struct{} {
	for u, m := range g.Edges {
		for v := range m {
			if v == id {
				deps[u] = struct{}{}
				DepsOf(u, g)
			}
		}
	}
	return deps
}
