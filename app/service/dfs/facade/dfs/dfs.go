package dfs

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"math/rand"

	"crypto/md5"
	"hash"
	"io/ioutil"
	"os"

	"open.chat/app/service/dfs/dfspb"
	dfs_facade "open.chat/app/service/dfs/facade"
	"open.chat/app/service/dfs/internal/dao"
	"open.chat/app/service/dfs/internal/gif2mp4"
	"open.chat/app/service/dfs/internal/imaging"
	"open.chat/app/service/dfs/internal/video2mp4"
	"open.chat/app/service/dfs/model"
	idgen "open.chat/app/service/idgen/client"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type dfsFacade struct {
	*dao.Dao
}

func New() dfs_facade.DfsFacade {
	idgen.NewUUID()
	return &dfsFacade{
		Dao: dao.New(),
	}
}

func (c *dfsFacade) WriteFilePartData(ctx context.Context, creatorId, fileId int64, filePart int32, bytes []byte) error {
	err := c.Dao.WriteFilePartData(ctx, creatorId, fileId, filePart, bytes)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}

func (c *dfsFacade) uploadPhotoFile(ctx context.Context, creatorId int64, file *mtproto.InputFile, isABC bool) (fileMDList []*dfspb.PhotoFileMetadata, err error) {
	var (
		fileSize   int64
		fileParts  int32
		partLength int32
		cacheData  []byte
		isBig      bool
		md5Hash    hash.Hash
	)

	if err = model2.CheckFileParts(file.Parts); err != nil {
		return
	}

	switch file.PredicateName {
	case mtproto.Predicate_inputFile:
		if len(file.Md5Checksum) != 0 {
			md5Hash = md5.New()
		}
	case mtproto.Predicate_inputFileBig:
		isBig = true
		_ = isBig
	default:
		err = mtproto.ErrTypeConstructorInvalid
		return
	}

	fileParts, err = c.Dao.GetFileParts(ctx, creatorId, file.Id)
	if fileParts != file.Parts {
		err = mtproto.ErrFilePartsInvalid
		return
	}

	// partLength, cacheData,
	err = c.Dao.ReadFileCB(ctx, creatorId, file.Id, file.Parts, func(part int32, bytes []byte) (err2 error) {
		if part == 0 {
			cacheData = make([]byte, 0, int(fileParts)*len(bytes))
			if fileParts > 1 {
				if err2 = model2.CheckFilePartSize(int32(len(bytes))); err2 != nil {
					return
				}
			}
		}
		cacheData = append(cacheData, bytes...)
		if md5Hash != nil {
			md5Hash.Write(bytes)
		}
		return
	})
	_ = partLength

	if err != nil {
		log.Error(err.Error())
		err = mtproto.ErrMediaInvalid
		return
	}

	if md5Hash != nil {
		if fmt.Sprintf("%x", md5Hash.Sum(nil)) != file.Md5Checksum {
			err = mtproto.ErrCheckSumInvalid
			return
		}
	}

	photoId := idgen.GetUUID()
	ext := model.GetFileExtName(file.GetName())

	extType := model.GetStorageFileTypeConstructor(ext)
	err = imaging.ReSizeImage(cacheData, ext, isABC, func(szType int, w, h int32, b []byte) error {
		secretId := int64(extType)<<32 | int64(rand.Uint32())
		path := fmt.Sprintf("%s/%d.%d.dat", imaging.GetSizeType(szType), photoId, secretId)
		log.Debugf("path: %s", path)
		fileSize, err = c.Dao.PutPhotoFile(ctx, path, b)
		if err != nil {
			return err
		}
		fileMD := &dfspb.PhotoFileMetadata{
			FileId:   file.GetId(),
			PhotoId:  photoId,
			DcId:     2,
			VolumeId: photoId,
			LocalId:  int32(szType),
			SecretId: secretId,
			Width:    w,
			Height:   h,
			FileSize: int32(fileSize),
			FilePath: path,
			Ext:      ext,
		}
		fileMDList = append(fileMDList, fileMD)
		return nil
	})

	if err != nil {
		log.Error(err.Error())
		err = mtproto.ErrImageProcessFailed
	}
	return
}

func (c *dfsFacade) UploadPhotoFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) ([]*dfspb.PhotoFileMetadata, error) {
	return c.uploadPhotoFile(ctx, creatorId, file, false)
}

func (c *dfsFacade) UploadProfilePhotoFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) ([]*dfspb.PhotoFileMetadata, error) {
	return c.uploadPhotoFile(ctx, creatorId, file, true)
}

