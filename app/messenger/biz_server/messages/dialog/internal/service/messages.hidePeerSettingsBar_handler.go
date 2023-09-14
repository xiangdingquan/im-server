package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesHidePeerSettingsBar(ctx context.Context, request *mtproto.TLMessagesHidePeerSettingsBar) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.hidePeerSettingsBar - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	peer := model.FromInputPeer2(md.UserId, request.Peer)
	s.UserFacade.DeletePeerSettings(ctx, md.UserId, peer)

	log.Debugf("messages.hidePeerSettingsBar - reply {true}")
	return mtproto.BoolTrue, nil
}
