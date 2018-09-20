package nec

import (
	"os"
	"os/exec"
)

func doJob(job *Job) error {
	cmd := exec.Command(job.Name, job.Args...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Dir = job.Path

	return cmd.Run()
}
