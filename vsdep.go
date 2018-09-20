package nec

import (
	"path"

	"github.com/atakanozceviz/vsdep"
)

func findPaths(settings *Settings, lastcommit string) error {
	paths, err := vsdep.FindOut(lastcommit)
	if err != nil {
		return err
	}
	for _, pth := range paths.Tests {
		if settings.Paths == nil {
			settings.Paths = make(map[string][]string)
		}
		settings.Paths["test"] = append(settings.Paths["test"], path.Dir(pth))
	}

	for _, pth := range paths.Solutions {
		if settings.Paths == nil {
			settings.Paths = make(map[string][]string)
		}
		settings.Paths["build"] = append(settings.Paths["build"], path.Dir(pth))
	}

	return nil
}
