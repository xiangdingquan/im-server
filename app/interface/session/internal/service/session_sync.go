package service

import (
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (c *session) onSyncData(obj mtproto.TLObject) {
	log.Debugf("onSyncData>> - session: %s, syncData: %s", c, logger.JsonDebugData(obj))

	if c.isAndroidPush {
		pusMsgId := c.authSessions.getNextNotifyId()
		c.sendPushToQueue(pusMsgId, androidPushTooLong)
	} else {
		pusMsgId := c.authSessions.getNextPushId()
		c.sendPushToQueue(pusMsgId, obj)
	}

	if c.sessionOnline() {
		gatewayId := c.getGatewayId()
		if gatewayId == "" {
			log.Errorf("gatewayId is empty, send delay...")
		} else {
			c.sendQueueToGateway(gatewayId)
		}
	}
}

func (c *session) onSyncRpcResultData(reqMsgId int64, data []byte) {
	c.pendingQueue.Remove(reqMsgId)
	c.sendPushRpcResultToQueue(reqMsgId, data)
}

func (c *session) onSyncSessionData(obj mtproto.TLObject) {
	log.Debugf("onSyncSessionData>> - session: %s, syncData: %s", c, logger.JsonDebugData(obj))
	pusMsgId := c.authSessions.getNextPushId()
	c.sendPushToQueue(pusMsgId, obj)
	c.sendQueueToGateway(c.getGatewayId())
}
