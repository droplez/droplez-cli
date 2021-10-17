package projects

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/droplez/droplez-cli/pkg/grpc/studio"
	"github.com/droplez/droplez-cli/tools/constants"
	"github.com/droplez/droplez-go-proto/pkg/studio/projects"
	"github.com/droplez/droplez-go-proto/pkg/studio/versions"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/manifoldco/promptui"
)

// Project types
const (
	EmptyProject        = "Empty Droplez project"
	EmptyDawProject     = "Default DAW project"
	TemplatedDawProject = "DAW project from template"
)

type CreateOpts struct {
	Name        string
	Path        string
	DAW         string
	Template    string
	Public      bool
	Empty       bool
	Metadata    bool
	Description string
	Key         string
	Genre       string
	Bpm         int
	Message     string
	Templated   bool
	Default     bool
}

// Errors
const (
	ErrEmptyName   = "name can not be empty"
	ErrNotValidBPM = "bpm must me a number"
)

// Create a Droplez project
func Create(opts *CreateOpts) (err error) {
	var (
		result string
		dawMap = make(map[string]string)
	)

	/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
	* 1. Set the name if not provided via args
	* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
	if len(opts.Name) == 0 {
		prompt := promptui.Prompt{
			Label: "Name your project",
			Validate: func(s string) error {
				if len(s) < 1 {
					return fmt.Errorf(ErrEmptyName)
				}
				return nil
			},
		}

		result, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		opts.Name = result
	}

	/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
	* 2. Check for empty project (CreateOpts.Empty)
	* -> If empty, skip choosing DAW and templates
	* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
	if !opts.Empty {
		if opts.Templated {
			result = TemplatedDawProject
		} else if opts.Default {
			result = EmptyDawProject
		} else {
			prompt := promptui.Select{
				Label: "What do you wanna create?",
				Items: []string{
					EmptyProject,
					EmptyDawProject,
					TemplatedDawProject,
				},
			}
			_, result, err = prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}
		switch result {
		case EmptyDawProject:
			opts.Template = ""
		case TemplatedDawProject:
			prompt := promptui.Prompt{
				Label: "Provide template ID",
				Validate: func(s string) (err error) {
					if err = validation.Validate(s, is.UUID); err != nil {
						return
					}
					// TODO: Add getting template from studio service and check that it exosts
					return
				},
			}

			result, err = prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			opts.Description = result

		default:
			opts.Empty = true
		}

		if result != EmptyProject {
			// Choose DAW
			var daws []string

			// List supported DAW
			for _, daw := range projects.DAW_name {
				if daw != projects.DAW_DAW_UNSPECIFIED.String() {
					dawBeatyfied := strings.Replace(daw, "DAW_", "", -1)
					dawBeatyfied = strings.Replace(dawBeatyfied, "_", " ", -1)
					dawBeatyfied = strings.Title(strings.ToLower(dawBeatyfied))
					daws = append(daws, dawBeatyfied)
					dawMap[dawBeatyfied] = daw
				}
			}

			if len(opts.DAW) < 1 {
				prompt := promptui.Select{
					Label: "What DAW are you using?",
					Items: daws,
				}
				_, result, err = prompt.Run()

				opts.DAW = result

				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}
			}
		}
	}
	/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
	* 2. Check for metadata input (CreateOpts.Metadata)
	* -> If false, skip this step
	* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

	if opts.Metadata {
		// Project description
		if len(opts.Description) < 1 {
			prompt := promptui.Prompt{
				Label: "Describe your project",
			}

			result, err = prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			opts.Description = result
		}

		// Project BPM
		if opts.Bpm == 0 {
			prompt := promptui.Prompt{
				Label: "Choose BPM",
				Validate: func(s string) error {
					if _, err := strconv.Atoi(s); err != nil {
						return fmt.Errorf(ErrNotValidBPM)
					}
					return nil
				},
				Default: "0",
			}

			result, err = prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			opts.Bpm, _ = strconv.Atoi(result)
		}

		//Project Key
		if len(opts.Key) < 1 {
			prompt := promptui.Prompt{
				Label: "Provide project key",
			}

			result, err = prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			opts.Key = result
		}

		// Project Genre
		if len(opts.Genre) < 1 {
			prompt := promptui.Prompt{
				Label:   "Specify genre",
				Default: "none",
			}

			result, err = prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			opts.Genre = result
		}
	}

	// Create a project
	projectsClient, err := studio.ProjectsClient()
	if err != nil {
		return err
	}

	projectInfo := &projects.ProjectMeta{
		Daw:         projects.DAW(projects.DAW_value[dawMap[opts.DAW]]),
		Name:        opts.Name,
		Description: opts.Description,
		Bpm:         int32(opts.Bpm),
		Key:         opts.Key,
		Genre:       opts.Genre,
		Public:      opts.Public,
	}

	project, err := projectsClient.Create(context.Background(), projectInfo)
	if err != nil {
		return
	}
	if !opts.Empty {
		if opts.Default {
			// Getdefault project id here 
		} else {
			// UUID is already defined
		}
		
		// GetVersion
		// Create a version
		versionsClient, err := studio.VersionsClient()
		if err != nil {
			return err
		}

		versionInfo := &versions.VersionMeta{
			Version: 1,
			Message: func(s string) string {
				if len(s) < 1 {
					return "first version"
				}
				return s
			}(opts.Message),
			UploadedAt: time.Now().Format(constants.TimeFormat),
		}
		versionsClient.Create(context.Background(), versionInfo)
	}

	fmt.Printf("Successfully created a new project! \nTo get information run '$ droplez studio projects get %s \n", project.GetId().GetId())
	return
}
