package authsessionpb

import (
	"fmt"

	"github.com/gogo/protobuf/types"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

var _ *types.Int32Value
var _ *mtproto.Bool

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor
	-793297679: func() mtproto.TLObject { // 0xd0b73cf1
		o := MakeTLAuthKeyInfo(nil)
		o.Data2.Constructor = -793297679
		return o
	},
	167005722: func() mtproto.TLObject { // 0x9f44e1a
		o := MakeTLClientSessionInfo(nil)
		o.Data2.Constructor = 167005722
		return o
	},

	// Method
	-1513913311: func() mtproto.TLObject { // 0xa5c38421
		return &TLSessionSetClientSessionInfo{
			Constructor: -1513913311,
		}
	},
	848027106: func() mtproto.TLObject { // 0x328bdde2
		return &TLSessionGetAuthorizations{
			Constructor: 848027106,
		}
	},
	-1038977694: func() mtproto.TLObject { // 0xc2127562
		return &TLSessionResetAuthorization{
			Constructor: -1038977694,
		}
	},
	-238328911: func() mtproto.TLObject { // 0xf1cb63b1
		return &TLSessionGetLayer{
			Constructor: -238328911,
		}
	},
	-1213481174: func() mtproto.TLObject { // 0xb7abbf2a
		return &TLSessionGetLangCode{
			Constructor: -1213481174,
		}
	},
	-798477825: func() mtproto.TLObject { // 0xd06831ff
		return &TLSessionGetUserId{
			Constructor: -798477825,
		}
	},
	-1731520768: func() mtproto.TLObject { // 0x98cb1700
		return &TLSessionGetPushSessionId{
			Constructor: -1731520768,
		}
	},
	-364935027: func() mtproto.TLObject { // 0xea3f888d
		return &TLSessionGetFutureSalts{
			Constructor: -364935027,
		}
	},
	1798174801: func() mtproto.TLObject { // 0x6b2df851
		return &TLSessionQueryAuthKey{
			Constructor: 1798174801,
		}
	},
	-1302981520: func() mtproto.TLObject { // 0xb2561470
		return &TLSessionSetAuthKey{
			Constructor: -1302981520,
		}
	},
	-1721267986: func() mtproto.TLObject { // 0x996788ee
		return &TLSessionBindAuthKeyUser{
			Constructor: -1721267986,
		}
	},
	-359222990: func() mtproto.TLObject { // 0xea96b132
		return &TLSessionUnbindAuthKeyUser{
			Constructor: -359222990,
		}
	},
}

func NewTLObjectByClassID(classId int32) mtproto.TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}

func (m *AuthKeyInfo) Encode(layer int32) []byte {
	switch m.PredicateName {
	case Predicate_authKeyInfo:
		t := m.To_AuthKeyInfo()
		return t.Encode(layer)

	default:
		err := fmt.Errorf("invalid predicate error: %s", m.PredicateName)
		log.Errorf(err.Error())
		return []byte{}
	}
}

func (m *AuthKeyInfo) CalcByteSize(layer int32) int {
	return 0
}

func (m *AuthKeyInfo) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xd0b73cf1:
		m2 := MakeTLAuthKeyInfo(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *AuthKeyInfo) DebugString() string {
	switch m.PredicateName {
	case Predicate_authKeyInfo:
		t := m.To_AuthKeyInfo()
		return t.DebugString()

	default:
		return "{}"
	}
}

func (m *AuthKeyInfo) To_AuthKeyInfo() *TLAuthKeyInfo {
	m.PredicateName = Predicate_authKeyInfo
	return &TLAuthKeyInfo{
		Data2: m,
	}
}

func MakeTLAuthKeyInfo(data2 *AuthKeyInfo) *TLAuthKeyInfo {
	if data2 == nil {
		return &TLAuthKeyInfo{Data2: &AuthKeyInfo{
			PredicateName: Predicate_authKeyInfo,
		}}
	} else {
		data2.PredicateName = Predicate_authKeyInfo
		return &TLAuthKeyInfo{Data2: data2}
	}
}

func (m *TLAuthKeyInfo) To_AuthKeyInfo() *AuthKeyInfo {
	m.Data2.PredicateName = Predicate_authKeyInfo
	return m.Data2
}

func (m *TLAuthKeyInfo) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLAuthKeyInfo) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLAuthKeyInfo) SetAuthKey(v []byte) { m.Data2.AuthKey = v }
func (m *TLAuthKeyInfo) GetAuthKey() []byte  { return m.Data2.AuthKey }

func (m *TLAuthKeyInfo) GetPredicateName() string {
	return Predicate_authKeyInfo
}

func (m *TLAuthKeyInfo) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xd0b73cf1: func() []byte {
			x.UInt(0xd0b73cf1)
			var flags uint32 = 0
			x.UInt(flags)

			x.Long(m.GetAuthKeyId())
			x.StringBytes(m.GetAuthKey())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_authKeyInfo, int(layer))
	if clazzId == 0 {
		return x.GetBuf()
	}

	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		return nil
	}

	return x.GetBuf()
}

