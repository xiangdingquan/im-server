package service

import (
	"context"

	"open.chat/app/bots/botpb"
	sync_client "open.chat/app/messenger/sync/client"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetInlineBotResults(ctx context.Context, request *mtproto.TLMessagesGetInlineBotResults) (*mtproto.Messages_BotResults, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getInlineBotResults - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_INLINE_DISABLED	This bot can't be used in inline mode
	// 400	BOT_INVALID	This is not a valid bot
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CHANNEL_PRIVATE	You haven't joined this channel/supergroup
	// -503	Timeout	Timeout while fetching data
	//

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getInlineBotResults - error: %v", err)
		return nil, err
	}

	if request.GetBot().GetPredicateName() != mtproto.Predicate_inputUser {
		err := mtproto.ErrBotInvalid
		log.Errorf("messages.getInlineBotResults - error: %v", err)
		return nil, err
	}

	var (
		err        error
		botId      = request.Bot.UserId
		botResults *mtproto.Messages_BotResults
		peer       = model.FromInputPeer2(md.UserId, request.Peer)
	)

	if model.IsBotGif(botId) {
		botResults, err = s.GifClient.MessagesGetInlineBotResults(ctx, &botpb.GetInlineBotResults{
			UserId:    md.UserId,
			AuthKeyId: md.AuthId,
			Bot:       request.Bot,
			Peer:      request.Peer,
			GeoPoint:  request.GeoPoint,
			Query:     request.Query,
			Offset:    request.Offset,
		})

		if err != nil {
			log.Errorf("messages.getInlineBotResults - error: %v", err)
			return nil, err
		}

		if len(botResults.Results) > 0 {
			botResults.QueryId, _ = s.Dao.PutCacheInlineBotResults(ctx, botId, botResults.Results, 300)
		} else {
			botResults.QueryId = idgen.GetUUID()
		}

	} else if model.IsBotFoursquare(botId) {
		err = mtproto.ErrBotMethodInvalid
	} else if model.IsBotPic(botId) {
		err = mtproto.ErrBotMethodInvalid
	} else if model.IsBotBing(botId) {
		err = mtproto.ErrBotMethodInvalid
	} else {
		queryId, err := s.Dao.PutCacheRpcMetadata(ctx, md)
		if err != nil {
			//
			log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
			return nil, err
		}

		go func() {
			ctx2 := context.Background()

			pushUpdates := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateInlineBotCallbackQuery(&mtproto.Update{
				QueryId:                           queryId,
				UserId:                            md.UserId,
				Peer_PEER:                         model.MakePeerUser(md.UserId),
				MsgId_FLAGINPUTBOTINLINEMESSAGEID: nil,
				ChatInstance:                      0,
				Data_FLAGBYTES:                    nil,
				GameShortName:                     nil,
			}).To_Update())
			pushUpdates.Users = s.UserFacade.GetUserListByIdList(ctx2, peer.PeerId, []int32{md.UserId, peer.PeerId})
			sync_client.PushUpdates(ctx2, peer.PeerId, pushUpdates)
		}()

		err = mtproto.ErrNotReturnClient
		log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
		return nil, err
	}

	if err != nil {
		log.Errorf("messages.getInlineBotResults - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.getInlineBotResults - result: %v", botResults.DebugString())
	return botResults, nil
}
