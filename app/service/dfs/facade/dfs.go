package dfs_facade

import (
	"context"
	"fmt"

	"open.chat/app/service/dfs/dfspb"
	"open.chat/mtproto"
)

type DfsFacade interface {
	WriteFilePartData(ctx context.Context, creatorId, fileId int64, filePart int32, bytes []byte) error
	UploadPhotoFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) ([]*dfspb.PhotoFileMetadata, error)
	UploadProfilePhotoFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) ([]*dfspb.PhotoFileMetadata, error)
	UploadDocumentFile(ctx context.Context, creatorId int64, file *mtproto.InputFile) (*dfspb.DocumentFileMetadata, error)
	UploadEncryptedFile(ctx context.Context, creatorId int64, file *mtproto.InputEncryptedFile) (*dfspb.EncryptedFileMetadata, error)
	DownloadFile(ctx context.Context, location *mtproto.InputFileLocation, offset, limit int32) (*mtproto.Upload_File, error)
	UploadGifDocumentMedia(ctx context.Context, creatorId int64, media *mtproto.InputMedia) (*mtproto.Document, error)
	UploadVideoDocument(ctx context.Context, creatorId int64, file *mtproto.InputFile) (*mtproto.Document, error)
}

type Instance func() DfsFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewDfsFacade(name string) (inst DfsFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
