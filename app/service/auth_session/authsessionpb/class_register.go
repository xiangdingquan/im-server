package authsessionpb

const (
	Predicate_authKeyInfo                  = "authKeyInfo"
	Predicate_clientSessionInfo            = "clientSessionInfo"
	Predicate_session_setClientSessionInfo = "session_setClientSessionInfo"
	Predicate_session_getAuthorizations    = "session_getAuthorizations"
	Predicate_session_resetAuthorization   = "session_resetAuthorization"
	Predicate_session_getLayer             = "session_getLayer"
	Predicate_session_getLangCode          = "session_getLangCode"
	Predicate_session_getUserId            = "session_getUserId"
	Predicate_session_getPushSessionId     = "session_getPushSessionId"
	Predicate_session_getFutureSalts       = "session_getFutureSalts"
	Predicate_session_queryAuthKey         = "session_queryAuthKey"
	Predicate_session_setAuthKey           = "session_setAuthKey"
	Predicate_session_bindAuthKeyUser      = "session_bindAuthKeyUser"
	Predicate_session_unbindAuthKeyUser    = "session_unbindAuthKeyUser"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_authKeyInfo: {
		0:  -793297679, // 0xd0b73cf1
		85: -793297679, // 0xd0b73cf1

	},
	Predicate_clientSessionInfo: {
		0:  167005722, // 0x9f44e1a
		85: 167005722, // 0x9f44e1a

	},
	Predicate_session_setClientSessionInfo: {
		0:  -1513913311, // 0xa5c38421
		85: -1513913311, // 0xa5c38421

	},
	Predicate_session_getAuthorizations: {
		0:  848027106, // 0x328bdde2
		85: 848027106, // 0x328bdde2

	},
	Predicate_session_resetAuthorization: {
		0:  -1038977694, // 0xc2127562
		85: -1038977694, // 0xc2127562

	},
	Predicate_session_getLayer: {
		0:  -238328911, // 0xf1cb63b1
		85: -238328911, // 0xf1cb63b1

	},
	Predicate_session_getLangCode: {
		0:  -1213481174, // 0xb7abbf2a
		85: -1213481174, // 0xb7abbf2a

	},
	Predicate_session_getUserId: {
		0:  -798477825, // 0xd06831ff
		85: -798477825, // 0xd06831ff

	},
	Predicate_session_getPushSessionId: {
		0:  -1731520768, // 0x98cb1700
		85: -1731520768, // 0x98cb1700

	},
	Predicate_session_getFutureSalts: {
		0:  -364935027, // 0xea3f888d
		85: -364935027, // 0xea3f888d

	},
	Predicate_session_queryAuthKey: {
		0:  1798174801, // 0x6b2df851
		85: 1798174801, // 0x6b2df851

	},
	Predicate_session_setAuthKey: {
		0:  -1302981520, // 0xb2561470
		85: -1302981520, // 0xb2561470

	},
	Predicate_session_bindAuthKeyUser: {
		0:  -1721267986, // 0x996788ee
		85: -1721267986, // 0x996788ee

	},
	Predicate_session_unbindAuthKeyUser: {
		0:  -359222990, // 0xea96b132
		85: -359222990, // 0xea96b132

	},
}

func init() {
}

func GetClazzID(clazzName string, layer int) int32 {
	if m, ok := clazzNameRegisters2[clazzName]; ok {
		if m2, ok := m[layer]; ok {
			return m2
		}
	}
	return 0
}
