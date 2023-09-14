package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UsersGetUsers(ctx context.Context, request *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("users.getUsers - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	idList := make([]int32, 0, len(request.Id))
	for _, inputUser := range request.Id {
		peer := model.FromInputUser(md.UserId, inputUser)
		switch peer.PeerType {
		case model.PEER_SELF, model.PEER_USER:
			idList = append(idList, peer.PeerId)
		default:
			log.Errorf("invalid userId")
		}
	}

	mUsers := s.UserFacade.GetMutableUsers(ctx, append([]int32{md.UserId}, idList...)...)
	users := &mtproto.Vector_User{
		Datas: mUsers.GetUsersByIdList(md.UserId, idList),
	}

	log.Debugf("users.getUsers - reply: {%s}", users.DebugString())
	return users, nil
}
