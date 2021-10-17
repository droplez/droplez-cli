package cmd

import (
	"os"

	studiocmd "github.com/droplez/droplez-cli/cmd/studio"

	"github.com/spf13/cobra"
)

// flags
var (
	fPath string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "droplez-cli",
	Short:   "just for help",
	Long:    `droplez-cli`,
	Version: "v0.1.0",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(setDefaultVars, initConfig)
	RootCmd.AddCommand(studiocmd.StudioCmd)

}

// set defaults values for viper variables