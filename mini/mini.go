package mini

import (
	"errors"
	"io"

	minio "github.com/minio/minio-go"
)

/*
Mini is Temporal's interface with our Minio object storage backend
It provides helper methods to temporal store file ojects that are uploaded to Temporal.
Once objects have been stored in Minio, it can then be uploaded to IPFS

TODO: Add in encryption module that mkaes use of minio server side encryption, allowing a user provided pasword
*/

// MinioManager is our helper methods to interface with minio
type MinioManager struct {
	Client *minio.Client
}

func NewMinioManager(endpoint, accessKeyID, secretAccessKey string, secure bool) (*MinioManager, error) {
	mm := &MinioManager{}
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, secure)
	if err != nil {
		return nil, err
	}
	mm.Client = client
	return mm, nil
}

func (mm *MinioManager) ListBuckets() ([]minio.BucketInfo, error) {
	return mm.Client.ListBuckets()
}

// PutObject is a wrapper for the mini PutObject method, returning the number of bytes put or an error
func (mm *MinioManager) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (int64, error) {
	bucketExists, err := mm.CheckIfBucketExists(bucketName)
	if err != nil {
		return 0, err
	}
	if !bucketExists {
		return 0, errors.New("bucket does not exist")
	}
	return mm.Client.PutObject(bucketName, objectName, reader, objectSize, opts)
}

func (mm *MinioManager) CheckIfBucketExists(bucketName string) (bool, error) {
	return mm.Client.BucketExists(bucketName)
}


func (mm)