package service

import (
	"context"
	"fmt"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/pkg/naming"
	"open.chat/app/interface/session/client"
	"open.chat/app/interface/session/sessionpb"
	"open.chat/pkg/log"
)

type SessionOptions struct {
	RoutineSize uint64
	RoutineChan uint64
}

type Session struct {
	serverID       string
	client         sessionpb.RPCPushClient
	sessionChan    []chan interface{}
	sessionChanNum uint64
	options        SessionOptions
	ctx            context.Context
	cancel         context.CancelFunc
}

// process
func (c *Session) process(sessionChan chan interface{}) {
	var err error
	for {
		select {
		case updates, ok := <-sessionChan:
			if !ok {
				log.Error("process error")
				return
			}

			switch r := updates.(type) {
			case *sessionpb.PushSessionUpdatesData:
				_, err = c.client.PushSessionUpdates(context.Background(), r)
				if err != nil {
					log.Error("c.client.PushSessionUpdates(%s, %v, reply) serverId:%d error(%v)", r, c.serverID, err)
				}
			case *sessionpb.PushUpdatesData:
				_, err = c.client.PushUpdates(context.Background(), r)
				if err != nil {
					log.Error("c.client.PushUpdates(%s, %v, reply) serverId:%d error(%v)", r, c.serverID, err)
				}
			case *sessionpb.PushRpcResultData:
				_, err = c.client.PushRpcResult(context.Background(), r)
				if err != nil {
					log.Error("c.client.PushRpcResult(%s, %v, reply) serverId:%d error(%v)", r, c.serverID, err)
				}
			default:
				log.Errorf("")
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Session) Close() (err error) {
	finish := make(chan bool)
	go func() {
		for {
			n := len(c.sessionChan)
			for _, ch := range c.sessionChan {
				n += len(ch)
			}
			if n == 0 {
				finish <- true
				return
			}
			time.Sleep(time.Second)
		}
	}()
	select {
	case <-finish:
		log.Info("close gateway finish")
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("close gateway(server:%s push:%d) timeout", c.serverID, len(c.sessionChan))
	}
	c.cancel()
	return
}

func (c *Session) PushUpdates(ctx context.Context, msg *sessionpb.PushUpdatesData) (err error) {
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- msg
	return
}

func (c *Session) PushSessionUpdates(ctx context.Context, msg *sessionpb.PushSessionUpdatesData) (err error) {
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- msg
	return
}

func (c *Session) PushRpcResult(ctx context.Context, msg *sessionpb.PushRpcResultData) (err error) {
	idx := atomic.AddUint64(&c.sessionChanNum, 1) % c.options.RoutineSize
	c.sessionChan[idx] <- msg
	return
}

func NewSession(data *naming.Instance, conf *Config, options SessionOptions) (*Session, error) {
	c := &Session{
		serverID:    data.Hostname,
		sessionChan: make([]chan interface{}, options.RoutineSize),
		options:     options,
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
	if c.client, err = session_client.NewPushRpcClient(grpcAddr, conf.RPC); err != nil {
		return nil, err
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())

	for i := uint64(0); i < options.RoutineSize; i++ {
		c.sessionChan[i] = make(chan interface{}, options.RoutineChan)
		go c.process(c.sessionChan[i])
	}
	return c, nil
}
