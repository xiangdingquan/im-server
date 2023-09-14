package service

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetConfig(ctx context.Context, request *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getConfig#c4f9186b - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	CONNECTION_DEVICE_MODEL_EMPTY	Device model empty
	// -503	Timeout	Timeout while fetching data

	helpConfig, _ := proto.Clone(&config).(*mtproto.TLConfig)
	now := int32(time.Now().Unix())
	helpConfig.SetDate(now)
	helpConfig.SetExpires(now + expiresTimeout)

	if md.GetLayer() >= 1202 {
		helpConfig.SetDcOptions(nil)
	}

	reply := helpConfig.To_Config()

	log.Debugf("help.getConfig#c4f9186b - reply: %s", reply.DebugString())
	return reply, nil
}
