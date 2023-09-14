package authsession_client

import (
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/pkg/env2"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/pkg/grpc_util/client"
	"open.chat/pkg/log"
)

func New(c *warden.ClientConfig) (authsessionpb.RPCSessionClient, error) {
	conn, err := client.NewClient(env2.ServiceAuthSessionId, c)
	if err != nil {
		log.Errorf("fail to dial: %v", err)
		return nil, err
	}

	return authsessionpb.NewRPCSessionClient(conn), nil
}
