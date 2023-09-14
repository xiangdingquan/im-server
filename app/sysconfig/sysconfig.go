package sysconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"open.chat/app/infra/databus/pkg/cache/redis"
	"open.chat/app/pkg/mysql_util"
	"open.chat/app/pkg/redis_util"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

type ConfigKeys string

// call
const (
	// App名字
	ConfigKeysAppName ConfigKeys = "APP_name"
	// App名字
	ConfigKeysWebSite ConfigKeys = "web_site"
	// 短信签名
	ConfigKeysSmsSigned ConfigKeys = "sms_qianming"
	// 群内敏感词
	ConfigKeysBanWords ConfigKeys = "ban_words"

	// 充值客服标识
	ConfigKeysPayServiceUId ConfigKeys = "pay_service_uid"

	// 加为好友才能聊天
	ConfigKeysFriendChat ConfigKeys = "friend_chat"
	// 群内禁止发送图片
	ConfigKeysBanImages ConfigKeys = "ban_imgs"
	// 群内禁止发送二维码
	ConfigKeysBanQrcode ConfigKeys = "ban_qrcode"
	// 群内禁止发送网址链接
	ConfigKeysBanLinks ConfigKeys = "ban_links"
	// 群内限制发送间隔
	ConfigKeysGroupChatSendInterval ConfigKeys = "group_chat_time_limit"

	// 验证码登录
	ConfigKeysPhoneCodeLogin ConfigKeys = "phone_code_login"

	// 注册需要验证码
	ConfigKeysRegisterNeedPhoneCode ConfigKeys = "register_need_phone_code"
	// 注册需要邀请人
	ConfigKeysRegisterNeedInviter ConfigKeys = "register_need_inviter"
	// 注册添加邀请人为好友
	ConfigKeysRegisterAddInviterAsFriend ConfigKeys = "register_add_inviter_as_friend"
	// 允许修改用户名
	ConfigKeysPermitModifyUserName ConfigKeys = "permit_modify_user_name"

	// 验证码
	ConfigKeysCommonSmsCode ConfigKeys = "common_sms_code"

	// 阿里云oss配置
	ConfigKeysUsingOss          ConfigKeys = "using_oss"
	ConfigKeyOssBucketName      ConfigKeys = "oss_bucket_name"
	ConfigKeyOssEndpoint        ConfigKeys = "oss_endpoint"
	ConfigKeyOssAccessKeyId     ConfigKeys = "oss_access_key_id"
	ConfigKeyOssAccessKeySecret ConfigKeys = "oss_access_key_secret"

	// 同IP注册限制
	ConfigKeyRegisterLimitOfIp ConfigKeys = "register_limit_of_ip"
	// 客服消息
	ConfigKeyCustomerServiceMessageForRegister ConfigKeys = "customer_service_message_for_register"
	// 密码错误次数
	ConfigKeyPasswordFloodLimit ConfigKeys = "password_flood_limit"
	// 密码错误锁定时间
	ConfigKeyPasswordFloodInterval ConfigKeys = "password_flood_interval"

	// 各种开关控制
	//发文件
	ConfigKeysCanSendFile ConfigKeys = "can_send_file"
	//位置
	ConfigKeysCanSendLocation ConfigKeys = "can_send_location"
	//红包
	ConfigKeysCanSendRedpacket ConfigKeys = "can_send_redpacket"
	//转账
	ConfigKeysCanRemit ConfigKeys = "can_remit"
	//联系人页面的通讯录
	ConfigKeysCanSeeAddressBook ConfigKeys = "can_see_address_book"
	//发现页面的朋友圈
	ConfigKeysCanSeeBlog ConfigKeys = "can_see_blog"
	//邀请好友
	ConfigKeysCanInviteFriend ConfigKeys = "can_invite_friend"
	//附近的人
	ConfigKeysCanSeeNearby ConfigKeys = "can_see_nearby"
	//公开的群
	ConfigKeysCanSeePublicGroup ConfigKeys = "can_see_public_group"
	//二维码
	ConfigKeysCanSeeQrCode ConfigKeys = "can_see_qr_code"
	//我的钱包
	ConfigKeysCanSeeWallet ConfigKeys = "can_see_wallet"
	//交易记录
	ConfigKeysCanSeeWalletRecords ConfigKeys = "can_see_wallet_records"
	//表情商店
	ConfigKeysCanSeeEmojiShop ConfigKeys = "can_see_emoji_shop"
)

