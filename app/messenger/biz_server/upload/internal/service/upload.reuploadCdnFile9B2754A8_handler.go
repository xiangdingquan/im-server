package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadReuploadCdnFile9B2754A8(ctx context.Context, request *mtproto.TLUploadReuploadCdnFile9B2754A8) (*mtproto.Vector_FileHash, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.reuploadCdnFile9B2754A8 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("upload.reuploadCdnFile9B2754A8 - not imp UploadReuploadCdnFile9B2754A8")
}
