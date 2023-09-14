package relay_client

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"

	"open.chat/app/interface/relay/relaypb"
	"open.chat/app/pkg/env2"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util/client"
)

var (
	_self relaypb.RPCRelayClient
)

func New() {
	if _self == nil {
		var (
			c   = &warden.ClientConfig{}
			err error
		)

		if err := paladin.Get("relay.toml").UnmarshalTOML(&c); err != nil {
			if err != paladin.ErrNotExist {
				panic(err)
			}
		}

		conn, err := client.NewClient(env2.InterfaceRelayId, c)
		if err != nil {
			panic(err)
		}

		_self = relaypb.NewRPCRelayClient(conn)
	}
}

func CreateCallSession(ctx context.Context, callId int64, opts ...grpc.CallOption) (*relaypb.CallConnections, error) {
	request := &relaypb.RelayCreateCallRequest{
		Id: callId,
	}
	return _self.RelayCreateCall(ctx, request, opts...)
}

func DiscardCallSession(ctx context.Context, callId int64, opts ...grpc.CallOption) (bool, error) {
	request := &relaypb.RelaydiscardCallRequest{
		Id: callId,
	}
	discarded, err := _self.RelayDiscardCall(ctx, request, opts...)
	return mtproto.FromBool(discarded), err
}