type (
	configs struct {
		*redis_util.Redis
		*sqlx.DB
	}

	sysConfigDO struct {
		ID      uint32 `db:"id"`
		Key     string `db:"key"`
		Value   string `db:"value"`
		Date    int32  `db:"date"`
		Deleted bool   `db:"deleted"`
	}
)

const (
	configCacheTimeout = 1 // salt timeout
	configCachePrefix  = "system_config_cache"
)

var g_Configs *configs = nil

func getGConfigs() *configs {
	if g_Configs == nil {
		g_Configs = &configs{
			Redis: redis_util.GetSingletonRedis(),
			DB:    mysql_util.GetSingletonSqlxDB(),
		}
	}
	return g_Configs
}

func configCacheKey(key ConfigKeys) string {
	return fmt.Sprintf("%s_%s", configCachePrefix, key)
}

// SelectByKey .
func selectByKey(ctx context.Context, key string) (rValue *sysConfigDO, err error) {
	var (
		query = "SELECT id, `key`, value, `date`, deleted FROM `sys_configs` WHERE `key` = ?"
		rows  *sqlx.Rows
	)
	rows, err = getGConfigs().DB.Query(ctx, query, key)

	if err != nil {
		log.Errorf("queryx in SelectByID(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &sysConfigDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByID(_), error: %v", err)
		} else {
			rValue = do
		}
	}

	return
}

func getConfig(ctx context.Context, key ConfigKeys, def interface{}, timeout uint) bool {
	if timeout == 0 {
		timeout = configCacheTimeout
	}
	cacheKey := configCacheKey(key)
	conn := getGConfigs().Redis.Redis.Get(ctx)
	defer conn.Close()
	value := ""
	v, err := redis.String(conn.Do("GET", cacheKey))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", cacheKey, err)
		} else {
			err = nil
		}
	} else {
		value = v
	}
	if value == "" {
		scdo, err := selectByKey(ctx, string(key))
		if scdo == nil || err != nil {
			return false
		}
		value = scdo.Value
		if _, err = conn.Do("SETEX", cacheKey, timeout, value); err != nil {
			log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
		}
	}

	valueType := reflect.ValueOf(def)
	if valueType.Kind() != reflect.Ptr {
		return false
	}
	valueType = valueType.Elem()
	switch valueType.Kind() {
	case reflect.Bool:
		intValue, _ := util.StringToInt32(value)
		valueType.SetBool(intValue != 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, _ := strconv.ParseInt(value, 10, 64)
		valueType.SetInt(intValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uintValue, _ := strconv.ParseUint(value, 10, 64)
		valueType.SetUint(uintValue)
	case reflect.Float32, reflect.Float64:
		floatValue, _ := strconv.ParseFloat(value, 64)
		valueType.SetFloat(floatValue)
	case reflect.Complex64, reflect.Complex128:
		complexValue, _ := strconv.ParseComplex(value, 128)
		valueType.SetComplex(complexValue)
	case reflect.String:
		valueType.SetString(value)
	case reflect.Slice:
		switch reflect.TypeOf(def).Elem().Elem().Kind() {
		case reflect.String:
			json.Unmarshal([]byte(value), &def)
		default:
			return false
		}
	default:
		return false
	}
	return true
}

func GetConfig2Bool(ctx context.Context, key ConfigKeys, def bool, timeout uint) bool {
	getConfig(ctx, key, &def, timeout)
	return def
}

func GetConfig2Int32(ctx context.Context, key ConfigKeys, def int32, timeout uint) int32 {
	getConfig(ctx, key, &def, timeout)
	return def
}

func GetConfig2Uint32(ctx context.Context, key ConfigKeys, def uint32, timeout uint) uint32 {
	getConfig(ctx, key, &def, timeout)
	return def
}

func GetConfig2float32(ctx context.Context, key ConfigKeys, def float32, timeout uint) float32 {
	getConfig(ctx, key, &def, timeout)
	return def
}

func GetConfig2String(ctx context.Context, key ConfigKeys, def string, timeout uint) string {
	getConfig(ctx, key, &def, timeout)
	return def
}

func GetConfig2StringArray(ctx context.Context, key ConfigKeys, def []string, timeout uint) []string {
	if def == nil {
		def = make([]string, 0)
	}
	getConfig(ctx, key, &def, timeout)
	return def
}
