package session_client

import (
	"context"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
	"open.chat/app/interface/session/sessionpb"
)

func NewSessionRpcClient(target string, cfg *warden.ClientConfig, opts ...grpc.DialOption) (sessionpb.RPCSessionClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), target, warden.WithDialLogFlag(warden.LogFlagDisableArgs))
	if err != nil {
		return nil, err
	}
	return sessionpb.NewRPCSessionClient(cc), nil
}

func NewPushRpcClient(target string, cfg *warden.ClientConfig, opts ...grpc.DialOption) (sessionpb.RPCPushClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), target, warden.WithDialLogFlag(warden.LogFlagDisableArgs))
	if err != nil {
		return nil, err
	}
	return sessionpb.NewRPCPushClient(cc), nil
}
