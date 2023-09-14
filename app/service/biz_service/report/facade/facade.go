package report_facade

import (
	"context"
	"fmt"
)

const (
	MinUsernameLen = 5
	MaxUsernameLen = 32
)

type ReportFacade interface {
	Report(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId, messageId, reason int32, text string) (bool, error)
	ReportIdList(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId int32, messageIdList []int32, reason int32, text string) (bool, error)
}

type Instance func() ReportFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewReportFacade(name string) (inst ReportFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