func (c *dfsFacade) UploadDocumentFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) (fileMD *dfspb.DocumentFileMetadata, err error) {
	var (
		fileSize   = int64(0)
		documentId = idgen.GetUUID()
	)
	ext := model.GetFileExtName(file.GetName())
	extType := model.GetStorageFileTypeConstructor(ext)
	accessHash := int64(extType)<<32 | int64(rand.Uint32())

	r, _ := c.Dao.OpenFile(ctx, creatorId, file.GetId(), file.GetParts())
	path := fmt.Sprintf("%d.%d.dat", documentId, accessHash)
	log.Debugf("path: %s", path)

	fileSize, err = c.Dao.PutDocumentFile(ctx, path, r)
	if err != nil {
		return
	}

	fileMD = &dfspb.DocumentFileMetadata{
		FileId:           file.GetId(),
		DocumentId:       documentId,
		AccessHash:       accessHash,
		DcId:             2,
		FileSize:         int32(fileSize),
		FilePath:         path,
		UploadedFileName: file.GetName(),
		Ext:              ext,
	}
	return
}

func (c *dfsFacade) UploadEncryptedFile(ctx context.Context, creatorId int64, file *mtproto.InputEncryptedFile) (fileMD *dfspb.EncryptedFileMetadata, err error) {
	var (
		fileSize = int64(0)
	)

	encryptedFileId := idgen.GetUUID()
	accessHash := int64(rand.Uint32())

	r, _ := c.Dao.OpenFile(ctx, creatorId, file.GetId(), file.GetParts())
	path := fmt.Sprintf("%d.%d.dat", encryptedFileId, accessHash)
	log.Debugf("path: %s", path)
	if fileSize, err = c.Dao.PutEncryptedFile(ctx, path, r); err != nil {
		return
	}

	fileMD = &dfspb.EncryptedFileMetadata{
		FileId:          file.GetId(),
		EncryptedFileId: encryptedFileId,
		AccessHash:      accessHash,
		DcId:            2,
		FileSize:        int32(fileSize),
		FilePath:        path,
		//Md5Hash:         md5Hash.Sum(nil),
	}
	return
}

func (c *dfsFacade) DownloadFile(ctx context.Context, location *mtproto.InputFileLocation, offset, limit int32) (file *mtproto.Upload_File, err error) {
	var (
		bytes []byte
		sType int32
	)
	switch location.PredicateName {
	case mtproto.Predicate_inputFileLocation:
		fileLocation := location.To_InputFileLocation()
		path := fmt.Sprintf("%s/%d.%d.dat", imaging.GetSizeType(int(fileLocation.GetLocalId())), fileLocation.GetVolumeId(), fileLocation.GetSecret())
		log.Debugf("path: %s", path)
		bytes, err = c.Dao.GetFile(ctx, "photos", path, offset, limit)
		if err != nil {
			log.Warnf("download file: %v", err)
			err = nil
			bytes = []byte{}
		}
		sType = int32(fileLocation.GetSecret() >> 32)
	case mtproto.Predicate_inputEncryptedFileLocation:
		fileLocation := location.To_InputEncryptedFileLocation()
		path := fmt.Sprintf("%d.%d.dat", fileLocation.GetId(), fileLocation.GetAccessHash())
		log.Debugf("path: %s", path)
		bytes, err = c.Dao.GetFile(ctx, "encryptedfiles", path, offset, limit)
		if err != nil {
			log.Warnf("download file: %v", err)
			err = nil
			bytes = []byte{}
		}
		sType = int32(fileLocation.GetAccessHash() >> 32)
	case mtproto.Predicate_inputDocumentFileLocation:
		if location.ThumbSize == "" {
			fileLocation := location.To_InputDocumentFileLocation()
			path := fmt.Sprintf("%d.%d.dat", fileLocation.GetId(), fileLocation.GetAccessHash())
			log.Debugf("path: %s", path)
			bytes, err = c.Dao.GetFile(ctx, "documents", path, offset, limit)
			if err != nil {
				log.Warnf("download file: %v", err)
				err = nil
				bytes = []byte{}
			}
			sType = int32(fileLocation.GetAccessHash() >> 32)
		} else {
			path := fmt.Sprintf("%s/%d.%d.dat", imaging.GetSizeType(int(location.LocalId)), location.VolumeId, location.Secret)
			log.Debugf("path: %s", path)
			bytes, err = c.Dao.GetFile(ctx, "photos", path, offset, limit)
			if err != nil {
				log.Warnf("download file: %v", err)
				err = nil
				bytes = []byte{}
			}
			sType = int32(location.Secret >> 32)
		}
	case mtproto.Predicate_inputSecureFileLocation:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		log.Error(err.Error())
		return
	case mtproto.Predicate_inputTakeoutFileLocation:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		log.Error(err.Error())
		return
	case mtproto.Predicate_inputPhotoFileLocation:
		path := fmt.Sprintf("%s/%d.%d.dat", imaging.GetSizeType(int(location.LocalId)), location.VolumeId, location.Secret)
		log.Debugf("path: %s", path)
		sType = int32(location.Secret >> 32)
		bytes, err = c.Dao.GetFile(ctx, "photos", path, offset, limit)
		if err != nil {
			log.Warnf("download file: %v", err)
			err = nil
			bytes = []byte{}
		}
	case mtproto.Predicate_inputPeerPhotoFileLocation:
		path := fmt.Sprintf("%s/%d.%d.dat", imaging.GetSizeType(int(location.LocalId)), location.VolumeId, location.Secret)
		log.Debugf("path: %s", path)
		sType = int32(location.Secret >> 32)
		bytes, err = c.Dao.GetFile(ctx, "photos", path, offset, limit)
		if err != nil {
			log.Warnf("download file: %v", err)
			err = nil
			bytes = []byte{}
		}
	case mtproto.Predicate_inputStickerSetThumb:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		log.Error(err.Error())
		return
	default:
		err = fmt.Errorf("invalid InputFileLocation type: {%v}", location)
		log.Error(err.Error())
		return
	}

	uploadFile := mtproto.MakeTLUploadFile(&mtproto.Upload_File{
		Type:  model.MakeStorageFileType(sType),
		Mtime: int32(time.Now().Unix()),
		Bytes: bytes,
	})

	file = uploadFile.To_Upload_File()
	return
}

func getFileSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func (c *dfsFacade) saveThumbFiles(ctx context.Context, data *bytes.Buffer) ([]*mtproto.PhotoSize, error) {
	var (
		szList   = make([]*mtproto.PhotoSize, 0)
		photoId  = idgen.GetUUID()
		extType  = model.GetStorageFileTypeConstructor(".jpg")
		fileSize = int64(0)
		err      error
	)
	secretId := int64(extType)<<32 | int64(rand.Uint32())
	path := fmt.Sprintf("%s/%d.%d.dat", imaging.GetSizeType(2), photoId, secretId)
	log.Debugf("path: %s", path)
	fileSize, err = c.Dao.PutPhotoFile(ctx, path, data.Bytes())
	if err != nil {
		return szList, err
	}
	w, h := imaging.GetImgSize(data)
	szList = append(szList, mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
		Type:  imaging.GetSizeType(2),
		W:     w, //320
		H:     h, //268
		Size2: int32(fileSize),
		Location: mtproto.MakeTLFileLocation(&mtproto.FileLocation{
			VolumeId: photoId,
			LocalId:  int32(2),
			Secret:   secretId,
		}).To_FileLocation(),
	}).To_PhotoSize())
	return szList, nil
}

func (c *dfsFacade) UploadVideoDocument(ctx context.Context, creatorId int64, file *mtproto.InputFile) (document *mtproto.Document, err error) {
	var (
		documentId  = idgen.GetUUID()
		extName     = model.GetFileExtName(file.GetName())
		extType     = model.GetStorageFileTypeConstructor(extName)
		accessHash  = int64(extType)<<32 | int64(rand.Uint32())
		cacheData   []byte
		tmpFileName = fmt.Sprintf("/opt/nbfs/tmp/%d.%d%s", uint64(creatorId), uint64(documentId), extName)
	)

	err = c.Dao.ReadFileCB(ctx, creatorId, file.Id, file.Parts, func(part int32, bytes []byte) (err2 error) {
		cacheData = append(cacheData, bytes...)
		return
	})
	if err != nil {
		return
	}

	err = ioutil.WriteFile(tmpFileName, cacheData, 06444)
	if err != nil {
		log.Errorf("writeFile - error: %v", err)
	}
	mp4FileName := tmpFileName
	b, err := video2mp4.GetVideoFirstFrame(mp4FileName) //, 320, 268
	if err != nil {
		return
	}

	szList, err := c.saveThumbFiles(ctx, b)
	if err != nil {
		return
	}

	w := int32(320)
	h := int32(268)
	if len(szList) > 0 {
		w = szList[0].W
		h = szList[0].H
	}
	path := fmt.Sprintf("%d.%d.dat", documentId, accessHash)
	_, _ = c.Dao.FPutDocumentFile(ctx, path, mp4FileName)
	attributes := []*mtproto.DocumentAttribute{
		mtproto.MakeTLDocumentAttributeImageSize(&mtproto.DocumentAttribute{
			W: w,
			H: h,
		}).To_DocumentAttribute(),
		mtproto.MakeTLDocumentAttributeVideo(&mtproto.DocumentAttribute{
			RoundMessage:      false,
			SupportsStreaming: true,
			Duration:          6,
			W:                 w,
			H:                 h,
		}).To_DocumentAttribute(),
		mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
			FileName: file.GetName() + ".mp4",
		}).To_DocumentAttribute(),
	}
	document = mtproto.MakeTLDocument(&mtproto.Document{
		Id:            documentId,
		AccessHash:    accessHash,
		FileReference: []byte{},
		Date:          int32(time.Now().Unix()),
		MimeType:      "video/mp4",
		Size2:         int32(getFileSize(mp4FileName)),
		Thumbs:        szList,
		DcId:          2,
		Attributes:    attributes,
	}).To_Document()
	return
}

