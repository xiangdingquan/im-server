package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadGetCdnFileHashesF715C87B(ctx context.Context, request *mtproto.TLUploadGetCdnFileHashesF715C87B) (*mtproto.Vector_CdnFileHash, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.getCdnFileHashesF715C87B - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("upload.getCdnFileHashesF715C87B - not imp UploadGetCdnFileHashesF715C87B")
}
