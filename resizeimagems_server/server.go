package main

import (
	"context"
	"fmt"
	"image"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/bensooraj/go-image-resizer-grpc-ms/resizeimagemspb"
	"github.com/bensooraj/go-image-resizer-grpc-ms/s3upload"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var wg sync.WaitGroup
var mut sync.Mutex

type server struct {
}

func (*server) ResizeImage(ctx context.Context, req *resizeimagemspb.ResizeImageRequest) (*resizeimagemspb.ResizeImageResponse, error) {
	// Get the IMAGE_UPLOAD_DIRECTORY environment variable
	imageUploadDirName, _ := os.LookupEnv("IMAGE_UPLOAD_DIRECTORY")

	fmt.Printf("Image resize request received: %v\n", req)
	imageFileExtension := filepath.Ext(req.ImageFilename)

	img, err := imgio.Open(imageUploadDirName + req.ImageFilename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	scaleToPrefixMap := map[float64]string{
		0.75: "_medium",
		0.50: "_small",
	}
	for scale, prefix := range scaleToPrefixMap {
		fmt.Printf("scale: %v -> prefix: %v\n", scale, prefix)
		wg.Add(1)
		go resizeAndUpload(img, req.ImageId, imageFileExtension, scale, prefix)
	}

	wg.Wait()

	return &resizeimagemspb.ResizeImageResponse{
		ImagesResized: 2,
	}, nil
}

func resizeAndUpload(baseImage image.Image, imageID, imageFileExtension string, scale float64, prefix string) {
	defer wg.Done()
	// Get the IMAGE_UPLOAD_DIRECTORY environment variable
	imageUploadDirName, _ := os.LookupEnv("IMAGE_UPLOAD_DIRECTORY")

	imageSize := baseImage.Bounds().Size()

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

	mut.Lock()
	resized75 := transform.Resize(baseImage, int(scale*float64(imageSize.X)), int(scale*float64(imageSize.Y)), transform.Linear)
	mut.Unlock()
	if err := imgio.Save(imageUploadDirName+imageID+prefix+imageFileExtension, resized75, imageEncodingType); err != nil {
		fmt.Println(err)
	}
	s3upload.UploadImageToS3(imageID, imageID+prefix+imageFileExtension, imageUploadDirName+imageID+prefix+imageFileExtension)
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
