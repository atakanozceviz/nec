package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

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
finding dependencies.
`,
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nec.yaml)")
	rootCmd.PersistentFlags().StringVarP(&settingsPath, "s", "s", "nec.json", "Path to settings file.")
	rootCmd.PersistentFlags().StringVarP(&commitID, "c", "c", "HEAD^", "Git commit id for getting changes.")
	rootCmd.PersistentFlags().StringVarP(&walkpath, "w", "w", ".", "The path to start the search for .sln files.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".nec" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".nec")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
