package service

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/gorilla/websocket"

	"open.chat/pkg/log"
)

var (
	addr  string
	faddr string
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    65536,
	WriteBufferSize:   65536,
	Subprotocols:      []string{"binary"},
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	flag.StringVar(&addr, "addr", "0.0.0.0:10443", "http service address")
	flag.StringVar(&faddr, "faddr", "127.0.0.1:12443", "forward server address")
}

type Service struct {
	c *Config
}

func New() (s *Service) {
	var (
		ac  = &Config{}
		err error
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s = &Service{
		c: ac,
	}

	go func() {
		http.HandleFunc("/apiws", apiws)
		http.ListenAndServe(addr, nil)
	}()

	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return nil
}

// Close close the resource.
func (s *Service) Close() {
}

func apiws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("upgrade: %v", err)
		return
	}
	defer c.Close()

	dconn, err := net.Dial("tcp", faddr)
	if err != nil {
		log.Errorf("connect forward server error: %v", err)
		return
	}

	exitChan := make(chan bool, 1)
	go func(sconn *websocket.Conn, dconn net.Conn, Exit chan bool) {
		for {
			_, message, err := sconn.ReadMessage()
			if err != nil {
				log.Errorf("recv websocket data error: %v", err)
				break
			}
			if _, err := dconn.Write(message); err != nil {
				log.Errorf("forward to server data error: %v", err)
				break
			}
		}
		exitChan <- true
	}(c, dconn, exitChan)

	go func(sconn *websocket.Conn, dconn net.Conn, Exit chan bool) {
		var (
			bytes = make([]byte, 4096)
			n     int
			err   error
		)

		for {
			if n, err = dconn.Read(bytes); err != nil {
				log.Errorf("recv forward data error: %v", err)
				break
			}

			if err = sconn.WriteMessage(websocket.BinaryMessage, bytes[:n]); err != nil {
				log.Errorf("forward to websocket client data error: %v", err)
				break
			}
		}
		exitChan <- true
	}(c, dconn, exitChan)

	<-exitChan
	dconn.Close()
}
