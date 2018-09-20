package nec

import (
	"fmt"
	"os"
	"sync"
)

func worker(job <-chan *Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range job {

		if err := doJob(j); err != nil {
			fmt.Printf("Worker error\nJob path: %s\nCmd returns: %s\n", j.Path, err)
			fmt.Println("Command:")
			if err := prettyPrint(j.Command); err != nil {
				fmt.Println(err)
			}
			if j.OnErr == "exit" {
				os.Exit(1)
			}
		}
	}
}
