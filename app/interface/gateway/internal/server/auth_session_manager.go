package server

import (
	"container/list"
	"sync"

	"open.chat/pkg/log"
)

type session struct {
	sessionId           int64
	isWebsocket         bool
	connIdList          *list.List
	pendingHttpDataList *list.List
}

type authSessionManager struct {
	rw       sync.RWMutex
	sessions map[int64]map[int64]session
}

func NewAuthSessionManager() *authSessionManager {
	return &authSessionManager{
		sessions: make(map[int64]map[int64]session),
	}
}

func (m *authSessionManager) PushBackHttpData(authKeyId, sessionId int64, raw []byte) {
	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if v2, ok2 := v[sessionId]; ok2 {
			v2.pendingHttpDataList.PushFront(raw)
		}
	}
}

func (m *authSessionManager) PopFrontHttpData(authKeyId, sessionId int64) []byte {
	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if v2, ok2 := v[sessionId]; ok2 {
			if e := v2.pendingHttpDataList.Front(); e != nil {
				v2.pendingHttpDataList.Remove(e)
				return e.Value.([]byte)
			}
		}
	}
	return nil
}

func (m *authSessionManager) AddNewSession(authKeyId, sessionId int64, isWebsocket bool, connId uint64) (bNew bool) {
	log.Infof("addNewSession: auth_key_id: %d, session_id: %d, conn_id: %d",
		authKeyId,
		sessionId,
		connId)

	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKeyId]; ok {
		var (
			cExisted = false
		)
		if v2, ok2 := v[sessionId]; ok2 {
			for e := v2.connIdList.Front(); e != nil; e = e.Next() {
				if e.Value.(uint64) == connId {
					cExisted = true
					break
				}
			}
			if !cExisted {
				v2.connIdList.PushBack(connId)
			}
		} else {
			s := session{
				sessionId:           sessionId,
				isWebsocket:         isWebsocket,
				connIdList:          list.New(),
				pendingHttpDataList: list.New(),
			}
			s.connIdList.PushBack(connId)
			v[sessionId] = s
			bNew = true
		}
	} else {
		s := session{
			sessionId:           sessionId,
			isWebsocket:         isWebsocket,
			connIdList:          list.New(),
			pendingHttpDataList: list.New(),
		}
		s.connIdList.PushBack(connId)

		m.sessions[authKeyId] = map[int64]session{
			sessionId: s,
		}
		bNew = true
	}
	return
}

func (m *authSessionManager) RemoveSession(authKeyId, sessionId int64, connId uint64) (bDeleted bool) {
	log.Infof("removeSession: auth_key_id: %d, session_id: %d, conn_id: %d",
		authKeyId,
		sessionId,
		connId)

	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if v2, ok2 := v[sessionId]; ok2 {
			for e := v2.connIdList.Front(); e != nil; e = e.Next() {
				if e.Value.(uint64) == connId {
					v2.connIdList.Remove(e)
					break
				}
			}
			if v2.connIdList.Len() == 0 {
				delete(v, sessionId)
				bDeleted = true
			}
			if len(v) == 0 {
				delete(m.sessions, authKeyId)
			}
		}
	}

	return
}

func (m *authSessionManager) FoundSessionConnIdList(authKeyId, sessionId int64) (bool, []uint64) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if v2, ok2 := v[sessionId]; ok2 {
			connIdList := make([]uint64, 0, v2.connIdList.Len())
			for e := v2.connIdList.Back(); e != nil; e = e.Prev() {
				connIdList = append(connIdList, e.Value.(uint64))
			}
			return v2.isWebsocket, connIdList
		}
	}

	return false, nil
}
