package studio

import (
	"context"
	"log"

	proto_uploader "github.com/droplez/droplez-go-proto/pkg/uploader"
	"google.golang.org/grpc"
)

func StudioClient() (err error) {
	var opt []grpc.DialOption
	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
	if err != nil {
		return nil
	}
	defer conn.Close()
	client := proto_uploader.NewUploaderClient(conn)
	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
}
