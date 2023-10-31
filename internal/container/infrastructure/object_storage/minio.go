package objectstorage

import (
	"context"
	"fww-core/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Initialize(cfg *config.MinioConfig) (*minio.Client, error) {
	endpoint := cfg.Endpoint
	accessKeyID := cfg.AccessKeyID
	secretAccessKey := cfg.SecretAccessKey
	useSSL := cfg.UseSSL

	options := &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	}

	minioClient, err := minio.New(endpoint, options)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func CheckBucket(client *minio.Client, bucketName string) error {
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		errorMessage := "Error checking bucket, " + err.Error()
		panic(errorMessage)
	}
	if !exists {
		err = CreateBucket(client, bucketName)
		if err != nil {
			errorMessage := "Error creating bucket, " + err.Error()
			panic(errorMessage)
		}
	}

	return nil
}

func CreateBucket(minioClient *minio.Client, bucketName string) error {
	ctx := context.Background()
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return nil
		}
		return err
	}

	return nil
}
