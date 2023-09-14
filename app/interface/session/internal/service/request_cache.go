package service

import (
	"errors"
	"fmt"
	"sync"

	"open.chat/mtproto"
)

const mapNum = 8

type requestMap struct {
	mutex    sync.Mutex
	respChan map[int64]chan mtproto.TLObject
}

type RequestManager struct {
	requestMaps [mapNum]requestMap
	current     uint32
}

func NewRequestManager() *RequestManager {
	manager := &RequestManager{}
	for i := 0; i < mapNum; i++ {
		manager.requestMaps[i].respChan = make(map[int64]chan mtproto.TLObject)
	}

	return manager
}

func (r *RequestManager) cache(id int64, msg chan mtproto.TLObject) {
	m := &r.requestMaps[uint64(id)%mapNum]
	m.mutex.Lock()
	m.respChan[id] = msg
	m.mutex.Unlock()
}

func (r *RequestManager) shoot(id int64, msg mtproto.TLObject) (err error) {
	m := &r.requestMaps[uint64(id)%mapNum]

	m.mutex.Lock()
	respChans, exist := m.respChan[id]
	if exist {
		delete(m.respChan, id)
		m.mutex.Unlock()
		select {
		case respChans <- msg:
			close(respChans)
		default:
			err = errors.New(fmt.Sprint("Default fail !!!!! request id : ", id))
		}
	} else {
		m.mutex.Unlock()
		err = errors.New("default fail ")
	}
	return
}

func (r *RequestManager) dispose(id int64) {
	m := &r.requestMaps[uint64(id)%mapNum]
	m.mutex.Lock()
	respChans, exist := m.respChan[id]
	if exist {
		delete(m.respChan, id)
		m.mutex.Unlock()
		close(respChans)
	} else {
		m.mutex.Unlock()
	}
}
