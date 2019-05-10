package nec

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetAffectedExt returns list of affected projects since "lastcommit" (using git diff)
// Specified folders will be ignored
func GetAffectedProjects(lastcommit string, ignore ...string) (map[string]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dir := ""
	if list := strings.Split(lastcommit, "/"); len(list) > 1 {
		dir = strings.Join(list[:len(list)-1], "/")
		lastcommit = list[len(list)-1]
	}
	wd = filepath.Join(wd, dir)

	cmd := exec.Command("git", "diff", "--name-only", lastcommit, "HEAD")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Dir = wd

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	dirs := make(map[string]struct{})

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		dirs[filepath.Dir(scanner.Text())] = struct{}{}
	}

	projects := make(map[string]string)

dirs:
	for p := range dirs {
		for p != "." {
			files, err := ioutil.ReadDir(filepath.Join(wd, p))
			if err != nil && os.IsExist(err) {
				return nil, err
			}

			for _, f := range files {
				if filepath.Ext(f.Name()) == ".csproj" && !isInside(p, ignore) {
					projects[f.Name()] = filepath.Join(p, f.Name())
					continue dirs
				}
			}
			p = filepath.Clean(p + "/..")
		}
	}
	return projects, nil
}

func isInside(p string, sx []string) bool {
	p = strings.ReplaceAll(p, "\\", "/")
	for _, s := range sx {
		if strings.HasPrefix(p, s) {
			return true
		}
		continue
	}
	return false
}
