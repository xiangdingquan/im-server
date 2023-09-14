package service

import (
	"context"
	"crypto/sha256"

	media_client "open.chat/app/service/media/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadGetFileHashes(ctx context.Context, request *mtproto.TLUploadGetFileHashes) (*mtproto.Vector_FileHash, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.getFileHashes#c7025931 - metadata: %s, request: %s", md.DebugString(), request.DebugString())
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

	var (
		offset     int32                    = request.GetOffset()
		count      int32                    = 10
		size       int32                    = 10240
		fileHashes *mtproto.Vector_FileHash = &mtproto.Vector_FileHash{
			Datas: make([]*mtproto.FileHash, 0, count),
		}
	)

	fd, err := s.DfsFacade.DownloadFile(ctx, request.GetLocation(), offset, size*count)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	data := fd.GetBytes()
	datalen := int32(len(data))
	for idx := int32(0); idx < datalen; idx += size {
		if idx+size > datalen {
			size = datalen - idx
		}
		sum := sha256.Sum256(data[idx : idx+size])
		fileHashes.Datas = append(fileHashes.Datas, mtproto.MakeTLFileHash(&mtproto.FileHash{
			Offset: offset + idx,
			Limit:  size,
			Hash:   sum[:],
		}).To_FileHash())
	}

	log.Debugf("upload.getFileHashes#c7025931 - reply: %s", fileHashes.DebugString())
	return fileHashes, nil
}
