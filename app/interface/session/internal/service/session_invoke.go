package service

import (
	"context"
	"reflect"

	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (c *session) onInvokeWithLayer(clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithLayer) {
	log.Debugf("onInvokeWithLayer - request data: {sess: %s, md: %s, msg: %s, req: {%s}}",
		c, clientIp, msgId.DebugString(), reflect.TypeOf(request))

	if request.GetQuery() == nil {
		log.Errorf("invokeWithLayer Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		log.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		log.Errorf("decode buf is nil, query: %v", query)
		return
	}

	initConnection, ok := query.(*mtproto.TLInitConnection)
	if !ok {
		log.Errorf("need initConnection, but query is : %v", query)
	}

	c.cb.setLayer(request.Layer)
	c.cb.setIpAddress(clientIp)
	c.cb.setClient(initConnection.LangPack)

	c.PutUploadInitConnection(context.Background(), c.cb.getAuthKeyId(), request.Layer, clientIp, initConnection)

	dBuf = mtproto.NewDecodeBuf(initConnection.GetQuery())
	query = dBuf.Object()
	if dBuf.GetError() != nil {
		log.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		log.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(clientIp, msgId, query)
}

func (c *session) onInvokeAfterMsg(clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeAfterMsg) {
	log.Debugf("onInvokeAfterMsg - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), reflect.TypeOf(request))

	if request.GetQuery() == nil {
		log.Errorf("invokeAfterMsg Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		log.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		log.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(clientIp, msgId, query)
}

func (c *session) onInvokeAfterMsgs(clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeAfterMsgs) {
	log.Debugf("onInvokeAfterMsgs - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), reflect.TypeOf(request))

	if request.GetQuery() == nil {
		log.Errorf("invokeAfterMsgs Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		log.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		log.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(clientIp, msgId, query)
}

func (c *session) onInvokeWithoutUpdates(clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithoutUpdates) {
	log.Debugf("onInvokeWithoutUpdates - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), reflect.TypeOf(request))

	if request.GetQuery() == nil {
		log.Errorf("invokeWithoutUpdates Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		log.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		log.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(clientIp, msgId, query)
}

func (c *session) onInvokeWithMessagesRange(clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithMessagesRange) {
	log.Debugf("onInvokeWithMessagesRange - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), reflect.TypeOf(request))

	if request.GetQuery() == nil {
		log.Errorf("invokeWithMessagesRange Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		log.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		log.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(clientIp, msgId, query)
}

func (c *session) onInvokeWithTakeout(clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithTakeout) {
	log.Debugf("onInvokeWithTakeout - request data: {sess: %s, msg: %s, req: {%s}}",
		c, msgId.DebugString(), reflect.TypeOf(request))

	if request.GetQuery() == nil {
		log.Errorf("invokeWithTakeout Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		log.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		log.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(clientIp, msgId, query)
}

func (c *session) onRpcRequest(clientIp string, msgId *inboxMsg, query mtproto.TLObject) bool {
	log.Debugf("onRpcRequest - request data: {sess: %s, md: %s, msg: %s, req: {%s}}",
		c, clientIp, msgId.DebugString(), reflect.TypeOf(query))
	switch q := query.(type) {
	case *mtproto.TLAccountRegisterDevice:
		if q.TokenType == 7 {
			pushSessionId, err := util.StringToUint64(q.GetToken())
			if err == nil {
				c.cb.onBindPushSessionId(int64(pushSessionId))
				c.PutCachePushSessionId(context.Background(), c.cb.getAuthKeyId(), int64(pushSessionId))
			}
		}
	case *mtproto.TLUpdatesGetState:
		if !c.isGeneric {
			c.isGeneric = true
			c.cb.setOnline()
		}
	case *mtproto.TLUpdatesGetDifference:
		if !c.isGeneric {
			c.isGeneric = true
			c.cb.setOnline()
		}
	case *mtproto.TLUpdatesGetChannelDifference:
		if !c.isGeneric {
			c.isGeneric = true
			c.cb.setOnline()
		}
	}

	if c.cb.getUserId() == 0 {
		if !checkRpcWithoutLogin(query) {
			authUserId, _ := c.GetCacheUserID(context.Background(), c.cb.getAuthKeyId())
			if authUserId == 0 {
				log.Errorf("not found authUserId by authKeyId: %d", c.cb.getAuthKeyId())
				rpcError := &mtproto.TLRpcError{Data2: &mtproto.RpcError{
					ErrorCode:    401,
					ErrorMessage: "AUTH_KEY_INVALID",
				}}
				c.sendRpcResultToQueue(msgId.msgId, rpcError)
				msgId.state = RECEIVED | RESPONSE_GENERATED
				return false
			} else {
				c.cb.setUserId(authUserId)
			}
		}
	}

	msgId.state = RECEIVED | RPC_PROCESSING
	c.cb.sendToRpcQueue(&rpcApiMessage{
		sessionId: c.sessionId,
		clientIp:  clientIp,
		reqMsgId:  msgId.msgId,
		reqMsg:    query,
	})

	return true
}

func (c *session) onRpcResult(rpcResult *rpcApiMessage) {
	log.Debugf("onRpcResult - result data: {sess: %s, reqMsgId: %d, reqMsg: {%s}, resMsg: {%s}}",
		c, rpcResult.reqMsgId, reflect.TypeOf(rpcResult.reqMsg), reflect.TypeOf(rpcResult.rpcResult.Result))
	defer func() {
		if _, ok := rpcResult.reqMsg.(*mtproto.TLAuthLogOut); ok {
			c.DeleteByAuthKeyId(c.cb.getAuthKeyId())
		}
	}()

	if rpcErr, ok := rpcResult.rpcResult.Result.(*mtproto.TLRpcError); ok {
		if rpcErr.GetErrorCode() == int32(mtproto.TLRpcErrorCodes_NOTRETURN_CLIENT) {
			log.Debugf("recv NOTRETURN_CLIENT")
			c.pendingQueue.Add(rpcResult.reqMsgId)
			return
		}
	}

	c.sendRpcResult(rpcResult.rpcResult)
}

func (c *session) sendRpcResult(rpcResult *mtproto.TLRpcResult) {
	log.Debugf("onRpcResult - result data: {sess: %s, reqMsgId: %d, resMsg: {%s}}",
		c, rpcResult.ReqMsgId, reflect.TypeOf(rpcResult.Result))
	msgId := c.inQueue.Lookup(rpcResult.ReqMsgId)
	if msgId == nil {
		log.Errorf("not found msgId, maybe removed: %d", rpcResult.ReqMsgId)
		return
	}
	c.sendRpcResultToQueue(msgId.msgId, rpcResult.Result)
	msgId.state = RECEIVED | ACKNOWLEDGED

	gatewayId := c.getGatewayId()
	if gatewayId == "" {
		log.Errorf("gatewayId is empty, send delay...")
	} else {
		c.sendQueueToGateway(gatewayId)
	}
}
