package studiocmd

import (
	"fmt"
	"log"
	"os"

	"github.com/droplez/droplez-cli/pkg/cli/studio/projects"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"p", "project"},
	Short:   "Manage Droplez projects",
	//TODO: Provide a normal description for droplez-studio
	Long: `

	`,
}

// For flags
var (
	// Common
	fPath string
	// Create
	fDaw         string
	fTemplate    string
	fDescription string
	fGenre       string
	fKey         string
	fEmpty       bool
	fMetadata    bool
	fPublic      bool
	fBpm         int
	fMessage     string
	fTemplated   bool
	fDefault     bool
)

func init() {
	// Helpers
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Register commands
	projectsCmd.AddCommand(createCmd)
	projectsCmd.AddCommand(initCmd)
	projectsCmd.AddCommand(updateCmd)
	projectsCmd.AddCommand(getCmd)
	projectsCmd.AddCommand(listCmd)
	projectsCmd.AddCommand(deleteCmd)

	// Register flags
	projectsCmd.PersistentFlags().StringVarP(&fPath, "path", "P", currentPath, "provide a path to the project")

	createCmd.Flags().StringVarP(&fDaw, "daw", "D", "", "choose a DAW")
	createCmd.Flags().StringVarP(&fTemplate, "template", "t", "", "choose a template")
	createCmd.Flags().StringVarP(&fDescription, "description", "d", "", "provide project description")
	createCmd.Flags().StringVar(&fKey, "key", "", "provide project key")
	createCmd.Flags().StringVar(&fGenre, "genre", "", "provide project genre")
	createCmd.Flags().StringVarP(&fMessage, "message", "m", "", "provide a message describing your version")

	createCmd.Flags().BoolVarP(&fEmpty, "empty", "e", false, "create an empty Droplez project")
	createCmd.Flags().BoolVarP(&fMetadata, "metadata", "M", false, "fill metadata interactively")
	createCmd.Flags().BoolVar(&fPublic, "public", false, "make project public")
	createCmd.Flags().BoolVar(&fTemplated, "templated", false, "create project from template")
	createCmd.Flags().BoolVar(&fDefault, "default", false, "create default DAW project")

	createCmd.Flags().IntVar(&fBpm, "bpm", 0, "choose project bpm")

}

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c", "cr"},
	Short:   "Create new Droplez project",
	Example: "$ droplez studio projects create my-first-project",
	Long: `🔥 Create new Droplez project

Here you can choose from options:

- Create a diretctory with a Droplez config 
	In this case your chages won't be synced with the server until you create and push the first version.

- Create an empty DAW project 
	After choosing a DAW, you will get a droplez config that will use on of our default empty templates as a first version.

- Create a project from the template 
	You can choose on of your personal or any public template to initizalize a project. the already existing one will be choosen for the first version.

The command is taking the first argument as a project name
`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			name string
		)
		if len(args) > 0 {
			name = args[0]
		} else {
			name = ""
		}

		if err := projects.Create(&projects.CreateOpts{
			Name:        name,
			Path:        fPath,
			DAW:         fDaw,
			Template:    fTemplate,
			Empty:       fEmpty,
			Metadata:    fMetadata,
			Description: fDescription,
			Bpm:         fBpm,
			Key:         fKey,
			Genre:       fGenre,
			Public:      fPublic,
			Message:     fMessage,
			Templated:   fTemplated,
			Default:     fDefault,
		}); err != nil {
			fmt.Println(err)
		}
	},
}

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Start using Droplez in already existing project",
	Example: "$ droplez studio projects init my-first-project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("droplez init project")
	},
}
var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "Update Droplez project",
	Example: "$ droplez studio projects update my-first-project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("droplez update project")
	},
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get Droplez project",
	Example: "$ droplez studio projects get my-first-project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("droplez get project ")
	},
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List Droplez projects",
	Example: "$ droplez studio projects list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("droplez list projects")
	},
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove"},
	Short:   "Delete Droplez project (can't be undone)",
	Example: "$ droplez studio projects delete ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("droplez delete project")
	},
}
