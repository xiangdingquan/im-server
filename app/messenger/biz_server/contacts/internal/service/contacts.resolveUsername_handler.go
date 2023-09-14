package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.resolveUsername#f93ccba3 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 401	AUTH_KEY_PERM_EMPTY	The temporary auth key must be binded to the permanent auth key to use these methods.
	// 401	SESSION_PASSWORD_NEEDED	2FA is enabled, use a password to login
	// 400	USERNAME_INVALID	The provided username is not valid
	// 400	USERNAME_NOT_OCCUPIED	The provided username is not occupied
	//
	var peer *model.PeerUtil

	id := model.GetBotIdByName(request.GetUsername())
	if id > 0 {
		peer = model.MakeUserPeerUtil(id)
	} else {
		peerType, peerId, err := s.UsernameFacade.ResolveUsername(ctx, request.GetUsername())
		if err != nil {
			log.Errorf("contacts.resolveUsername#f93ccba3 - reply: {%v}", err)
			return nil, err
		}

		peer = &model.PeerUtil{
			PeerType: peerType,
			PeerId:   peerId,
		}
	}

	resolvedPeer := &mtproto.TLContactsResolvedPeer{Data2: &mtproto.Contacts_ResolvedPeer{
		Peer:  peer.ToPeer(),
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}}

	switch peer.PeerType {
	case model.PEER_USER:
		user, _ := s.UserFacade.GetUserById(ctx, md.UserId, peer.PeerId)
		if user != nil {
			resolvedPeer.SetUsers([]*mtproto.User{user})
		}
	case model.PEER_CHANNEL:
		chat := s.ChannelFacade.GetChannelById(ctx, md.UserId, peer.PeerId)
		if chat != nil {
			resolvedPeer.SetChats([]*mtproto.Chat{chat})
		}
	}

	log.Debugf("contacts.resolveUsername#f93ccba3 - reply: {%s}", resolvedPeer.DebugString())
	return resolvedPeer.To_Contacts_ResolvedPeer(), nil
}
