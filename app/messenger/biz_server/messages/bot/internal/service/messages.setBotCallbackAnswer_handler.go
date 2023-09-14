package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSetBotCallbackAnswer(ctx context.Context, request *mtproto.TLMessagesSetBotCallbackAnswer) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.setBotCallbackAnswer - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	rpcMd, err := s.Dao.GetCacheRpcMetadata(ctx, request.QueryId)
	if err != nil {
		log.Errorf("messages.setBotCallbackAnswer - error: %v", err)
		return mtproto.BoolFalse, nil
	} else if rpcMd == nil {
		log.Errorf("messages.setBotCallbackAnswer - error: miss query_id(%d)", request.QueryId)
		return mtproto.BoolFalse, nil
	}

	rpcResult := &mtproto.TLRpcResult{
		ReqMsgId: rpcMd.ClientMsgId,
		Result: mtproto.MakeTLMessagesBotCallbackAnswer(&mtproto.Messages_BotCallbackAnswer{
			Alert:     request.Alert,
			HasUrl:    false,
			NativeUi:  true,
			Message:   request.Message,
			Url:       request.Url,
			CacheTime: request.CacheTime,
		}).To_Messages_BotCallbackAnswer(),
	}
	sync_client.PushRpcResult(ctx, rpcMd.AuthId, rpcMd.ServerId, rpcMd.SessionId, rpcMd.ClientMsgId, rpcResult.Encode(rpcMd.Layer))

	log.Debugf("messages.setBotCallbackAnswer - reply {true}")
	return mtproto.BoolTrue, nil
}
