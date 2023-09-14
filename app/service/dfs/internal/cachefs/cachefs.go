package cachefs

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"open.chat/pkg/log"
)

var rootDataPath = "/opt/nbfs"
var subPaths = []string{"0", "a", "b", "c", "s", "m", "x", "y"}

// var uuidgen idgen.UUIDGen

func InitCacheFS(dataPath string) error {
	if dataPath != "" {
		rootDataPath = dataPath
	}
	for _, p := range subPaths {
		err := os.MkdirAll(rootDataPath+"/"+p, 0755)
		if err != nil {
			log.Errorf("init cache fs error: %v", err)
			return err
		}
	}
	return nil
}

func pathExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getFileSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	return strings.ToLower(ext)
}

type cacheFile struct {
	creatorId int64
	fileId    int64
}

func NewCacheFile(creatorId, fileId int64) *cacheFile {
	return &cacheFile{creatorId, fileId}
}

func (f *cacheFile) WriteFilePartData(filePart int32, bytes []byte) error {
	filePath := fmt.Sprintf("%s/0/%d.%d.parts", rootDataPath, f.creatorId, f.fileId)

	exist, err := pathExists(filePath)
	if err != nil {
		log.Errorf("pathExists error![%v]", err)
		return err
	}

	if !exist {
		err := os.Mkdir(filePath, 0755)
		if err != nil {
			log.Errorf("mkdir failed![%v]\n", err)
			return err
		}
	}

	fileName := fmt.Sprintf("%s/%d.part", filePath, filePart)
	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (f *cacheFile) CheckFileParts(fileParts int32) bool {
	for i := 0; i < int(fileParts); i++ {
		filePath := fmt.Sprintf("%s/0/%d.%d.parts/%d.part", rootDataPath, f.creatorId, f.fileId, i)
		filePartInfo, err := os.Stat(filePath)
		if err != nil {
			log.Error(err.Error())
			return false
		}

		if filePartInfo.IsDir() {
			err = fmt.Errorf("exist dir - %s", filePath)
			return false
		}
	}
	return true
}

func (f *cacheFile) Md5Checksum(fileParts int32) (string, error) {
	md5Hash := md5.New()
	for i := 0; i < int(fileParts); i++ {
		filePath := fmt.Sprintf("%s/0/%d.%d.parts/%d.part", rootDataPath, f.creatorId, f.fileId, i)
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			err = fmt.Errorf("path not exists: %s", filePath)
			return "", err
		}
		md5Hash.Write(b)
	}
	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}

func (f *cacheFile) ReadFileParts(fileParts int32, cb func(int, []byte)) (err error) {
	if cb == nil {
		return nil
	}

	for i := 0; i < int(fileParts); i++ {
		filePath := fmt.Sprintf("%s/0/%d.%d.parts/%d.part", rootDataPath, f.creatorId, f.fileId, i)
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			err = fmt.Errorf("read %s error: %v", filePath, err)
			return err
		}
		cb(i, b)
	}
	return nil
}

func ReadFileOffsetData(filePath string, offset int32, limit int32) ([]byte, error) {
	fileSize := getFileSize(filePath)

	if int64(offset) > fileSize {
		limit = 0
	} else if int64(offset+limit) > fileSize {
		limit = int32(fileSize - int64(offset))
	}

	f2, err := os.Open(filePath)
	if err != nil {
		log.Errorf("open %s error: %v", filePath, err)
		return nil, err
	}
	defer f2.Close()

	bytes := make([]byte, limit)
	_, err = f2.ReadAt(bytes, int64(offset))
	if err != nil {
		log.Errorf("read file %s error: ", filePath, err)
		return nil, err
	}
	return bytes, nil
}
