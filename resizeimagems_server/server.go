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

	mysqldb "github.com/bensooraj/go-image-resizer-grpc-ms/database"
	"github.com/bensooraj/go-image-resizer-grpc-ms/s3upload"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/bensooraj/go-image-resizer-grpc-ms/resizeimagemspb"
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

	fmt.Println("Done Waiting For The Image Resizing Operation.")

	return &resizeimagemspb.ResizeImageResponse{
		ImagesResized: 2,
	}, nil
}

func resizeAndUpload(baseImage image.Image, imageID, imageFileExtension string, scale float64, prefix string) {
	defer wg.Done()
	// Get the IMAGE_UPLOAD_DIRECTORY environment variable
	imageUploadDirName, _ := os.LookupEnv("IMAGE_UPLOAD_DIRECTORY")
	imageHostDomainName, _ := os.LookupEnv("IMAGE_HOST_DOMAIN_NAME")

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
	defer mut.Unlock()
	resized75 := transform.Resize(baseImage, int(scale*float64(imageSize.X)), int(scale*float64(imageSize.Y)), transform.Linear)
	if err := imgio.Save(imageUploadDirName+imageID+prefix+imageFileExtension, resized75, imageEncodingType); err != nil {
		fmt.Println(err)
	}
	s3upload.UploadImageToS3(imageID, imageID+prefix+imageFileExtension, imageUploadDirName+imageID+prefix+imageFileExtension)

	// Upload to MySQL
	statement, err := mysqldb.Db.Prepare("INSERT INTO images(image_id, scale, image_url) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatalf("Error preparing the DB query: %v", err)
	}

	imageURL := fmt.Sprintf("%s/%s/%s", imageHostDomainName, imageID, imageID+prefix+imageFileExtension)
	result, err := statement.Exec(imageID, scale, imageURL)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n[%s] Last ID inserted: %d\nRows affected: %d\n", imageID, lastID, rowCount)

}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	fmt.Println("Resize Image Golang GRPC Server")

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
