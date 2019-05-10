package cmd

import (
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Run build command.",
	Long:  `Run build command specified in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := run("build")
		if err != nil {
			er(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

}
