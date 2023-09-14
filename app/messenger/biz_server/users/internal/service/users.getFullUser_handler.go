package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (s *Service) UsersGetFullUser(ctx context.Context, request *mtproto.TLUsersGetFullUser) (*mtproto.UserFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("users.getFullUser - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peerId int32
		id     = model.FromInputUser(md.UserId, request.Id)
	)

	switch id.PeerType {
	case model.PEER_SELF, model.PEER_USER:
		peerId = id.PeerId
	default:
		err := mtproto.ErrUserIdInvalid
		log.Error("users.getFullUser - error: %v", err)
		return nil, err
	}

	userFull, err := s.UserFacade.GetFullUser(ctx, md.UserId, peerId)
	if err != nil {
		log.Error("users.getFullUser - error: %v", err)
		return nil, err
	}

	// pinned_msg_id
	if id.PeerType == model.PEER_SELF {
		pinnedMsgId := s.PrivateFacade.GetUserPinnedMessage(ctx, md.UserId, md.UserId)
		if pinnedMsgId != 0 {
			userFull.PinnedMsgId = &types.Int32Value{Value: pinnedMsgId}
		}
	}

	if md.UserId != peerId {
		chatIdList := s.ChatFacade.GetUsersChatIdList(ctx, []int32{md.UserId, peerId})
		log.Debugf("users.getFullUser - chatIdList: %v", chatIdList)
		if len(chatIdList) == 2 {
			commonChats := util.Int32Intersect(chatIdList[md.UserId], chatIdList[peerId])
			log.Debugf("users.getFullUser - chatIdList: %v", commonChats)
			userFull.CommonChatsCount = int32(len(commonChats))
		}
	}

	log.Debugf("users.getFullUser - reply: %s", userFull.DebugString())
	return userFull, nil
}
