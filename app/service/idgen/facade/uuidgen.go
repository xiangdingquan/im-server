package id_facade

import "fmt"

type UUIDGen interface {
	GetUUID() (int64, error)
}

type UUIDGenInstance func() UUIDGen

var uuidGenAdapters = make(map[string]UUIDGenInstance)

func UUIDGenRegister(name string, adapter UUIDGenInstance) {
	if adapter == nil {
		panic("uuidgen: Register adapter is nil")
	}
	if _, ok := uuidGenAdapters[name]; ok {
		panic("uuidgen: Register called twice for adapter " + name)
	}
	uuidGenAdapters[name] = adapter
}

func NewUUIDGen(adapterName string) (adapter UUIDGen, err error) {
	instanceFunc, ok := uuidGenAdapters[adapterName]
	if !ok {
		err = fmt.Errorf("uuidgen: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	return
}
