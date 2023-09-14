// Copyright (c) 2019 Andy Pan
// Copyright (c) 2018 Joshua J Baker
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

//go:build linux || freebsd || dragonfly || darwin
// +build linux freebsd dragonfly darwin

package gnet

import (
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/panjf2000/gnet/errors"
	"github.com/panjf2000/gnet/internal/logging"
	"github.com/panjf2000/gnet/internal/netpoll"
)

type server struct {
	lns             []*listener        // all the listeners
	wg              sync.WaitGroup     // event-loop close WaitGroup
	opts            *Options           // options with server
	once            sync.Once          // make sure only signalShutdown once
	cond            *sync.Cond         // shutdown signaler
	codec           ICodec             // codec for TCP stream
	logger          logging.Logger     // customized logger for logging info
	ticktock        chan time.Duration // ticker channel
	mainLoop        *eventloop         // main event-loop for accepting connections
	eventHandler    EventHandler       // user eventHandler
	subEventLoopSet loadBalancer       // event-loops for handling events
}

// waitForShutdown waits for a signal to shutdown.
func (svr *server) waitForShutdown() {
	svr.cond.L.Lock()
	svr.cond.Wait()
	svr.cond.L.Unlock()
}

// signalShutdown signals a shutdown an begins server closing.
func (svr *server) signalShutdown() {
	svr.once.Do(func() {
		svr.cond.L.Lock()
		svr.cond.Signal()
		svr.cond.L.Unlock()
	})
}

func (svr *server) startEventLoops() {
	svr.subEventLoopSet.iterate(func(i int, el *eventloop) bool {
		svr.wg.Add(1)
		go func() {
			el.loopRun()
			svr.wg.Done()
		}()
		return true
	})
}

func (svr *server) closeEventLoops() {
	svr.subEventLoopSet.iterate(func(i int, el *eventloop) bool {
		_ = el.poller.Close()
		return true
	})
}

func (svr *server) startSubReactors() {
	svr.subEventLoopSet.iterate(func(i int, el *eventloop) bool {
		svr.wg.Add(1)
		go func() {
			svr.activateSubReactor(el)
			svr.wg.Done()
		}()
		return true
	})
}

func (svr *server) activateEventLoops(numEventLoop int) (err error) {
	// Create loops locally and bind the listeners.
	for i := 0; i < numEventLoop; i++ {
		lns := svr.lns
		if i > 0 && svr.opts.ReusePort {
			lns = make([]*listener, len(svr.lns))
			for j := 0; j < len(lns); j++ {
				if lns[j], err = initListener(svr.lns[j].network, svr.lns[j].addr, svr.lns[j].reusePort); err != nil {
					log.Println(err)
					return
				}
			}
		}

		var p *netpoll.Poller
		if p, err = netpoll.OpenPoller(); err == nil {
			el := &eventloop{
				lns:               lns,
				svr:               svr,
				codec:             svr.codec,
				poller:            p,
				packet:            make([]byte, 0x10000),
				connections:       make(map[int]*conn),
				eventHandler:      svr.eventHandler,
				calibrateCallback: svr.subEventLoopSet.calibrate,
			}
			for _, ln := range el.lns {
				_ = el.poller.AddRead(ln.fd)
			}
			svr.subEventLoopSet.register(el)
		} else {
			return err
		}
	}

	// Start event-loops in background.
	svr.startEventLoops()

	return nil
}

func (svr *server) activateReactors(numEventLoop int) error {
	for i := 0; i < numEventLoop; i++ {
		if p, err := netpoll.OpenPoller(); err == nil {
			el := &eventloop{
				lns:               svr.lns,
				svr:               svr,
				codec:             svr.codec,
				poller:            p,
				packet:            make([]byte, 0x10000),
				connections:       make(map[int]*conn),
				eventHandler:      svr.eventHandler,
				calibrateCallback: svr.subEventLoopSet.calibrate,
			}
			svr.subEventLoopSet.register(el)
		} else {
			return err
		}
	}

	// Start sub reactors in background.
	svr.startSubReactors()

	if p, err := netpoll.OpenPoller(); err == nil {
		el := &eventloop{
			lns:    svr.lns,
			idx:    -1,
			poller: p,
			svr:    svr,
		}
		for _, ln := range el.lns {
			_ = el.poller.AddRead(ln.fd)
		}
		svr.mainLoop = el

		// Start main reactor in background.
		svr.wg.Add(1)
		go func() {
			svr.activateMainReactor()
			svr.wg.Done()
		}()
	} else {
		return err
	}
	return nil
}

func (svr *server) start(numEventLoop int) error {
	if svr.opts.ReusePort || isUdp(svr.lns) {
		return svr.activateEventLoops(numEventLoop)
	}
	return svr.activateReactors(numEventLoop)
}

func (svr *server) stop() {
	// Wait on a signal for shutdown
	svr.waitForShutdown()

	// Notify all loops to close by closing all listeners
	svr.subEventLoopSet.iterate(func(i int, el *eventloop) bool {
		sniffErrorAndLog(el.poller.Trigger(func() error {
			return errors.ErrServerShutdown
		}))
		return true
	})

	if svr.mainLoop != nil {
		for _, ln := range svr.lns {
			ln.close()
		}
		sniffErrorAndLog(svr.mainLoop.poller.Trigger(func() error {
			return errors.ErrServerShutdown
		}))
	}

	// Wait on all loops to complete reading events
	svr.wg.Wait()

	svr.closeEventLoops()

	if svr.mainLoop != nil {
		sniffErrorAndLog(svr.mainLoop.poller.Close())
	}
}

func serve(eventHandler EventHandler, listeners []*listener, options *Options) error {
	// Figure out the correct number of loops/goroutines to use.
	numEventLoop := 1
	if options.Multicore {
		numEventLoop = runtime.NumCPU()
	}
	if options.NumEventLoop > 0 {
		numEventLoop = options.NumEventLoop
	}

	svr := new(server)
	svr.opts = options
	svr.eventHandler = eventHandler
	svr.lns = listeners

	switch options.LB {
	case RoundRobin:
		svr.subEventLoopSet = new(roundRobinEventLoopSet)
	case LeastConnections:
		svr.subEventLoopSet = new(leastConnectionsEventLoopSet)
	case SourceAddrHash:
		svr.subEventLoopSet = new(sourceAddrHashEventLoopSet)
	}

	svr.cond = sync.NewCond(&sync.Mutex{})
	svr.ticktock = make(chan time.Duration, 1)
	svr.logger = logging.DefaultLogger
	svr.codec = func() ICodec {
		if options.Codec == nil {
			return new(BuiltInFrameCodec)
		}
		return options.Codec
	}()

	server := Server{
		svr:          svr,
		Multicore:    options.Multicore,
		Addrs:        nil,
		NumEventLoop: numEventLoop,
		ReusePort:    options.ReusePort,
		TCPKeepAlive: options.TCPKeepAlive,
	}

	server.Addrs = make([]net.Addr, len(listeners))
	for i, ln := range listeners {
		server.Addrs[i] = ln.lnaddr
	}

	switch svr.eventHandler.OnInitComplete(server) {
	case None:
	case Shutdown:
		return nil
	}
	defer svr.eventHandler.OnShutdown(server)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer close(shutdown)

	go func() {
		if <-shutdown == nil {
			return
		}
		svr.signalShutdown()
	}()

	if err := svr.start(numEventLoop); err != nil {
		svr.closeEventLoops()
		svr.logger.Errorf("gnet server is stopping with error: %v", err)
		return err
	}
	defer svr.stop()

	return nil
}
