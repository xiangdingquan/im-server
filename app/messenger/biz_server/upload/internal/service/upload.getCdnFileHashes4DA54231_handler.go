package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadGetCdnFileHashes4DA54231(ctx context.Context, request *mtproto.TLUploadGetCdnFileHashes4DA54231) (*mtproto.Vector_FileHash, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.getCdnFileHashes4DA54231 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("upload.getCdnFileHashes4DA54231 - not imp UploadGetCdnFileHashes4DA54231")
}
