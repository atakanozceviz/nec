package job

import (
	"os"
	"os/exec"

	"github.com/atakanozceviz/nec/internal/config"
)

type Job struct {
	*config.Command
	Path    string
	JobName string
}

func (j *Job) ExecJob() error {
	cmd := exec.Command(j.Name, j.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Dir = j.Path
	return cmd.Run()
}
