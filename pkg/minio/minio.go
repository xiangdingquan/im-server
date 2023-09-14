package minio_util

import (
	"github.com/minio/minio-go"
)

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func NewMinioClient(c *MinioConfig) (*minio.Core, error) {
	return minio.NewCore(c.Endpoint, c.AccessKeyID, c.SecretAccessKey, c.UseSSL)
}
