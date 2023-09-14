package service

import (
	"container/list"
	"math/rand"
	"sync"
	"time"
)

type botSessionManager struct {
	mutex    sync.Mutex
	sessions map[int64]*cacheBotSession
	reqCache *RequestManager
}

func NewBotSessionManager() *botSessionManager {
	return &botSessionManager{
		sessions: make(map[int64]*cacheBotSession),
		reqCache: NewRequestManager(),
	}
}

func (c *botSessionManager) Get(id int64) *cacheBotSession {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.sessions[id]
}

func (c *botSessionManager) Put(s *cacheBotSession) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.sessions[s.authKeyId] = s
}

type cacheBotSession struct {
	userId       int32
	authKeyId    int64
	sessionId    int64
	layer        int32
	mutex        sync.Mutex
	lastPingTime int64
	updatesList  *list.List
}

func newCacheBotSession(userId int32, authKeyId int64, layer int32) *cacheBotSession {
	return &cacheBotSession{
		userId:       userId,
		authKeyId:    authKeyId,
		sessionId:    rand.Int63(),
		layer:        layer,
		lastPingTime: 0,
		updatesList:  list.New(),
	}
}

func (cv *cacheBotSession) Size() int {
	return 1
}

func (cv *cacheBotSession) UserId() int32 {
	return cv.userId
}

func (cv *cacheBotSession) AuthKeyId() int64 {
	return cv.authKeyId
}

func (cv *cacheBotSession) Layer() int32 {
	return cv.layer
}

func (cv *cacheBotSession) SessionId() int64 {
	return cv.sessionId
}

func (cv *cacheBotSession) FetchUpdates(walk func(v interface{})) bool {
	cv.mutex.Lock()
	defer cv.mutex.Unlock()

	if cv.updatesList.Len() > 0 {
		for e := cv.updatesList.Front(); e != nil; e = e.Next() {
			if walk != nil {
				walk(e.Value)
			}
		}
		cv.updatesList.Init()
		return true
	}

	return false
}

func (cv *cacheBotSession) PushUpdates(v interface{}) {
	cv.mutex.Lock()
	defer cv.mutex.Unlock()

	cv.updatesList.PushBack(v)
}

func (cv *cacheBotSession) TrySetLastTime(timeout int64) bool {
	cv.mutex.Lock()
	defer cv.mutex.Unlock()

	date := time.Now().Unix()
	if date >= cv.lastPingTime+timeout {
		cv.lastPingTime = date
		return true
	}

	return false
}
