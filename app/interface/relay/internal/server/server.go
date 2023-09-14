package server

import (
	"context"
	"net"
	"sync"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/service/idgen/facade"
	_ "open.chat/app/service/idgen/facade/snowflake"
	"open.chat/pkg/log"
	"open.chat/pkg/net2"
)

type Config struct {
	RelayIp     string
	RelayPort   int32
	RelayServer net2.ServerConfig
	SendBuf     int
	ReceiveBuf  int
}

type udpDataBuf struct {
	addr *net.UDPAddr
	b    []byte
}

type Server struct {
	c          *Config
	sendMutex  sync.RWMutex
	sendChan   chan udpDataBuf
	closeChan  chan int
	relayConn  *net.UDPConn
	running    bool
	tableMutex sync.RWMutex
	relayTable map[string]*RelayTable
	idTable    map[int64]string
	uuidGen    id_facade.UUIDGen
}

func New() *Server {
	var (
		ac  = &Config{}
		err error
		s   = new(Server)
	)

	if err := paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	log.Info("config: %#v", ac)

	if ac.ReceiveBuf == 0 {
		ac.ReceiveBuf = 4096
	}
	if ac.SendBuf == 0 {
		ac.SendBuf = 4096
	}

	s.c = ac
	s.relayTable = map[string]*RelayTable{}
	s.idTable = map[int64]string{}
	s.sendChan = make(chan udpDataBuf, 1024)
	s.closeChan = make(chan int, 1)
	s.relayConn, err = listenUdp(ac)
	if err != nil {
		panic(err)
	}

	s.uuidGen, _ = id_facade.NewUUIDGen("snowflake")

	go s.sendLoop()
	go s.readLoop()

	return s
}

func (s *Server) Close() {
	log.Infof("Close...")
}

func (s *Server) Ping(ctx context.Context) (err error) {
	return nil
}

func (s *Server) GetConnectionID() (id int64) {
	id, _ = s.uuidGen.GetUUID()
	return
}
