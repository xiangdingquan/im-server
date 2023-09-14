package cachefs

import (
	"context"
	"crypto/md5"
	"fmt"
	"hash"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"open.chat/app/service/dfs/dfspb"
	dfs_facade "open.chat/app/service/dfs/facade"
	"open.chat/app/service/dfs/internal/cachefs"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type localNbfsFacade struct {
}

func New() dfs_facade.DfsFacade {
	idgen.NewUUID()
	return &localNbfsFacade{}
}

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	return strings.ToLower(ext)
}

func getStorageFileTypeConstructor(extName string) int32 {
	var c mtproto.TLConstructor
	switch extName {
	case ".partial":
		c = mtproto.CRC32_storage_filePartial
	case ".jpeg", ".jpg":
		c = mtproto.CRC32_storage_fileJpeg
	case ".gif":
		c = mtproto.CRC32_storage_fileGif
	case ".png":
		c = mtproto.CRC32_storage_filePng
	case ".pdf":
		c = mtproto.CRC32_storage_filePdf
	case ".mp3":
		c = mtproto.CRC32_storage_fileMp3
	case ".mov":
		c = mtproto.CRC32_storage_fileMov
	case ".mp4":
		c = mtproto.CRC32_storage_fileMp4
	case ".webp":
		c = mtproto.CRC32_storage_fileWebp
	default:
		c = mtproto.CRC32_storage_filePartial
	}
	return int32(c)
}

func makeStorageFileType(c int32) *mtproto.Storage_FileType {
	fileType := &mtproto.Storage_FileType{}

	switch mtproto.TLConstructor(c) {
	case mtproto.CRC32_storage_filePartial:
		fileType.PredicateName = mtproto.Predicate_storage_filePartial
	case mtproto.CRC32_storage_fileJpeg:
		fileType.PredicateName = mtproto.Predicate_storage_fileJpeg
	case mtproto.CRC32_storage_fileGif:
		fileType.PredicateName = mtproto.Predicate_storage_fileGif
	case mtproto.CRC32_storage_filePng:
		fileType.PredicateName = mtproto.Predicate_storage_filePng
	case mtproto.CRC32_storage_filePdf:
		fileType.PredicateName = mtproto.Predicate_storage_filePdf
	case mtproto.CRC32_storage_fileMp3:
		fileType.PredicateName = mtproto.Predicate_storage_fileMp3
	case mtproto.CRC32_storage_fileMov:
		fileType.PredicateName = mtproto.Predicate_storage_fileMov
	case mtproto.CRC32_storage_fileMp4:
		fileType.PredicateName = mtproto.Predicate_storage_fileMp4
	case mtproto.CRC32_storage_fileWebp:
		fileType.PredicateName = mtproto.Predicate_storage_fileWebp
	default:
		fileType.PredicateName = mtproto.Predicate_storage_fileWebp
	}
	return fileType
}

func (c *localNbfsFacade) WriteFilePartData(ctx context.Context, creatorId, fileId int64, filePart int32, bytes []byte) error {
	f := cachefs.NewCacheFile(creatorId, fileId)
	err := f.WriteFilePartData(filePart, bytes)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}

