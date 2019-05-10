package cmd

import (
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run test command.",
	Long:  `Run test command specified in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := run("test")
		if err != nil {
			er(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
