package minio

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"time"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// NewMinioClient - returns new minio client
func NewMinioClient(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool) *minio.Client {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}

// UploadFileToMinio - uploads file to minio
func UploadFileToMinio(client *minio.Client, bucketName string, objectName string, reader io.Reader) (minio.UploadInfo, error) {
	uploadInfo, err := client.PutObject(context.Background(), bucketName, objectName, reader, -1, minio.PutObjectOptions{ContentType: "text/plain"})
	if err != nil {
		log.Fatalln(err)
		return uploadInfo, err
	}
	return uploadInfo, nil
}

// GetObjectFromMinio - get file from minio
func GetObjectFromMinio(client *minio.Client, bucketName string, objectName string) (*minio.Object, error) {
	object, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return object, nil
	}
	return object, nil
}

// DeleteObjectInMinio - delete object in minio
func DeleteObjectInMinio(client *minio.Client, bucketName string, objectName string) error {
	err := client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	return err
}

// GetPresignedURLForObject - Get signed url for object
func GetPresignedURLForObject(client *minio.Client, bucketName string, objectName string, expiresIn time.Duration) *url.URL {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	// reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
	// Generates a presigned url which expires in a day.
	presignedURL, err := client.PresignedGetObject(context.Background(), bucketName, objectName, expiresIn, reqParams)
	if err != nil {
		fmt.Println(err)
		return presignedURL
	}
	return presignedURL
}

//Check if a bucket already exists
func CheckIfBucketExists(minioClient *minio.Client, bucketName string) (bool, error) {

	// Check to see if we already own this bucket
	exists, errBucketExists := minioClient.BucketExists(bucketName)

	if errBucketExists != nil {
		log.Fatalln(errBucketExists)
	}

	return exists, errBucketExists
}

//Create new bucket
func CreateNewBucket(minioClient *minio.Client, bucketName string, location string) error {

	err := minioClient.MakeBucket(bucketName, location)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}