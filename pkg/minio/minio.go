package minio

import (
	"github.com/minio/minio-go"
	"github.com/tkanos/gonfig"
	"log"
	"net/url"
	"time"
)

//Minio configuration struct
type Configuration struct {
	ENDPOINT        string
	ACCESSKEYID     string
	SECRETACCESSKEY string
	USESSL          bool
}

var config = Configuration{}

//Map config values
func init() {
	gonfig.GetConf("config.minio.json", &config)
}

//Create new minio instance
func NewInstance() (*minio.Client, error) {

	Client, err := minio.New(config.ENDPOINT, config.ACCESSKEYID, config.SECRETACCESSKEY, config.USESSL)

	if err != nil {
		log.Fatalln(err)
	}

	return Client, err
}

//Upload file to mino
func UploadFile(minioClient *minio.Client, bucketName string, filePath string, objectName string, contentType string) (int64, error) {

	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		log.Fatalln(err)
	}

	return n, err
}

//Generate signed url for file
func GenerateSignedUrl(minioClient *minio.Client, bucketName string, objectName string) (*url.URL, error) {

	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename="+objectName+"")

	// Generates a pre-signed url which expires in a day.
	presignedUrl, err := minioClient.PresignedGetObject(bucketName, objectName, time.Second*24*60*60, reqParams)

	if err != nil {
		log.Fatalln(err)
	}

	return presignedUrl, err
}

//Check if a bucket already exists
func CheckIfBucketExists(minioClient *minio.Client, bucketName string) (bool, error) {

	// Check to see if we already own this bucket
	exists, errBucketExists := minioClient.BucketExists(bucketName)

	return exists, errBucketExists
}
