package main

import (
	"bytes"
	"log"

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
	fileName := "file.pdf"
	//ZIPNAME := "file.zip"

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
