package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetAppConfig(ctx context.Context, request *mtproto.TLHelpGetAppConfig) (*mtproto.JSONValue, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getAppConfig - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getAppConfig - error: %v", err)
		return nil, err
	}

	r := s.Dao.GetAppConfigs(ctx)
	r.Value_VECTORJSONOBJECTVALUE = append(r.Value_VECTORJSONOBJECTVALUE,
		mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
			Key: "emojies_send_dice",
			Value: mtproto.MakeTLJsonArray(&mtproto.JSONValue{
				Value_VECTORJSONVALUE: []*mtproto.JSONValue{
					mtproto.MakeTLJsonString(&mtproto.JSONValue{
						Value_STRING: "🎲",
					}).To_JSONValue(),
					mtproto.MakeTLJsonString(&mtproto.JSONValue{
						Value_STRING: "🎯",
					}).To_JSONValue(),
					mtproto.MakeTLJsonString(&mtproto.JSONValue{
						Value_STRING: "🏀",
					}).To_JSONValue(),
				},
			}).To_JSONValue(),
		}).To_JSONObjectValue())
	r.Value_VECTORJSONOBJECTVALUE = append(r.Value_VECTORJSONOBJECTVALUE,
		mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
			Key: "emojies_send_dice_success",
			Value: mtproto.MakeTLJsonObject(&mtproto.JSONValue{
				Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{
					mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
						Key: "🎯",
						Value: mtproto.MakeTLJsonObject(&mtproto.JSONValue{
							Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{
								mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
									Key: "value",
									Value: mtproto.MakeTLJsonNumber(&mtproto.JSONValue{
										Value_FLOAT64: 6,
									}).To_JSONValue(),
								}).To_JSONObjectValue(),
								mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
									Key: "frame_start",
									Value: mtproto.MakeTLJsonNumber(&mtproto.JSONValue{
										Value_FLOAT64: 62,
									}).To_JSONValue(),
								}).To_JSONObjectValue(),
							},
						}).To_JSONValue(),
					}).To_JSONObjectValue(),
					mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
						Key: "🏀",
						Value: mtproto.MakeTLJsonObject(&mtproto.JSONValue{
							Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{
								mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
									Key: "value",
									Value: mtproto.MakeTLJsonNumber(&mtproto.JSONValue{
										Value_FLOAT64: 5,
									}).To_JSONValue(),
								}).To_JSONObjectValue(),
								mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
									Key: "frame_start",
									Value: mtproto.MakeTLJsonNumber(&mtproto.JSONValue{
										Value_FLOAT64: 110,
									}).To_JSONValue(),
								}).To_JSONObjectValue(),
							},
						}).To_JSONValue(),
					}).To_JSONObjectValue(),
				},
			}).To_JSONValue(),
		}).To_JSONObjectValue())

	log.Debugf("help.getAppConfig#98914110 - reply: %s", r.DebugString())
	return r, nil
}
