package server

import (
	"net"
)

type Endpoint struct {
	addr             *net.UDPAddr
	lastReceivedTime int64
}

type RelayTable struct {
	peerTag string
	peers   []Endpoint
}
