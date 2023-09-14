package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UploadReuploadCdnFile1AF91C09(ctx context.Context, request *mtproto.TLUploadReuploadCdnFile1AF91C09) (*mtproto.Vector_CdnFileHash, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("upload.reuploadCdnFile1AF91C09 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("upload.reuploadCdnFile1AF91C09 - not imp UploadReuploadCdnFile1AF91C09")
}
