package main

import (
	"bytes"
	"log"

	minio "github.com/k8-proxy/k8-go-comm/pkg/minio"
)

func main() {
	client := minio.NewMinioClient("play.minio.io", "Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", false)
	fileContent := []byte("test")
	uploadInfo, err := minio.UploadFileToMinio(client, "test-bucket", "test.txt", bytes.NewReader(fileContent))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully uploaded bytes: ", uploadInfo)
	return
}
