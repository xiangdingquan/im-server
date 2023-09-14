package model

import (
	"path"
	"strings"

	"open.chat/mtproto"
)

func GetFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	if ext == "" {
		ext = "partial"
	}
	return strings.ToLower(ext)
}

func GetStorageFileTypeConstructor(extName string) int32 {
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

func MakeStorageFileType(c int32) *mtproto.Storage_FileType {
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

var (
	IMAGE_EXTENSIONS = [5]string{".jpg", ".jpeg", ".gif", ".bmp", ".png"}
	IMAGE_MIME_TYPES = map[string]string{".jpg": "image/jpeg", ".jpeg": "image/jpeg", ".gif": "image/gif", ".bmp": "image/bmp", ".png": "image/png", ".tiff": "image/tiff"}
)

func IsFileExtImage(ext string) bool {
	ext = strings.ToLower(ext)
	for _, imgExt := range IMAGE_EXTENSIONS {
		if ext == imgExt {
			return true
		}
	}
	return false
}

func GetImageMimeType(ext string) string {
	ext = strings.ToLower(ext)
	if len(IMAGE_MIME_TYPES[ext]) == 0 {
		return "image"
	} else {
		return IMAGE_MIME_TYPES[ext]
	}
}
