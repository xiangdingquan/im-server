package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetDialogUnreadMarks(ctx context.Context, request *mtproto.TLMessagesGetDialogUnreadMarks) (*mtproto.Vector_DialogPeer, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getDialogUnreadMarks - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getDialogUnreadMarks - error: %v", err)
		return nil, err
	}

	var (
		dialogPeerList = &mtproto.Vector_DialogPeer{Datas: []*mtproto.DialogPeer{}}
	)

	dialogPeerList.Datas = append(dialogPeerList.Datas, s.PrivateFacade.GetDialogUnreadMarkList(ctx, md.UserId)...)
	dialogPeerList.Datas = append(dialogPeerList.Datas, s.ChatFacade.GetDialogUnreadMarkList(ctx, md.UserId)...)
	dialogPeerList.Datas = append(dialogPeerList.Datas, s.ChannelFacade.GetDialogUnreadMarkList(ctx, md.UserId)...)

	log.Debugf("messages.getDialogUnreadMarks#22e24e22 - reply: %s", dialogPeerList.DebugString())
	return dialogPeerList, nil
}
