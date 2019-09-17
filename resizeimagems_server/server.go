package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/bensooraj/go-image-resizer-grpc-ms/resizeimagemspb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) ResizeImage(ctx context.Context, req *resizeimagemspb.ResizeImageRequest) (*resizeimagemspb.ResizeImageResponse, error) {
	// Get the GITHUB_USERNAME environment variable
	imageUploadDirName, _ := os.LookupEnv("IMAGE_UPLOAD_DIRECTORY")

	fmt.Printf("Image resize request received: %v\n", req)
	imageFileExtension := filepath.Ext(req.ImageFilename)

	img, err := imgio.Open(imageUploadDirName + req.ImageFilename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	imageSize := img.Bounds().Size()

	fmt.Printf("X: %v\n", imageSize.X)
	fmt.Printf("Y: %v\n", imageSize.Y)

	var imageEncodingType imgio.Format
	switch imageFileExtension {
	case ".jpg":
	case ".jpeg":
		imageEncodingType = imgio.JPEG
		break
	case ".png":
		imageEncodingType = imgio.PNG
		break
	}

	resized75 := transform.Resize(img, int(0.75*float64(imageSize.X)), int(0.75*float64(imageSize.Y)), transform.Linear)
	if err := imgio.Save(imageUploadDirName+req.ImageId+"_medium"+imageFileExtension, resized75, imageEncodingType); err != nil {
		fmt.Println(err)
	}

	resized50 := transform.Resize(img, int(0.50*float64(imageSize.X)), int(0.50*float64(imageSize.Y)), transform.Linear)
	if err := imgio.Save(imageUploadDirName+req.ImageId+"_small"+imageFileExtension, resized50, imageEncodingType); err != nil {
		fmt.Println(err)
	}

	return &resizeimagemspb.ResizeImageResponse{
		ImagesResized: 1,
	}, nil
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
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
