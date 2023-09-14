package service

import (
	"context"
	"time"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsEditBanned(ctx context.Context, request *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.editBanned - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		channelId    int32
		bannedRights model.ChatBannedRights
		userId       *model.PeerUtil
		deleted      bool
	)

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.editAdmin - error: %v", err)
		return nil, err
	} else {
		channelId = request.Channel.ChannelId
	}

	userId = model.FromInputUser(md.UserId, request.UserId)
	switch userId.PeerType {
	case model.PEER_SELF:
	case model.PEER_USER:
	default:
		err := mtproto.ErrUserIdInvalid
		log.Errorf("channels.editBanned - error: %v", err)
		return nil, err
	}

	switch request.Constructor {
	case mtproto.CRC32_channels_editBanned_72796912:
		bannedRights = model.MakeChatBannedRights(request.BannedRights_CHATBANNEDRIGHTS)
	case mtproto.CRC32_channels_editBanned_bfd915cd:
		bannedRights = model.MakeChannelBannedRights(request.BannedRights_CHANNELBANNEDRIGHTS)
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("channels.editBanned - error: %v", err)
		return nil, err
	}

	channel, deleted, err := s.ChannelFacade.EditBanned(ctx, channelId, md.UserId, userId.PeerId, bannedRights)
	if err != nil {
		log.Errorf("channels.editBanned - error: %v", err)
		return nil, err
	}

	if deleted && channel.Channel.IsMegagroup() {
	}

	result := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()

	log.Debugf("channels.editBanned - reply: {%s}", result.DebugString())
	return model.WrapperGoFunc(result, func() {
		sync_client.PushUpdates(context.Background(), userId.PeerId, mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{model.MakeUpdateChannel(channel.GetChannelId())},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{channel.ToUnsafeChat(userId.PeerId)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates())
	}).(*mtproto.Updates), nil
}
