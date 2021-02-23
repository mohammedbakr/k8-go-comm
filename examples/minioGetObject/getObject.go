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
	client := minio.NewMinioClient(endpoint, accessKey, secretKey, false)
	object, err := minio.GetObjectFromMinio(client, "test-bucket", "test.txt")
	if err != nil {
		log.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(object)
	log.Println(buf.String())
	return
}
