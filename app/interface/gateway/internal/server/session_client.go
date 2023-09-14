package server

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-kratos/kratos/pkg/naming"

	"open.chat/app/interface/session/client"
	"open.chat/app/interface/session/sessionpb"
	"open.chat/mtproto"
)

type sessionClient struct {
	serverID string
	client   sessionpb.RPCSessionClient
	ctx      context.Context
	cancel   context.CancelFunc
}

func newSession(data *naming.Instance, conf *Config) (*sessionClient, error) {
	c := &sessionClient{
		serverID: data.Hostname,
	}

	var gRpcAddr string
	for _, addr := range data.Addrs {
		u, err := url.Parse(addr)
		if err == nil && u.Scheme == "grpc" {
			gRpcAddr = u.Host
		}
	}
	if gRpcAddr == "" {
		return nil, fmt.Errorf("invalid grpc address:%v", data.Addrs)
	}
	var err error
	if c.client, err = session_client.NewSessionRpcClient(gRpcAddr, conf.WardenClient); err != nil {
		return nil, err
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())

	return c, nil
}

func (c *sessionClient) Close() (err error) {
	c.cancel()
	return
}

func (c *sessionClient) CreateSession(ctx context.Context, r *sessionpb.SessionClientEvent) (*mtproto.Bool, error) {
	return c.client.CreateSession(ctx, r)
}

func (c *sessionClient) CloseSession(ctx context.Context, r *sessionpb.SessionClientEvent) (*mtproto.Bool, error) {
	return c.client.CloseSession(ctx, r)
}

func (c *sessionClient) SendAsyncDataToSession(ctx context.Context, r *sessionpb.SessionClientData) (*mtproto.Bool, error) {
	return c.client.SendAsyncSessionData(ctx, r)
}

func (c *sessionClient) SendSyncDataToSession(ctx context.Context, r *sessionpb.SessionClientData) (*sessionpb.SessionData, error) {
	return c.client.SendSyncSessionData(ctx, r)
}

func (c *sessionClient) Heartbeat(ctx context.Context, r *sessionpb.SessionClientEvent) (*mtproto.Bool, error) {
	return c.client.Heartbeat(ctx, r)
}

func (c *sessionClient) DestroySessions(ctx context.Context, authKeyId int64) (*mtproto.Bool, error) {
	return c.client.DestroySessions(ctx, &sessionpb.AuthId{AuthKeyId: authKeyId})
}

func (c *sessionClient) ImportSessions(ctx context.Context, r *sessionpb.AuthSessionIdList) (*mtproto.Bool, error) {
	return c.client.ImportSessions(ctx, r)
}
