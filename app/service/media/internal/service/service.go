package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/service/dfs/dfspb"
	"open.chat/app/service/dfs/model"
	"open.chat/app/service/media/internal/core"
	"open.chat/app/service/media/internal/dao"
	"open.chat/app/service/media/mediapb"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

type Service struct {
	ac *paladin.Map
	*dao.Dao
	*core.MediaCore
}

func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}

	dao := dao.New()
	s = &Service{
		ac:        ac,
		Dao:       dao,
		MediaCore: core.New(dao),
	}
	return s
}

func (s *Service) NbfsUploadPhotoFile(ctx context.Context, request *mediapb.TLNbfsUploadPhotoFile) (reply *mediapb.PhotoDataRsp, err error) {
	log.Debugf("nbfs.uploadPhotoFile - request: %s", logger.JsonDebugData(request))

	var (
		inputFile  = request.GetFile()
		fileMDList []*dfspb.PhotoFileMetadata
	)

	if request.GetFile() == nil {
		log.Errorf("file is nil")
		err = mtproto.ErrMediaInvalid
		return
	}

	if request.GetIsProfile() {
		fileMDList, err = s.DfsFacade.UploadProfilePhotoFile(ctx, request.OwnerId, inputFile)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	} else {
		fileMDList, err = s.DfsFacade.UploadPhotoFile(ctx, request.OwnerId, inputFile)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	photoId, accessHash, szList, err := s.MediaCore.UploadPhotoFile2(ctx, fileMDList)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	reply = &mediapb.PhotoDataRsp{
		PhotoId:    photoId,
		AccessHash: accessHash,
		Date:       int32(time.Now().Unix()),
		SizeList:   szList,
	}

	log.Debugf("nbfs.uploadPhotoFile - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

func (s *Service) NbfsUploadVideoFile(ctx context.Context, request *mediapb.TLNbfsUploadVideoFile) (document *mtproto.Document, err error) {
	log.Debugf("nbfs.NbfsUploadedVideoFile - request: %s", logger.JsonDebugData(request))

	inputFile := request.GetFile()
	document, err = s.DfsFacade.UploadVideoDocument(ctx, request.OwnerId, inputFile)
	if err != nil {
		log.Errorf("nbfs.uploadedDocumentMedia - error: %v", err)
		return
	}

	fileName := inputFile.GetName()
	fileExtName := model.GetFileExtName(fileName)
	s.MediaCore.SaveDocument(ctx, fileName, fileExtName, document)

	log.Debugf("nbfs.uploadedVideoFile - document: %s", logger.JsonDebugData(document))
	return document, nil
}

func (s *Service) NbfsGetPhotoFileData(ctx context.Context, request *mediapb.TLNbfsGetPhotoFileData) (*mediapb.PhotoDataRsp, error) {
	log.Debugf("nbfs.getPhotoFileData - request: %s", logger.JsonDebugData(request))

	var photoId = request.GetPhotoId()
	szList := s.MediaCore.GetPhotoSizeList(ctx, photoId)
	reply := &mediapb.PhotoDataRsp{
		PhotoId:  photoId,
		SizeList: szList,
	}

	log.Debugf("nbfs.getPhotoFileData - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

func (s *Service) NbfsUploadedPhotoMedia(ctx context.Context, request *mediapb.TLNbfsUploadedPhotoMedia) (*mtproto.MessageMedia, error) {
	log.Debugf("nbfs.uploadedPhotoMedia - request: %s", logger.JsonDebugData(request))

	var (
		inputFile  = request.GetMedia().GetFile()
		fileMDList []*dfspb.PhotoFileMetadata
		err        error
	)

	if request.IsProfile {
		fileMDList, err = s.DfsFacade.UploadProfilePhotoFile(ctx, request.OwnerId, inputFile)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	} else {
		fileMDList, err = s.DfsFacade.UploadPhotoFile(ctx, request.OwnerId, inputFile)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	photoId, accessHash, szList, err := s.MediaCore.UploadPhotoFile2(ctx, fileMDList)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	photo := mtproto.MakeTLPhoto(&mtproto.Photo{
		Id:          photoId,
		HasStickers: false,
		AccessHash:  accessHash,
		Date:        int32(time.Now().Unix()),
		Sizes:       szList,
		DcId:        2,
	})

	var reply = mtproto.MakeTLMessageMediaPhoto(&mtproto.MessageMedia{
		Photo_FLAGPHOTO: photo.To_Photo(),
		TtlSeconds:      request.GetMedia().GetTtlSeconds(),
	})

	log.Debugf("nbfs.uploadedPhotoMedia - reply: %s", logger.JsonDebugData(reply))
	return reply.To_MessageMedia(), nil

}

func (s *Service) NbfsUploadedDocumentMedia(ctx context.Context, request *mediapb.TLNbfsUploadedDocumentMedia) (*mtproto.MessageMedia, error) {
	log.Debugf("nbfs.uploadedDocumentMedia - request: %s", logger.JsonDebugData(request))

	var (
		inputFile  = request.GetMedia().GetFile()
		inputThumb = request.GetMedia().GetThumb()
		media      = request.GetMedia()
		thumb      *mtproto.PhotoSize
		thumbId    int64
		isGif      = false
		document   *mtproto.Document
		err        error
	)

	for _, attr := range request.GetMedia().GetAttributes() {
		if attr.PredicateName == mtproto.Predicate_documentAttributeAnimated {
			isGif = true
			break
		}
	}
	if isGif && request.GetMedia().GetMimeType() == "image/gif" {
		document, err = s.DfsFacade.UploadGifDocumentMedia(ctx, request.OwnerId, request.GetMedia())
		if err != nil {
			log.Errorf("nbfs.uploadedDocumentMedia - error: %v", err)
			return nil, err
		}

		fileName := inputFile.GetName()
		fileExtName := model.GetFileExtName(fileName)

		s.MediaCore.SaveDocument(ctx, fileName, fileExtName, document)
	} else {
		if media.GetThumb() != nil {
			thumbFileMDList, err := s.DfsFacade.UploadPhotoFile(ctx, request.OwnerId, inputThumb)
			if err != nil {
				log.Error(err.Error())
				return nil, err
			}

			var szList []*mtproto.PhotoSize
			thumbId, _, szList, err = s.MediaCore.UploadPhotoFile2(ctx, thumbFileMDList)
			if err != nil {
				log.Error(err.Error())
				return nil, err
			}

			thumb = szList[0]
			if thumb.Size2 == 0 {
				thumb.Size2 = int32(len(thumb.Bytes))
			}
		} else {
			thumb = &mtproto.PhotoSize{
				Constructor: mtproto.CRC32_photoSizeEmpty,
				Type:        "s",
			}
		}

		fileMD, err := s.DfsFacade.UploadDocumentFile(ctx, request.OwnerId, inputFile)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		fileMD.MimeType = request.GetMedia().GetMimeType()
		mediaAttributes, _ := json.Marshal(media.GetAttributes())
		data, err := s.MediaCore.DoUploadedDocumentFile2(ctx, fileMD, thumbId, mediaAttributes)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		document = mtproto.MakeTLDocument(&mtproto.Document{
			Id:         data.DocumentId,
			AccessHash: data.AccessHash,
			Date:       int32(time.Now().Unix()),
			MimeType:   data.MimeType,
			Size2:      int32(data.FileSize),
			Thumb:      thumb,
			DcId:       fileMD.DcId,
			Attributes: media.GetAttributes(),
		}).To_Document()

		if thumb.Size2 > 0 {
			document.Thumbs = []*mtproto.PhotoSize{thumb}
		}
	}
	var reply = mtproto.MakeTLMessageMediaDocument(&mtproto.MessageMedia{
		Document:   document,
		TtlSeconds: request.GetMedia().GetTtlSeconds(),
	})

	log.Debugf("nbfs.uploadedDocumentMedia - reply: %s", logger.JsonDebugData(reply))
	return reply.To_MessageMedia(), nil
}

func (s *Service) NbfsGetDocument(ctx context.Context, request *mediapb.TLNbfsGetDocument) (*mtproto.Document, error) {
	log.Debugf("nbfs_getDocument - request: %s", logger.JsonDebugData(request))

	id := request.GetDocumentId()
	reply := s.MediaCore.GetDocument(ctx, id.Id, id.AccessHash, id.Version)

	log.Debugf("nbfs_getDocument - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

func (s *Service) NbfsGetDocumentList(ctx context.Context, request *mediapb.TLNbfsGetDocumentList) (*mediapb.DocumentList, error) {
	log.Debugf("nbfs_getDocumentList - request: %s", logger.JsonDebugData(request))

	idList := make([]int64, len(request.IdList))
	for i := 0; i < len(idList); i++ {
		idList[i] = request.IdList[i].Id
	}

	documents := s.MediaCore.GetDocumentList(ctx, idList)
	log.Debugf("nbfs_getDocumentList - reply: %s", logger.JsonDebugData(documents))
	return &mediapb.DocumentList{Documents: documents}, nil
}

func (s *Service) NbfsUploadEncryptedFile(ctx context.Context, request *mediapb.TLNbfsUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	log.Debugf("nbfs.uploadEncryptedFile - request: %s", logger.JsonDebugData(request))

	inputFile := request.File
	if inputFile == nil {
		return nil, fmt.Errorf("bad request")
	}

	fileMD, err := s.DfsFacade.UploadEncryptedFile(ctx, request.OwnerId, inputFile)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	encryptedFile, err := s.MediaCore.DoUploadedEncryptedFile2(ctx, fileMD, inputFile.GetKeyFingerprint())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debugf("nbfs.uploadEncryptedFile - reply: %s", logger.JsonDebugData(encryptedFile))
	return encryptedFile, nil
}

func (s *Service) NbfsGetEncryptedFile(ctx context.Context, request *mediapb.TLNbfsGetEncryptedFile) (*mtproto.EncryptedFile, error) {
	log.Debugf("nbfs.getEncryptedFile - request: %s", logger.JsonDebugData(request))
	encryptedFile, err := s.MediaCore.GetEncryptedFile(ctx, request.Id, request.AccessHash)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debugf("nbfs.getEncryptedFile - reply: %s", logger.JsonDebugData(encryptedFile))
	return encryptedFile, nil
}

func (s *Service) NbfsGetFileLocationSecret(ctx context.Context, request *mediapb.TLNbfsGetFileLocationSecret) (*mediapb.FileLocationSecret, error) {
	log.Debugf("nbfs.getFileLocationSecret - request: %s", logger.JsonDebugData(request))

	secret, err := s.PhotoDatasDAO.SelectAccessHash(ctx, request.VolumeId, request.LocalId)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	res := &mediapb.FileLocationSecret{
		Secret: secret,
	}

	log.Debugf("nbfs.getFileLocationSecret - reply: %s", logger.JsonDebugData(res))
	return res, nil
}
