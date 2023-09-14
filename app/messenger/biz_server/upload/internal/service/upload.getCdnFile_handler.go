package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadGetCdnFile(ctx context.Context, request *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("UploadGetCdnFile - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("Not impl UploadGetCdnFile")
}
