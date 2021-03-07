package main

import (
	"log"
	"os"

	minio "github.com/k8-proxy/k8-go-comm/pkg/minio"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	endpoint := "localhost:9000"
	accessKey := "minioadmin"
	secretKey := "minioadmin"
	bucketName := "test"
	//PATHPDF := "./sampledata/file.pdf"
	//PATHZIP := "./sampledata/file.zip"
	fileName := "Screenshot"
	//ZIPNAME := "file.zip"

	client, err := minio.NewMinioClient(endpoint, accessKey, secretKey, false)
	if err != nil {
		log.Println(err)
		return
	}
	fileContent, err := os.Open("/home/ibrahim/Screenshot")
	if err != nil {
		log.Println(err)

	}
	uploadInfo, err := minio.UploadFileToMinio(client, bucketName, fileName, fileContent)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully uploaded bytes: ", uploadInfo)
	return
}
