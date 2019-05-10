package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nec",
	Short: "Build and test necessary VS projects.",
	Long: `Nec helps you to speed up your CI for 
Visual Studio projects by looking up 
changes (using git diff) and finds out 
which solutions needs to build and tests 
needs to run. After that executes user-defined 
commands for the test projects and solutions.

Nec parses all the solutions (.sln) and 
projects (.csproj) in a folder and creates 
dependency graph, then uses that graph for 
finding dependencies.`,
	//Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&CfgFile, "settings", "s", "", "settings file (default is nec.json)")
	rootCmd.PersistentFlags().StringP("commit", "c", "HEAD^", "git commit id to find affected projects")
	rootCmd.PersistentFlags().StringP("walk-path", "w", ".", "the path to start the search for .sln files")
	rootCmd.PersistentFlags().StringP("ignore", "i", "", "ignore list file")
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		er(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Search config in home directory with name ".nec" (without extension).
		viper.AddConfigPath("./config")
		viper.SetConfigName("nec")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		er(err)
	}
}
