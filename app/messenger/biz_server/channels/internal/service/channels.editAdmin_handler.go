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

func (s *Service) ChannelsEditAdmin(ctx context.Context, request *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.editAdmin#20b88214 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		channelId   int32
		adminRights model.ChatAdminRights
		rank        string
		userId      *model.PeerUtil
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
		log.Errorf("channels.editAdmin - error: %v", err)
		return nil, err
	}

	switch request.Constructor {
	case mtproto.CRC32_channels_editAdmin_20b88214:
		adminRights = model.MakeChannelAdminRights(request.AdminRights_CHANNELADMINRIGHTS)
	case mtproto.CRC32_channels_editAdmin_70f893ba:
		adminRights = model.MakeChatAdminRights(request.AdminRights_CHATADMINRIGHTS)
	case mtproto.CRC32_channels_editAdmin_d33c8902:
		adminRights = model.MakeChatAdminRights(request.AdminRights_CHATADMINRIGHTS)
		rank = request.Rank
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("channels.editAdmin - error: %v", err)
		return nil, err
	}

	users := s.UserFacade.GetMutableUsers(ctx, md.UserId, userId.PeerId)
	me, _ := users.GetImmutableUser(md.UserId)
	_ = me
	user, _ := users.GetImmutableUser(userId.PeerId)
	_ = user

	channel, added, err := s.ChannelFacade.EditAdminRights(ctx, channelId, md.UserId, userId.PeerId, int32(adminRights), rank)
	if err != nil {
		log.Errorf("channels.editAdmin#20b88214 - error: %v", err)
		return nil, err
	}

	_ = added

	result := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()

	log.Debugf("channels.editTitle#566decd0 - reply: {%s}", result.DebugString())
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
