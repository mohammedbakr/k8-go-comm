package main

import (
	"bytes"
	"log"
	"os"

	minio "github.com/k8-proxy/k8-go-comm/pkg/minio"
)

func main() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	client, err := minio.NewMinioClient(endpoint, accessKey, secretKey, false)
	if err != nil {
		log.Println(err)
		return
	}
	fileContent := []byte("test")
	uploadInfo, err := minio.UploadFileToMinio(client, "test-bucket", "test.txt", bytes.NewReader(fileContent))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully uploaded bytes: ", uploadInfo)
	return
}
