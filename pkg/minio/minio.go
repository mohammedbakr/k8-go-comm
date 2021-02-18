package minio

import (
	"github.com/minio/minio-go"
	"github.com/tkanos/gonfig"
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

	return Client, err
}
