package grpc_util

import (
	"encoding/base64"
	"fmt"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/metadata"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util/metautils"
	"open.chat/pkg/log"
)

var (
	headerRpcError = "rpc_error"
)

func RpcErrorFromMD(md metadata.MD) (rpcErr *mtproto.TLRpcError) {
	log.Debugf("rpc error from md: %s", md)
	val := metautils.NiceMD(md).Get(headerRpcError)
	if val == "" {
		rpcErr = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL),
			fmt.Sprintf("Unknown error"))
		log.Errorf("%v", rpcErr)
		return
	}

	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		rpcErr = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL),
			fmt.Sprintf("Base64 decode error, rpc_error: %s, error: %v", val, err))
		log.Errorf("%v", rpcErr)
		return
	}

	rpcErr = &mtproto.TLRpcError{}
	err = proto.Unmarshal(buf, rpcErr)
	if err != nil {
		rpcErr = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL),
			fmt.Sprintf("RpcError unmarshal error, rpc_error: %s, error: %v", val, err))
		log.Errorf("%v", rpcErr)
		return
	}

	return rpcErr
}

func RpcErrorToMD(md *mtproto.TLRpcError) (metadata.MD, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		log.Errorf("Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.Pairs(headerRpcError, base64.StdEncoding.EncodeToString(buf)), nil
}
