package main

import (
	"bytes"
	"log"

	minio "github.com/k8-proxy/k8-go-comm/pkg/minio"
)

func main() {
	client := minio.NewMinioClient("play.minio.io", "Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", false)
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
