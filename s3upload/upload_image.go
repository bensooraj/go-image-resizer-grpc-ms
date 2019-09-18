package s3upload

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsSession session.Session
var awsS3Service *s3.S3

func init() {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	accessKeyID, _ := os.LookupEnv("AWS_S3_USER_ACCESS_KEY_ID")
	secretAccessKey, _ := os.LookupEnv("AWS_S3_USER_SECRET_ACCESS_KEY")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"",
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	awsS3Service = s3.New(awsSession)
}

// UploadImageToS3 ...
func UploadImageToS3(imageID, filename, filepath string) {

	// open the file for use
	imageFile, err := os.Open(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer imageFile.Close()

	// get the file size and read
	// the file content into a buffer
	imageFileInfo, _ := imageFile.Stat()
	var imageFileSize = imageFileInfo.Size()
	imageBuffer := make([]byte, imageFileSize)
	imageFile.Read(imageBuffer)

	bucketName, _ := os.LookupEnv("AWS_S3_BUCKET_NAME")

	putObjectInput := &s3.PutObjectInput{
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(imageBuffer),
		Bucket:        aws.String(bucketName),
		Key:           aws.String(imageID + "/" + filename),
		ContentLength: aws.Int64(imageFileSize),
		ContentType:   aws.String(http.DetectContentType(imageBuffer)),
	}

	result, err := awsS3Service.PutObject(putObjectInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}
