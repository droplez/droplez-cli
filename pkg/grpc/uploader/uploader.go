package uploader

import (
	proto_uploader "github.com/droplez/droplez-go-proto/pkg/uploader"
	"google.golang.org/grpc"
)

func UploaderClient() (client proto_uploader.UploaderClient, err error) {
	// var opt []grpc.DialOption
	conn, err := grpc.Dial("0.0.0.0:9090", grpc.WithInsecure())
	if err != nil {
		return
	}
	// defer conn.Close()
	client = proto_uploader.NewUploaderClient(conn)
	return
}
