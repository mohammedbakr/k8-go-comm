package main

import (
	"log"
	"time"

	minio "github.com/k8-proxy/k8-go-comm/pkg/minio"
)

func main() {
	client := minio.NewMinioClient("play.minio.io", "Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", false)
	presignedURL := minio.GetPresignedURLForObject(client, "test-bucket", "test.txt", time.Second*60*60*24)
	log.Println(presignedURL)
	return
}
