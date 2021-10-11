package archiver

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

/* Archive creates a zip archive with folders
* TODO: Refactor and better description @allanger
 */
func Archive(path string) (archive string) {
	archive = fmt.Sprintf("%s/.droplez/%s.zip", path, uuid.New().String())
	file, err := os.Create(archive)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(fullPath string, info os.FileInfo, err error) error {
		//ignore .droplez dir
		zipPath := strings.Replace(fullPath, path, "", -1)
		if strings.HasPrefix(zipPath, "/.droplez") {
			return nil
		}
		// ignore dirs
		if info.IsDir() {
			return nil
		}

		// handle files
		file, err := os.Open(fullPath)
		if err != nil {
			return err
		}
		defer file.Close()

		fmt.Println(zipPath)
		f, err := w.Create(zipPath)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}
		return nil
	}

	err = filepath.Walk(path, walker)
	if err != nil {
		panic(err)
	}

	return archive
}
