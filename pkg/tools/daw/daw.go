package daw

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/droplez/droplez-go-proto/pkg/studio/projects"
)

func isLogicPackage(projectPath string) bool {
	var pathToProjectInformation = fmt.Sprintf("%s/Resources/ProjectInformation.plist", projectPath)

	if strings.Contains(projectPath, ".logicx") {
		if _, err := os.Stat(pathToProjectInformation); os.IsNotExist(err) {
			return false
		}
		return true
	}

	return false
}

func isLogicFolder(files []fs.FileInfo) bool {
	for _, f := range files {
		if strings.Contains(f.Name(), ".logicx") {
			return true
		}
	}
	return false
}

// GetProjectDAW returns daw name if it is supported by droplez, otherwise it returns an error
func GetProjectDAW(projectPath string) projects.DAW {
	// Logic Pro
	if isLogicPackage(projectPath) {
		return projects.DAW_DAW_LOGIC_PRO
	}
	files, err := ioutil.ReadDir(projectPath)
	if err != nil {
		log.Fatal(err)
	}
	if isLogicFolder(files) {
		return projects.DAW_DAW_LOGIC_PRO
	}
	return projects.DAW_DAW_UNSPECIFIED
}
