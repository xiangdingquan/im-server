package service

import (
	"context"

	media_client "open.chat/app/service/media/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

const (
	kPhotoSizeOriginalType = "0" // client upload original photo
	kPhotoSizeSmallType    = "s"
	kPhotoSizeMediumType   = "m"
	kPhotoSizeXLargeType   = "x"
	kPhotoSizeYLargeType   = "y"
	kPhotoSizeAType        = "a"
	kPhotoSizeBType        = "b"
	kPhotoSizeCType        = "c"
)

func getLocalIdBySizeType(t string) int32 {
	switch t {
	case kPhotoSizeOriginalType:
		return 0
	case kPhotoSizeSmallType:
		return 1
	case kPhotoSizeMediumType:
		return 2
	case kPhotoSizeXLargeType:
		return 3
	case kPhotoSizeYLargeType:
		return 4
	case kPhotoSizeAType:
		return 5
	case kPhotoSizeBType:
		return 6
	case kPhotoSizeCType:
		return 7
	}

	return 0
}

func (s *Service) UploadGetFile(ctx context.Context, request *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.getFile#e3a6cfb5 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	switch request.Location.PredicateName {
	case mtproto.Predicate_inputPeerPhotoFileLocation:
		secret, err := media_client.GetFileLocationSecret(request.Location.VolumeId, request.Location.LocalId)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		request.Location.Secret = secret
	case mtproto.Predicate_inputPhotoFileLocation:
		localId := getLocalIdBySizeType(request.Location.ThumbSize)
		secret, err := media_client.GetFileLocationSecret(request.Location.Id, localId)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		request.Location.VolumeId = request.Location.Id
		request.Location.LocalId = localId
		request.Location.Secret = secret
	case mtproto.Predicate_inputDocumentFileLocation:
		if request.Location.ThumbSize != "" {
			document, err := media_client.GetDocumentById(request.Location.Id, request.Location.AccessHash)
			if err != nil {
				log.Error(err.Error())
				return nil, err
			}

			found := false
			for _, thumb2 := range document.Thumbs {
				if thumb2.GetType() == request.Location.ThumbSize {
					secret, err := media_client.GetFileLocationSecret(thumb2.Location.VolumeId, thumb2.Location.LocalId)
					if err != nil {
						log.Error(err.Error())
						return nil, err
					}
					request.Location.VolumeId = thumb2.Location.VolumeId
					request.Location.LocalId = thumb2.Location.LocalId
					request.Location.Secret = secret
					found = true
				}
			}
			if !found {
				err = mtproto.ErrLocationInvalid
				log.Errorf(err.Error())
				return nil, err
			}
		}
	default:
	}

	uploadFile, err := s.DfsFacade.DownloadFile(ctx, request.GetLocation(), request.GetOffset(), request.GetLimit())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debugf("upload.getFile#e3a6cfb5 - reply: {type: %v, mime: %d, len_bytes: %d}",
		uploadFile.GetType(),
		uploadFile.GetMtime(),
		len(uploadFile.GetBytes()))
	return uploadFile, nil
}
