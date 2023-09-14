package syncpb

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TLConstructor int32

const (
	CRC32_UNKNOWN               TLConstructor = 0
	CRC32_sync_syncUpdates      TLConstructor = 1199756177
	CRC32_sync_pushUpdates      TLConstructor = 1549870665
	CRC32_sync_pushRpcResult    TLConstructor = 477907040
	CRC32_updates_getState      TLConstructor = 911585209
	CRC32_updates_getDifference TLConstructor = -1240060382
)

var TLConstructor_name = map[int32]string{
	0:           "CRC32_UNKNOWN",
	1199756177:  "CRC32_sync_syncUpdates",
	1549870665:  "CRC32_sync_pushUpdates",
	477907040:   "CRC32_sync_pushRpcResult",
	911585209:   "CRC32_updates_getState",
	-1240060382: "CRC32_updates_getDifference",
}
var TLConstructor_value = map[string]int32{
	"CRC32_UNKNOWN":               0,
	"CRC32_sync_syncUpdates":      1199756177,
	"CRC32_sync_pushUpdates":      1549870665,
	"CRC32_sync_pushRpcResult":    477907040,
	"CRC32_updates_getState":      911585209,
	"CRC32_updates_getDifference": -1240060382,
}

func (x TLConstructor) String() string {
	return proto.EnumName(TLConstructor_name, int32(x))
}
func (TLConstructor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_syncpb_tl_eb697f42b5d4823d, []int{0}
}

// --------------------------------------------------------------------------------------------
// sync.syncUpdates flags:# user_id:int auth_key_id:long server_id:flags.0?int updates:Updates = Bool;
type TLSyncSyncUpdates struct {
	Constructor          TLConstructor      `protobuf:"varint,1,opt,name=constructor,proto3,enum=syncpb.TLConstructor" json:"constructor,omitempty"`
	UserId               int32              `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthKeyId            int64              `protobuf:"varint,4,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	ServerId             *types.StringValue `protobuf:"bytes,5,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	SessionId            *types.Int64Value  `protobuf:"bytes,6,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Updates              *mtproto.Updates   `protobuf:"bytes,7,opt,name=updates" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TLSyncSyncUpdates) Reset()         { *m = TLSyncSyncUpdates{} }
func (m *TLSyncSyncUpdates) String() string { return proto.CompactTextString(m) }
func (*TLSyncSyncUpdates) ProtoMessage()    {}
func (*TLSyncSyncUpdates) Descriptor() ([]byte, []int) {
	return fileDescriptor_syncpb_tl_eb697f42b5d4823d, []int{0}
}
func (m *TLSyncSyncUpdates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLSyncSyncUpdates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLSyncSyncUpdates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLSyncSyncUpdates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSyncSyncUpdates.Merge(dst, src)
}
func (m *TLSyncSyncUpdates) XXX_Size() int {
	return m.Size()
}
func (m *TLSyncSyncUpdates) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSyncSyncUpdates.DiscardUnknown(m)
}

var xxx_messageInfo_TLSyncSyncUpdates proto.InternalMessageInfo

func (m *TLSyncSyncUpdates) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLSyncSyncUpdates) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TLSyncSyncUpdates) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *TLSyncSyncUpdates) GetServerId() *types.StringValue {
	if m != nil {
		return m.ServerId
	}
	return nil
}

func (m *TLSyncSyncUpdates) GetSessionId() *types.Int64Value {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *TLSyncSyncUpdates) GetUpdates() *mtproto.Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// sync.pushUpdates user_id:int updates:Updates = Bool;
type TLSyncPushUpdates struct {
	Constructor          TLConstructor    `protobuf:"varint,1,opt,name=constructor,proto3,enum=syncpb.TLConstructor" json:"constructor,omitempty"`
	UserId               int32            `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsBot                bool             `protobuf:"varint,4,opt,name=is_bot,json=isBot,proto3" json:"is_bot,omitempty"`
	Updates              *mtproto.Updates `protobuf:"bytes,5,opt,name=updates" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TLSyncPushUpdates) Reset()         { *m = TLSyncPushUpdates{} }
