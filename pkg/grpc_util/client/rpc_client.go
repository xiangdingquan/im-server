package client

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/go-kratos/kratos/pkg/naming/etcd"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden/resolver"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func init() {
	resolver.Register(etcd.Builder(nil))
}

func NewClient(appId string, cfg *warden.ClientConfig, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(),
		"etcd://default/"+appId,
		warden.WithDialLogFlag(warden.LogFlagDisableInfo))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type RPCClient struct {
	conn *grpc.ClientConn
}

func NewRPCClient(appId string, cfg *warden.ClientConfig, opts ...grpc.DialOption) (*RPCClient, error) {
	conn, err := NewClient(appId, cfg, opts...)
	if err != nil {
		return nil, err
	}
	c := &RPCClient{
		conn: conn,
	}
	return c, nil
}

func (c *RPCClient) GetClientConn() *grpc.ClientConn {
	return c.conn
}

func (c *RPCClient) Invoke(rpcMetaData *grpc_util.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	t := mtproto.FindRPCContextTuple(object)
	if t == nil {
		err := fmt.Errorf("Invoke error: %v not regist!\n", object)
		log.Error(err.Error())
		return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
	}

	r := t.NewReplyFunc()
	var header, trailer metadata.MD

	ctxWithTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, _ := grpc_util.RpcMetadataToOutgoing(ctxWithTimeout, rpcMetaData)

	rt := time.Now()

	err := c.conn.Invoke(ctx, t.Method, object, r, grpc.Header(&header), grpc.Trailer(&trailer))
	log.Debugf("rpc Invoke: {method: %s, metadata: %s,  result: {%s}, error: {%v}}, cost = %v",
		t.Method,
		rpcMetaData.DebugString(),
		reflect.TypeOf(r),
		err,
		time.Since(rt))

	if err != nil {
		log.Errorf("RPC method: %s,  >> %v.Invoke(_) = _, %v: %#v", t.Method, c.conn, err, reflect.TypeOf(err))
		switch nErr := errors.Cause(err).(type) {
		case ecode.Codes:
			return nil, mtproto.MakeTLRpcError(&mtproto.RpcError{
				ErrorCode:    int32(nErr.Code()),
				ErrorMessage: nErr.Message(),
			})
		default:
			rpcErr := new(mtproto.TLRpcError)
			if err2 := jsonpb.UnmarshalString(err.Error(), rpcErr); err2 == nil {
				return nil, rpcErr
			} else {
				return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR - "+err.Error())
			}
		}
	} else {
		reply, ok := r.(mtproto.TLObject)

		if !ok {
			err = fmt.Errorf("Invalid reply type, maybe server side bug, %v\n", reply)
			return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		}

		return reply, nil
	}
}

func (c *RPCClient) InvokeContext(ctx2 context.Context, rpcMetaData *grpc_util.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	t := mtproto.FindRPCContextTuple(object)
	if t == nil {
		err := fmt.Errorf("Invoke error: %v not regist!\n", object)
		log.Error(err.Error())
		return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
	}

	r := t.NewReplyFunc()

	var header, trailer metadata.MD
	ctxWithTimeout, _ := context.WithTimeout(ctx2, 5*time.Second)
	ctx, _ := grpc_util.RpcMetadataToOutgoing(ctxWithTimeout, rpcMetaData)
	rt := time.Now()
	err := c.conn.Invoke(ctx, t.Method, object, r, grpc.Header(&header), grpc.Trailer(&trailer))
	log.Debugf("rpc Invoke: {method: %s, metadata: %s,  result: {%s}, error: {%v}}, cost = %v",
		t.Method,
		rpcMetaData.DebugString(),
		reflect.TypeOf(r),
		err,
		time.Since(rt))
	if err != nil {
		log.Errorf("RPC method: %s,  >> %v.Invoke(_) = _, %v: %#v", t.Method, c.conn, err, reflect.TypeOf(err))
		switch nErr := errors.Cause(err).(type) {
		case ecode.Codes:
			return nil, mtproto.MakeTLRpcError(&mtproto.RpcError{
				ErrorCode:    int32(nErr.Code()),
				ErrorMessage: nErr.Message(),
			})
		default:
			rpcErr := new(mtproto.TLRpcError)
			if err2 := jsonpb.UnmarshalString(err.Error(), rpcErr); err2 == nil {
				log.Debugf("%v", rpcErr)
				return nil, rpcErr
			} else {
				log.Debugf("error")
				return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR - "+err.Error())
			}
		}
	} else {
		reply, ok := r.(mtproto.TLObject)

		log.Debugf("Invoke %s time: %d", t.Method, time.Now().Unix()-rpcMetaData.ReceiveTime)

		if !ok {
			err = fmt.Errorf("Invalid reply type, maybe server side bug, %v\n", reply)
			log.Error(err.Error())
			return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		}

		return reply, nil
	}
}
