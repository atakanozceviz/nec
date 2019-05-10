package sln

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/atakanozceviz/nec/pkg/csproj"
)

var projectPattern = regexp.MustCompile(`(?m)Project\("(.*?)"\)\s+=\s+"(.*?)",\s+"(.*?)",\s+"(.*?)"`)

// Parse sln file and return projects.
func Parse(p string) (*Sln, error) {
	p = strings.ReplaceAll(p, "\\", "/")
	slnData, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	var sln Sln
	for _, v := range projectPattern.FindAllStringSubmatch(string(slnData), -1) {
		if len(v) >= 4 {
			if !strings.Contains(v[3], ".csproj") {
				continue
			}
			csprojPath := filepath.Join(filepath.Dir(p), v[3])

			project, err := csproj.Parse(csprojPath)
			if err != nil {
				err = fmt.Errorf("cannot parse csproj file referenced in %s: %v", p, err)
				return nil, err
			}
			sln.Projects = append(sln.Projects, project)
		}
	}
	sln.FilePath = p
	sln.Name = filepath.Base(p)
	return &sln, nil
}
