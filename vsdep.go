package nec

import (
	"os"
	"path"
	"strings"

	"github.com/atakanozceviz/vsdep"
)

func findPaths(settings *Settings, lastcommit string, walkpath ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	projectFolderName := "/" + path.Base(wd) + "/"

	paths, err := vsdep.FindOut(lastcommit, walkpath...)
	if err != nil {
		return err
	}
	for _, pth := range paths.Tests {
		if settings.Paths == nil {
			settings.Paths = make(map[string][]string)
		}
		if len(walkpath) >= 1 && walkpath[0] != "." {
			if strings.Contains(path.Dir(pth), projectFolderName) {
				settings.Paths["test"] = append(settings.Paths["test"], path.Dir(pth))
			}
			continue
		}
		settings.Paths["test"] = append(settings.Paths["test"], path.Dir(pth))
	}

	for _, pth := range paths.Solutions {
		if settings.Paths == nil {
			settings.Paths = make(map[string][]string)
		}
		if len(walkpath) >= 1 && walkpath[0] != "." {
			if strings.Contains(path.Dir(pth), projectFolderName) {
				settings.Paths["build"] = append(settings.Paths["build"], path.Dir(pth))
			}
			continue
		}
		settings.Paths["build"] = append(settings.Paths["build"], path.Dir(pth))
	}

	return nil
}
