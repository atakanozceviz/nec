package sln

import "github.com/atakanozceviz/nec/pkg/csproj"

type Sln struct {
	Name     string
	FilePath string
	Projects []*csproj.Csproj
}