func (m *TLSyncPushUpdates) String() string { return proto.CompactTextString(m) }
func (*TLSyncPushUpdates) ProtoMessage()    {}
func (*TLSyncPushUpdates) Descriptor() ([]byte, []int) {
	return fileDescriptor_syncpb_tl_eb697f42b5d4823d, []int{1}
}
func (m *TLSyncPushUpdates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLSyncPushUpdates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLSyncPushUpdates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLSyncPushUpdates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSyncPushUpdates.Merge(dst, src)
}
func (m *TLSyncPushUpdates) XXX_Size() int {
	return m.Size()
}
func (m *TLSyncPushUpdates) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSyncPushUpdates.DiscardUnknown(m)
}

var xxx_messageInfo_TLSyncPushUpdates proto.InternalMessageInfo

func (m *TLSyncPushUpdates) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLSyncPushUpdates) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TLSyncPushUpdates) GetIsBot() bool {
	if m != nil {
		return m.IsBot
	}
	return false
}

func (m *TLSyncPushUpdates) GetUpdates() *mtproto.Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// sync.broadcastUpdates broadcast_type:int chat_id:int exclude_ids:Vector<int> updates:Updates = Bool;
type TLSyncBroadcastUpdates struct {
	Constructor          TLConstructor    `protobuf:"varint,1,opt,name=constructor,proto3,enum=syncpb.TLConstructor" json:"constructor,omitempty"`
	BroadcastType        int32            `protobuf:"varint,2,opt,name=broadcast_type,json=broadcastType,proto3" json:"broadcast_type,omitempty"`
	ChatId               int32            `protobuf:"varint,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	ExcludeIds           []int32          `protobuf:"varint,4,rep,packed,name=exclude_ids,json=excludeIds" json:"exclude_ids,omitempty"`
	Updates              *mtproto.Updates `protobuf:"bytes,5,opt,name=updates" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TLSyncBroadcastUpdates) Reset()         { *m = TLSyncBroadcastUpdates{} }
func (m *TLSyncBroadcastUpdates) String() string { return proto.CompactTextString(m) }
func (*TLSyncBroadcastUpdates) ProtoMessage()    {}
func (*TLSyncBroadcastUpdates) Descriptor() ([]byte, []int) {
	return fileDescriptor_syncpb_tl_eb697f42b5d4823d, []int{2}
}
func (m *TLSyncBroadcastUpdates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLSyncBroadcastUpdates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLSyncBroadcastUpdates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLSyncBroadcastUpdates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSyncBroadcastUpdates.Merge(dst, src)
}
func (m *TLSyncBroadcastUpdates) XXX_Size() int {
	return m.Size()
}
func (m *TLSyncBroadcastUpdates) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSyncBroadcastUpdates.DiscardUnknown(m)
}

var xxx_messageInfo_TLSyncBroadcastUpdates proto.InternalMessageInfo

func (m *TLSyncBroadcastUpdates) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLSyncBroadcastUpdates) GetBroadcastType() int32 {
	if m != nil {
		return m.BroadcastType
	}
	return 0
}

func (m *TLSyncBroadcastUpdates) GetChatId() int32 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *TLSyncBroadcastUpdates) GetExcludeIds() []int32 {
	if m != nil {
		return m.ExcludeIds
	}
	return nil
}

func (m *TLSyncBroadcastUpdates) GetUpdates() *mtproto.Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// sync.pushRpcResult server_id:int auth_key_id:long req_msg_id:long result:bytes = Bool;
type TLSyncPushRpcResult struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=syncpb.TLConstructor" json:"constructor,omitempty"`
	ServerId             string        `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	AuthKeyId            int64         `protobuf:"varint,4,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	SessionId            int64         `protobuf:"varint,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ReqMsgId             int64         `protobuf:"varint,6,opt,name=req_msg_id,json=reqMsgId,proto3" json:"req_msg_id,omitempty"`
	Result               []byte        `protobuf:"bytes,7,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLSyncPushRpcResult) Reset()         { *m = TLSyncPushRpcResult{} }
func (m *TLSyncPushRpcResult) String() string { return proto.CompactTextString(m) }
func (*TLSyncPushRpcResult) ProtoMessage()    {}
func (*TLSyncPushRpcResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_syncpb_tl_eb697f42b5d4823d, []int{3}
}
func (m *TLSyncPushRpcResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLSyncPushRpcResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLSyncPushRpcResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLSyncPushRpcResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSyncPushRpcResult.Merge(dst, src)
}
func (m *TLSyncPushRpcResult) XXX_Size() int {
	return m.Size()
}
func (m *TLSyncPushRpcResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSyncPushRpcResult.DiscardUnknown(m)
}

var xxx_messageInfo_TLSyncPushRpcResult proto.InternalMessageInfo

func (m *TLSyncPushRpcResult) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLSyncPushRpcResult) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *TLSyncPushRpcResult) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *TLSyncPushRpcResult) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *TLSyncPushRpcResult) GetReqMsgId() int64 {
	if m != nil {
		return m.ReqMsgId
	}
	return 0
}

func (m *TLSyncPushRpcResult) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// updates.getState auth_key_id:long user_id:int = updates.State;
type TLUpdatesGetState struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=syncpb.TLConstructor" json:"constructor,omitempty"`
	AuthKeyId            int64         `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	UserId               int32         `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLUpdatesGetState) Reset()         { *m = TLUpdatesGetState{} }
