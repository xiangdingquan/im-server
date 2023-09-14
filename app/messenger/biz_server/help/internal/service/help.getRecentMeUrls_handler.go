package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetRecentMeUrls(ctx context.Context, request *mtproto.TLHelpGetRecentMeUrls) (*mtproto.Help_RecentMeUrls, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getRecentMeUrls - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getRecentMeUrls - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLHelpRecentMeUrls(&mtproto.Help_RecentMeUrls{
		Urls:  []*mtproto.RecentMeUrl{},
		Users: []*mtproto.User{},
		Chats: []*mtproto.Chat{},
	}).To_Help_RecentMeUrls()

	log.Debugf("help.getRecentMeUrls - reply: %s", reply.DebugString())
	return reply, nil
}
