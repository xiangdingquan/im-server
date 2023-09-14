package cachefs

import (
	"fmt"
	"os"

	"open.chat/pkg/log"
)

type EncryptedFile struct {
	fileId     int64
	accessHash int64
	*os.File
}

func CreateEncryptedFile(fileId, accessHash int64) (d *EncryptedFile, err error) {
	d = &EncryptedFile{fileId: fileId, accessHash: accessHash}
	d.File, err = os.Create(d.ToFilePath())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return d, nil
}

func NewEncryptedFile(fileId, accessHash int64) *EncryptedFile {
	return &EncryptedFile{fileId: fileId, accessHash: accessHash}
}

func (f *EncryptedFile) ToFilePath() string {
	return fmt.Sprintf("%s/0/%d.%d.dat", rootDataPath, f.fileId, f.accessHash)
}

func (f *EncryptedFile) ToFilePath2() string {
	return fmt.Sprintf("/0/%d.%d.dat", f.fileId, f.accessHash)
}

func (f *EncryptedFile) Write(b []byte) (int, error) {
	if f.File == nil {
		return 0, fmt.Errorf("file not open")
	}

	return f.File.Write(b)
}

func (f *EncryptedFile) Sync() {
	if f.File != nil {
		f.File.Sync()
	}
}

func (f *EncryptedFile) Close() {
	if f.File != nil {
		f.File.Close()
	}
}

func (f *EncryptedFile) ReadData(offset int32, limit int32) ([]byte, error) {
	return ReadFileOffsetData(f.ToFilePath(), offset, limit)
}
