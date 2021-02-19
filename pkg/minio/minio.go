package minio

import (
	"github.com/minio/minio-go"
	"github.com/tkanos/gonfig"
	"log"
)

//Minio configuration struct
type Configuration struct {
	ENDPOINT        string
	ACCESSKEYID     string
	SECRETACCESSKEY string
	USESSL          bool
}

var config = Configuration{}

//Map config values
func init() {
	gonfig.GetConf("config.minio.json", &config)
}

//Create new minio instance
func NewInstance() (*minio.Client, error) {

	Client, err := minio.New(config.ENDPOINT, config.ACCESSKEYID, config.SECRETACCESSKEY, config.USESSL)

	if err != nil {
		log.Fatalln(err)
	}

	return Client, err
}

//Upload file to mino
func uploadFile(minoClient *minio.Client, bucketName string, filePath string, objectName string, contentType string) (int64, error) {

	n, err := minoClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})

	return n, err
}
