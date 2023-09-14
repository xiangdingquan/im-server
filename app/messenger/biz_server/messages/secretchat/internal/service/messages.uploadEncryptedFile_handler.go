package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesUploadEncryptedFile(ctx context.Context, request *mtproto.TLMessagesUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("MessagesUploadEncryptedFile - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("not MessagesUploadEncryptedFile")
}