func (m *TLAuthKeyInfo) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLAuthKeyInfo) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xd0b73cf1: func() error {
			flags := dBuf.UInt()
			_ = flags

			m.SetAuthKeyId(dBuf.Long())
			m.SetAuthKey(dBuf.StringBytes())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLAuthKeyInfo) DebugString() string {
	return "{}"
}

func (m *ClientSession) Encode(layer int32) []byte {
	switch m.PredicateName {
	case Predicate_clientSessionInfo:
		t := m.To_ClientSessionInfo()
		return t.Encode(layer)

	default:
		err := fmt.Errorf("invalid predicate error: %s", m.PredicateName)
		log.Errorf(err.Error())
		return []byte{}
	}
}

func (m *ClientSession) CalcByteSize(layer int32) int {
	return 0
}

func (m *ClientSession) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x9f44e1a:
		m2 := MakeTLClientSessionInfo(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *ClientSession) DebugString() string {
	switch m.PredicateName {
	case Predicate_clientSessionInfo:
		t := m.To_ClientSessionInfo()
		return t.DebugString()

	default:
		return "{}"
	}
}

func (m *ClientSession) To_ClientSessionInfo() *TLClientSessionInfo {
	m.PredicateName = Predicate_clientSessionInfo
	return &TLClientSessionInfo{
		Data2: m,
	}
}

func MakeTLClientSessionInfo(data2 *ClientSession) *TLClientSessionInfo {
	if data2 == nil {
		return &TLClientSessionInfo{Data2: &ClientSession{
			PredicateName: Predicate_clientSessionInfo,
		}}
	} else {
		data2.PredicateName = Predicate_clientSessionInfo
		return &TLClientSessionInfo{Data2: data2}
	}
}

func (m *TLClientSessionInfo) To_ClientSession() *ClientSession {
	m.Data2.PredicateName = Predicate_clientSessionInfo
	return m.Data2
}

func (m *TLClientSessionInfo) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLClientSessionInfo) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLClientSessionInfo) SetIp(v string) { m.Data2.Ip = v }
func (m *TLClientSessionInfo) GetIp() string  { return m.Data2.Ip }

func (m *TLClientSessionInfo) SetLayer(v int32) { m.Data2.Layer = v }
func (m *TLClientSessionInfo) GetLayer() int32  { return m.Data2.Layer }

func (m *TLClientSessionInfo) SetApiId(v int32) { m.Data2.ApiId = v }
func (m *TLClientSessionInfo) GetApiId() int32  { return m.Data2.ApiId }

func (m *TLClientSessionInfo) SetDeviceModel(v string) { m.Data2.DeviceModel = v }
func (m *TLClientSessionInfo) GetDeviceModel() string  { return m.Data2.DeviceModel }

func (m *TLClientSessionInfo) SetSystemVersion(v string) { m.Data2.SystemVersion = v }
func (m *TLClientSessionInfo) GetSystemVersion() string  { return m.Data2.SystemVersion }

func (m *TLClientSessionInfo) SetAppVersion(v string) { m.Data2.AppVersion = v }
func (m *TLClientSessionInfo) GetAppVersion() string  { return m.Data2.AppVersion }

func (m *TLClientSessionInfo) SetSystemLangCode(v string) { m.Data2.SystemLangCode = v }
func (m *TLClientSessionInfo) GetSystemLangCode() string  { return m.Data2.SystemLangCode }

func (m *TLClientSessionInfo) SetLangPack(v string) { m.Data2.LangPack = v }
func (m *TLClientSessionInfo) GetLangPack() string  { return m.Data2.LangPack }

func (m *TLClientSessionInfo) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLClientSessionInfo) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLClientSessionInfo) GetPredicateName() string {
	return Predicate_clientSessionInfo
}

func (m *TLClientSessionInfo) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x9f44e1a: func() []byte {
			x.UInt(0x9f44e1a)
			x.Long(m.GetAuthKeyId())
			x.String(m.GetIp())
			x.Int(m.GetLayer())
			x.Int(m.GetApiId())
			x.String(m.GetDeviceModel())
			x.String(m.GetSystemVersion())
			x.String(m.GetAppVersion())
			x.String(m.GetSystemLangCode())
			x.String(m.GetLangPack())
			x.String(m.GetLangCode())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_clientSessionInfo, int(layer))
	if clazzId == 0 {
		return x.GetBuf()
	}

	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		return nil
	}

	return x.GetBuf()
}

func (m *TLClientSessionInfo) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLClientSessionInfo) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x9f44e1a: func() error {
			m.SetAuthKeyId(dBuf.Long())
			m.SetIp(dBuf.String())
			m.SetLayer(dBuf.Int())
			m.SetApiId(dBuf.Int())
			m.SetDeviceModel(dBuf.String())
			m.SetSystemVersion(dBuf.String())
			m.SetAppVersion(dBuf.String())
			m.SetSystemLangCode(dBuf.String())
			m.SetLangPack(dBuf.String())
			m.SetLangCode(dBuf.String())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLClientSessionInfo) DebugString() string {
	return "{}"
}

