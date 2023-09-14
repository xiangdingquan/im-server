package dfsutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"open.chat/pkg/log"
	"os"
	"path/filepath"
)

const (
	testFilePath = "/testfile"
)

type cacheFs struct {
	directory string
}

func writeFileLocally(fr io.Reader, path string) (int64, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0774); err != nil {
		directory, _ := filepath.Abs(filepath.Dir(path))
		return 0, fmt.Errorf("writeFile: %s error: %v", directory, err)
	}
	fw, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	written, err := io.Copy(fw, fr)
	if err != nil {
		return written, err
	}
	return written, nil
}

func (m *cacheFs) TestConnection() error {
	f2 := bytes.NewReader([]byte("testingwrite"))
	if _, err := writeFileLocally(f2, filepath.Join(m.directory, testFilePath)); err != nil {
		return err
	}
	os.Remove(filepath.Join(m.directory, testFilePath))
	log.Info("Able to write files to local storage.")
	return nil
}

func (m *cacheFs) Reader(path string) (io.ReadCloser, error) {
	f, err := os.Open(filepath.Join(m.directory, path))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (m *cacheFs) ReadFile(path string) ([]byte, error) {
	f, err := ioutil.ReadFile(filepath.Join(m.directory, path))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (m *cacheFs) FileExists(path string) (bool, error) {
	_, err := os.Stat(filepath.Join(m.directory, path))

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// todo
func (m *cacheFs) CopyFile(oldPath, newPath string) error {
	//if err := util.CopyFile(filepath.Join(m.directory, oldPath), filepath.Join(m.directory, newPath)); err != nil {
	//	return err
	//}
	return nil
}

func (m *cacheFs) MoveFile(oldPath, newPath string) error {
	if err := os.MkdirAll(filepath.Dir(filepath.Join(m.directory, newPath)), 0774); err != nil {
		return err
	}

	if err := os.Rename(filepath.Join(m.directory, oldPath), filepath.Join(m.directory, newPath)); err != nil {
		return err
	}

	return nil
}

func (m *cacheFs) WriteFile(fr io.Reader, path string) (int64, error) {
	return writeFileLocally(fr, filepath.Join(m.directory, path))
}

func (m *cacheFs) RemoveFile(path string) error {
	if err := os.Remove(filepath.Join(m.directory, path)); err != nil {
		return err
	}
	return nil
}

func (m *cacheFs) ListDirectory(path string) (*[]string, error) {
	var paths []string
	fileInfos, err := ioutil.ReadDir(filepath.Join(m.directory, path))
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			paths = append(paths, filepath.Join(path, fileInfo.Name()))
		}
	}
	return &paths, nil
}

func (m *cacheFs) RemoveDirectory(path string) error {
	if err := os.RemoveAll(filepath.Join(m.directory, path)); err != nil {
		return err
	}
	return nil
}
