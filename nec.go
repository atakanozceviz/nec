package nec

import (
	"fmt"
	"os"
	"path"
	"sync"
)

func Run(settingsPath, lastCommit, command string, walkpath ...string) error {
	settings, err := loadSettings(settingsPath)
	if err != nil {
		return fmt.Errorf("cannot load settings: %s", err)
	}

	if err := findPaths(settings, lastCommit, walkpath...); err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	jobs := make(chan *Job)

	for i := 0; i < settings.Commands[command].Workers; i++ {
		wg.Add(1)
		go worker(jobs, wg)
	}

	rootPath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get working directory: %s", err)
	}

	fmt.Println("Running command:")
	prettyPrint(settings.Commands[command])
	fmt.Println("Paths:")
	prettyPrint(settings.Paths[command])

	for _, dir := range settings.Paths[command] {
		jobs <- &Job{
			Path:    path.Join(rootPath, dir),
			OnErr:   settings.Commands[command].OnErr,
			Command: settings.Commands[command],
		}
	}
	close(jobs)

	wg.Wait()
	return nil
}

func ParseCommands(settingsPath string) (map[string]*Command, error) {
	settings, err := loadSettings(settingsPath)
	if err != nil {
		return nil, fmt.Errorf("cannot load settings: %s", err)
	}
	return settings.Commands, nil
}
