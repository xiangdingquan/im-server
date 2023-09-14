package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

const (
	supportUserID      = int32(424000)
	supportPhoneNumber = "42400"
	supportName        = "Volunteer Support"
)

var (
	supportUser = mtproto.MakeTLUser(&mtproto.User{
		Self:           false,
		Contact:        false,
		MutualContact:  false,
		Deleted:        false,
		Bot:            true,
		BotChatHistory: false,
		BotNochats:     true,
		Verified:       false,
		Restricted:     false,
		Min:            false,
		BotInlineGeo:   false,
		Support:        true,
		Scam:           false,
		Id:             supportUserID,
		AccessHash:     &types.Int64Value{Value: 6599886787491911852},
		FirstName:      &types.StringValue{Value: supportName},
		LastName:       nil,
		Username:       nil,
		Phone:          nil,
		Photo:          nil,
		Status:         nil,
		BotInfoVersion: &types.Int32Value{Value: 0},
		RestrictionReason_FLAGVECTORRESTRICTIONREASON: nil,
		RestrictionReason_FLAGSTRING:                  nil,
		BotInlinePlaceholder:                          nil,
		LangCode:                                      nil,
	}).To_User()
)

func (s *Service) HelpGetSupport(ctx context.Context, request *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getSupport - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getSupport - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLHelpSupport(&mtproto.Help_Support{
		PhoneNumber: supportPhoneNumber,
		User:        supportUser,
	}).To_Help_Support()

	log.Debugf("help.getSupport - reply: {%v}\n", reply)
	return reply, nil
}