func (c *localNbfsFacade) UploadPhotoFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) (fileMDList []*dfspb.PhotoFileMetadata, err error) {
	var (
		inputFile = file
		md5Hash   hash.Hash
		fileSize  = int64(0)
	)

	var cacheData []byte

	if inputFile.GetMd5Checksum() != "" {
		md5Hash = md5.New()
	}
	cacheFile := cachefs.NewCacheFile(creatorId, inputFile.GetId())
	err = cacheFile.ReadFileParts(inputFile.GetParts(), func(i int, bytes []byte) {
		if md5Hash != nil {
			md5Hash.Write(bytes)
		}
		cacheData = append(cacheData, bytes...)
		fileSize += int64(len(bytes))
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if md5Hash != nil {
		if fmt.Sprintf("%x", md5Hash.Sum(nil)) != inputFile.GetMd5Checksum() {
			err = fmt.Errorf("invalid md5")
			return nil, err
		}
	}

	photoId := idgen.GetUUID()
	photoFile2 := cachefs.NewPhotoFile(photoId, 0, 0)

	ext := getFileExtName(inputFile.GetName())
	extType := getStorageFileTypeConstructor(ext)

	err = cachefs.DoUploadedPhotoFile(photoFile2, ext, cacheData, false, func(pi *cachefs.PhotoInfo) {
		secretId := int64(extType)<<32 | int64(rand.Uint32())

		srcFile := cachefs.NewPhotoFile(photoId, pi.LocalId, 0)
		dstFile := cachefs.NewPhotoFile(photoId, pi.LocalId, secretId)
		os.Rename(srcFile.ToFilePath(), dstFile.ToFilePath())

		fileMD := &dfspb.PhotoFileMetadata{
			FileId:   inputFile.GetId(),
			PhotoId:  photoId,
			DcId:     2,
			VolumeId: photoId,
			LocalId:  pi.LocalId,
			SecretId: secretId,
			Width:    pi.Width,
			Height:   pi.Height,
			FileSize: int32(pi.FileSize),
			FilePath: dstFile.ToFilePath2(),
			Ext:      ext,
		}
		fileMDList = append(fileMDList, fileMD)
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return
}

func (c *localNbfsFacade) UploadProfilePhotoFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) (fileMDList []*dfspb.PhotoFileMetadata, err error) {
	var (
		inputFile = file
		md5Hash   hash.Hash
		fileSize  = int64(0)
	)

	var cacheData []byte

	if inputFile.GetMd5Checksum() != "" {
		md5Hash = md5.New()
	}
	cacheFile := cachefs.NewCacheFile(creatorId, inputFile.GetId())
	err = cacheFile.ReadFileParts(inputFile.GetParts(), func(i int, bytes []byte) {
		if md5Hash != nil {
			md5Hash.Write(bytes)
		}
		cacheData = append(cacheData, bytes...)
		fileSize += int64(len(bytes))
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if md5Hash != nil {
		if fmt.Sprintf("%x", md5Hash.Sum(nil)) != inputFile.GetMd5Checksum() {
			err = fmt.Errorf("invalid md5")
			return nil, err
		}
	}

	photoId := idgen.GetUUID()
	photoFile2 := cachefs.NewPhotoFile(photoId, 0, 0)

	ext := getFileExtName(inputFile.GetName())
	extType := getStorageFileTypeConstructor(ext)

	err = cachefs.DoUploadedPhotoFile(photoFile2, ext, cacheData, false, func(pi *cachefs.PhotoInfo) {
		secretId := int64(extType)<<32 | int64(rand.Uint32())

		srcFile := cachefs.NewPhotoFile(photoId, pi.LocalId, 0)
		dstFile := cachefs.NewPhotoFile(photoId, pi.LocalId, secretId)
		os.Rename(srcFile.ToFilePath(), dstFile.ToFilePath())

		fileMD := &dfspb.PhotoFileMetadata{
			FileId:   inputFile.GetId(),
			PhotoId:  photoId,
			DcId:     2,
			VolumeId: photoId,
			LocalId:  pi.LocalId,
			SecretId: secretId,
			Width:    pi.Width,
			Height:   pi.Height,
			FileSize: int32(pi.FileSize),
			FilePath: dstFile.ToFilePath2(),
			Ext:      ext,
		}
		fileMDList = append(fileMDList, fileMD)
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return
}

func (c *localNbfsFacade) UploadDocumentFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) (fileMD *dfspb.DocumentFileMetadata, err error) {
	var (
		inputFile    = file
		fileSize     = int64(0)
		documentFile *cachefs.DocumentFile
	)

	documentId := idgen.GetUUID()
	ext := getFileExtName(inputFile.GetName())
	extType := getStorageFileTypeConstructor(ext)
	accessHash := int64(extType)<<32 | int64(rand.Uint32())

	documentFile, err = cachefs.CreateDocumentFile(documentId, accessHash)
	if err != nil {
		return nil, err
	}
	defer documentFile.Close()

	// var cacheData []byte
	cacheFile := cachefs.NewCacheFile(creatorId, inputFile.GetId())
	err = cacheFile.ReadFileParts(inputFile.GetParts(), func(i int, bytes []byte) {
		// cacheData = append(cacheData, bytes...)
		fileSize += int64(len(bytes))
		documentFile.Write(bytes)
		documentFile.Sync()
	})
	if err != nil {
		log.Error(err.Error())
		return
	}

	fileMD = &dfspb.DocumentFileMetadata{
		FileId:           inputFile.GetId(),
		DocumentId:       documentId,
		AccessHash:       accessHash,
		DcId:             2,
		FileSize:         int32(fileSize),
		FilePath:         documentFile.ToFilePath2(),
		UploadedFileName: inputFile.GetName(),
		Ext:              ext,
		// MimeType:         inputFile.GetName(),
	}
	return
}

func (c *localNbfsFacade) UploadEncryptedFile(ctx context.Context, creatorId int64, file *mtproto.InputEncryptedFile) (fileMD *dfspb.EncryptedFileMetadata, err error) {
	var (
		inputFile     = file
		fileSize      = int64(0)
		md5Hash       hash.Hash
		encryptedFile *cachefs.EncryptedFile
	)

	encryptedFileId := idgen.GetUUID()
	accessHash := int64(rand.Uint32())

	encryptedFile, err = cachefs.CreateEncryptedFile(encryptedFileId, accessHash)
	if err != nil {
		return nil, err
	}
	defer encryptedFile.Close()

	if inputFile.GetMd5Checksum() != "" {
		md5Hash = md5.New()
	}

	// var cacheData []byte
	cacheFile := cachefs.NewCacheFile(creatorId, inputFile.GetId())
	err = cacheFile.ReadFileParts(inputFile.GetParts(), func(i int, bytes []byte) {
		if md5Hash != nil {
			md5Hash.Write(bytes)
		}

		fileSize += int64(len(bytes))
		encryptedFile.Write(bytes)
		encryptedFile.Sync()
	})
	if err != nil {
		log.Error(err.Error())
		return
	}

	fileMD = &dfspb.EncryptedFileMetadata{
		FileId:          inputFile.GetId(),
		EncryptedFileId: encryptedFileId,
		AccessHash:      accessHash,
		DcId:            2,
		FileSize:        int32(fileSize),
		FilePath:        encryptedFile.ToFilePath2(),
		//Md5Hash:         md5Hash.Sum(nil),
	}

	if md5Hash != nil {
		fileMD.Md5Hash = fmt.Sprintf("%x", md5Hash.Sum(nil))
		if fileMD.Md5Hash != inputFile.GetMd5Checksum() {
			err = fmt.Errorf("invalid md5")
			return nil, err
		}
	}

	return
}

func (c *localNbfsFacade) DownloadFile(ctx context.Context, location *mtproto.InputFileLocation, offset, limit int32) (file *mtproto.Upload_File, err error) {
	var (
		bytes []byte
		sType int32
	)

	switch location.PredicateName {
	case mtproto.Predicate_inputFileLocation:
		fileLocation := location.To_InputFileLocation()
		file := cachefs.NewPhotoFile(fileLocation.GetVolumeId(), fileLocation.GetLocalId(), fileLocation.GetSecret())
		bytes, err = file.ReadData(offset, limit)
		sType = int32(fileLocation.GetSecret() >> 32)
	case mtproto.Predicate_inputEncryptedFileLocation:
		fileLocation := location.To_InputEncryptedFileLocation()
		file := cachefs.NewEncryptedFile(fileLocation.GetId(), fileLocation.GetAccessHash())
		bytes, err = file.ReadData(offset, limit)
		sType = int32(fileLocation.GetAccessHash() >> 32)
	case mtproto.Predicate_inputDocumentFileLocation:
		fileLocation := location.To_InputDocumentFileLocation()
		file := cachefs.NewDocumentFile(fileLocation.GetId(), fileLocation.GetAccessHash())
		bytes, err = file.ReadData(offset, limit)
		sType = int32(fileLocation.GetAccessHash() >> 32)
	case mtproto.Predicate_inputSecureFileLocation:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		log.Error(err.Error())
		return nil, err
	case mtproto.Predicate_inputTakeoutFileLocation:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		log.Error(err.Error())
		return nil, err
	case mtproto.Predicate_inputPhotoFileLocation:
		file := cachefs.NewPhotoFile(location.VolumeId, location.LocalId, location.Secret)
		bytes, err = file.ReadData(offset, limit)
		sType = int32(location.Secret >> 32)
	case mtproto.Predicate_inputPeerPhotoFileLocation:
		file := cachefs.NewPhotoFile(location.VolumeId, location.LocalId, location.Secret)
		bytes, err = file.ReadData(offset, limit)
		sType = int32(location.Secret >> 32)
	case mtproto.Predicate_inputStickerSetThumb:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		log.Error(err.Error())
		return nil, err
	default:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		log.Error(err.Error())
		return nil, err
	}

	uploadFile := mtproto.MakeTLUploadFile(&mtproto.Upload_File{
		Type:  makeStorageFileType(sType),
		Mtime: int32(time.Now().Unix()),
		Bytes: bytes,
	})

	file = uploadFile.To_Upload_File()
	return
}

func (c *localNbfsFacade) UploadVideoDocument(ctx context.Context, creatorId int64, file *mtproto.InputFile) (document *mtproto.Document, err error) {
	return nil, mtproto.ErrMethodNotImpl
}

func (c *localNbfsFacade) UploadGifDocumentMedia(ctx context.Context, creatorId int64, media *mtproto.InputMedia) (*mtproto.Document, error) {
	return nil, mtproto.ErrMethodNotImpl
}

func init() {
	dfs_facade.Register("local", New)
}