func (m *TLSessionSetClientSessionInfo) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0xa5c38421:
		x.UInt(0xa5c38421)
		x.Bytes(m.GetSession().Encode(layer))

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionSetClientSessionInfo) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionSetClientSessionInfo) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa5c38421:
		m1 := &ClientSession{}
		m1.Decode(dBuf)
		m.Session = m1

		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionSetClientSessionInfo) DebugString() string {
	return "{}"
}

func (m *TLSessionGetAuthorizations) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0x328bdde2:
		x.UInt(0x328bdde2)
		x.Int(m.GetUserId())
		x.Long(m.GetExcludeAuthKeyId())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionGetAuthorizations) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionGetAuthorizations) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x328bdde2:
		m.UserId = dBuf.Int()
		m.ExcludeAuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionGetAuthorizations) DebugString() string {
	return "{}"
}

func (m *TLSessionResetAuthorization) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0xc2127562:
		x.UInt(0xc2127562)
		x.Int(m.GetUserId())
		x.Long(m.GetHash())
	default:
	}

	return x.GetBuf()
}

func (m *TLSessionResetAuthorization) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionResetAuthorization) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc2127562:
		m.UserId = dBuf.Int()
		m.Hash = dBuf.Long()
		return dBuf.GetError()
	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionResetAuthorization) DebugString() string {
	return "{}"
}

func (m *TLSessionGetLayer) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0xf1cb63b1:
		x.UInt(0xf1cb63b1)
		x.Long(m.GetAuthKeyId())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionGetLayer) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionGetLayer) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf1cb63b1:
		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionGetLayer) DebugString() string {
	return "{}"
}

func (m *TLSessionGetLangCode) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0xb7abbf2a:
		x.UInt(0xb7abbf2a)
		x.Long(m.GetAuthKeyId())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionGetLangCode) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionGetLangCode) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb7abbf2a:
		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionGetLangCode) DebugString() string {
	return "{}"
}

func (m *TLSessionGetUserId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0xd06831ff:
		x.UInt(0xd06831ff)
		x.Long(m.GetAuthKeyId())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionGetUserId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionGetUserId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd06831ff:
		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionGetUserId) DebugString() string {
	return "{}"
}

func (m *TLSessionGetPushSessionId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0x98cb1700:
		x.UInt(0x98cb1700)
		x.Int(m.GetUserId())
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetTokenType())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionGetPushSessionId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionGetPushSessionId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x98cb1700:
		m.UserId = dBuf.Int()
		m.AuthKeyId = dBuf.Long()
		m.TokenType = dBuf.Int()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionGetPushSessionId) DebugString() string {
	return "{}"
}

func (m *TLSessionGetFutureSalts) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0xea3f888d:
		x.UInt(0xea3f888d)
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetNum())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionGetFutureSalts) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionGetFutureSalts) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xea3f888d:
		m.AuthKeyId = dBuf.Long()
		m.Num = dBuf.Int()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionGetFutureSalts) DebugString() string {
	return "{}"
}

func (m *TLSessionQueryAuthKey) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0x6b2df851:
		x.UInt(0x6b2df851)
		x.Long(m.GetAuthKeyId())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionQueryAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionQueryAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6b2df851:
		m.AuthKeyId = dBuf.Long()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionQueryAuthKey) DebugString() string {
	return "{}"
}

func (m *TLSessionSetAuthKey) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0xb2561470:
		x.UInt(0xb2561470)
		x.Bytes(m.GetAuthKey().Encode(layer))

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionSetAuthKey) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionSetAuthKey) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb2561470:
		m1 := &AuthKeyInfo{}
		m1.Decode(dBuf)
		m.AuthKey = m1

		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionSetAuthKey) DebugString() string {
	return "{}"
}

func (m *TLSessionBindAuthKeyUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0x996788ee:
		x.UInt(0x996788ee)
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetUserId())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionBindAuthKeyUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionBindAuthKeyUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x996788ee:
		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Int()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionBindAuthKeyUser) DebugString() string {
	return "{}"
}

func (m *TLSessionUnbindAuthKeyUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	switch uint32(m.Constructor) {
	case 0xea96b132:
		x.UInt(0xea96b132)
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetUserId())

	default:
	}

	return x.GetBuf()
}

func (m *TLSessionUnbindAuthKeyUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSessionUnbindAuthKeyUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xea96b132:
		m.AuthKeyId = dBuf.Long()
		m.UserId = dBuf.Int()
		return dBuf.GetError()

	default:
	}
	return dBuf.GetError()
}

func (m *TLSessionUnbindAuthKeyUser) DebugString() string {
	return "{}"
}
