package main

import (
	"fmt"
	"k8-go-comm/pkg/minio"
	"log"
)

func main() {
	//created new minio client
	client, err := minio.NewInstance()

	if err != nil {
		log.Fatalln(err)
	} else {
		//Check if bucket already exists
		exists, errBucketExists := minio.CheckIfBucketExists(client, "mymusic")

		if exists != true && errBucketExists == nil {

			//If bucket doesn't exists create a bucket
			minio.CreateNewBucket(client, "mymusic", "us-east-1")

		}

		//Upload file
		minio.UploadFile(client, "mymusic", "../test.zip", "test.zip", "application/zip")

		//Generate Signed url
		presignedUrl, err := minio.GenerateSignedUrl(client, "mymusic", "test.zip")

		fmt.Println(presignedUrl, err)
	}

}
