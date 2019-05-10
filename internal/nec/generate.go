package nec

import (
	"sync"

	"github.com/atakanozceviz/nec/pkg/csproj"
	"github.com/atakanozceviz/nec/pkg/graph"
	"github.com/atakanozceviz/nec/pkg/sln"
	"github.com/panjf2000/ants"
)

func ParseAndGenerate(paths ...string) ([]*sln.Sln, *graph.Graph, error) {
	allSln := make([]*sln.Sln, 0)
	g := &graph.Graph{}
	wg := sync.WaitGroup{}
	defer wg.Wait()

	pool, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		defer wg.Done()
		var err error
		g, err = addToGraph(i.(*csproj.Csproj))
		if err != nil {
			panic(err)
		}
	})

	for _, path := range paths {
		s, err := sln.Parse(path)
		if err != nil {
			return nil, nil, err
		}
		for _, p := range s.Projects {
			wg.Add(1)
			_ = pool.Invoke(p)
		}
		allSln = append(allSln, s)
	}
	return allSln, g, nil
}
