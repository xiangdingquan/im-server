package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetDocumentByHash(ctx context.Context, request *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getDocumentByHash#338e2464 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	document := mtproto.MakeTLDocumentEmpty(nil).To_Document()

	log.Debugf("messages.getDocumentByHash#338e2464 - reply: %s", document.DebugString())
	return document, nil

}
