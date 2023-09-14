package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadSaveBigFilePart(ctx context.Context, request *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.saveBigFilePart#de7b673d - metadata: %s, request: {file_id: %d, file_part: %d, bytes_len: %d}",
		md.DebugString(),
		request.FileId,
		request.FilePart,
		len(request.Bytes))

	// 400	FILE_PARTS_INVALID	The number of file parts is invalid
	// 400	FILE_PART_EMPTY	The provided file part is empty
	// 400	FILE_PART_INVALID	The file part number is invalid
	// 400	FILE_PART_SIZE_CHANGED	Provided file part size has changed
	// 400	FILE_PART_SIZE_INVALID	The provided file part size is invalid
	// -503	Timeout	Timeout while fetching data
	err := s.DfsFacade.WriteFilePartData(ctx, md.AuthId, request.FileId, request.FilePart, request.Bytes)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debugf("upload.saveBigFilePart#de7b673d - reply: {true}")
	return mtproto.ToBool(true), nil

}
