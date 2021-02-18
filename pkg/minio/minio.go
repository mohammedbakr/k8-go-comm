package minio

import (
	"context"
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
)

//Mino settings and credentials
const (
	Endpoint        = "play.min.io"                              //to be replaced with real ones
	AccessKeyID     = ""                     //to be replaced with real ones
	SecretAccessKey = "" //to be replaced with real ones
	UseSSL          = true
	BucketName      = "" //to be replaced with real ones
)

//Get mini client
func MinioInstance() (*minio.Client, error) {

	Client, err := minio.New(Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyID, SecretAccessKey, ""),
		Secure: UseSSL,
	})

	return Client, err
}

//Upload file to mino
func uploadFile(minoClient *minio.Client, filePath string, objectName string, contentType string) (minio.UploadInfo, error) {
	ctx := context.Background()

	n, err := minoClient.FPutObject(ctx, BucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})

	return n, err
}
