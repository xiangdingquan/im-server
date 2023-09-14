package core

import (
	"context"

	"open.chat/app/service/dfs/dfspb"
	"open.chat/app/service/media/internal/dal/dataobject"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type encryptedFileData struct {
	*dataobject.EncryptedFilesDO
}

func (m *MediaCore) DoUploadedEncryptedFile2(ctx context.Context, fileMD *dfspb.EncryptedFileMetadata, keyFingerPrint int32) (*mtproto.EncryptedFile, error) {
	do := &dataobject.EncryptedFilesDO{
		EncryptedFileId: fileMD.EncryptedFileId,
		AccessHash:      fileMD.AccessHash,
		DcId:            fileMD.DcId,
		FilePath:        fileMD.FilePath,
		FileSize:        fileMD.FileSize,
		KeyFingerprint:  keyFingerPrint,
		Md5Checksum:     fileMD.Md5Hash,
	}

	var err error
	do.Id, _, err = m.EncryptedFilesDAO.Insert(ctx, do)
	if err != nil {
		return nil, err
	}

	encryptedFile := mtproto.MakeTLEncryptedFile(&mtproto.EncryptedFile{
		Id:             do.EncryptedFileId,
		AccessHash:     do.AccessHash,
		Size2:          do.FileSize,
		DcId:           do.DcId,
		KeyFingerprint: do.KeyFingerprint,
	})
	return encryptedFile.To_EncryptedFile(), nil
}

func (m *MediaCore) GetEncryptedFile(ctx context.Context, id, accessHash int64) (*mtproto.EncryptedFile, error) {
	do, err := m.EncryptedFilesDAO.SelectByFileLocation(ctx, id, accessHash)
	if err != nil {
		log.Errorf("error - %v", err)
		return nil, err
	}

	if do == nil {
		return mtproto.MakeTLEncryptedFileEmpty(nil).To_EncryptedFile(), nil
	} else {
		encryptedFile := mtproto.MakeTLEncryptedFile(&mtproto.EncryptedFile{
			Id:             do.EncryptedFileId,
			AccessHash:     do.AccessHash,
			Size2:          do.FileSize,
			DcId:           do.DcId,
			KeyFingerprint: do.KeyFingerprint,
		})
		return encryptedFile.To_EncryptedFile(), nil
	}
}
