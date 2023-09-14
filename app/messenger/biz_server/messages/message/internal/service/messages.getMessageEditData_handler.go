package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetMessageEditData(ctx context.Context, request *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getMessageEditData#fda68d36 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err     error
		peer    = model.FromInputPeer2(md.UserId, request.GetPeer())
		boxList model.MessageBoxList
	)

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getMessageEditData - error: %v", err)
		return nil, err
	}

	if peer.PeerType == model.PEER_CHANNEL {
		boxList = s.MessageFacade.GetChannelMessageList(ctx, md.UserId, peer.PeerId, []int32{request.Id})
	} else {
		boxList = s.MessageFacade.GetUserMessageList(ctx, md.UserId, []int32{request.Id})
	}

	if len(boxList) < 1 {
		err = mtproto.ErrMessageAuthorRequired
		log.Errorf("messages.getMessageEditData - error: %v", err)
		return nil, err
	}
	box := boxList[0]

	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		switch box.Message.GetToId().GetPredicateName() {
		case mtproto.Predicate_peerUser:
			if peer.PeerType != model.PEER_SELF && peer.PeerType != model.PEER_USER {
				err = mtproto.ErrMessageAuthorRequired
				log.Errorf("messages.getMessageEditData - error: %v", err)
				return nil, err
			} else if peer.PeerId != boxList[0].Message.ToId.UserId {
				err = mtproto.ErrMessageAuthorRequired
				log.Errorf("messages.getMessageEditData - error: %v", err)
				return nil, err
			}
		case mtproto.Predicate_peerChat:
			if peer.PeerType != model.PEER_CHAT {
				err = mtproto.ErrMessageAuthorRequired
				log.Errorf("messages.getMessageEditData - error: %v", err)
				return nil, err
			} else if peer.PeerId != box.Message.ToId.ChatId {
				err = mtproto.ErrMessageAuthorRequired
				log.Errorf("messages.getMessageEditData - error: %v", err)
				return nil, err
			}
		default:
			err = mtproto.ErrMessageAuthorRequired
			log.Errorf("messages.getMessageEditData - error: %v", err)
			return nil, err
		}
	case model.PEER_CHANNEL:
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("messages.getMessageEditData - error: %v", err)
		return nil, err
	}

	editData := mtproto.MakeTLMessagesMessageEditData(&mtproto.Messages_MessageEditData{
		Caption: false,
	}).To_Messages_MessageEditData()

	log.Debugf("messages.getMessageEditData - reply: %s", editData.DebugString())
	return editData, nil
}
