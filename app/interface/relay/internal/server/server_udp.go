package server

import (
	"bytes"
	"encoding/hex"
	"net"
	"time"

	"open.chat/pkg/log"
	"open.chat/pkg/net2"
	"open.chat/pkg/util"
)

const (
	TLID_DECRYPTED_AUDIO_BLOCK        = 0xDBF948C1
	TLID_SIMPLE_AUDIO_BLOCK           = 0xCC0D0E76
	TLID_UDP_REFLECTOR_PEER_INFO      = 0x27D9371C
	TLID_UDP_REFLECTOR_PEER_INFO_IPV6 = 0x83fc73b1
	TLID_UDP_REFLECTOR_SELF_INFO      = 0xc01572c7
)

func listenUdp(c *Config) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", c.RelayServer.Addr)
	if err != nil {
		log.Errorf("net.ResolveUDPAddr fail - %v", err)
		return nil, err
	}
	log.Info(addr.String())
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Errorf("net.ListenUDP fail - %v", err)
		return nil, err
	}
	conn.SetReadBuffer(c.ReceiveBuf)
	conn.SetWriteBuffer(c.SendBuf)
	return conn, nil
}

func (s *Server) sendLoop() {
	for {
		select {
		case b, ok := <-s.sendChan:
			if ok {
				if _, err := s.relayConn.WriteToUDP(b.b, b.addr); err != nil {
					return
				}
			} else {
				return
			}
		case <-s.closeChan:
			return
		}
	}
}

func (s *Server) readLoop() {
	for {
		buf := make([]byte, 4096)
		bufLen, remote, err := s.relayConn.ReadFromUDP(buf)
		if err != nil {
			log.Errorf("readFromUDP error - %v", err)
			break
		} else {
			s.onRelayData(remote, buf[:bufLen])
		}
	}
}

func (s *Server) Send(addr *net.UDPAddr, b []byte) error {
	s.sendMutex.RLock()
	if !s.running {
		s.sendMutex.RUnlock()
		return net2.ConnectionClosedError
	}

	select {
	case s.sendChan <- udpDataBuf{addr, b}:
		s.sendMutex.RUnlock()
		return nil
	default:
		s.sendMutex.RUnlock()
		s.running = false
		return net2.ConnectionBlockedError
	}
}

func (s *Server) onRelayData(cAddr *net.UDPAddr, b []byte) {
	if len(b) < 32 {
		log.Errorf("invalid data len %d, addr %s", len(b), cAddr.String())
	}

	log.Debugf("onRelayData - receive from %s data: %s", cAddr.String(), hex.EncodeToString(b[:32]))

	buf := util.NewBufferInput(b)
	peerTag := string(buf.Bytes(16))
	packetTypes := make([]uint32, 4)
	packetTypes[0] = buf.UInt32()
	packetTypes[1] = buf.UInt32()
	packetTypes[2] = buf.UInt32()
	packetTypes[3] = buf.UInt32()

	if packetTypes[0] == 0xFFFFFFFF &&
		packetTypes[1] == 0xFFFFFFFF &&
		packetTypes[2] == 0xFFFFFFFF &&
		packetTypes[3] == 0xFFFFFFFF {
		s.onPublicEndpointsRequest(cAddr, peerTag)
	} else if packetTypes[0] == 0xFFFFFFFF &&
		packetTypes[1] == 0xFFFFFFFF &&
		packetTypes[2] == 0xFFFFFFFF &&
		packetTypes[3] == 0xFFFFFFFE {

		s.onUdpPing(cAddr, peerTag, buf.UInt64())
	} else {
		s.onRelayDataAvailable(cAddr, peerTag, b)
	}
}

func ipEmptyString(ip net.IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}

func (s *Server) onUdpPing(cAddr *net.UDPAddr, peerTag string, queryId uint64) {
	log.Infof("onUdpPing - recv from %s, peer_tag: %s, query_id: %d", cAddr.String(), peerTag, queryId)
	oBuf := util.NewBufferOutput(1024)
	oBuf.Bytes([]byte(peerTag))
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(TLID_UDP_REFLECTOR_SELF_INFO)
	oBuf.UInt32(uint32(time.Now().Unix()))
	oBuf.UInt64(queryId)

	ip := []byte(ipEmptyString(cAddr.IP))
	oBuf.Bytes(ip)
	oBuf.Bytes(make([]byte, 16-len(ip)))
	oBuf.UInt32(uint32(cAddr.Port))

	s.relayConn.WriteToUDP(oBuf.Buf(), cAddr)
}

func (s *Server) onPublicEndpointsRequest(cAddr *net.UDPAddr, peerTag string) {
	log.Debugf("onPublicEndpointsRequest - recv from %s, peer_tag: %s", cAddr.String(), peerTag)
}

func equalAddr(a1, a2 *net.UDPAddr) bool {
	if a1 == nil || a2 == nil {
		return false
	}

	return bytes.Equal(a1.IP, a2.IP) && a1.Port == a2.Port && a1.Zone == a2.Zone
}

func (s *Server) onRelayDataAvailable(cAddr *net.UDPAddr, peerTag string, buf []byte) {
	log.Debugf("onRelayDataAvailable - receive from %s, peer_tag: %s, len: %d", cAddr.String(), peerTag, len(buf))

	var (
		tb *RelayTable
		ok bool
	)

	s.tableMutex.RLock()
	if tb, ok = s.relayTable[peerTag]; !ok {
		s.tableMutex.RUnlock()
		log.Errorf("not found {peer_tag: %s} by {addr: %s}", peerTag, cAddr.String())
		return
	} else {
		s.tableMutex.RUnlock()
	}

	found := false
	for _, e := range tb.peers {
		if equalAddr(cAddr, e.addr) {
			found = true
			break
		}
	}
	if !found {
		tb.peers = append(tb.peers, Endpoint{cAddr, 0})
	}

	for _, e := range tb.peers {
		if !equalAddr(e.addr, cAddr) {
			s.relayConn.WriteToUDP(buf, e.addr)
		}
	}
}
