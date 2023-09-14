package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on [%s] (multi-cores: %t, loops: %d)\n",
		srv.AddrsString(), srv.Multicore, srv.NumEventLoop)
	return
}

func (es *echoServer) React(frame interface{}, c gnet.Conn) (out interface{}, action gnet.Action) {
	// Echo synchronously.
	out = frame
	return

	/*
		// Echo asynchronously.
		data := append([]byte{}, frame...)
		go func() {
			time.Sleep(time.Second)
			c.AsyncWrite(data)
		}()
		return
	*/
}

func main() {
	var port, port2 int
	var multicore, reuseport bool

	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 10443, "--port 10443")
	flag.IntVar(&port2, "port2", 10444, "--port2 10444")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.BoolVar(&reuseport, "reuseport", false, "--reuseport true")
	flag.Parse()
	echo := new(echoServer)
	log.Fatal(gnet.Serve(echo,
		[]string{fmt.Sprintf("tcp://:%d", port), fmt.Sprintf("tcp://:%d", port2)},
		gnet.WithMulticore(multicore),
		gnet.WithReusePort(reuseport),
		gnet.WithReadBuffer(4096),
		gnet.WithWriteBuffer(4096)))
}
