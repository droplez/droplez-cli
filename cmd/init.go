package cmd

import (
	"fmt"

	"github.com/droplez/droplez-cli/pkg/service/studio"
	proto_projects "github.com/droplez/droplez-go-proto/pkg/studio/projects"
	"github.com/spf13/cobra"
)

var (
	pName        string
	pDescription string
	pPublic      bool
	pBpm         int32
	pKey         string
	pGenre       string
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "start using droplez with the project specified",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {

			var projectInfo = &proto_projects.ProjectInfo{
				Id: &proto_projects.ProjectId{
					Id: "",
				},
				Metadata: &proto_projects.ProjectMeta{
					Name:        pName,
					Description: pDescription,
					Public:      pPublic,
					Bpm:         pBpm,
					Key:         pKey,
					Genre:       pGenre,
				},
			}
			studio.Init(projectInfo, args[0])
		} else {
			fmt.Println("no path specified")
		}

	},
}

func init() {
	initCmd.Flags().StringVarP(&pName, "name", "n", "", "specify project name") //Default is the file name
	initCmd.Flags().StringVarP(&pDescription, "description", "d", "", "provice a project description")
	initCmd.Flags().BoolVarP(&pPublic, "public", "p", false, "choose true to let other users see the project")
	initCmd.Flags().Int32Var(&pBpm, "bpm", 120, "provide a project BPM")
	initCmd.Flags().StringVarP(&pKey, "key", "k", "C", "provice a project key")
	initCmd.Flags().StringVarP(&pGenre, "genre", "g", "", "provice a project genre")
}
