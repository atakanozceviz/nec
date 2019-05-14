package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func prettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Printf("%s\n", b)
	}
}

func slnPaths(dir ...string) ([]string, error) {
	wp := "."
	if len(dir) == 1 {
		wp = dir[0]
	}
	var paths []string
	if err := filepath.Walk(wp, func(pth string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ".sln" {
			paths = append(paths, pth)
		}
		return err
	}); err != nil {
		return nil, err
	}
	return paths, nil
}
