package csproj

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Parse(p string) (*Csproj, error) {
	p = strings.ReplaceAll(p, "\\", "/")
	csproj := &Csproj{}
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(data, csproj)
	if err != nil {
		return nil, err
	}
	csproj.FilePath = p
	csproj.Name = filepath.Base(p)
	return csproj, nil
}
