package dao

import (
	"bytes"
	"context"
	"io"
	"path/filepath"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/encrypt"

	"open.chat/app/service/dfs/model"
	"open.chat/pkg/log"
)

func s3PutOptions(encrypted bool, contentType string) minio.PutObjectOptions {
	options := minio.PutObjectOptions{}
	if encrypted {
		options.ServerSideEncryption = encrypt.NewSSE()
	}
	options.ContentType = contentType

	return options
}

func (d *Dao) GetFile(ctx context.Context, bucket, path string, offset, limit int32) (bytes []byte, err error) {
	var (
		object *minio.Object
		n      int
	)

	object, err = d.minio.Client.GetObject(bucket, path, minio.GetObjectOptions{})
	if err != nil {
		log.Errorf("GetFile error: %v")
		return
	}

	bytes = make([]byte, limit)
	n, err = object.ReadAt(bytes, int64(offset))
	bytes = bytes[:n]
	if n > 0 {
		err = nil
	} else {
		log.Errorf("GetFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) PutPhotoFile(ctx context.Context, path string, buf []byte) (n int64, err error) {
	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = d.minio.Client.PutObject("photos", path, bytes.NewReader(buf), int64(len(buf)), options)
	if err != nil {
		log.Errorf("PutPhotoFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) PutDocumentFile(ctx context.Context, path string, r io.Reader) (n int64, err error) {
	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = d.minio.Client.PutObject("documents", path, r, -1, options)
	if err != nil {
		log.Errorf("PutDocumentFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) FPutDocumentFile(ctx context.Context, path string, r string) (n int64, err error) {
	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(false, contentType)
	n, err = d.minio.Client.FPutObject("documents", path, r, options)
	if err != nil {
		log.Errorf("PutDocumentFile (%s) error: %v", path, err)
	}
	return
}

func (d *Dao) PutEncryptedFile(ctx context.Context, path string, r io.Reader) (n int64, err error) {
	options := s3PutOptions(false, "binary/octet-stream")
	n, err = d.minio.Client.PutObject("encryptedfiles", path, r, -1, options)
	if err != nil {
		log.Errorf("PutEncryptedFile (%s) error: %v", path, err)
	}
	return
}
