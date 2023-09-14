package id_facade

import (
	"context"
	"fmt"
)

type SeqIDGen interface {
	GetCurrentSeqID(ctx context.Context, key string) (int64, error)
	SetCurrentSeqID(ctx context.Context, key string, v int64) error
	GetNextSeqID(ctx context.Context, key string) (int64, error)
	GetNextNSeqID(ctx context.Context, key string, n int) (seq int64, err error)
	GetNextPhoneNumber(ctx context.Context, key string) (string, error)
}

type SeqIDGenInstance func() SeqIDGen

var seqIDGenAdapters = make(map[string]SeqIDGenInstance)

func SeqIDGenRegister(name string, adapter SeqIDGenInstance) {
	if adapter == nil {
		panic("seqidgen: Register adapter is nil")
	}
	if _, ok := seqIDGenAdapters[name]; ok {
		panic("seqidgen: Register called twice for adapter " + name)
	}
	seqIDGenAdapters[name] = adapter
}

func NewSeqIDGen(adapterName string) (adapter SeqIDGen, err error) {
	instanceFunc, ok := seqIDGenAdapters[adapterName]
	if !ok {
		err = fmt.Errorf("seqidgen: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	return
}
