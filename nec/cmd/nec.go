package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/atakanozceviz/nec"
	"github.com/spf13/cobra"
)

var settingsPath string
var commitID string
var walkpath string

func init() {
	flag.StringVar(&settingsPath, "s", "nec.json", "Path to settings file.")
	flag.StringVar(&commitID, "c", "HEAD^", "Git commit id for getting changes.")
	flag.StringVar(&walkpath, "w", ".", "Path to search for sln files.")
	flag.Parse()

	for _, cmd := range commands(settingsPath) {
		cmd.Run = func(cmd *cobra.Command, args []string) {
			err := nec.Run(settingsPath, commitID, cmd.Use, walkpath)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		}

		rootCmd.AddCommand(cmd)
	}

}

func commands(settingsPath string) []*cobra.Command {
	var cmds []*cobra.Command

	commands, err := nec.ParseCommands(settingsPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for commandName, command := range commands {
		var cmd = &cobra.Command{
			Use:   commandName,
			Short: command.Description,
		}
		cmds = append(cmds, cmd)
	}

	return cmds
}
