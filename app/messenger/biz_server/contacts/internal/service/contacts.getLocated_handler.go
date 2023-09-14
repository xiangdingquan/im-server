package service

import (
	"context"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsGetLocated(ctx context.Context, request *mtproto.TLContactsGetLocated) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.getLocated - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var err error
	if request.GeoPoint.GetPredicateName() == mtproto.Predicate_inputGeoPointEmpty {
		log.Errorf("ContactsGetLocated, invalid geo point, %s", request.GeoPoint.GetPredicateName())
		err = mtproto.ErrButtonTypeInvalid
		return nil, err
	}

	var nearByUsers []*model.NearByUser
	if request.GetBackground() {
		err = s.ChannelFacade.SetLocated(ctx, request.GeoPoint.GetLat(), request.GeoPoint.GetLong(), md.GetUserId(), request.GetSelfExpires() != nil, request.GetSelfExpires().GetValue())
	} else {
		_ = s.ChannelFacade.ClearExpiredLocation(ctx)

		radius := int32(500)
		if request.GeoPoint.GetAccuracyRadius() != nil {
			radius = request.GeoPoint.GetAccuracyRadius().GetValue()
		}
		nearByUsers, err = s.SearchNearBy(ctx, request.GeoPoint.GetLat(), request.GeoPoint.GetLong(), radius, 500)
	}

	if err != nil {
		log.Errorf("contacts.getLocated, call facade failed, error: %v", err)
		return nil, err
	}

	l := make([]*mtproto.PeerLocated, 0, len(nearByUsers))
	for _, n := range nearByUsers {
		l = append(l, mtproto.MakeTLPeerLocated(&mtproto.PeerLocated{
			Peer:     model.MakePeerUser(n.Id),
			Expires:  n.Expiration,
			Distance: int32(n.Distance),
		}).To_PeerLocated())
	}

	contactUpdates := model.NewUpdatesLogic(md.GetUserId())
	contactUpdates.AddUpdate(mtproto.MakeTLUpdatePeerLocated(&mtproto.Update{
		Peers: l,
	}).To_Update())

	ups := contactUpdates.ToUpdates()
	return ups, nil
}
