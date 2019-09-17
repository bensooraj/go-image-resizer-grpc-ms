package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bensooraj/go-image-resizer-grpc-ms/resizeimagemspb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) ResizeImage(context.Context, *resizeimagemspb.ResizeImageRequest) (*resizeimagemspb.ResizeImageResponse, error) {
	return nil, nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}

	s := grpc.NewServer()
	resizeimagemspb.RegisterResizeImageMicroServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
