package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"open.chat/pkg/http_client"
	"open.chat/pkg/math2"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadGetWebFile(ctx context.Context, request *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.getWebFile - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		webfile *mtproto.Upload_WebFile
	)

	switch request.GetLocation().GetPredicateName() {
	case mtproto.Predicate_inputWebFileLocation:
		err := mtproto.ErrLocationInvalid
		log.Errorf("upload.getWebFile - error: %v", err)
		return nil, err
	case mtproto.Predicate_inputWebFileGeoPointLocation:
		webFilePath := fmt.Sprintf("https://static-maps.yandex.ru/1.x/?lang=en_US&ll=%f,%f&size=400,400&z=%d&l=map",
			request.GetLocation().GetGeoPoint().GetLong(),
			request.GetLocation().GetGeoPoint().GetLat(),
			request.GetLocation().GetZoom())
		log.Debugf("webfile: %s", webFilePath)
		bytes, err := http_client.Get(webFilePath).
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
			SetTimeout(5*time.Second, 5*time.Second).
			Bytes()
		log.Debugf("bytes: %d", len(bytes))
		size2 := int32(len(bytes))
		if err != nil {
			log.Errorf("upload.getWebFile - error: %v", err)
			err := mtproto.ErrLocationInvalid
			return nil, err
		}

		if request.GetOffset() > int32(len(bytes)) {
			bytes = bytes[0:0]
		} else {
			bytes = bytes[int(request.GetOffset()) : request.GetOffset()+math2.Int32Min(request.GetLimit(), int32(len(bytes))-request.GetOffset())]
		}

		webfile = mtproto.MakeTLUploadWebFile(&mtproto.Upload_WebFile{
			Size2:    size2,
			MimeType: "img/png",
			FileType: mtproto.MakeTLStorageFileUnknown(nil).To_Storage_FileType(),
			Mtime:    int32(time.Now().Unix()),
			Bytes:    bytes,
		}).To_Upload_WebFile()
	default:
		err := mtproto.ErrLocationInvalid
		log.Errorf("upload.getWebFile - error: %v", err)
		return nil, err
	}

	log.Debugf("upload.getWebFile - reply: {size: %d, mime_type: %, file_type: %s, len_bytes: %d}",
		webfile.GetSize2(),
		webfile.GetMimeType(),
		webfile.GetFileType().DebugString(),
		len(webfile.GetBytes()))
	return webfile, nil
}
