package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReadHistoryB04F2510(ctx context.Context, request *mtproto.TLMessagesReadHistoryB04F2510) (*mtproto.Messages_AffectedHistory, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.readHistory#b04f2510 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	peer := model.FromInputPeer2(md.UserId, request.Peer)
	if peer.PeerType == model.PEER_EMPTY ||
		peer.PeerType == model.PEER_CHANNEL {
		log.Errorf("invalid peer: %v", request.Peer)
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	affected, err := s.readHistory(ctx, md, peer, request.MaxId, request.Offset)
	if err != nil {
		log.Errorf("messages.readHistory#b04f2510 - readHistory error: %v", err)
		return nil, err
	}

	affectedHistory := &mtproto.TLMessagesAffectedHistory{Data2: &mtproto.Messages_AffectedHistory{
		Pts:      affected.Pts,
		PtsCount: affected.PtsCount,
		Offset:   request.GetOffset(),
	}}

	log.Debugf("messages.readHistory#b04f2510 - reply: {%s}", affectedHistory.DebugString())
	return affectedHistory.To_Messages_AffectedHistory(), nil
}
