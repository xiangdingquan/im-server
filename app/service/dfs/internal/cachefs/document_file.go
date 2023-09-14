package cachefs

import (
	"fmt"
	"os"

	"open.chat/pkg/log"
)

type DocumentFile struct {
	fileId     int64
	accessHash int64
	*os.File
}

func CreateDocumentFile(fileId, accessHash int64) (d *DocumentFile, err error) {
	d = &DocumentFile{fileId: fileId, accessHash: accessHash}
	d.File, err = os.Create(d.ToFilePath())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return d, nil
}

func NewDocumentFile(fileId, accessHash int64) *DocumentFile {
	return &DocumentFile{fileId: fileId, accessHash: accessHash}
}

func (f *DocumentFile) ToFilePath() string {
	return fmt.Sprintf("%s/0/%d.%d.dat", rootDataPath, f.fileId, f.accessHash)
}

func (f *DocumentFile) ToFilePath2() string {
	return fmt.Sprintf("/0/%d.%d.dat", f.fileId, f.accessHash)
}

func (f *DocumentFile) Write(b []byte) (int, error) {
	if f.File == nil {
		return 0, fmt.Errorf("file not open")
	}

	return f.File.Write(b)
}

func (f *DocumentFile) Sync() {
	if f.File != nil {
		f.File.Sync()
	}
}

func (f *DocumentFile) Close() {
	if f.File != nil {
		f.File.Close()
	}
}

func (f *DocumentFile) ReadData(offset int32, limit int32) ([]byte, error) {
	return ReadFileOffsetData(f.ToFilePath(), offset, limit)
}
