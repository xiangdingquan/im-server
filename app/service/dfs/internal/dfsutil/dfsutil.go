package dfsutil

import (
	"io"
)

type DfsFileBackend interface {
	TestConnection() error

	Reader(path string) (io.ReadCloser, error)
	ReadFile(path string) ([]byte, error)
	FileExists(path string) (bool, error)
	CopyFile(oldPath, newPath string) error
	MoveFile(oldPath, newPath string) error
	WriteFile(fr io.Reader, path string) (int64, error)
	RemoveFile(path string) error

	ListDirectory(path string) (*[]string, error)
	RemoveDirectory(path string) error
}

func New() DfsFileBackend {
	return nil
}
