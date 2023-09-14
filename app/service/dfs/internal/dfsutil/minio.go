package dfsutil

import (
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	s3 "github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/encrypt"

	"open.chat/app/service/dfs/model"
	"open.chat/pkg/log"
	"open.chat/pkg/minio"
)

type minioFs struct {
	bucket string
	region string
	*minio_util.MinioConfig
}

func (m *minioFs) s3New() (*s3.Core, error) {
	s3Clnt, err := minio_util.NewMinioClient(m.MinioConfig)
	if err != nil {
		return nil, err
	}
	return s3Clnt, nil
}

func (m *minioFs) TestConnection() error {
	s3Clnt, err := m.s3New()
	if err != nil {
		return err
	}

	exists, err := s3Clnt.BucketExists(m.bucket)
	if err != nil {
		return err
	}

	if !exists {
		log.Warn("Bucket specified does not exist. Attempting to create...")
		err := s3Clnt.MakeBucket(m.bucket, m.region)
		if err != nil {
			log.Error("Unable to create bucket.")
			return err
		}
	}
	log.Info("Connection to S3 or minio is good. Bucket exists.")
	return nil
}

func (m *minioFs) Reader(path string) (io.ReadCloser, error) {
	s3Clnt, err := m.s3New()
	if err != nil {
		return nil, err
	}
	minioObject, _, err := s3Clnt.GetObject(m.bucket, path, s3.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return minioObject, nil
}

func (m *minioFs) ReadFile(path string) ([]byte, error) {
	s3Clnt, err := m.s3New()
	if err != nil {
		return nil, err
	}
	minioObject, _, err := s3Clnt.GetObject(m.bucket, path, s3.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer minioObject.Close()
	if f, err := ioutil.ReadAll(minioObject); err != nil {
		return nil, err
	} else {
		return f, nil
	}
}

func (m *minioFs) FileExists(path string) (bool, error) {
	s3Clnt, err := m.s3New()

	if err != nil {
		return false, err
	}
	_, err = s3Clnt.StatObject(m.bucket, path, s3.StatObjectOptions{})

	if err == nil {
		return true, nil
	}

	if err.(s3.ErrorResponse).Code == "NoSuchKey" {
		return false, nil
	}

	return false, err
}

func (m *minioFs) CopyFile(oldPath, newPath string) error {
	s3Clnt, err := m.s3New()
	if err != nil {
		return err
	}

	source := s3.NewSourceInfo(m.bucket, oldPath, nil)
	destination, err := s3.NewDestinationInfo(m.bucket, newPath, encrypt.NewSSE(), nil)
	if err != nil {
		return err
	}
	if err = s3Clnt.Client.CopyObject(destination, source); err != nil {
		return err
	}
	return nil
}

func (m *minioFs) MoveFile(oldPath, newPath string) error {
	s3Clnt, err := m.s3New()
	if err != nil {
		return err
	}

	source := s3.NewSourceInfo(m.bucket, oldPath, nil)
	destination, err := s3.NewDestinationInfo(m.bucket, newPath, encrypt.NewSSE(), nil)
	if err != nil {
		return err
	}
	if err = s3Clnt.Client.CopyObject(destination, source); err != nil {
		return err
	}
	if err = s3Clnt.RemoveObject(m.bucket, oldPath); err != nil {
		return err
	}
	return nil
}

func (m *minioFs) WriteFile(fr io.Reader, path string) (int64, error) {
	s3Clnt, err := m.s3New()
	if err != nil {
		return 0, err
	}

	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(m.MinioConfig.UseSSL, contentType)
	var buf bytes.Buffer
	_, err = buf.ReadFrom(fr)
	if err != nil {
		return 0, err
	}

	written, err := s3Clnt.Client.PutObject(m.bucket, path, &buf, int64(buf.Len()), options)
	if err != nil {
		return written, err
	}

	return written, nil
}

func (m *minioFs) RemoveFile(path string) error {
	s3Clnt, err := m.s3New()
	if err != nil {
		return err
	}

	if err := s3Clnt.RemoveObject(m.bucket, path); err != nil {
		return err
	}

	return nil
}

func getPathsFromObjectInfos(in <-chan s3.ObjectInfo) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer close(out)

		for {
			info, done := <-in

			if !done {
				break
			}

			out <- info.Key
		}
	}()

	return out
}

func (m *minioFs) ListDirectory(path string) (*[]string, error) {
	var paths []string

	s3Clnt, err := m.s3New()
	if err != nil {
		return nil, err
	}

	doneCh := make(chan struct{})

	defer close(doneCh)

	for object := range s3Clnt.Client.ListObjects(m.bucket, path, false, doneCh) {
		if object.Err != nil {
			return nil, err
		}
		paths = append(paths, strings.Trim(object.Key, "/"))
	}

	return &paths, nil
}

func (m *minioFs) RemoveDirectory(path string) error {
	s3Clnt, err := m.s3New()
	if err != nil {
		return err
	}

	doneCh := make(chan struct{})

	for err := range s3Clnt.RemoveObjects(m.bucket, getPathsFromObjectInfos(s3Clnt.Client.ListObjects(m.bucket, path, true, doneCh))) {
		if err.Err != nil {
			doneCh <- struct{}{}
			return err.Err
		}
	}

	close(doneCh)
	return nil
}

func s3PutOptions(encrypted bool, contentType string) s3.PutObjectOptions {
	options := s3.PutObjectOptions{}
	if encrypted {
		options.ServerSideEncryption = encrypt.NewSSE()
	}
	options.ContentType = contentType

	return options
}
