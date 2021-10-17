package studio

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	grpc_uploader "github.com/droplez/droplez-cli/pkg/grpc/uploader"
	"github.com/droplez/droplez-cli/pkg/tools/archiver"
	"github.com/droplez/droplez-cli/pkg/tools/config"
	proto_projects "github.com/droplez/droplez-go-proto/pkg/studio/projects"
	proto_versions "github.com/droplez/droplez-go-proto/pkg/studio/versions"
	proto_uploader "github.com/droplez/droplez-go-proto/pkg/uploader"

	"github.com/google/uuid"
)

/* Init: start using droplez in the already existing project
*
* TODO: write better description
 */
func Init(project *proto_projects.ProjectInfo, path string) {
	// prepare project information
	project.Id = &proto_projects.ProjectId{
		Id: uuid.New().String(),
	}

	// project.Daw = daw.GetProjectDAW(path)
	// if len(project.Metadata.GetName()) == 0 {
		// project.Metadata.Name = filepath.Base(filepath.Dir(path))
	// }

	// create config struct
	conf := &config.Config{
		Project: project,
	}

	// Create a hidden dir
	err := os.Mkdir(fmt.Sprintf("%s/.droplez", path), 0755)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create a yaml file with project properties
	if err = config.WriteConfig(path, conf); err != nil {
		return
	}
}

func Version(path string, message string) (err error) {
	// Try to open config file
	conf, err := config.ReadConfig(path)
	if err != nil {
		return
	}
	fmt.Println("ARCHIVING")
	archive := archiver.Archive(path)
	// Upload a new version

	client, err := grpc_uploader.UploaderClient()
	if err != nil {
		return err
	}

	fmt.Println("INIT GRPC")
	stream, err := client.Upload(context.Background())
	if err != nil {
		return err
	}

	arch, err := os.Open(archive)
	if err != nil {
		return err
	}

	fileSize, err := arch.Stat()
	if err != nil {
		return err
	}

	// First sending metadata without content
	stream.Send(&proto_uploader.Chunk{
		Content: nil,
		FileMetadata: &proto_uploader.Metadata{
			// Name:        conf.Project.GetObjectName(),
			ContentType: proto_uploader.Metadata_CONTENT_TYPE_ARCHIVE,
			LocalName:   archive,
			FileSize:    fileSize.Size(),
			UserId:      "user id",
		},
	})

	r := bufio.NewReader(arch)
	for {
		bytes, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		stream.Send(&proto_uploader.Chunk{
			Content: bytes,
		})
	}

	// data, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	version := &proto_versions.VersionInfo{
		Id: &proto_versions.VersionId{
			Id: uuid.New().String(),
		},
		Metadata: &proto_versions.VersionMeta{
			Version: func(version []*proto_versions.VersionInfo) int32 {
				if len(version) > 0 {
					return int32(len(version) + 1)
				}
				return 1
			}(conf.Versions),
			Message:    message,
			// VersionId:  data.Version,
			UploadedAt: time.Now().Format("01-02-2006 15:04:05"),
		},
	}
	conf.Versions = append(conf.Versions, version)
	// conf.Project.ObjectName = data.GetObject()
	err = config.WriteConfig(path, conf)
	if err != nil {
		return
	}

	if err = os.Remove(archive); err != nil {
		return err
	}

	return
}

/*
* Create new droplez project:
* Choose a DAW or template
* Generate a yaml file
* Download a latest project version
*
* TODO:
 */
