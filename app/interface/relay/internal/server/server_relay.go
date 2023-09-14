package server

import (
	"context"
	"open.chat/app/interface/relay/relaypb"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
	"open.chat/pkg/random2"
)

func (s *Server) RelayCreateCall(ctx context.Context, request *relaypb.RelayCreateCallRequest) (resp *relaypb.CallConnections, err error) {
	log.Debugf("relayCreateCall request - %s", logger.JsonDebugData(request))

	s.tableMutex.Lock()
	defer s.tableMutex.Unlock()

	if peerTag, ok := s.idTable[request.Id]; ok {
		log.Debugf("relayCreateCall - peerTag(%s) found: %d", peerTag, request.Id)
		resp = &relaypb.CallConnections{
			Id:                     request.Id,
			PeerTag:                []byte(peerTag),
			Connection:             makeConnection(s.GetConnectionID(), s.c.RelayIp, s.c.RelayPort, peerTag),
			AlternativeConnections: []*mtproto.PhoneConnection{},
		}
	} else {
		for {
			peerTag = random2.RandomAlphabetic(16)
			if _, ok := s.relayTable[peerTag]; !ok {
				break
			}
		}
		log.Debugf("relayCreateCall - peerTag(%s) create: %d", peerTag, request.Id)
		s.idTable[request.Id] = peerTag
		s.relayTable[peerTag] = &RelayTable{peerTag, []Endpoint{}}
		resp = &relaypb.CallConnections{
			Id:                     request.Id,
			PeerTag:                []byte(peerTag),
			Connection:             makeConnection(s.GetConnectionID(), s.c.RelayIp, s.c.RelayPort, peerTag),
			AlternativeConnections: []*mtproto.PhoneConnection{},
		}
	}

	log.Debugf("relayCreateCall reply - %s", logger.JsonDebugData(resp))
	return
}

func (s *Server) RelayDiscardCall(ctx context.Context, request *relaypb.RelaydiscardCallRequest) (*mtproto.Bool, error) {
	log.Debugf("relayDiscardCall request - %s", logger.JsonDebugData(request))

	s.tableMutex.Lock()
	defer s.tableMutex.Unlock()

	discardOk := true
	if peerTag, ok := s.idTable[request.Id]; ok {
		delete(s.relayTable, peerTag)
		delete(s.idTable, request.Id)
	} else {
		log.Debugf("relayDiscardCall - request not found: %d", request.Id)
		discardOk = false
	}

	log.Debugf("relayDiscardCall reply - {true}")
	return mtproto.ToBool(discardOk), nil
}

func makeConnection(id int64, relayIp string, relayPort int32, peerTag string) *mtproto.PhoneConnection {
	return mtproto.MakeTLPhoneConnection(&mtproto.PhoneConnection{
		Id:      id,
		Ip:      relayIp,
		Ipv6:    "",
		Port:    relayPort,
		PeerTag: []byte(peerTag),
	}).To_PhoneConnection()
}
