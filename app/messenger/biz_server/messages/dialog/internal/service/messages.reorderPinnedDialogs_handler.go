package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReorderPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.reorderPinnedDialogs - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peerUsers    []int32
		peerChats    []int32
		peerChannels []int32
	)

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.reorderPinnedDialogs - error: %v", err)
		return nil, err
	}

	switch request.Constructor {
	case mtproto.CRC32_messages_reorderPinnedDialogs_959ff644:
		for _, peer := range request.Order_VECTORINPUTPEER {
			p := model.FromInputPeer2(md.UserId, peer)
			switch p.PeerType {
			case model.PEER_SELF, model.PEER_USER:
				peerUsers = append(peerUsers, p.PeerId)
			case model.PEER_CHAT:
				peerChats = append(peerChats, p.PeerId)
			case model.PEER_CHANNEL:
				peerChannels = append(peerChannels, p.PeerId)
			default:
				err := mtproto.ErrPeerIdInvalid
				log.Errorf("messages.reorderPinnedDialogs - error: %v", err)
				return nil, err
			}
		}
	case mtproto.CRC32_messages_reorderPinnedDialogs_5b51d63f:
		for _, peer := range request.Order_VECTORINPUTDIALOGPEER {
			switch peer.PredicateName {
			case mtproto.Predicate_inputDialogPeer:
				p := model.FromInputPeer2(md.UserId, peer.Peer)
				switch p.PeerType {
				case model.PEER_SELF, model.PEER_USER:
					peerUsers = append(peerUsers, p.PeerId)
				case model.PEER_CHAT:
					peerChats = append(peerChats, p.PeerId)
				case model.PEER_CHANNEL:
					peerChannels = append(peerChannels, p.PeerId)
				default:
					err := mtproto.ErrPeerIdInvalid
					log.Errorf("messages.reorderPinnedDialogs - error: %v", err)
					return nil, err
				}
			default:
				err := mtproto.ErrPeerIdInvalid
				log.Errorf("messages.reorderPinnedDialogs - error: %v", err)
				return nil, err
			}
		}
	case mtproto.CRC32_messages_reorderPinnedDialogs_3b1adf37:
		for _, peer := range request.Order_VECTORINPUTDIALOGPEER {
			switch peer.PredicateName {
			case mtproto.Predicate_inputDialogPeer:
				p := model.FromInputPeer2(md.UserId, peer.Peer)
				switch p.PeerType {
				case model.PEER_SELF, model.PEER_USER:
					peerUsers = append(peerUsers, p.PeerId)
				case model.PEER_CHAT:
					peerChats = append(peerChats, p.PeerId)
				case model.PEER_CHANNEL:
					peerChannels = append(peerChannels, p.PeerId)
				default:
					err := mtproto.ErrPeerIdInvalid
					log.Errorf("messages.reorderPinnedDialogs - error: %v", err)
					return nil, err
				}
			default:
				err := mtproto.ErrPeerIdInvalid
				log.Errorf("messages.reorderPinnedDialogs - error: %v", err)
				return nil, err
			}
		}
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("messages.reorderPinnedDialogs - error: %v", err)
		return nil, err
	}

	s.PrivateFacade.ReorderPinnedDialogs(ctx, md.UserId, request.Force, request.FolderId, peerUsers)
	s.ChatFacade.ReorderPinnedDialogs(ctx, md.UserId, request.Force, request.FolderId, peerUsers)
	s.ChannelFacade.ReorderPinnedDialogs(ctx, md.UserId, request.Force, request.FolderId, peerUsers)

	log.Debugf("messages.reorderPinnedDialogs - reply {true}")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
		ctx := context.Background()

		syncUpdates := model.NewUpdatesLogic(md.UserId)
		updatePinnedDialogs := mtproto.MakeTLUpdatePinnedDialogs(&mtproto.Update{
			Order_FLAGVECTORDIALOGPEER: make([]*mtproto.DialogPeer, 0, len(peerUsers)+len(peerChats)+len(peerChannels)),
		}).To_Update()
		if request.FolderId != 0 {
			updatePinnedDialogs.FolderId = &types.Int32Value{Value: request.FolderId}
		}
		syncUpdates.AddUpdate(updatePinnedDialogs)

		for _, id := range peerUsers {
			updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER = append(updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
				Peer: model.MakePeerUser(id),
			}).To_DialogPeer())
		}
		users := s.GetUserListByIdList(ctx, md.UserId, peerUsers)
		syncUpdates.AddUsers(users)

		for _, id := range peerChats {
			updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER = append(updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
				Peer: model.MakePeerChat(id),
			}).To_DialogPeer())
		}
		chats := s.ChatFacade.GetChatListByIdList(ctx, md.UserId, peerChats)
		syncUpdates.AddChats(chats)

		for _, id := range peerChannels {
			updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER = append(updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
				Peer: model.MakePeerChannel(id),
			}).To_DialogPeer())
		}
		syncUpdates.AddChats(s.ChannelFacade.GetChannelListByIdList(ctx, md.UserId, peerChannels...))

		sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, syncUpdates.ToUpdates())
	}).(*mtproto.Bool), nil
}
