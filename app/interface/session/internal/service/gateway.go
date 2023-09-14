package service

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-kratos/kratos/pkg/naming"

	"open.chat/app/interface/gateway/client"
	"open.chat/app/interface/gateway/egatepb"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type gatewayOptions struct {
	RoutineSize uint64
	RoutineChan uint64
}

type Gateway struct {
	serverID string
	client   egatepb.EGateClient
}

func (c *Gateway) Close() (err error) {
	return
}

func (c *Gateway) SendDataToGate(ctx context.Context, authKeyId, sessionId int64, payload []byte) (b bool, err error) {
	var (
		res *mtproto.Bool
	)

	res, err = c.client.ReceiveData(context.Background(), &egatepb.SessionRawData{
		AuthKeyId: authKeyId,
		SessionId: sessionId,
		Payload:   payload,
	})

	if err != nil {
		log.Errorf("sendDataToGate error: %v", err)
		b = false
		return
	}

	b = mtproto.FromBool(res)
	return
}

func NewGateway(data *naming.Instance, conf *Config, options gatewayOptions) (*Gateway, error) {
	c := &Gateway{
		serverID: data.Hostname,
	}
	var grpcAddr string
	for _, addrs := range data.Addrs {
		u, err := url.Parse(addrs)
		if err == nil && u.Scheme == "grpc" {
			grpcAddr = u.Host
		}
	}
	if grpcAddr == "" {
		return nil, fmt.Errorf("invalid grpc address:%v", data.Addrs)
	}
	var err error
	if c.client, err = gateway_client.NewClient(grpcAddr, conf.WardenClient); err != nil {
		return nil, err
	}
	return c, nil
}
