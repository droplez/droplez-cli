package studiocmd

import (
	"github.com/spf13/cobra"
)

var StudioCmd = &cobra.Command{
	Use:     "studio",
	Aliases: []string{"s", "st"},
	Short:   "Interact with droplez studio service",
	//TODO: Provide a normal description for droplez-studio
	Long: `

	`,
}

func init() {
	StudioCmd.AddCommand(projectsCmd)
}
