package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/atakanozceviz/nec/internal/config"
	"github.com/atakanozceviz/nec/internal/job"
	"github.com/atakanozceviz/nec/internal/nec"
	"github.com/panjf2000/ants"
	"github.com/spf13/viper"
)

func run(typ string) error {
	var c config.Config
	err := viper.Unmarshal(&c)
	if err != nil {
		er(err)
	}

	ignore := make([]string, 0)
	ignorePath := viper.GetString("ignore")
	if ignorePath != "" {
		ignore, err = parseIgnore(ignorePath)
		if err != nil {
			return err
		}
	}

	command := c.Commands[typ]

	paths, err := slnPaths(c.WalkPath)
	if err != nil {
		return err
	}

	affectedProjects, err := nec.GetAffectedProjects(c.Commit, ignore...)
	if err != nil {
		return err
	}

	sln, g, err := nec.ParseAndGenerate(paths...)
	if err != nil {
		return err
	}

	deps := make(map[string]struct{})
	for project := range affectedProjects {
		deps = nec.DepsOf(project, g)
	}

	testsNeedsToRun := make(map[string]string)
	solutionsNeedsToBuild := make(map[string]string)
	allProjects := make(map[string]map[string]string)

	for projectName := range deps {
		for _, solution := range sln {
			for _, project := range solution.Projects {
				if projectName == project.Name {
					if project.IsTest() {
						testsNeedsToRun[project.Name] = project.FilePath
						continue
					}
					solutionsNeedsToBuild[solution.Name] = solution.FilePath
				}
			}
		}
	}
	allProjects["build"] = solutionsNeedsToBuild
	allProjects["test"] = testsNeedsToRun

	wg := sync.WaitGroup{}
	defer wg.Wait()
	defer ants.Release()
	pool, _ := ants.NewPoolWithFunc(command.Workers, func(i interface{}) {
		defer wg.Done()
		j, ok := i.(*job.Job)
		if !ok {
			return
		}
		err := j.ExecJob()
		if err != nil {
			fmt.Printf("Error while working on %s\nCommand Description: %s\nMessage: %v\n", j.JobName, j.Description, err)
			if j.Onerror == "exit" {
				os.Exit(1)
			}
		}
	})

	fmt.Println("Running command:")
	prettyPrint(command)
	fmt.Println("Paths:")
	prettyPrint(allProjects[typ])

	for name, path := range allProjects[typ] {
		wg.Add(1)
		err := pool.Invoke(&job.Job{
			JobName: name,
			Path:    filepath.Dir(path),
			Command: command,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func parseIgnore(p string) ([]string, error) {
	v := viper.New()
	v.SetConfigFile(p)
	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	parsed := make(map[string]interface{})
	err := v.Unmarshal(&parsed)
	if err != nil {
		er(err)
	}
	ignore := make([]string, 0, len(parsed))
	for p := range parsed {
		ignore = append(ignore, p)
	}

	return ignore, nil
}
