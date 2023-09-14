package dao

import (
	"context"
	"strconv"

	"open.chat/app/sysconfig"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func makeBoolConfig(key string, value bool) *mtproto.JSONObjectValue {
	valueBOOL := mtproto.BoolFalse
	if value {
		valueBOOL = mtproto.BoolTrue
	}
	return mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: key,
		Value: mtproto.MakeTLJsonBool(&mtproto.JSONValue{
			Value_BOOL: valueBOOL,
		}).To_JSONValue(),
	}).To_JSONObjectValue()
}

func makeInt32Config(key string, value int32) *mtproto.JSONObjectValue {
	return mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: key,
		Value: mtproto.MakeTLJsonNumber(&mtproto.JSONValue{
			Value_FLOAT64: float64(value),
		}).To_JSONValue(),
	}).To_JSONObjectValue()
}

func addSysConfigsBool(ctx context.Context, values []*mtproto.JSONObjectValue, key sysconfig.ConfigKeys, defaultValue bool, timeout uint) []*mtproto.JSONObjectValue {
	v := sysconfig.GetConfig2Bool(ctx, key, defaultValue, timeout)
	return append(values, makeBoolConfig(string(key), v))
}

func (d *Dao) GetAppConfigs(ctx context.Context) *mtproto.JSONValue {
	appConfigs, err := d.AppConfigsDAO.SelectList(ctx)
	if err != nil {
		log.Errorf("getAppConfigs - error: %v", err)
	}

	values := make([]*mtproto.JSONObjectValue, 0, len(appConfigs))
	for i := 0; i < len(appConfigs); i++ {
		switch appConfigs[i].Type2 {
		case "number":
			v, err := strconv.ParseFloat(appConfigs[i].Value2, 64)
			if err != nil {
				log.Errorf("getAppConfigs - error(%v): %v", appConfigs[i], err)
				continue
			}
			values = append(values, mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
				Key: appConfigs[i].Key2,
				Value: mtproto.MakeTLJsonNumber(&mtproto.JSONValue{
					Value_FLOAT64: v,
				}).To_JSONValue(),
			}).To_JSONObjectValue())
		case "string":
			values = append(values, mtproto.MakeTLJsonObjectValue(&mtproto.JSONObjectValue{
				Key: appConfigs[i].Key2,
				Value: mtproto.MakeTLJsonString(&mtproto.JSONValue{
					Value_STRING: appConfigs[i].Value2,
				}).To_JSONValue(),
			}).To_JSONObjectValue())
		case "bool":
			v := appConfigs[i].Value2 == "true" || appConfigs[i].Value2 == "1"
			values = append(values, makeBoolConfig(appConfigs[i].Key2, v))
		case "array":
		case "object":
		}
	}

	phoneCodeLogin := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysPhoneCodeLogin, false, 0)
	values = append(values, makeBoolConfig(string(sysconfig.ConfigKeysPhoneCodeLogin), phoneCodeLogin))

	needPhoneCode := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterNeedPhoneCode, false, 0)
	values = append(values, makeBoolConfig(string(sysconfig.ConfigKeysRegisterNeedPhoneCode), needPhoneCode))

	needInviter := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterNeedInviter, false, 0)
	values = append(values, makeBoolConfig(string(sysconfig.ConfigKeysRegisterNeedInviter), needInviter))

	addInviterAsFriend := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterAddInviterAsFriend, false, 0)
	values = append(values, makeBoolConfig(string(sysconfig.ConfigKeysRegisterAddInviterAsFriend), addInviterAsFriend))

	usingOss := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysUsingOss, false, 0)
	values = append(values, makeBoolConfig(string(sysconfig.ConfigKeysUsingOss), usingOss))

	passwordFloodInterval := sysconfig.GetConfig2Int32(ctx, sysconfig.ConfigKeyPasswordFloodInterval, 0, 0)
	values = append(values, makeInt32Config(string(sysconfig.ConfigKeyPasswordFloodInterval), passwordFloodInterval))

	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSendFile, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSendLocation, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSendRedpacket, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanRemit, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSeeAddressBook, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSeeBlog, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanInviteFriend, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSeeNearby, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSeePublicGroup, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSeeQrCode, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSeeWallet, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSeeWalletRecords, false, 0)
	values = addSysConfigsBool(ctx, values, sysconfig.ConfigKeysCanSeeEmojiShop, false, 0)

	return mtproto.MakeTLJsonObject(&mtproto.JSONValue{
		Value_VECTORJSONOBJECTVALUE: values,
	}).To_JSONValue()
}
