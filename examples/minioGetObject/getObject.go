package main

import (
	"bytes"
	"log"
	"os"

	minio "github.com/k8-proxy/k8-go-comm/pkg/minio"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	bucketName := os.Getenv("BUCKET_NAME")
	fileName := os.Getenv("FILE_NAME")

	client, err := minio.NewMinioClient(endpoint, accessKey, secretKey, false)
	if err != nil {
		log.Println(err)
		return
	}
	object, err := minio.GetObjectFromMinio(client, bucketName, fileName)
	if err != nil {
		log.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(object)
	log.Println(buf.String())
	return
}
