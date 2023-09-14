package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsCheckUsername(ctx context.Context, request *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.checkUsername#10e6bd2c - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.checkUsername - error: %v", err)
		return nil, err
	}

	if request.Username == ".bad." {
		return nil, mtproto.ErrUsernameInvalid
	}

	if !model.CheckUsernameInvalid(request.Username) {
		return nil, mtproto.ErrUsernameInvalid
	}

	switch request.Channel.GetPredicateName() {
	case mtproto.Predicate_inputChannelEmpty:
		checked, err := s.UsernameFacade.CheckUsername(ctx, request.Username)
		if err != nil {
			log.Errorf("unknown error: %v", err)
			return nil, err
		}

		if checked == model.UsernameNotExisted {
			log.Debugf("channels.checkUsername - reply: {true}")
			return mtproto.BoolTrue, nil
		} else {
			log.Debugf("channels.checkUsername - reply: {false}")
			return mtproto.BoolFalse, nil
		}
	case mtproto.Predicate_inputChannel:
		checked, err := s.UsernameFacade.CheckChannelUsername(ctx, request.Channel.ChannelId, request.Username)
		if err != nil {
			log.Errorf("channels.checkUsername - error: %v", err)
			return nil, err
		}

		if checked == model.UsernameExistedNotMe {
			log.Debugf("channels.checkUsername - reply: {false}")
			return mtproto.BoolFalse, nil
		} else {
			log.Debugf("channels.checkUsername - reply: {true}")
			return mtproto.BoolTrue, nil
		}
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("channels.checkUsername - error: %v", err)
		return nil, err
	}
}