func (m *TLUpdatesGetState) String() string { return proto.CompactTextString(m) }
func (*TLUpdatesGetState) ProtoMessage()    {}
func (*TLUpdatesGetState) Descriptor() ([]byte, []int) {
	return fileDescriptor_syncpb_tl_eb697f42b5d4823d, []int{4}
}
func (m *TLUpdatesGetState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLUpdatesGetState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLUpdatesGetState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLUpdatesGetState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLUpdatesGetState.Merge(dst, src)
}
func (m *TLUpdatesGetState) XXX_Size() int {
	return m.Size()
}
func (m *TLUpdatesGetState) XXX_DiscardUnknown() {
	xxx_messageInfo_TLUpdatesGetState.DiscardUnknown(m)
}

var xxx_messageInfo_TLUpdatesGetState proto.InternalMessageInfo

func (m *TLUpdatesGetState) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLUpdatesGetState) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *TLUpdatesGetState) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// --------------------------------------------------------------------------------------------
// updates.getDifference flags:# auth_key_id:long user_id:int pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
type TLUpdatesGetDifference struct {
	Constructor          TLConstructor     `protobuf:"varint,1,opt,name=constructor,proto3,enum=syncpb.TLConstructor" json:"constructor,omitempty"`
	AuthKeyId            int64             `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	UserId               int32             `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Pts                  int32             `protobuf:"varint,5,opt,name=pts,proto3" json:"pts,omitempty"`
	PtsTotalLimit        *types.Int32Value `protobuf:"bytes,6,opt,name=pts_total_limit,json=ptsTotalLimit" json:"pts_total_limit,omitempty"`
	Date                 int32             `protobuf:"varint,7,opt,name=date,proto3" json:"date,omitempty"`
	Qts                  int32             `protobuf:"varint,8,opt,name=qts,proto3" json:"qts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TLUpdatesGetDifference) Reset()         { *m = TLUpdatesGetDifference{} }
func (m *TLUpdatesGetDifference) String() string { return proto.CompactTextString(m) }
func (*TLUpdatesGetDifference) ProtoMessage()    {}
func (*TLUpdatesGetDifference) Descriptor() ([]byte, []int) {
	return fileDescriptor_syncpb_tl_eb697f42b5d4823d, []int{5}
}
func (m *TLUpdatesGetDifference) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLUpdatesGetDifference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLUpdatesGetDifference.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLUpdatesGetDifference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLUpdatesGetDifference.Merge(dst, src)
}
func (m *TLUpdatesGetDifference) XXX_Size() int {
	return m.Size()
}
func (m *TLUpdatesGetDifference) XXX_DiscardUnknown() {
	xxx_messageInfo_TLUpdatesGetDifference.DiscardUnknown(m)
}

var xxx_messageInfo_TLUpdatesGetDifference proto.InternalMessageInfo

func (m *TLUpdatesGetDifference) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLUpdatesGetDifference) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *TLUpdatesGetDifference) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TLUpdatesGetDifference) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

func (m *TLUpdatesGetDifference) GetPtsTotalLimit() *types.Int32Value {
	if m != nil {
		return m.PtsTotalLimit
	}
	return nil
}

func (m *TLUpdatesGetDifference) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *TLUpdatesGetDifference) GetQts() int32 {
	if m != nil {
		return m.Qts
	}
	return 0
}

func init() {
	proto.RegisterType((*TLSyncSyncUpdates)(nil), "syncpb.TL_sync_syncUpdates")
	proto.RegisterType((*TLSyncPushUpdates)(nil), "syncpb.TL_sync_pushUpdates")
	proto.RegisterType((*TLSyncBroadcastUpdates)(nil), "syncpb.TL_sync_broadcastUpdates")
	proto.RegisterType((*TLSyncPushRpcResult)(nil), "syncpb.TL_sync_pushRpcResult")
	proto.RegisterType((*TLUpdatesGetState)(nil), "syncpb.TL_updates_getState")
	proto.RegisterType((*TLUpdatesGetDifference)(nil), "syncpb.TL_updates_getDifference")
	proto.RegisterEnum("syncpb.TLConstructor", TLConstructor_name, TLConstructor_value)
}
func (this *TLSyncSyncUpdates) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&syncpb.TLSyncSyncUpdates{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",")
	s = append(s, "AuthKeyId: "+fmt.Sprintf("%#v", this.AuthKeyId)+",")
	if this.ServerId != nil {
		s = append(s, "ServerId: "+fmt.Sprintf("%#v", this.ServerId)+",")
	}
	if this.SessionId != nil {
		s = append(s, "SessionId: "+fmt.Sprintf("%#v", this.SessionId)+",")
	}
	if this.Updates != nil {
		s = append(s, "Updates: "+fmt.Sprintf("%#v", this.Updates)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLSyncPushUpdates) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&syncpb.TLSyncPushUpdates{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",")
	s = append(s, "IsBot: "+fmt.Sprintf("%#v", this.IsBot)+",")
	if this.Updates != nil {
		s = append(s, "Updates: "+fmt.Sprintf("%#v", this.Updates)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLSyncBroadcastUpdates) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&syncpb.TLSyncBroadcastUpdates{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",")
	s = append(s, "BroadcastType: "+fmt.Sprintf("%#v", this.BroadcastType)+",")
	s = append(s, "ChatId: "+fmt.Sprintf("%#v", this.ChatId)+",")
	s = append(s, "ExcludeIds: "+fmt.Sprintf("%#v", this.ExcludeIds)+",")
	if this.Updates != nil {
		s = append(s, "Updates: "+fmt.Sprintf("%#v", this.Updates)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLSyncPushRpcResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&syncpb.TLSyncPushRpcResult{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",")
	s = append(s, "ServerId: "+fmt.Sprintf("%#v", this.ServerId)+",")
	s = append(s, "AuthKeyId: "+fmt.Sprintf("%#v", this.AuthKeyId)+",")
	s = append(s, "SessionId: "+fmt.Sprintf("%#v", this.SessionId)+",")
	s = append(s, "ReqMsgId: "+fmt.Sprintf("%#v", this.ReqMsgId)+",")
	s = append(s, "Result: "+fmt.Sprintf("%#v", this.Result)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLUpdatesGetState) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&syncpb.TLUpdatesGetState{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",")
	s = append(s, "AuthKeyId: "+fmt.Sprintf("%#v", this.AuthKeyId)+",")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLUpdatesGetDifference) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&syncpb.TLUpdatesGetDifference{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",")
	s = append(s, "AuthKeyId: "+fmt.Sprintf("%#v", this.AuthKeyId)+",")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",")
	s = append(s, "Pts: "+fmt.Sprintf("%#v", this.Pts)+",")
	if this.PtsTotalLimit != nil {
		s = append(s, "PtsTotalLimit: "+fmt.Sprintf("%#v", this.PtsTotalLimit)+",")
	}
	s = append(s, "Date: "+fmt.Sprintf("%#v", this.Date)+",")
	s = append(s, "Qts: "+fmt.Sprintf("%#v", this.Qts)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSyncpbTl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TLSyncSyncUpdates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLSyncSyncUpdates) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.UserId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.UserId))
	}
	if m.AuthKeyId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.AuthKeyId))
	}
	if m.ServerId != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.ServerId.Size()))
		n1, err := m.ServerId.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.SessionId != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.SessionId.Size()))
		n2, err := m.SessionId.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Updates != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Updates.Size()))
		n3, err := m.Updates.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLSyncPushUpdates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLSyncPushUpdates) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.UserId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.UserId))
	}
	if m.IsBot {
		dAtA[i] = 0x20
		i++
		if m.IsBot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Updates != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Updates.Size()))
		n4, err := m.Updates.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLSyncBroadcastUpdates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLSyncBroadcastUpdates) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.BroadcastType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.BroadcastType))
	}
	if m.ChatId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.ChatId))
	}
	if len(m.ExcludeIds) > 0 {
		dAtA6 := make([]byte, len(m.ExcludeIds)*10)
		var j5 int
		for _, num1 := range m.ExcludeIds {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(j5))
		i += copy(dAtA[i:], dAtA6[:j5])
	}
	if m.Updates != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Updates.Size()))
		n7, err := m.Updates.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLSyncPushRpcResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLSyncPushRpcResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Constructor))
	}
	if len(m.ServerId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(len(m.ServerId)))
		i += copy(dAtA[i:], m.ServerId)
	}
	if m.AuthKeyId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.AuthKeyId))
	}
	if m.SessionId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.SessionId))
	}
	if m.ReqMsgId != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.ReqMsgId))
	}
	if len(m.Result) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(len(m.Result)))
		i += copy(dAtA[i:], m.Result)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLUpdatesGetState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLUpdatesGetState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.AuthKeyId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.AuthKeyId))
	}
	if m.UserId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.UserId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLUpdatesGetDifference) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLUpdatesGetDifference) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.AuthKeyId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.AuthKeyId))
	}
	if m.UserId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.UserId))
	}
	if m.Pts != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Pts))
	}
	if m.PtsTotalLimit != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.PtsTotalLimit.Size()))
		n8, err := m.PtsTotalLimit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.Date != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Date))
	}
	if m.Qts != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintSyncpbTl(dAtA, i, uint64(m.Qts))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSyncpbTl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TLSyncSyncUpdates) Size() (n int) {
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Constructor))
	}
	if m.UserId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.UserId))
	}
	if m.AuthKeyId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.AuthKeyId))
	}
	if m.ServerId != nil {
		l = m.ServerId.Size()
		n += 1 + l + sovSyncpbTl(uint64(l))
	}
	if m.SessionId != nil {
		l = m.SessionId.Size()
		n += 1 + l + sovSyncpbTl(uint64(l))
	}
	if m.Updates != nil {
		l = m.Updates.Size()
		n += 1 + l + sovSyncpbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLSyncPushUpdates) Size() (n int) {
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Constructor))
	}
	if m.UserId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.UserId))
	}
	if m.IsBot {
		n += 2
	}
	if m.Updates != nil {
		l = m.Updates.Size()
		n += 1 + l + sovSyncpbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLSyncBroadcastUpdates) Size() (n int) {
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Constructor))
	}
	if m.BroadcastType != 0 {
		n += 1 + sovSyncpbTl(uint64(m.BroadcastType))
	}
	if m.ChatId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.ChatId))
	}
	if len(m.ExcludeIds) > 0 {
		l = 0
		for _, e := range m.ExcludeIds {
			l += sovSyncpbTl(uint64(e))
		}
		n += 1 + sovSyncpbTl(uint64(l)) + l
	}
	if m.Updates != nil {
		l = m.Updates.Size()
		n += 1 + l + sovSyncpbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLSyncPushRpcResult) Size() (n int) {
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Constructor))
	}
	l = len(m.ServerId)
	if l > 0 {
		n += 1 + l + sovSyncpbTl(uint64(l))
	}
	if m.AuthKeyId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.AuthKeyId))
	}
	if m.SessionId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.SessionId))
	}
	if m.ReqMsgId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.ReqMsgId))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovSyncpbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLUpdatesGetState) Size() (n int) {
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Constructor))
	}
	if m.AuthKeyId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.AuthKeyId))
	}
	if m.UserId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.UserId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLUpdatesGetDifference) Size() (n int) {
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Constructor))
	}
	if m.AuthKeyId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.AuthKeyId))
	}
	if m.UserId != 0 {
		n += 1 + sovSyncpbTl(uint64(m.UserId))
	}
	if m.Pts != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Pts))
	}
	if m.PtsTotalLimit != nil {
		l = m.PtsTotalLimit.Size()
		n += 1 + l + sovSyncpbTl(uint64(l))
	}
	if m.Date != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Date))
	}
	if m.Qts != 0 {
		n += 1 + sovSyncpbTl(uint64(m.Qts))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSyncpbTl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSyncpbTl(x uint64) (n int) {
	return sovSyncpbTl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TLSyncSyncUpdates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncpbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_sync_syncUpdates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_sync_syncUpdates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthKeyId", wireType)
			}
			m.AuthKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthKeyId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ServerId == nil {
				m.ServerId = &types.StringValue{}
			}
			if err := m.ServerId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SessionId == nil {
				m.SessionId = &types.Int64Value{}
			}
			if err := m.SessionId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Updates == nil {
				m.Updates = &mtproto.Updates{}
			}
			if err := m.Updates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSyncpbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLSyncPushUpdates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncpbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_sync_pushUpdates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_sync_pushUpdates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsBot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsBot = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Updates == nil {
				m.Updates = &mtproto.Updates{}
			}
			if err := m.Updates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSyncpbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLSyncBroadcastUpdates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncpbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_sync_broadcastUpdates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_sync_broadcastUpdates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BroadcastType", wireType)
			}
			m.BroadcastType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BroadcastType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChatId", wireType)
			}
			m.ChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSyncpbTl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ExcludeIds = append(m.ExcludeIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSyncpbTl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSyncpbTl
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSyncpbTl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ExcludeIds = append(m.ExcludeIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ExcludeIds", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Updates == nil {
				m.Updates = &mtproto.Updates{}
			}
			if err := m.Updates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSyncpbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLSyncPushRpcResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncpbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_sync_pushRpcResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_sync_pushRpcResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthKeyId", wireType)
			}
			m.AuthKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthKeyId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			m.SessionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqMsgId", wireType)
			}
			m.ReqMsgId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqMsgId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSyncpbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLUpdatesGetState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncpbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_updates_getState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_updates_getState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthKeyId", wireType)
			}
			m.AuthKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthKeyId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSyncpbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLUpdatesGetDifference) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncpbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_updates_getDifference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_updates_getDifference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthKeyId", wireType)
			}
			m.AuthKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthKeyId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pts", wireType)
			}
			m.Pts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pts |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PtsTotalLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PtsTotalLimit == nil {
				m.PtsTotalLimit = &types.Int32Value{}
			}
			if err := m.PtsTotalLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			m.Date = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Date |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qts", wireType)
			}
			m.Qts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Qts |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSyncpbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncpbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSyncpbTl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSyncpbTl
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSyncpbTl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSyncpbTl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSyncpbTl
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSyncpbTl(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSyncpbTl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSyncpbTl   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("syncpb.tl.proto", fileDescriptor_syncpb_tl_eb697f42b5d4823d) }

var fileDescriptor_syncpb_tl_eb697f42b5d4823d = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcf, 0x6b, 0x13, 0x4b,
	0x1c, 0xef, 0x36, 0xd9, 0x34, 0x99, 0xbc, 0xb4, 0xfb, 0xe6, 0xbd, 0xf6, 0x2d, 0x6d, 0xdf, 0xbe,
	0x10, 0x78, 0x10, 0x0a, 0xdd, 0x40, 0xfa, 0x78, 0xa2, 0x17, 0xa1, 0xf1, 0xb2, 0xb4, 0x56, 0xd8,
	0xa6, 0x0a, 0x5e, 0x96, 0xcd, 0xee, 0x74, 0xb3, 0xb8, 0xd9, 0xdd, 0xcc, 0xcc, 0xaa, 0xb9, 0x0b,
	0x2a, 0x5e, 0xf4, 0x5f, 0xd0, 0x83, 0xd0, 0x93, 0x78, 0x13, 0xfc, 0x03, 0x14, 0x3c, 0x78, 0x10,
	0xf1, 0x20, 0x58, 0xd3, 0x7f, 0x40, 0xcf, 0x1e, 0x22, 0x33, 0x93, 0x34, 0xdb, 0x14, 0x2c, 0x56,
	0xc4, 0x4b, 0x98, 0xef, 0x8f, 0xcf, 0xce, 0xe7, 0xf3, 0x9d, 0x4f, 0xbe, 0x60, 0x8e, 0xf4, 0x42,
	0x27, 0x6e, 0xe9, 0x34, 0xd0, 0x63, 0x1c, 0xd1, 0x08, 0xe6, 0x44, 0x62, 0x71, 0xd5, 0xf3, 0x69,
	0x3b, 0x69, 0xe9, 0x4e, 0xd4, 0xa9, 0x79, 0x91, 0x17, 0xd5, 0x78, 0xb9, 0x95, 0xec, 0xf2, 0x88,
	0x07, 0xfc, 0x24, 0x60, 0x8b, 0x9a, 0x17, 0x45, 0x5e, 0x80, 0xc6, 0x5d, 0x37, 0xb0, 0x1d, 0xc7,
	0x08, 0x93, 0x61, 0xfd, 0x4f, 0xe2, 0xb4, 0x51, 0xc7, 0x66, 0xf7, 0xb0, 0x0b, 0x44, 0xb6, 0xb2,
	0x37, 0x0d, 0xfe, 0x68, 0x6e, 0x5a, 0x2c, 0xc3, 0x7f, 0x76, 0x62, 0xd7, 0xa6, 0x88, 0xc0, 0x33,
	0xa0, 0xe8, 0x44, 0x21, 0xa1, 0x38, 0x71, 0x68, 0x84, 0x55, 0xa9, 0x2c, 0x55, 0x67, 0xeb, 0xf3,
	0xfa, 0x90, 0x6b, 0x73, 0xb3, 0x31, 0x2e, 0x9a, 0xe9, 0x4e, 0xf8, 0x17, 0x98, 0x49, 0x08, 0xc2,
	0x96, 0xef, 0xaa, 0x99, 0xb2, 0x54, 0x95, 0xcd, 0x1c, 0x0b, 0x0d, 0x17, 0x6a, 0xa0, 0x68, 0x27,
	0xb4, 0x6d, 0x5d, 0x43, 0x3d, 0x56, 0xcc, 0x96, 0xa5, 0x6a, 0xc6, 0x2c, 0xb0, 0xd4, 0x06, 0xea,
	0x19, 0x2e, 0x3c, 0x0b, 0x0a, 0x04, 0xe1, 0xeb, 0x02, 0x2a, 0x97, 0xa5, 0x6a, 0xb1, 0xbe, 0xac,
	0x0b, 0x4d, 0xfa, 0x48, 0x93, 0xbe, 0x4d, 0xb1, 0x1f, 0x7a, 0x97, 0xed, 0x20, 0x41, 0x66, 0x5e,
	0xb4, 0x1b, 0x2e, 0x3c, 0x07, 0x00, 0x41, 0x84, 0xf8, 0x51, 0xc8, 0xb0, 0x39, 0x8e, 0x5d, 0x3a,
	0x86, 0x35, 0x42, 0xfa, 0xff, 0x7f, 0x02, 0x5a, 0x18, 0xb6, 0x1b, 0x2e, 0x5c, 0x01, 0x33, 0x89,
	0xd0, 0xac, 0xce, 0x70, 0xa0, 0xa2, 0x77, 0x28, 0xc7, 0xe8, 0xc3, 0x59, 0x98, 0xa3, 0x86, 0xca,
	0x9e, 0x34, 0x1e, 0x56, 0x9c, 0x90, 0xf6, 0xcf, 0x1b, 0xd6, 0x3c, 0xc8, 0xf9, 0xc4, 0x6a, 0x45,
	0x94, 0xcf, 0x29, 0x6f, 0xca, 0x3e, 0x59, 0x8f, 0x68, 0x9a, 0xac, 0x7c, 0x12, 0xd9, 0x7d, 0x09,
	0xa8, 0x23, 0xb2, 0x2d, 0x1c, 0xd9, 0xae, 0x63, 0x13, 0xfa, 0xc3, 0x8c, 0xff, 0x05, 0xb3, 0x87,
	0x1f, 0xb3, 0x68, 0x2f, 0x46, 0xea, 0x34, 0x27, 0x5e, 0x3a, 0xcc, 0x36, 0x7b, 0x31, 0x62, 0xc2,
	0x9c, 0xb6, 0x4d, 0x53, 0xc2, 0x58, 0x68, 0xb8, 0xf0, 0x1f, 0x50, 0x44, 0x37, 0x9d, 0x20, 0x71,
	0x91, 0xe5, 0xbb, 0x44, 0xcd, 0x96, 0x33, 0x55, 0xd9, 0x04, 0xc3, 0x94, 0xe1, 0x92, 0xef, 0x92,
	0xd8, 0x97, 0xc0, 0x7c, 0xfa, 0x3d, 0xcc, 0xd8, 0x31, 0x11, 0x49, 0x02, 0x7a, 0x7a, 0x7d, 0x4b,
	0x69, 0x17, 0x32, 0xea, 0x85, 0x94, 0xcf, 0x4e, 0xb2, 0xf0, 0xdf, 0x47, 0x7c, 0x28, 0x8b, 0xf2,
	0xd8, 0x6a, 0xcb, 0x00, 0x60, 0xd4, 0xb5, 0x3a, 0xc4, 0x1b, 0xd9, 0x34, 0x63, 0xe6, 0x31, 0xea,
	0x5e, 0x24, 0x9e, 0xe1, 0xc2, 0x05, 0x90, 0xc3, 0x9c, 0x3c, 0xf7, 0xe1, 0x6f, 0xe6, 0x30, 0xaa,
	0xdc, 0x16, 0xa6, 0x1b, 0x6a, 0xb6, 0x3c, 0x44, 0xb7, 0xa9, 0x4d, 0xd1, 0xe9, 0x25, 0x4e, 0xa8,
	0xc8, 0x4c, 0xaa, 0x48, 0x99, 0x32, 0x9b, 0x36, 0x65, 0xe5, 0xee, 0x34, 0x77, 0x54, 0x8a, 0xc9,
	0x05, 0x7f, 0x77, 0x17, 0x61, 0x14, 0x3a, 0xbf, 0x80, 0x0e, 0x54, 0x40, 0x26, 0xa6, 0xc2, 0x25,
	0xb2, 0xc9, 0x8e, 0xb0, 0x01, 0xe6, 0x62, 0x4a, 0x2c, 0x1a, 0x51, 0x3b, 0xb0, 0x02, 0xbf, 0xe3,
	0xd3, 0x6f, 0x2d, 0x83, 0xb5, 0xba, 0x58, 0x06, 0xa5, 0x98, 0x92, 0x26, 0x83, 0x6c, 0x32, 0x04,
	0x84, 0x20, 0xcb, 0xf4, 0xf1, 0x57, 0x90, 0x4d, 0x7e, 0x66, 0x57, 0x75, 0x29, 0x51, 0xf3, 0xe2,
	0xaa, 0x2e, 0x25, 0x2b, 0x6f, 0x24, 0x50, 0x3a, 0x22, 0x0a, 0xfe, 0x0e, 0x4a, 0x0d, 0xb3, 0xb1,
	0x56, 0xb7, 0x76, 0xb6, 0x36, 0xb6, 0x2e, 0x5d, 0xd9, 0x52, 0xa6, 0xa0, 0x06, 0x16, 0x44, 0x6a,
	0x72, 0xbd, 0x2a, 0x0f, 0x1e, 0xdf, 0x7b, 0x9e, 0x9d, 0xa8, 0xa7, 0x36, 0x8a, 0xf2, 0xf2, 0xd5,
	0xad, 0x03, 0x19, 0x96, 0x81, 0x3a, 0x51, 0x3f, 0x74, 0xb8, 0xb2, 0x7f, 0xff, 0xf3, 0x81, 0x34,
	0xfe, 0xc2, 0xa4, 0x3d, 0x94, 0x67, 0x1f, 0xde, 0x3e, 0xcd, 0xc0, 0x2a, 0x58, 0x3a, 0x56, 0x1f,
	0x3f, 0x9a, 0xf2, 0xf0, 0xfd, 0xbb, 0x27, 0x5f, 0x06, 0x83, 0xc1, 0x40, 0x5a, 0xcc, 0xde, 0x79,
	0xa4, 0x4d, 0xad, 0x9f, 0xff, 0xf4, 0x51, 0x93, 0x5e, 0xf4, 0x35, 0xe9, 0x75, 0x5f, 0x93, 0xf6,
	0xfb, 0x9a, 0x74, 0x75, 0x35, 0x44, 0xad, 0x24, 0xb0, 0x75, 0xf6, 0xff, 0xad, 0x39, 0x41, 0x42,
	0x28, 0xc2, 0x35, 0x3b, 0x8e, 0x6b, 0x1d, 0x44, 0x08, 0x0a, 0x3d, 0x84, 0x6b, 0x8c, 0x5d, 0x4d,
	0x3c, 0x74, 0x2b, 0xc7, 0x27, 0xbc, 0xf6, 0x35, 0x00, 0x00, 0xff, 0xff, 0x45, 0x69, 0x7c, 0x4b,
	0xd6, 0x06, 0x00, 0x00,
}
