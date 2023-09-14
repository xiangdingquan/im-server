// Copyright (c) 2019 Andy Pan
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
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func (svr *server) acceptNewConnection(fd int) error {
	for i, ln := range svr.lns {
		if ln.fd == fd {
			nfd, sa, err := unix.Accept(fd)
			if err != nil {
				if err == unix.EAGAIN {
					return nil
				}
				return os.NewSyscallError("accept", err)
			}
			if err := unix.SetNonblock(nfd, true); err != nil {
				return os.NewSyscallError("fcntl nonblock", err)
			}
			if svr.opts.ReceiveBuf > 0 {
				if err := unix.SetsockoptInt(nfd,
					syscall.SOL_SOCKET,
					syscall.SO_RCVBUF,
					svr.opts.ReceiveBuf); err != nil {
					return os.NewSyscallError("setsockopt so_rcvbuf", err)
				}
			}
			if svr.opts.SendBuf > 0 {
				if err := unix.SetsockoptInt(nfd,
					syscall.SOL_SOCKET,
					syscall.SO_SNDBUF,
					svr.opts.SendBuf); err != nil {
					return os.NewSyscallError("setsockopt so_sndbuf", err)
				}
			}
			el := svr.subEventLoopSet.next(nfd)
			c := newTCPConn(nfd, el, sa)
			_ = el.poller.Trigger(func() (err error) {
				if err = el.poller.AddRead(nfd); err != nil {
					return
				}
				el.connections[nfd] = c
				el.calibrateCallback(el, 1)
				err = el.loopOpen(c, i)
				return
			})
		}
	}
	return nil
}
