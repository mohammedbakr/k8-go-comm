package main

import (
	"log"
	"os"
	"time"

	minio "github.com/k8-proxy/k8-go-comm/pkg/minio"
)

func main() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	client := minio.NewMinioClient(endpoint, accessKey, secretKey, false)
	presignedURL := minio.GetPresignedURLForObject(client, "test-bucket", "test.txt", time.Second*60*60*24)
	log.Println(presignedURL)
	return
}
