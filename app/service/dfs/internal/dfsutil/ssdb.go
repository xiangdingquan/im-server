package dfsutil

import "io"

type ssdbFs struct {
}

func (f *ssdbFs) TestConnection() error {
	return nil
}

func (f *ssdbFs) Reader(path string) (io.ReadCloser, error) {
	return nil, nil
}

func (f *ssdbFs) ReadFile(path string) ([]byte, error) {
	return nil, nil
}

func (f *ssdbFs) FileExists(path string) (bool, error) {
	return true, nil
}

func (f *ssdbFs) CopyFile(oldPath, newPath string) error {
	return nil
}

func (f *ssdbFs) MoveFile(oldPath, newPath string) error {
	return nil
}

func (f *ssdbFs) WriteFile(fr io.Reader, path string) (int64, error) {
	return 0, nil
}

func (f *ssdbFs) RemoveFile(path string) error {
	return nil
}

func (f *ssdbFs) ListDirectory(path string) (*[]string, error) {
	return nil, nil
}

func (f *ssdbFs) RemoveDirectory(path string) error {
	return nil
}
