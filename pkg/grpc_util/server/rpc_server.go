package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/env"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/naming"
	"github.com/go-kratos/kratos/pkg/naming/etcd"
	"github.com/go-kratos/kratos/pkg/net/ip"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

var (
	//etcdPrefix is a etcd globe key prefix
	endpoints string
)

func init() {
	endpoints = os.Getenv("ETCD_ENDPOINTS")
	if endpoints == "" {
		panic(fmt.Errorf("invalid etcd config endpoints:%+v", endpoints))
	}
}

type RPCServer struct {
	ws     *warden.Server
	cancel context.CancelFunc
}

type RegisterRPCServerFunc func(s *grpc.Server)

func NewRpcServer(appId string, regFunc RegisterRPCServerFunc) *RPCServer {
	var rc struct {
		Server *warden.ServerConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	return NewRpcServer2(appId, rc.Server, regFunc)
}

func NewRpcServer2(appId string, server *warden.ServerConfig, regFunc RegisterRPCServerFunc) *RPCServer {
	_, port, _ := net.SplitHostPort(server.Addr)

	ws := warden.NewServer(server)
	ws.Use(errorHandler())
	if regFunc != nil {
		regFunc(ws.Server())
	}

	ws, err := ws.Start()
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

	// register discover
	dis, err := etcd.New(&clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: time.Second * 3,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

	var (
		cancel context.CancelFunc
	)

	ipAddr := ip.InternalIP()
	ins := &naming.Instance{
		Zone:     env.Zone,
		Env:      env.DeployEnv,
		Hostname: env.Hostname,
		AppID:    appId, // AppID:    "session.service.auth_session",
		Addrs: []string{
			"grpc://" + ipAddr + ":" + port,
		},
	}
	cancel, err = dis.Register(context.Background(), ins)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return &RPCServer{ws: ws, cancel: cancel}
}

func (s *RPCServer) Stop() {
	if s.cancel != nil {
		s.cancel()
	}
	if err := s.ws.Shutdown(context.Background()); err != nil {
		log.Errorf("grpcSrv.Shutdown error(%v)", err)
	}
}

func errorHandler() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			switch err.(type) {
			case *mtproto.TLRpcError:
				md, _ := grpc_util.RpcErrorToMD(err.(*mtproto.TLRpcError))
				grpc.SetTrailer(ctx, md)
			}
		}
		return
	}
}
