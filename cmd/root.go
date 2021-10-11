package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/droplez/droplez-cli/pkg/service/studio"

	"github.com/spf13/cobra"
)

var cfgFile string

// flags
var (
	fPath string
	fMessage string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "help",
	Short: "just for help",
	Long:  `droplez-cli`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(setDefaultVars, initConfig)
	// register commands 
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(versionCmd)
	// register flags
	currentPath, err := os.Getwd()
	if err != nil {
	  log.Fatal(err)
	}
	RootCmd.PersistentFlags().StringVarP(&fPath, "path", "P", currentPath, "provide a path to the project. If empty, the current dir will be set")
	
	versionCmd.Flags().StringVarP(&fMessage, "message", "m", "", "Provide a message that would describe changes you've made in this version")
	versionCmd.MarkFlagRequired("message")

}

// set defaults values for viper variables
func setDefaultVars() {
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Create a version with current changes",
	Run: func(cmd *cobra.Command, args []string) {
		if err := studio.Version(fPath, fMessage); err != nil {
			fmt.Println(err)
		}
	},
	
	Version: "0.0.1",
	Example: "droplez version -P ~/Music/Projects/MyBestProject.logicx -m 'Added the coolest drum kit ever'",
}