func (c *dfsFacade) UploadGifDocumentMedia(ctx context.Context, creatorId int64, media *mtproto.InputMedia) (*mtproto.Document, error) {
	var (
		fileSize    = int64(0)
		documentId  = idgen.GetUUID()
		err         error
		ext         = model.GetFileExtName(media.File.GetName())
		extType     = model.GetStorageFileTypeConstructor(ext)
		accessHash  = int64(extType)<<32 | int64(rand.Uint32())
		file        = media.File
		tmpFileName = fmt.Sprintf("/opt/nbfs/tmp/%d.%d.gif", uint64(creatorId), uint64(documentId))
	)

	var cacheData []byte
	err = c.Dao.ReadFileCB(ctx, creatorId, file.Id, file.Parts, func(part int32, bytes []byte) (err2 error) {
		cacheData = append(cacheData, bytes...)
		return
	})
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(tmpFileName, cacheData, 06444)
	if err != nil {
		log.Errorf("writeFile - error: %v", err)
	}
	defer os.Remove(tmpFileName)

	gifFileName, err := gif2mp4.Gif2Mp4Convert(tmpFileName, "mp4")
	if err != nil {
		return nil, err
	}
	defer os.Remove(gifFileName)

	b, err := gif2mp4.GetVideoFirstFrame(gifFileName, 320, 268)
	if err != nil {
		return nil, err
	}

	var (
		photoId  = idgen.GetUUID()
		ext2     = ".jpg"
		extType2 = model.GetStorageFileTypeConstructor(ext2)
	)
	szList := make([]*mtproto.PhotoSize, 0)

	secretId := int64(extType2)<<32 | int64(rand.Uint32())
	path := fmt.Sprintf("%s/%d.%d.dat", imaging.GetSizeType(2), photoId, secretId)
	log.Debugf("path: %s", path)
	fileSize, err = c.Dao.PutPhotoFile(ctx, path, b.Bytes())
	if err != nil {
		return nil, err
	}
	// if szType > 0 {
	szList = append(szList, mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
		Type:  imaging.GetSizeType(2),
		W:     320,
		H:     268,
		Size2: int32(fileSize),
		Location: mtproto.MakeTLFileLocation(&mtproto.FileLocation{
			VolumeId: photoId,
			LocalId:  int32(2),
			Secret:   secretId,
		}).To_FileLocation(),
	}).To_PhotoSize())
	// }

	path = fmt.Sprintf("%d.%d.dat", documentId, accessHash)
	_, _ = c.Dao.FPutDocumentFile(ctx, path, gifFileName)

	attributes := []*mtproto.DocumentAttribute{
		func() *mtproto.DocumentAttribute {
			for _, d := range media.GetAttributes() {
				if d.PredicateName == mtproto.Predicate_documentAttributeImageSize {
					return d
				}
			}
			return mtproto.MakeTLDocumentAttributeImageSize(&mtproto.DocumentAttribute{
				W: 705,
				H: 591,
			}).To_DocumentAttribute()
		}(),
		mtproto.MakeTLDocumentAttributeVideo(&mtproto.DocumentAttribute{
			RoundMessage:      false,
			SupportsStreaming: true,
			Duration:          6,
			W:                 320,
			H:                 268,
		}).To_DocumentAttribute(),
		mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
			FileName: media.GetFile().GetName() + ".mp4",
		}).To_DocumentAttribute(),
		mtproto.MakeTLDocumentAttributeAnimated(&mtproto.DocumentAttribute{}).To_DocumentAttribute(),
	}

	document := mtproto.MakeTLDocument(&mtproto.Document{
		Id:            documentId,
		AccessHash:    accessHash,
		FileReference: []byte{},
		Date:          int32(time.Now().Unix()),
		MimeType:      "video/mp4",
		Size2:         int32(getFileSize(gifFileName)),
		Thumbs:        szList,
		DcId:          2,
		Attributes:    attributes,
	}).To_Document()

	return document, nil
}

func init() {
	dfs_facade.Register("dfs", New)
}
