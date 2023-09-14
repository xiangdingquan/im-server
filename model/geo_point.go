package model

import (
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func MakeGeoPointByInput(geoPoint *mtproto.InputGeoPoint) (geo *mtproto.GeoPoint) {
	switch geoPoint.PredicateName {
	case mtproto.Predicate_inputGeoPointEmpty:
		geo = mtproto.MakeTLGeoPointEmpty(nil).To_GeoPoint()
	case mtproto.Predicate_inputGeoPoint:
		geo = mtproto.MakeTLGeoPoint(&mtproto.GeoPoint{
			Long: geoPoint.Long,
			Lat:  geoPoint.Lat,
		}).To_GeoPoint()
	default:
		log.Errorf("invalid predicateName - %v", geoPoint)
	}
	return
}
