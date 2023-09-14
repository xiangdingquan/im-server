package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadSaveFilePart(ctx context.Context, request *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.saveFilePart#b304a621 - metadata: %s, request: {file_id: %d, file_part: %d, bytes_len: %d}",
		md.DebugString(),
		request.FileId,
		request.FilePart,
		len(request.Bytes))

	// 400	FILE_PART_EMPTY	The provided file part is empty
	// 400	FILE_PART_INVALID	The file part number is invalid
	// 401	SESSION_PASSWORD_NEEDED	2FA is enabled, use a password to login
	err := s.DfsFacade.WriteFilePartData(ctx, md.AuthId, request.FileId, request.FilePart, request.Bytes)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debugf("upload.saveFilePart#b304a621 - reply: {true}")
	return mtproto.ToBool(true), nil
}
