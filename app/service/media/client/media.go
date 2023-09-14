package media_client

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/pkg/env2"
	"open.chat/app/service/media/mediapb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util/client"
	"open.chat/pkg/log"
)

var (
	_self mediapb.RPCNbfsClient
)

func New() {
	if _self == nil {
		var (
			mc struct {
				Wardenclient *warden.ClientConfig
			}
			err error
		)

		if err := paladin.Get("media.toml").UnmarshalTOML(&mc); err != nil {
			if err != paladin.ErrNotExist {
				panic(err)
			}
		}

		log.Debugf("media_client config: %v", mc.Wardenclient)
		conn, err := client.NewClient(env2.ServiceMediaId, mc.Wardenclient)
		if err != nil {
			panic(err)
		}

		_self = mediapb.NewRPCNbfsClient(conn)
	}
}

func UploadPhotoFile(ownerId int64, file *mtproto.InputFile) (*mediapb.PhotoDataRsp, error) {
	request := &mediapb.TLNbfsUploadPhotoFile{
		OwnerId: ownerId,
		File:    file,
	}
	reply, err := _self.NbfsUploadPhotoFile(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func UploadVideoFile(ownerId int64, file *mtproto.InputFile) (*mtproto.Document, error) {
	request := &mediapb.TLNbfsUploadVideoFile{
		OwnerId: ownerId,
		File:    file,
	}
	reply, err := _self.NbfsUploadVideoFile(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func UploadProfilePhotoFile(ownerId int64, file *mtproto.InputFile) (*mediapb.PhotoDataRsp, error) {
	request := &mediapb.TLNbfsUploadPhotoFile{
		OwnerId:   ownerId,
		File:      file,
		IsProfile: true,
	}
	reply, err := _self.NbfsUploadPhotoFile(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetPhotoSizeList(photoId int64) ([]*mtproto.PhotoSize, error) {
	request := &mediapb.TLNbfsGetPhotoFileData{
		PhotoId: photoId,
	}
	reply, err := _self.NbfsGetPhotoFileData(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply.SizeList, nil
}

func UploadedPhotoMedia(ownerId int64, media *mtproto.TLInputMediaUploadedPhoto) (*mtproto.TLMessageMediaPhoto, error) {
	request := &mediapb.TLNbfsUploadedPhotoMedia{
		OwnerId: ownerId,
		Media:   media.To_InputMedia(),
	}

	reply, err := _self.NbfsUploadedPhotoMedia(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply.To_MessageMediaPhoto(), nil
}

func UploadedDocumentMedia(ownerId int64, media *mtproto.TLInputMediaUploadedDocument) (*mtproto.TLMessageMediaDocument, error) {
	request := &mediapb.TLNbfsUploadedDocumentMedia{
		OwnerId: ownerId,
		Media:   media.To_InputMedia(),
	}

	reply, err := _self.NbfsUploadedDocumentMedia(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply.To_MessageMediaDocument(), nil
}

func GetDocumentById(id, accessHash int64) (*mtproto.Document, error) {
	request := &mediapb.TLNbfsGetDocument{
		DocumentId: &mediapb.DocumentId{
			Id:         id,
			AccessHash: accessHash,
			Version:    0,
		},
	}
	reply, err := _self.NbfsGetDocument(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetDocumentByIdList(idList []int64) ([]*mtproto.Document, error) {
	idList2 := make([]*mediapb.DocumentId, 0, len(idList))
	for _, id := range idList {
		idList2 = append(idList2, &mediapb.DocumentId{Id: id})
	}
	reply, err := _self.NbfsGetDocumentList(context.Background(), &mediapb.TLNbfsGetDocumentList{IdList: idList2})
	if err != nil {
		return nil, err
	}

	return reply.Documents, nil
}

func UploadEncryptedFile(ownerId int64, file *mtproto.InputEncryptedFile) (*mtproto.EncryptedFile, error) {
	request := &mediapb.TLNbfsUploadEncryptedFile{
		OwnerId: ownerId,
		File:    file,
	}
	reply, err := _self.NbfsUploadEncryptedFile(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func GetEncryptedFile(id, accessHash int64) (*mtproto.EncryptedFile, error) {
	request := &mediapb.TLNbfsGetEncryptedFile{
		Id:         id,
		AccessHash: accessHash,
	}
	reply, err := _self.NbfsGetEncryptedFile(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func GetFileLocationSecret(volumeId int64, localId int32) (int64, error) {
	request := &mediapb.TLNbfsGetFileLocationSecret{
		VolumeId: volumeId,
		LocalId:  localId,
	}
	reply, err := _self.NbfsGetFileLocationSecret(context.Background(), request)
	if err != nil {
		return 0, err
	}

	return reply.Secret, nil
}
