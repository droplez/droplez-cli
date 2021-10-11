package studio

import (
	"github.com/droplez/droplez-go-proto/pkg/studio/projects"
	"github.com/droplez/droplez-go-proto/pkg/studio/versions"
	"google.golang.org/grpc"
)

func ProjectsClient() (client projects.ProjectsClient, err error) {
	// var opt []grpc.DialOption
	conn, err := grpc.Dial("0.0.0.0:9090", grpc.WithInsecure())
	if err != nil {
		return
	}
	// defer conn.Close()
	client = projects.NewProjectsClient(conn)
	return
}

func VersionsClient() (client versions.VersionsClient, err error) {
	// var opt []grpc.DialOption
	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	client = versions.NewVersionsClient(conn)
	return
}
