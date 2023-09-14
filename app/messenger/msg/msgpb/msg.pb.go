package msgpb

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

// /////////////////////////////////////////////////////////////////////////////
type Sender struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	AuthKeyId            int64    `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sender) Reset()         { *m = Sender{} }
func (m *Sender) String() string { return proto.CompactTextString(m) }
func (*Sender) ProtoMessage()    {}
func (*Sender) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{0}
}
func (m *Sender) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sender) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sender.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Sender) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sender.Merge(dst, src)
}
func (m *Sender) XXX_Size() int {
	return m.Size()
}
func (m *Sender) XXX_DiscardUnknown() {
	xxx_messageInfo_Sender.DiscardUnknown(m)
}

var xxx_messageInfo_Sender proto.InternalMessageInfo

func (m *Sender) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Sender) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Sender) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

// /////////////////////////////////////////////////////////////////////////////
type UserMessage struct {
	From                 *Sender          `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerUserId           int32            `protobuf:"varint,2,opt,name=peer_user_id,json=peerUserId,proto3" json:"peer_user_id,omitempty"`
	RandomId             int64            `protobuf:"varint,3,opt,name=random_id,json=randomId,proto3" json:"random_id,omitempty"`
	Message              *mtproto.Message `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}
func (*UserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{1}
}
func (m *UserMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMessage.Merge(dst, src)
}
func (m *UserMessage) XXX_Size() int {
	return m.Size()
}
func (m *UserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMessage proto.InternalMessageInfo

func (m *UserMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *UserMessage) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *UserMessage) GetRandomId() int64 {
	if m != nil {
		return m.RandomId
	}
	return 0
}

func (m *UserMessage) GetMessage() *mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ChatMessage struct {
	From                 *Sender          `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChatId           int32            `protobuf:"varint,2,opt,name=peer_chat_id,json=peerChatId,proto3" json:"peer_chat_id,omitempty"`
	RandomId             int64            `protobuf:"varint,3,opt,name=random_id,json=randomId,proto3" json:"random_id,omitempty"`
	Message              *mtproto.Message `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{2}
}
func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(dst, src)
}
func (m *ChatMessage) XXX_Size() int {
	return m.Size()
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ChatMessage) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *ChatMessage) GetRandomId() int64 {
	if m != nil {
		return m.RandomId
	}
	return 0
}

func (m *ChatMessage) GetMessage() *mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ChannelMessage struct {
	From                 *Sender          `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChannelId        int32            `protobuf:"varint,2,opt,name=peer_channel_id,json=peerChannelId,proto3" json:"peer_channel_id,omitempty"`
	RandomId             int64            `protobuf:"varint,3,opt,name=random_id,json=randomId,proto3" json:"random_id,omitempty"`
	Message              *mtproto.Message `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChannelMessage) Reset()         { *m = ChannelMessage{} }
func (m *ChannelMessage) String() string { return proto.CompactTextString(m) }
func (*ChannelMessage) ProtoMessage()    {}
func (*ChannelMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{3}
}
func (m *ChannelMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChannelMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChannelMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ChannelMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelMessage.Merge(dst, src)
}
func (m *ChannelMessage) XXX_Size() int {
	return m.Size()
}
func (m *ChannelMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelMessage proto.InternalMessageInfo

func (m *ChannelMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ChannelMessage) GetPeerChannelId() int32 {
	if m != nil {
		return m.PeerChannelId
	}
	return 0
}

func (m *ChannelMessage) GetRandomId() int64 {
	if m != nil {
		return m.RandomId
	}
	return 0
}

func (m *ChannelMessage) GetMessage() *mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type UserMultiMessage struct {
	From                 *Sender            `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerUserId           int32              `protobuf:"varint,2,opt,name=peer_user_id,json=peerUserId,proto3" json:"peer_user_id,omitempty"`
	RandomId             []int64            `protobuf:"varint,3,rep,packed,name=random_id,json=randomId" json:"random_id,omitempty"`
	Message              []*mtproto.Message `protobuf:"bytes,4,rep,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UserMultiMessage) Reset()         { *m = UserMultiMessage{} }
func (m *UserMultiMessage) String() string { return proto.CompactTextString(m) }
func (*UserMultiMessage) ProtoMessage()    {}
func (*UserMultiMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{4}
}
func (m *UserMultiMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserMultiMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserMultiMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserMultiMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMultiMessage.Merge(dst, src)
}
func (m *UserMultiMessage) XXX_Size() int {
	return m.Size()
}
func (m *UserMultiMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMultiMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMultiMessage proto.InternalMessageInfo

func (m *UserMultiMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *UserMultiMessage) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *UserMultiMessage) GetRandomId() []int64 {
	if m != nil {
		return m.RandomId
	}
	return nil
}

func (m *UserMultiMessage) GetMessage() []*mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ChatMultiMessage struct {
	From                 *Sender            `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChatId           int32              `protobuf:"varint,2,opt,name=peer_chat_id,json=peerChatId,proto3" json:"peer_chat_id,omitempty"`
	RandomId             []int64            `protobuf:"varint,3,rep,packed,name=random_id,json=randomId" json:"random_id,omitempty"`
	Message              []*mtproto.Message `protobuf:"bytes,4,rep,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChatMultiMessage) Reset()         { *m = ChatMultiMessage{} }
func (m *ChatMultiMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMultiMessage) ProtoMessage()    {}
func (*ChatMultiMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{5}
}
func (m *ChatMultiMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatMultiMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatMultiMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ChatMultiMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMultiMessage.Merge(dst, src)
}
func (m *ChatMultiMessage) XXX_Size() int {
	return m.Size()
}
func (m *ChatMultiMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMultiMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMultiMessage proto.InternalMessageInfo

func (m *ChatMultiMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ChatMultiMessage) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *ChatMultiMessage) GetRandomId() []int64 {
	if m != nil {
		return m.RandomId
	}
	return nil
}

func (m *ChatMultiMessage) GetMessage() []*mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ChannelMultiMessage struct {
	From                 *Sender            `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChannelId        int32              `protobuf:"varint,2,opt,name=peer_channel_id,json=peerChannelId,proto3" json:"peer_channel_id,omitempty"`
	RandomId             []int64            `protobuf:"varint,3,rep,packed,name=random_id,json=randomId" json:"random_id,omitempty"`
	Message              []*mtproto.Message `protobuf:"bytes,4,rep,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChannelMultiMessage) Reset()         { *m = ChannelMultiMessage{} }
func (m *ChannelMultiMessage) String() string { return proto.CompactTextString(m) }
func (*ChannelMultiMessage) ProtoMessage()    {}
func (*ChannelMultiMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{6}
}
func (m *ChannelMultiMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChannelMultiMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChannelMultiMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ChannelMultiMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelMultiMessage.Merge(dst, src)
}
func (m *ChannelMultiMessage) XXX_Size() int {
	return m.Size()
}
func (m *ChannelMultiMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelMultiMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelMultiMessage proto.InternalMessageInfo

func (m *ChannelMultiMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ChannelMultiMessage) GetPeerChannelId() int32 {
	if m != nil {
		return m.PeerChannelId
	}
	return 0
}

func (m *ChannelMultiMessage) GetRandomId() []int64 {
	if m != nil {
		return m.RandomId
	}
	return nil
}

func (m *ChannelMultiMessage) GetMessage() []*mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type OutboxMessage struct {
	NoWebpage            bool              `protobuf:"varint,1,opt,name=no_webpage,json=noWebpage,proto3" json:"no_webpage,omitempty"`
	Background           bool              `protobuf:"varint,2,opt,name=background,proto3" json:"background,omitempty"`
	RandomId             int64             `protobuf:"varint,3,opt,name=random_id,json=randomId,proto3" json:"random_id,omitempty"`
	Message              *mtproto.Message  `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	ScheduleDate         *types.Int32Value `protobuf:"bytes,5,opt,name=schedule_date,json=scheduleDate" json:"schedule_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OutboxMessage) Reset()         { *m = OutboxMessage{} }
func (m *OutboxMessage) String() string { return proto.CompactTextString(m) }
func (*OutboxMessage) ProtoMessage()    {}
func (*OutboxMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{7}
}
func (m *OutboxMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutboxMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutboxMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *OutboxMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutboxMessage.Merge(dst, src)
}
func (m *OutboxMessage) XXX_Size() int {
	return m.Size()
}
func (m *OutboxMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OutboxMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OutboxMessage proto.InternalMessageInfo

func (m *OutboxMessage) GetNoWebpage() bool {
	if m != nil {
		return m.NoWebpage
	}
	return false
}

func (m *OutboxMessage) GetBackground() bool {
	if m != nil {
		return m.Background
	}
	return false
}

func (m *OutboxMessage) GetRandomId() int64 {
	if m != nil {
		return m.RandomId
	}
	return 0
}

func (m *OutboxMessage) GetMessage() *mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *OutboxMessage) GetScheduleDate() *types.Int32Value {
	if m != nil {
		return m.ScheduleDate
	}
	return nil
}

type OutgoingMessage struct {
	From                 *Sender        `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerType             int32          `protobuf:"varint,2,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	PeerId               int32          `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Message              *OutboxMessage `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OutgoingMessage) Reset()         { *m = OutgoingMessage{} }
func (m *OutgoingMessage) String() string { return proto.CompactTextString(m) }
func (*OutgoingMessage) ProtoMessage()    {}
func (*OutgoingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{8}
}
func (m *OutgoingMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *OutgoingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingMessage.Merge(dst, src)
}
func (m *OutgoingMessage) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingMessage proto.InternalMessageInfo

func (m *OutgoingMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *OutgoingMessage) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *OutgoingMessage) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *OutgoingMessage) GetMessage() *OutboxMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type OutgoingMultiMessage struct {
	From                 *Sender          `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerType             int32            `protobuf:"varint,2,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	PeerId               int32            `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	MultiMessage         []*OutboxMessage `protobuf:"bytes,4,rep,name=multi_message,json=multiMessage" json:"multi_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *OutgoingMultiMessage) Reset()         { *m = OutgoingMultiMessage{} }
func (m *OutgoingMultiMessage) String() string { return proto.CompactTextString(m) }
func (*OutgoingMultiMessage) ProtoMessage()    {}
func (*OutgoingMultiMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{9}
}
func (m *OutgoingMultiMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingMultiMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingMultiMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *OutgoingMultiMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingMultiMessage.Merge(dst, src)
}
func (m *OutgoingMultiMessage) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingMultiMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingMultiMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingMultiMessage proto.InternalMessageInfo

func (m *OutgoingMultiMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *OutgoingMultiMessage) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *OutgoingMultiMessage) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *OutgoingMultiMessage) GetMultiMessage() []*OutboxMessage {
	if m != nil {
		return m.MultiMessage
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
type InboxUserMessage struct {
	From                 *Sender          `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerUserId           int32            `protobuf:"varint,2,opt,name=peer_user_id,json=peerUserId,proto3" json:"peer_user_id,omitempty"`
	RandomId             int64            `protobuf:"varint,3,opt,name=random_id,json=randomId,proto3" json:"random_id,omitempty"`
	DialogMessageId      int32            `protobuf:"varint,4,opt,name=dialog_message_id,json=dialogMessageId,proto3" json:"dialog_message_id,omitempty"`
	MessageDataId        int64            `protobuf:"varint,5,opt,name=message_data_id,json=messageDataId,proto3" json:"message_data_id,omitempty"`
	Message              *mtproto.Message `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InboxUserMessage) Reset()         { *m = InboxUserMessage{} }
func (m *InboxUserMessage) String() string { return proto.CompactTextString(m) }
func (*InboxUserMessage) ProtoMessage()    {}
func (*InboxUserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{10}
}
func (m *InboxUserMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxUserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxUserMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxUserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxUserMessage.Merge(dst, src)
}
func (m *InboxUserMessage) XXX_Size() int {
	return m.Size()
}
func (m *InboxUserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxUserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InboxUserMessage proto.InternalMessageInfo

func (m *InboxUserMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxUserMessage) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *InboxUserMessage) GetRandomId() int64 {
	if m != nil {
		return m.RandomId
	}
	return 0
}

func (m *InboxUserMessage) GetDialogMessageId() int32 {
	if m != nil {
		return m.DialogMessageId
	}
	return 0
}

func (m *InboxUserMessage) GetMessageDataId() int64 {
	if m != nil {
		return m.MessageDataId
	}
	return 0
}

func (m *InboxUserMessage) GetMessage() *mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type InboxChatMessage struct {
	From                 *Sender          `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChatId           int32            `protobuf:"varint,2,opt,name=peer_chat_id,json=peerChatId,proto3" json:"peer_chat_id,omitempty"`
	RandomId             int64            `protobuf:"varint,3,opt,name=random_id,json=randomId,proto3" json:"random_id,omitempty"`
	DialogMessageId      int32            `protobuf:"varint,4,opt,name=dialog_message_id,json=dialogMessageId,proto3" json:"dialog_message_id,omitempty"`
	MessageDataId        int64            `protobuf:"varint,5,opt,name=message_data_id,json=messageDataId,proto3" json:"message_data_id,omitempty"`
	Message              *mtproto.Message `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InboxChatMessage) Reset()         { *m = InboxChatMessage{} }
func (m *InboxChatMessage) String() string { return proto.CompactTextString(m) }
func (*InboxChatMessage) ProtoMessage()    {}
func (*InboxChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{11}
}
func (m *InboxChatMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxChatMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxChatMessage.Merge(dst, src)
}
func (m *InboxChatMessage) XXX_Size() int {
	return m.Size()
}
func (m *InboxChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InboxChatMessage proto.InternalMessageInfo

func (m *InboxChatMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxChatMessage) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *InboxChatMessage) GetRandomId() int64 {
	if m != nil {
		return m.RandomId
	}
	return 0
}

func (m *InboxChatMessage) GetDialogMessageId() int32 {
	if m != nil {
		return m.DialogMessageId
	}
	return 0
}

func (m *InboxChatMessage) GetMessageDataId() int64 {
	if m != nil {
		return m.MessageDataId
	}
	return 0
}

func (m *InboxChatMessage) GetMessage() *mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type InboxUserMultiMessage struct {
	From                 *Sender            `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerUserId           int32              `protobuf:"varint,2,opt,name=peer_user_id,json=peerUserId,proto3" json:"peer_user_id,omitempty"`
	RandomId             []int64            `protobuf:"varint,3,rep,packed,name=random_id,json=randomId" json:"random_id,omitempty"`
	DialogMessageId      []int32            `protobuf:"varint,4,rep,packed,name=dialog_message_id,json=dialogMessageId" json:"dialog_message_id,omitempty"`
	MessageDataId        []int64            `protobuf:"varint,5,rep,packed,name=message_data_id,json=messageDataId" json:"message_data_id,omitempty"`
	Message              []*mtproto.Message `protobuf:"bytes,6,rep,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *InboxUserMultiMessage) Reset()         { *m = InboxUserMultiMessage{} }
func (m *InboxUserMultiMessage) String() string { return proto.CompactTextString(m) }
func (*InboxUserMultiMessage) ProtoMessage()    {}
func (*InboxUserMultiMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{12}
}
func (m *InboxUserMultiMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxUserMultiMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxUserMultiMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxUserMultiMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxUserMultiMessage.Merge(dst, src)
}
func (m *InboxUserMultiMessage) XXX_Size() int {
	return m.Size()
}
func (m *InboxUserMultiMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxUserMultiMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InboxUserMultiMessage proto.InternalMessageInfo

func (m *InboxUserMultiMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxUserMultiMessage) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *InboxUserMultiMessage) GetRandomId() []int64 {
	if m != nil {
		return m.RandomId
	}
	return nil
}

func (m *InboxUserMultiMessage) GetDialogMessageId() []int32 {
	if m != nil {
		return m.DialogMessageId
	}
	return nil
}

func (m *InboxUserMultiMessage) GetMessageDataId() []int64 {
	if m != nil {
		return m.MessageDataId
	}
	return nil
}

func (m *InboxUserMultiMessage) GetMessage() []*mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type InboxChatMultiMessage struct {
	From                 *Sender            `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChatId           int32              `protobuf:"varint,2,opt,name=peer_chat_id,json=peerChatId,proto3" json:"peer_chat_id,omitempty"`
	RandomId             []int64            `protobuf:"varint,3,rep,packed,name=random_id,json=randomId" json:"random_id,omitempty"`
	DialogMessageId      []int32            `protobuf:"varint,4,rep,packed,name=dialog_message_id,json=dialogMessageId" json:"dialog_message_id,omitempty"`
	MessageDataId        []int64            `protobuf:"varint,5,rep,packed,name=message_data_id,json=messageDataId" json:"message_data_id,omitempty"`
	Message              []*mtproto.Message `protobuf:"bytes,6,rep,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *InboxChatMultiMessage) Reset()         { *m = InboxChatMultiMessage{} }
func (m *InboxChatMultiMessage) String() string { return proto.CompactTextString(m) }
func (*InboxChatMultiMessage) ProtoMessage()    {}
func (*InboxChatMultiMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{13}
}
func (m *InboxChatMultiMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxChatMultiMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxChatMultiMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxChatMultiMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxChatMultiMessage.Merge(dst, src)
}
func (m *InboxChatMultiMessage) XXX_Size() int {
	return m.Size()
}
func (m *InboxChatMultiMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxChatMultiMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InboxChatMultiMessage proto.InternalMessageInfo

func (m *InboxChatMultiMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxChatMultiMessage) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *InboxChatMultiMessage) GetRandomId() []int64 {
	if m != nil {
		return m.RandomId
	}
	return nil
}

func (m *InboxChatMultiMessage) GetDialogMessageId() []int32 {
	if m != nil {
		return m.DialogMessageId
	}
	return nil
}

func (m *InboxChatMultiMessage) GetMessageDataId() []int64 {
	if m != nil {
		return m.MessageDataId
	}
	return nil
}

func (m *InboxChatMultiMessage) GetMessage() []*mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type InboxUserEditMessage struct {
	From                 *Sender          `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerUserId           int32            `protobuf:"varint,2,opt,name=peer_user_id,json=peerUserId,proto3" json:"peer_user_id,omitempty"`
	Message              *mtproto.Message `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InboxUserEditMessage) Reset()         { *m = InboxUserEditMessage{} }
func (m *InboxUserEditMessage) String() string { return proto.CompactTextString(m) }
func (*InboxUserEditMessage) ProtoMessage()    {}
func (*InboxUserEditMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{14}
}
func (m *InboxUserEditMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxUserEditMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxUserEditMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxUserEditMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxUserEditMessage.Merge(dst, src)
}
func (m *InboxUserEditMessage) XXX_Size() int {
	return m.Size()
}
func (m *InboxUserEditMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxUserEditMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InboxUserEditMessage proto.InternalMessageInfo

func (m *InboxUserEditMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxUserEditMessage) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *InboxUserEditMessage) GetMessage() *mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type InboxChatEditMessage struct {
	From                 *Sender          `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChatId           int32            `protobuf:"varint,2,opt,name=peer_chat_id,json=peerChatId,proto3" json:"peer_chat_id,omitempty"`
	Message              *mtproto.Message `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InboxChatEditMessage) Reset()         { *m = InboxChatEditMessage{} }
func (m *InboxChatEditMessage) String() string { return proto.CompactTextString(m) }
func (*InboxChatEditMessage) ProtoMessage()    {}
func (*InboxChatEditMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{15}
}
func (m *InboxChatEditMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxChatEditMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxChatEditMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxChatEditMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxChatEditMessage.Merge(dst, src)
}
func (m *InboxChatEditMessage) XXX_Size() int {
	return m.Size()
}
func (m *InboxChatEditMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxChatEditMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InboxChatEditMessage proto.InternalMessageInfo

func (m *InboxChatEditMessage) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxChatEditMessage) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *InboxChatEditMessage) GetMessage() *mtproto.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type DeleteMessagesRequest struct {
	From                 *Sender  `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerType             int32    `protobuf:"varint,2,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	PeerId               int32    `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Revoke               bool     `protobuf:"varint,4,opt,name=revoke,proto3" json:"revoke,omitempty"`
	Id                   []int32  `protobuf:"varint,5,rep,packed,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMessagesRequest) Reset()         { *m = DeleteMessagesRequest{} }
func (m *DeleteMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMessagesRequest) ProtoMessage()    {}
func (*DeleteMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{16}
}
func (m *DeleteMessagesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeleteMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteMessagesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DeleteMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMessagesRequest.Merge(dst, src)
}
func (m *DeleteMessagesRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeleteMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMessagesRequest proto.InternalMessageInfo

func (m *DeleteMessagesRequest) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *DeleteMessagesRequest) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *DeleteMessagesRequest) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *DeleteMessagesRequest) GetRevoke() bool {
	if m != nil {
		return m.Revoke
	}
	return false
}

func (m *DeleteMessagesRequest) GetId() []int32 {
	if m != nil {
		return m.Id
	}
	return nil
}

// channel_id = 0: messages.deleteHistory#1c015b09 flags:# just_clear:flags.0?true revoke:flags.1?true peer:InputPeer max_id:int = messages.AffectedHistory;
// channel_id > 0: channels.deleteUserHistory#d10dd71b channel:InputChannel user_id:InputUser = messages.AffectedHistory;
type DeleteHistoryRequest struct {
	From                 *Sender  `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	ChannelId            int32    `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PeerType             int32    `protobuf:"varint,3,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	PeerId               int32    `protobuf:"varint,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	JustClear            bool     `protobuf:"varint,5,opt,name=just_clear,json=justClear,proto3" json:"just_clear,omitempty"`
	Revoke               bool     `protobuf:"varint,6,opt,name=revoke,proto3" json:"revoke,omitempty"`
	MaxId                int32    `protobuf:"varint,7,opt,name=max_id,json=maxId,proto3" json:"max_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteHistoryRequest) Reset()         { *m = DeleteHistoryRequest{} }
func (m *DeleteHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteHistoryRequest) ProtoMessage()    {}
func (*DeleteHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{17}
}
func (m *DeleteHistoryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeleteHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteHistoryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DeleteHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteHistoryRequest.Merge(dst, src)
}
func (m *DeleteHistoryRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeleteHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteHistoryRequest proto.InternalMessageInfo

func (m *DeleteHistoryRequest) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *DeleteHistoryRequest) GetChannelId() int32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *DeleteHistoryRequest) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *DeleteHistoryRequest) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *DeleteHistoryRequest) GetJustClear() bool {
	if m != nil {
		return m.JustClear
	}
	return false
}

func (m *DeleteHistoryRequest) GetRevoke() bool {
	if m != nil {
		return m.Revoke
	}
	return false
}

func (m *DeleteHistoryRequest) GetMaxId() int32 {
	if m != nil {
		return m.MaxId
	}
	return 0
}

type InboxDeleteMessages struct {
	From                 *Sender  `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Id                   []int64  `protobuf:"varint,2,rep,packed,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InboxDeleteMessages) Reset()         { *m = InboxDeleteMessages{} }
func (m *InboxDeleteMessages) String() string { return proto.CompactTextString(m) }
func (*InboxDeleteMessages) ProtoMessage()    {}
func (*InboxDeleteMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{18}
}
func (m *InboxDeleteMessages) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxDeleteMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxDeleteMessages.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxDeleteMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxDeleteMessages.Merge(dst, src)
}
func (m *InboxDeleteMessages) XXX_Size() int {
	return m.Size()
}
func (m *InboxDeleteMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxDeleteMessages.DiscardUnknown(m)
}

var xxx_messageInfo_InboxDeleteMessages proto.InternalMessageInfo

func (m *InboxDeleteMessages) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxDeleteMessages) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

type InboxUserDeleteHistory struct {
	From                 *Sender  `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerUserId           int32    `protobuf:"varint,2,opt,name=peer_user_id,json=peerUserId,proto3" json:"peer_user_id,omitempty"`
	JustClear            bool     `protobuf:"varint,3,opt,name=just_clear,json=justClear,proto3" json:"just_clear,omitempty"`
	MaxId                int32    `protobuf:"varint,4,opt,name=max_id,json=maxId,proto3" json:"max_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InboxUserDeleteHistory) Reset()         { *m = InboxUserDeleteHistory{} }
func (m *InboxUserDeleteHistory) String() string { return proto.CompactTextString(m) }
func (*InboxUserDeleteHistory) ProtoMessage()    {}
func (*InboxUserDeleteHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{19}
}
func (m *InboxUserDeleteHistory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxUserDeleteHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxUserDeleteHistory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxUserDeleteHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxUserDeleteHistory.Merge(dst, src)
}
func (m *InboxUserDeleteHistory) XXX_Size() int {
	return m.Size()
}
func (m *InboxUserDeleteHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxUserDeleteHistory.DiscardUnknown(m)
}

var xxx_messageInfo_InboxUserDeleteHistory proto.InternalMessageInfo

func (m *InboxUserDeleteHistory) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxUserDeleteHistory) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *InboxUserDeleteHistory) GetJustClear() bool {
	if m != nil {
		return m.JustClear
	}
	return false
}

func (m *InboxUserDeleteHistory) GetMaxId() int32 {
	if m != nil {
		return m.MaxId
	}
	return 0
}

type InboxChatDeleteHistory struct {
	From                 *Sender  `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChatId           int32    `protobuf:"varint,2,opt,name=peer_chat_id,json=peerChatId,proto3" json:"peer_chat_id,omitempty"`
	MaxId                int32    `protobuf:"varint,3,opt,name=max_id,json=maxId,proto3" json:"max_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InboxChatDeleteHistory) Reset()         { *m = InboxChatDeleteHistory{} }
func (m *InboxChatDeleteHistory) String() string { return proto.CompactTextString(m) }
func (*InboxChatDeleteHistory) ProtoMessage()    {}
func (*InboxChatDeleteHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{20}
}
func (m *InboxChatDeleteHistory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxChatDeleteHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxChatDeleteHistory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxChatDeleteHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxChatDeleteHistory.Merge(dst, src)
}
func (m *InboxChatDeleteHistory) XXX_Size() int {
	return m.Size()
}
func (m *InboxChatDeleteHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxChatDeleteHistory.DiscardUnknown(m)
}

var xxx_messageInfo_InboxChatDeleteHistory proto.InternalMessageInfo

func (m *InboxChatDeleteHistory) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxChatDeleteHistory) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *InboxChatDeleteHistory) GetMaxId() int32 {
	if m != nil {
		return m.MaxId
	}
	return 0
}

type ContentMessage struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsMentioned          bool     `protobuf:"varint,2,opt,name=is_mentioned,json=isMentioned,proto3" json:"is_mentioned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentMessage) Reset()         { *m = ContentMessage{} }
func (m *ContentMessage) String() string { return proto.CompactTextString(m) }
func (*ContentMessage) ProtoMessage()    {}
func (*ContentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{21}
}
func (m *ContentMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ContentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentMessage.Merge(dst, src)
}
func (m *ContentMessage) XXX_Size() int {
	return m.Size()
}
func (m *ContentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ContentMessage proto.InternalMessageInfo

func (m *ContentMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContentMessage) GetIsMentioned() bool {
	if m != nil {
		return m.IsMentioned
	}
	return false
}

type ReadMessageContentsRequest struct {
	From                 *Sender           `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerType             int32             `protobuf:"varint,2,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	PeerId               int32             `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Id                   []*ContentMessage `protobuf:"bytes,4,rep,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReadMessageContentsRequest) Reset()         { *m = ReadMessageContentsRequest{} }
func (m *ReadMessageContentsRequest) String() string { return proto.CompactTextString(m) }
func (*ReadMessageContentsRequest) ProtoMessage()    {}
func (*ReadMessageContentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{22}
}
func (m *ReadMessageContentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReadMessageContentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReadMessageContentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ReadMessageContentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMessageContentsRequest.Merge(dst, src)
}
func (m *ReadMessageContentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ReadMessageContentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMessageContentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMessageContentsRequest proto.InternalMessageInfo

func (m *ReadMessageContentsRequest) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ReadMessageContentsRequest) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *ReadMessageContentsRequest) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *ReadMessageContentsRequest) GetId() []*ContentMessage {
	if m != nil {
		return m.Id
	}
	return nil
}

type InboxUserReadMediaUnread struct {
	From                 *Sender  `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Id                   []int32  `protobuf:"varint,4,rep,packed,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InboxUserReadMediaUnread) Reset()         { *m = InboxUserReadMediaUnread{} }
func (m *InboxUserReadMediaUnread) String() string { return proto.CompactTextString(m) }
func (*InboxUserReadMediaUnread) ProtoMessage()    {}
func (*InboxUserReadMediaUnread) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{23}
}
func (m *InboxUserReadMediaUnread) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxUserReadMediaUnread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxUserReadMediaUnread.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxUserReadMediaUnread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxUserReadMediaUnread.Merge(dst, src)
}
func (m *InboxUserReadMediaUnread) XXX_Size() int {
	return m.Size()
}
func (m *InboxUserReadMediaUnread) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxUserReadMediaUnread.DiscardUnknown(m)
}

var xxx_messageInfo_InboxUserReadMediaUnread proto.InternalMessageInfo

func (m *InboxUserReadMediaUnread) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxUserReadMediaUnread) GetId() []int32 {
	if m != nil {
		return m.Id
	}
	return nil
}

type InboxChatReadMediaUnread struct {
	From                 *Sender  `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	PeerChatId           int32    `protobuf:"varint,2,opt,name=peer_chat_id,json=peerChatId,proto3" json:"peer_chat_id,omitempty"`
	Id                   []int32  `protobuf:"varint,3,rep,packed,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InboxChatReadMediaUnread) Reset()         { *m = InboxChatReadMediaUnread{} }
func (m *InboxChatReadMediaUnread) String() string { return proto.CompactTextString(m) }
func (*InboxChatReadMediaUnread) ProtoMessage()    {}
func (*InboxChatReadMediaUnread) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_a2fb1cd7e36d0df3, []int{24}
}
func (m *InboxChatReadMediaUnread) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboxChatReadMediaUnread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboxChatReadMediaUnread.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboxChatReadMediaUnread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxChatReadMediaUnread.Merge(dst, src)
}
func (m *InboxChatReadMediaUnread) XXX_Size() int {
	return m.Size()
}
func (m *InboxChatReadMediaUnread) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxChatReadMediaUnread.DiscardUnknown(m)
}

var xxx_messageInfo_InboxChatReadMediaUnread proto.InternalMessageInfo

func (m *InboxChatReadMediaUnread) GetFrom() *Sender {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *InboxChatReadMediaUnread) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *InboxChatReadMediaUnread) GetId() []int32 {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*Sender)(nil), "msgpb.Sender")
	proto.RegisterType((*UserMessage)(nil), "msgpb.UserMessage")
	proto.RegisterType((*ChatMessage)(nil), "msgpb.ChatMessage")
	proto.RegisterType((*ChannelMessage)(nil), "msgpb.ChannelMessage")
	proto.RegisterType((*UserMultiMessage)(nil), "msgpb.UserMultiMessage")
	proto.RegisterType((*ChatMultiMessage)(nil), "msgpb.ChatMultiMessage")
	proto.RegisterType((*ChannelMultiMessage)(nil), "msgpb.ChannelMultiMessage")
	proto.RegisterType((*OutboxMessage)(nil), "msgpb.OutboxMessage")
	proto.RegisterType((*OutgoingMessage)(nil), "msgpb.OutgoingMessage")
	proto.RegisterType((*OutgoingMultiMessage)(nil), "msgpb.OutgoingMultiMessage")
	proto.RegisterType((*InboxUserMessage)(nil), "msgpb.InboxUserMessage")
	proto.RegisterType((*InboxChatMessage)(nil), "msgpb.InboxChatMessage")
	proto.RegisterType((*InboxUserMultiMessage)(nil), "msgpb.InboxUserMultiMessage")
	proto.RegisterType((*InboxChatMultiMessage)(nil), "msgpb.InboxChatMultiMessage")
	proto.RegisterType((*InboxUserEditMessage)(nil), "msgpb.InboxUserEditMessage")
	proto.RegisterType((*InboxChatEditMessage)(nil), "msgpb.InboxChatEditMessage")
	proto.RegisterType((*DeleteMessagesRequest)(nil), "msgpb.DeleteMessagesRequest")
	proto.RegisterType((*DeleteHistoryRequest)(nil), "msgpb.DeleteHistoryRequest")
	proto.RegisterType((*InboxDeleteMessages)(nil), "msgpb.InboxDeleteMessages")
	proto.RegisterType((*InboxUserDeleteHistory)(nil), "msgpb.InboxUserDeleteHistory")
	proto.RegisterType((*InboxChatDeleteHistory)(nil), "msgpb.InboxChatDeleteHistory")
	proto.RegisterType((*ContentMessage)(nil), "msgpb.ContentMessage")
	proto.RegisterType((*ReadMessageContentsRequest)(nil), "msgpb.ReadMessageContentsRequest")
	proto.RegisterType((*InboxUserReadMediaUnread)(nil), "msgpb.InboxUserReadMediaUnread")
	proto.RegisterType((*InboxChatReadMediaUnread)(nil), "msgpb.InboxChatReadMediaUnread")
}
func (this *Sender) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&msgpb.Sender{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",")
	s = append(s, "AuthKeyId: "+fmt.Sprintf("%#v", this.AuthKeyId)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UserMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.UserMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerUserId: "+fmt.Sprintf("%#v", this.PeerUserId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ChatMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.ChatMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChatId: "+fmt.Sprintf("%#v", this.PeerChatId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ChannelMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.ChannelMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChannelId: "+fmt.Sprintf("%#v", this.PeerChannelId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UserMultiMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.UserMultiMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerUserId: "+fmt.Sprintf("%#v", this.PeerUserId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ChatMultiMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.ChatMultiMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChatId: "+fmt.Sprintf("%#v", this.PeerChatId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ChannelMultiMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.ChannelMultiMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChannelId: "+fmt.Sprintf("%#v", this.PeerChannelId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *OutboxMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&msgpb.OutboxMessage{")
	s = append(s, "NoWebpage: "+fmt.Sprintf("%#v", this.NoWebpage)+",")
	s = append(s, "Background: "+fmt.Sprintf("%#v", this.Background)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.ScheduleDate != nil {
		s = append(s, "ScheduleDate: "+fmt.Sprintf("%#v", this.ScheduleDate)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *OutgoingMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.OutgoingMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerType: "+fmt.Sprintf("%#v", this.PeerType)+",")
	s = append(s, "PeerId: "+fmt.Sprintf("%#v", this.PeerId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *OutgoingMultiMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.OutgoingMultiMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerType: "+fmt.Sprintf("%#v", this.PeerType)+",")
	s = append(s, "PeerId: "+fmt.Sprintf("%#v", this.PeerId)+",")
	if this.MultiMessage != nil {
		s = append(s, "MultiMessage: "+fmt.Sprintf("%#v", this.MultiMessage)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxUserMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&msgpb.InboxUserMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerUserId: "+fmt.Sprintf("%#v", this.PeerUserId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	s = append(s, "DialogMessageId: "+fmt.Sprintf("%#v", this.DialogMessageId)+",")
	s = append(s, "MessageDataId: "+fmt.Sprintf("%#v", this.MessageDataId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxChatMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&msgpb.InboxChatMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChatId: "+fmt.Sprintf("%#v", this.PeerChatId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	s = append(s, "DialogMessageId: "+fmt.Sprintf("%#v", this.DialogMessageId)+",")
	s = append(s, "MessageDataId: "+fmt.Sprintf("%#v", this.MessageDataId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxUserMultiMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&msgpb.InboxUserMultiMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerUserId: "+fmt.Sprintf("%#v", this.PeerUserId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	s = append(s, "DialogMessageId: "+fmt.Sprintf("%#v", this.DialogMessageId)+",")
	s = append(s, "MessageDataId: "+fmt.Sprintf("%#v", this.MessageDataId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxChatMultiMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&msgpb.InboxChatMultiMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChatId: "+fmt.Sprintf("%#v", this.PeerChatId)+",")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",")
	s = append(s, "DialogMessageId: "+fmt.Sprintf("%#v", this.DialogMessageId)+",")
	s = append(s, "MessageDataId: "+fmt.Sprintf("%#v", this.MessageDataId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxUserEditMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&msgpb.InboxUserEditMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerUserId: "+fmt.Sprintf("%#v", this.PeerUserId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxChatEditMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&msgpb.InboxChatEditMessage{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChatId: "+fmt.Sprintf("%#v", this.PeerChatId)+",")
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DeleteMessagesRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&msgpb.DeleteMessagesRequest{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerType: "+fmt.Sprintf("%#v", this.PeerType)+",")
	s = append(s, "PeerId: "+fmt.Sprintf("%#v", this.PeerId)+",")
	s = append(s, "Revoke: "+fmt.Sprintf("%#v", this.Revoke)+",")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DeleteHistoryRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&msgpb.DeleteHistoryRequest{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "ChannelId: "+fmt.Sprintf("%#v", this.ChannelId)+",")
	s = append(s, "PeerType: "+fmt.Sprintf("%#v", this.PeerType)+",")
	s = append(s, "PeerId: "+fmt.Sprintf("%#v", this.PeerId)+",")
	s = append(s, "JustClear: "+fmt.Sprintf("%#v", this.JustClear)+",")
	s = append(s, "Revoke: "+fmt.Sprintf("%#v", this.Revoke)+",")
	s = append(s, "MaxId: "+fmt.Sprintf("%#v", this.MaxId)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxDeleteMessages) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&msgpb.InboxDeleteMessages{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxUserDeleteHistory) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.InboxUserDeleteHistory{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerUserId: "+fmt.Sprintf("%#v", this.PeerUserId)+",")
	s = append(s, "JustClear: "+fmt.Sprintf("%#v", this.JustClear)+",")
	s = append(s, "MaxId: "+fmt.Sprintf("%#v", this.MaxId)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxChatDeleteHistory) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&msgpb.InboxChatDeleteHistory{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChatId: "+fmt.Sprintf("%#v", this.PeerChatId)+",")
	s = append(s, "MaxId: "+fmt.Sprintf("%#v", this.MaxId)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ContentMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&msgpb.ContentMessage{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	s = append(s, "IsMentioned: "+fmt.Sprintf("%#v", this.IsMentioned)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReadMessageContentsRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&msgpb.ReadMessageContentsRequest{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerType: "+fmt.Sprintf("%#v", this.PeerType)+",")
	s = append(s, "PeerId: "+fmt.Sprintf("%#v", this.PeerId)+",")
	if this.Id != nil {
		s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxUserReadMediaUnread) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&msgpb.InboxUserReadMediaUnread{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboxChatReadMediaUnread) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&msgpb.InboxChatReadMediaUnread{")
	if this.From != nil {
		s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",")
	}
	s = append(s, "PeerChatId: "+fmt.Sprintf("%#v", this.PeerChatId)+",")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMsg(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCMsgClient is the client API for RPCMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCMsgClient interface {
	SendUserMessage(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	SendChatMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	SendChannelMessage(ctx context.Context, in *ChannelMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	SendUserMultiMessage(ctx context.Context, in *UserMultiMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	SendChatMultiMessage(ctx context.Context, in *ChatMultiMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	SendChannelMultiMessage(ctx context.Context, in *ChannelMultiMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	PushUserMessage(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*mtproto.Bool, error)
	EditUserMessage(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	EditChatMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	EditChannelMessage(ctx context.Context, in *ChannelMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	SendMessage(ctx context.Context, in *OutgoingMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	SendMultiMessage(ctx context.Context, in *OutgoingMultiMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	PushMessage(ctx context.Context, in *OutgoingMessage, opts ...grpc.CallOption) (*mtproto.Bool, error)
	EditMessage(ctx context.Context, in *OutgoingMessage, opts ...grpc.CallOption) (*mtproto.Updates, error)
	DeleteMessages(ctx context.Context, in *DeleteMessagesRequest, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error)
	DeleteHistory(ctx context.Context, in *DeleteHistoryRequest, opts ...grpc.CallOption) (*mtproto.Messages_AffectedHistory, error)
	ReadMessageContents(ctx context.Context, in *ReadMessageContentsRequest, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error)
}

type rPCMsgClient struct {
	cc *grpc.ClientConn
}

func NewRPCMsgClient(cc *grpc.ClientConn) RPCMsgClient {
	return &rPCMsgClient{cc}
}

func (c *rPCMsgClient) SendUserMessage(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/SendUserMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) SendChatMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/SendChatMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) SendChannelMessage(ctx context.Context, in *ChannelMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/SendChannelMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) SendUserMultiMessage(ctx context.Context, in *UserMultiMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/SendUserMultiMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) SendChatMultiMessage(ctx context.Context, in *ChatMultiMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/SendChatMultiMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) SendChannelMultiMessage(ctx context.Context, in *ChannelMultiMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/SendChannelMultiMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) PushUserMessage(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/PushUserMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) EditUserMessage(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/EditUserMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) EditChatMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/EditChatMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) EditChannelMessage(ctx context.Context, in *ChannelMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/EditChannelMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) SendMessage(ctx context.Context, in *OutgoingMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) SendMultiMessage(ctx context.Context, in *OutgoingMultiMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/SendMultiMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) PushMessage(ctx context.Context, in *OutgoingMessage, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/PushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) EditMessage(ctx context.Context, in *OutgoingMessage, opts ...grpc.CallOption) (*mtproto.Updates, error) {
	out := new(mtproto.Updates)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/EditMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) DeleteMessages(ctx context.Context, in *DeleteMessagesRequest, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error) {
	out := new(mtproto.Messages_AffectedMessages)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/DeleteMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) DeleteHistory(ctx context.Context, in *DeleteHistoryRequest, opts ...grpc.CallOption) (*mtproto.Messages_AffectedHistory, error) {
	out := new(mtproto.Messages_AffectedHistory)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/DeleteHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCMsgClient) ReadMessageContents(ctx context.Context, in *ReadMessageContentsRequest, opts ...grpc.CallOption) (*mtproto.Messages_AffectedMessages, error) {
	out := new(mtproto.Messages_AffectedMessages)
	err := c.cc.Invoke(ctx, "/msgpb.RPCMsg/ReadMessageContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCMsg service

type RPCMsgServer interface {
	SendUserMessage(context.Context, *UserMessage) (*mtproto.Updates, error)
	SendChatMessage(context.Context, *ChatMessage) (*mtproto.Updates, error)
	SendChannelMessage(context.Context, *ChannelMessage) (*mtproto.Updates, error)
	SendUserMultiMessage(context.Context, *UserMultiMessage) (*mtproto.Updates, error)
	SendChatMultiMessage(context.Context, *ChatMultiMessage) (*mtproto.Updates, error)
	SendChannelMultiMessage(context.Context, *ChannelMultiMessage) (*mtproto.Updates, error)
	PushUserMessage(context.Context, *UserMessage) (*mtproto.Bool, error)
	EditUserMessage(context.Context, *UserMessage) (*mtproto.Updates, error)
	EditChatMessage(context.Context, *ChatMessage) (*mtproto.Updates, error)
	EditChannelMessage(context.Context, *ChannelMessage) (*mtproto.Updates, error)
	SendMessage(context.Context, *OutgoingMessage) (*mtproto.Updates, error)
	SendMultiMessage(context.Context, *OutgoingMultiMessage) (*mtproto.Updates, error)
	PushMessage(context.Context, *OutgoingMessage) (*mtproto.Bool, error)
	EditMessage(context.Context, *OutgoingMessage) (*mtproto.Updates, error)
	DeleteMessages(context.Context, *DeleteMessagesRequest) (*mtproto.Messages_AffectedMessages, error)
	DeleteHistory(context.Context, *DeleteHistoryRequest) (*mtproto.Messages_AffectedHistory, error)
	ReadMessageContents(context.Context, *ReadMessageContentsRequest) (*mtproto.Messages_AffectedMessages, error)
}

func RegisterRPCMsgServer(s *grpc.Server, srv RPCMsgServer) {
	s.RegisterService(&_RPCMsg_serviceDesc, srv)
}

func _RPCMsg_SendUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).SendUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/SendUserMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).SendUserMessage(ctx, req.(*UserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_SendChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).SendChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/SendChatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).SendChatMessage(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_SendChannelMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).SendChannelMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/SendChannelMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).SendChannelMessage(ctx, req.(*ChannelMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_SendUserMultiMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMultiMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).SendUserMultiMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/SendUserMultiMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).SendUserMultiMessage(ctx, req.(*UserMultiMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_SendChatMultiMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMultiMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).SendChatMultiMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/SendChatMultiMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).SendChatMultiMessage(ctx, req.(*ChatMultiMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_SendChannelMultiMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelMultiMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).SendChannelMultiMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/SendChannelMultiMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).SendChannelMultiMessage(ctx, req.(*ChannelMultiMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_PushUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).PushUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/PushUserMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).PushUserMessage(ctx, req.(*UserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_EditUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).EditUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/EditUserMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).EditUserMessage(ctx, req.(*UserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_EditChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).EditChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/EditChatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).EditChatMessage(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_EditChannelMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).EditChannelMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/EditChannelMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).EditChannelMessage(ctx, req.(*ChannelMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutgoingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).SendMessage(ctx, req.(*OutgoingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_SendMultiMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutgoingMultiMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).SendMultiMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/SendMultiMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).SendMultiMessage(ctx, req.(*OutgoingMultiMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutgoingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).PushMessage(ctx, req.(*OutgoingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_EditMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutgoingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).EditMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/EditMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).EditMessage(ctx, req.(*OutgoingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_DeleteMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).DeleteMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/DeleteMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).DeleteMessages(ctx, req.(*DeleteMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_DeleteHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).DeleteHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/DeleteHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).DeleteHistory(ctx, req.(*DeleteHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCMsg_ReadMessageContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMessageContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCMsgServer).ReadMessageContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgpb.RPCMsg/ReadMessageContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCMsgServer).ReadMessageContents(ctx, req.(*ReadMessageContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCMsg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msgpb.RPCMsg",
	HandlerType: (*RPCMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUserMessage",
			Handler:    _RPCMsg_SendUserMessage_Handler,
		},
		{
			MethodName: "SendChatMessage",
			Handler:    _RPCMsg_SendChatMessage_Handler,
		},
		{
			MethodName: "SendChannelMessage",
			Handler:    _RPCMsg_SendChannelMessage_Handler,
		},
		{
			MethodName: "SendUserMultiMessage",
			Handler:    _RPCMsg_SendUserMultiMessage_Handler,
		},
		{
			MethodName: "SendChatMultiMessage",
			Handler:    _RPCMsg_SendChatMultiMessage_Handler,
		},
		{
			MethodName: "SendChannelMultiMessage",
			Handler:    _RPCMsg_SendChannelMultiMessage_Handler,
		},
		{
			MethodName: "PushUserMessage",
			Handler:    _RPCMsg_PushUserMessage_Handler,
		},
		{
			MethodName: "EditUserMessage",
			Handler:    _RPCMsg_EditUserMessage_Handler,
		},
		{
			MethodName: "EditChatMessage",
			Handler:    _RPCMsg_EditChatMessage_Handler,
		},
		{
			MethodName: "EditChannelMessage",
			Handler:    _RPCMsg_EditChannelMessage_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _RPCMsg_SendMessage_Handler,
		},
		{
			MethodName: "SendMultiMessage",
			Handler:    _RPCMsg_SendMultiMessage_Handler,
		},
		{
			MethodName: "PushMessage",
			Handler:    _RPCMsg_PushMessage_Handler,
		},
		{
			MethodName: "EditMessage",
			Handler:    _RPCMsg_EditMessage_Handler,
		},
		{
			MethodName: "DeleteMessages",
			Handler:    _RPCMsg_DeleteMessages_Handler,
		},
		{
			MethodName: "DeleteHistory",
			Handler:    _RPCMsg_DeleteHistory_Handler,
		},
		{
			MethodName: "ReadMessageContents",
			Handler:    _RPCMsg_ReadMessageContents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}

func (m *Sender) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sender) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Id))
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Type))
	}
	if m.AuthKeyId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.AuthKeyId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UserMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n1, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PeerUserId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerUserId))
	}
	if m.RandomId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.RandomId))
	}
	if m.Message != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n2, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ChatMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n3, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.PeerChatId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChatId))
	}
	if m.RandomId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.RandomId))
	}
	if m.Message != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n4, err := m.Message.MarshalTo(dAtA[i:])
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

func (m *ChannelMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n5, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.PeerChannelId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChannelId))
	}
	if m.RandomId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.RandomId))
	}
	if m.Message != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n6, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UserMultiMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserMultiMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n7, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.PeerUserId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerUserId))
	}
	if len(m.RandomId) > 0 {
		dAtA9 := make([]byte, len(m.RandomId)*10)
		var j8 int
		for _, num1 := range m.RandomId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA9[j8] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j8++
			}
			dAtA9[j8] = uint8(num)
			j8++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j8))
		i += copy(dAtA[i:], dAtA9[:j8])
	}
	if len(m.Message) > 0 {
		for _, msg := range m.Message {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMsg(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ChatMultiMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatMultiMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n10, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.PeerChatId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChatId))
	}
	if len(m.RandomId) > 0 {
		dAtA12 := make([]byte, len(m.RandomId)*10)
		var j11 int
		for _, num1 := range m.RandomId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA12[j11] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j11++
			}
			dAtA12[j11] = uint8(num)
			j11++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j11))
		i += copy(dAtA[i:], dAtA12[:j11])
	}
	if len(m.Message) > 0 {
		for _, msg := range m.Message {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMsg(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ChannelMultiMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelMultiMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n13, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	if m.PeerChannelId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChannelId))
	}
	if len(m.RandomId) > 0 {
		dAtA15 := make([]byte, len(m.RandomId)*10)
		var j14 int
		for _, num1 := range m.RandomId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA15[j14] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j14++
			}
			dAtA15[j14] = uint8(num)
			j14++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j14))
		i += copy(dAtA[i:], dAtA15[:j14])
	}
	if len(m.Message) > 0 {
		for _, msg := range m.Message {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMsg(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *OutboxMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutboxMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NoWebpage {
		dAtA[i] = 0x8
		i++
		if m.NoWebpage {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Background {
		dAtA[i] = 0x10
		i++
		if m.Background {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RandomId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.RandomId))
	}
	if m.Message != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n16, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n16
	}
	if m.ScheduleDate != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.ScheduleDate.Size()))
		n17, err := m.ScheduleDate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *OutgoingMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n18, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	if m.PeerType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerId))
	}
	if m.Message != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n19, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n19
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *OutgoingMultiMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingMultiMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n20, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n20
	}
	if m.PeerType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerId))
	}
	if len(m.MultiMessage) > 0 {
		for _, msg := range m.MultiMessage {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMsg(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxUserMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxUserMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n21, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n21
	}
	if m.PeerUserId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerUserId))
	}
	if m.RandomId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.RandomId))
	}
	if m.DialogMessageId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.DialogMessageId))
	}
	if m.MessageDataId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.MessageDataId))
	}
	if m.Message != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n22, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n22
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxChatMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxChatMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n23, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n23
	}
	if m.PeerChatId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChatId))
	}
	if m.RandomId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.RandomId))
	}
	if m.DialogMessageId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.DialogMessageId))
	}
	if m.MessageDataId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.MessageDataId))
	}
	if m.Message != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n24, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n24
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxUserMultiMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxUserMultiMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n25, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n25
	}
	if m.PeerUserId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerUserId))
	}
	if len(m.RandomId) > 0 {
		dAtA27 := make([]byte, len(m.RandomId)*10)
		var j26 int
		for _, num1 := range m.RandomId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA27[j26] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j26++
			}
			dAtA27[j26] = uint8(num)
			j26++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j26))
		i += copy(dAtA[i:], dAtA27[:j26])
	}
	if len(m.DialogMessageId) > 0 {
		dAtA29 := make([]byte, len(m.DialogMessageId)*10)
		var j28 int
		for _, num1 := range m.DialogMessageId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA29[j28] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j28++
			}
			dAtA29[j28] = uint8(num)
			j28++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j28))
		i += copy(dAtA[i:], dAtA29[:j28])
	}
	if len(m.MessageDataId) > 0 {
		dAtA31 := make([]byte, len(m.MessageDataId)*10)
		var j30 int
		for _, num1 := range m.MessageDataId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA31[j30] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j30++
			}
			dAtA31[j30] = uint8(num)
			j30++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j30))
		i += copy(dAtA[i:], dAtA31[:j30])
	}
	if len(m.Message) > 0 {
		for _, msg := range m.Message {
			dAtA[i] = 0x32
			i++
			i = encodeVarintMsg(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxChatMultiMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxChatMultiMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n32, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n32
	}
	if m.PeerChatId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChatId))
	}
	if len(m.RandomId) > 0 {
		dAtA34 := make([]byte, len(m.RandomId)*10)
		var j33 int
		for _, num1 := range m.RandomId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA34[j33] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j33++
			}
			dAtA34[j33] = uint8(num)
			j33++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j33))
		i += copy(dAtA[i:], dAtA34[:j33])
	}
	if len(m.DialogMessageId) > 0 {
		dAtA36 := make([]byte, len(m.DialogMessageId)*10)
		var j35 int
		for _, num1 := range m.DialogMessageId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA36[j35] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j35++
			}
			dAtA36[j35] = uint8(num)
			j35++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j35))
		i += copy(dAtA[i:], dAtA36[:j35])
	}
	if len(m.MessageDataId) > 0 {
		dAtA38 := make([]byte, len(m.MessageDataId)*10)
		var j37 int
		for _, num1 := range m.MessageDataId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA38[j37] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j37++
			}
			dAtA38[j37] = uint8(num)
			j37++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j37))
		i += copy(dAtA[i:], dAtA38[:j37])
	}
	if len(m.Message) > 0 {
		for _, msg := range m.Message {
			dAtA[i] = 0x32
			i++
			i = encodeVarintMsg(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxUserEditMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxUserEditMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n39, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n39
	}
	if m.PeerUserId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerUserId))
	}
	if m.Message != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n40, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n40
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxChatEditMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxChatEditMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n41, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n41
	}
	if m.PeerChatId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChatId))
	}
	if m.Message != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Message.Size()))
		n42, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n42
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DeleteMessagesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteMessagesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n43, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n43
	}
	if m.PeerType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerId))
	}
	if m.Revoke {
		dAtA[i] = 0x20
		i++
		if m.Revoke {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Id) > 0 {
		dAtA45 := make([]byte, len(m.Id)*10)
		var j44 int
		for _, num1 := range m.Id {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA45[j44] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j44++
			}
			dAtA45[j44] = uint8(num)
			j44++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j44))
		i += copy(dAtA[i:], dAtA45[:j44])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DeleteHistoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteHistoryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n46, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n46
	}
	if m.ChannelId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.ChannelId))
	}
	if m.PeerType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerId))
	}
	if m.JustClear {
		dAtA[i] = 0x28
		i++
		if m.JustClear {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Revoke {
		dAtA[i] = 0x30
		i++
		if m.Revoke {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.MaxId != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.MaxId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxDeleteMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxDeleteMessages) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n47, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n47
	}
	if len(m.Id) > 0 {
		dAtA49 := make([]byte, len(m.Id)*10)
		var j48 int
		for _, num1 := range m.Id {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA49[j48] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j48++
			}
			dAtA49[j48] = uint8(num)
			j48++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j48))
		i += copy(dAtA[i:], dAtA49[:j48])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxUserDeleteHistory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxUserDeleteHistory) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n50, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n50
	}
	if m.PeerUserId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerUserId))
	}
	if m.JustClear {
		dAtA[i] = 0x18
		i++
		if m.JustClear {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.MaxId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.MaxId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxChatDeleteHistory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxChatDeleteHistory) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n51, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n51
	}
	if m.PeerChatId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChatId))
	}
	if m.MaxId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.MaxId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ContentMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.Id))
	}
	if m.IsMentioned {
		dAtA[i] = 0x10
		i++
		if m.IsMentioned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ReadMessageContentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadMessageContentsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n52, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n52
	}
	if m.PeerType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerId))
	}
	if len(m.Id) > 0 {
		for _, msg := range m.Id {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMsg(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxUserReadMediaUnread) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxUserReadMediaUnread) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n53, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n53
	}
	if len(m.Id) > 0 {
		dAtA55 := make([]byte, len(m.Id)*10)
		var j54 int
		for _, num1 := range m.Id {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA55[j54] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j54++
			}
			dAtA55[j54] = uint8(num)
			j54++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j54))
		i += copy(dAtA[i:], dAtA55[:j54])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *InboxChatReadMediaUnread) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboxChatReadMediaUnread) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.From != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.From.Size()))
		n56, err := m.From.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n56
	}
	if m.PeerChatId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMsg(dAtA, i, uint64(m.PeerChatId))
	}
	if len(m.Id) > 0 {
		dAtA58 := make([]byte, len(m.Id)*10)
		var j57 int
		for _, num1 := range m.Id {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA58[j57] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j57++
			}
			dAtA58[j57] = uint8(num)
			j57++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMsg(dAtA, i, uint64(j57))
		i += copy(dAtA[i:], dAtA58[:j57])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Sender) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMsg(uint64(m.Id))
	}
	if m.Type != 0 {
		n += 1 + sovMsg(uint64(m.Type))
	}
	if m.AuthKeyId != 0 {
		n += 1 + sovMsg(uint64(m.AuthKeyId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerUserId != 0 {
		n += 1 + sovMsg(uint64(m.PeerUserId))
	}
	if m.RandomId != 0 {
		n += 1 + sovMsg(uint64(m.RandomId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChatMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChatId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChatId))
	}
	if m.RandomId != 0 {
		n += 1 + sovMsg(uint64(m.RandomId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChannelMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChannelId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChannelId))
	}
	if m.RandomId != 0 {
		n += 1 + sovMsg(uint64(m.RandomId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserMultiMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerUserId != 0 {
		n += 1 + sovMsg(uint64(m.PeerUserId))
	}
	if len(m.RandomId) > 0 {
		l = 0
		for _, e := range m.RandomId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.Message) > 0 {
		for _, e := range m.Message {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChatMultiMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChatId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChatId))
	}
	if len(m.RandomId) > 0 {
		l = 0
		for _, e := range m.RandomId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.Message) > 0 {
		for _, e := range m.Message {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChannelMultiMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChannelId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChannelId))
	}
	if len(m.RandomId) > 0 {
		l = 0
		for _, e := range m.RandomId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.Message) > 0 {
		for _, e := range m.Message {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OutboxMessage) Size() (n int) {
	var l int
	_ = l
	if m.NoWebpage {
		n += 2
	}
	if m.Background {
		n += 2
	}
	if m.RandomId != 0 {
		n += 1 + sovMsg(uint64(m.RandomId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.ScheduleDate != nil {
		l = m.ScheduleDate.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OutgoingMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerType != 0 {
		n += 1 + sovMsg(uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		n += 1 + sovMsg(uint64(m.PeerId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OutgoingMultiMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerType != 0 {
		n += 1 + sovMsg(uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		n += 1 + sovMsg(uint64(m.PeerId))
	}
	if len(m.MultiMessage) > 0 {
		for _, e := range m.MultiMessage {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxUserMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerUserId != 0 {
		n += 1 + sovMsg(uint64(m.PeerUserId))
	}
	if m.RandomId != 0 {
		n += 1 + sovMsg(uint64(m.RandomId))
	}
	if m.DialogMessageId != 0 {
		n += 1 + sovMsg(uint64(m.DialogMessageId))
	}
	if m.MessageDataId != 0 {
		n += 1 + sovMsg(uint64(m.MessageDataId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxChatMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChatId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChatId))
	}
	if m.RandomId != 0 {
		n += 1 + sovMsg(uint64(m.RandomId))
	}
	if m.DialogMessageId != 0 {
		n += 1 + sovMsg(uint64(m.DialogMessageId))
	}
	if m.MessageDataId != 0 {
		n += 1 + sovMsg(uint64(m.MessageDataId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxUserMultiMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerUserId != 0 {
		n += 1 + sovMsg(uint64(m.PeerUserId))
	}
	if len(m.RandomId) > 0 {
		l = 0
		for _, e := range m.RandomId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.DialogMessageId) > 0 {
		l = 0
		for _, e := range m.DialogMessageId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.MessageDataId) > 0 {
		l = 0
		for _, e := range m.MessageDataId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.Message) > 0 {
		for _, e := range m.Message {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxChatMultiMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChatId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChatId))
	}
	if len(m.RandomId) > 0 {
		l = 0
		for _, e := range m.RandomId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.DialogMessageId) > 0 {
		l = 0
		for _, e := range m.DialogMessageId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.MessageDataId) > 0 {
		l = 0
		for _, e := range m.MessageDataId {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if len(m.Message) > 0 {
		for _, e := range m.Message {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxUserEditMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerUserId != 0 {
		n += 1 + sovMsg(uint64(m.PeerUserId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxChatEditMessage) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChatId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChatId))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeleteMessagesRequest) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerType != 0 {
		n += 1 + sovMsg(uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		n += 1 + sovMsg(uint64(m.PeerId))
	}
	if m.Revoke {
		n += 2
	}
	if len(m.Id) > 0 {
		l = 0
		for _, e := range m.Id {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeleteHistoryRequest) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.ChannelId != 0 {
		n += 1 + sovMsg(uint64(m.ChannelId))
	}
	if m.PeerType != 0 {
		n += 1 + sovMsg(uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		n += 1 + sovMsg(uint64(m.PeerId))
	}
	if m.JustClear {
		n += 2
	}
	if m.Revoke {
		n += 2
	}
	if m.MaxId != 0 {
		n += 1 + sovMsg(uint64(m.MaxId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxDeleteMessages) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if len(m.Id) > 0 {
		l = 0
		for _, e := range m.Id {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxUserDeleteHistory) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerUserId != 0 {
		n += 1 + sovMsg(uint64(m.PeerUserId))
	}
	if m.JustClear {
		n += 2
	}
	if m.MaxId != 0 {
		n += 1 + sovMsg(uint64(m.MaxId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxChatDeleteHistory) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChatId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChatId))
	}
	if m.MaxId != 0 {
		n += 1 + sovMsg(uint64(m.MaxId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContentMessage) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMsg(uint64(m.Id))
	}
	if m.IsMentioned {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReadMessageContentsRequest) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerType != 0 {
		n += 1 + sovMsg(uint64(m.PeerType))
	}
	if m.PeerId != 0 {
		n += 1 + sovMsg(uint64(m.PeerId))
	}
	if len(m.Id) > 0 {
		for _, e := range m.Id {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxUserReadMediaUnread) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if len(m.Id) > 0 {
		l = 0
		for _, e := range m.Id {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InboxChatReadMediaUnread) Size() (n int) {
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.PeerChatId != 0 {
		n += 1 + sovMsg(uint64(m.PeerChatId))
	}
	if len(m.Id) > 0 {
		l = 0
		for _, e := range m.Id {
			l += sovMsg(uint64(e))
		}
		n += 1 + sovMsg(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMsg(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sender) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: Sender: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sender: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
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
					return ErrIntOverflowMsg
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
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *UserMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: UserMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerUserId", wireType)
			}
			m.PeerUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerUserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
			m.RandomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &mtproto.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *ChatMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: ChatMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChatId", wireType)
			}
			m.PeerChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
			m.RandomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &mtproto.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *ChannelMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: ChannelMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChannelMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChannelId", wireType)
			}
			m.PeerChannelId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChannelId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
			m.RandomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &mtproto.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *UserMultiMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: UserMultiMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserMultiMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerUserId", wireType)
			}
			m.PeerUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerUserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RandomId = append(m.RandomId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RandomId = append(m.RandomId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message, &mtproto.Message{})
			if err := m.Message[len(m.Message)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *ChatMultiMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: ChatMultiMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatMultiMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChatId", wireType)
			}
			m.PeerChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RandomId = append(m.RandomId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RandomId = append(m.RandomId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message, &mtproto.Message{})
			if err := m.Message[len(m.Message)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *ChannelMultiMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: ChannelMultiMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChannelMultiMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChannelId", wireType)
			}
			m.PeerChannelId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChannelId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RandomId = append(m.RandomId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RandomId = append(m.RandomId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message, &mtproto.Message{})
			if err := m.Message[len(m.Message)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *OutboxMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: OutboxMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutboxMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoWebpage", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
			m.NoWebpage = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Background", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
			m.Background = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
			m.RandomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &mtproto.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduleDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ScheduleDate == nil {
				m.ScheduleDate = &types.Int32Value{}
			}
			if err := m.ScheduleDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *OutgoingMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: OutgoingMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerType", wireType)
			}
			m.PeerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			m.PeerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &OutboxMessage{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *OutgoingMultiMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: OutgoingMultiMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingMultiMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerType", wireType)
			}
			m.PeerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			m.PeerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultiMessage = append(m.MultiMessage, &OutboxMessage{})
			if err := m.MultiMessage[len(m.MultiMessage)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxUserMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxUserMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxUserMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerUserId", wireType)
			}
			m.PeerUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerUserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
			m.RandomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DialogMessageId", wireType)
			}
			m.DialogMessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DialogMessageId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageDataId", wireType)
			}
			m.MessageDataId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageDataId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &mtproto.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxChatMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxChatMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxChatMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChatId", wireType)
			}
			m.PeerChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
			m.RandomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DialogMessageId", wireType)
			}
			m.DialogMessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DialogMessageId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageDataId", wireType)
			}
			m.MessageDataId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageDataId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &mtproto.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxUserMultiMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxUserMultiMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxUserMultiMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerUserId", wireType)
			}
			m.PeerUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerUserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RandomId = append(m.RandomId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RandomId = append(m.RandomId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
				m.DialogMessageId = append(m.DialogMessageId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
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
					m.DialogMessageId = append(m.DialogMessageId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DialogMessageId", wireType)
			}
		case 5:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.MessageDataId = append(m.MessageDataId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.MessageDataId = append(m.MessageDataId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageDataId", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message, &mtproto.Message{})
			if err := m.Message[len(m.Message)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxChatMultiMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxChatMultiMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxChatMultiMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChatId", wireType)
			}
			m.PeerChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RandomId = append(m.RandomId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RandomId = append(m.RandomId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
				m.DialogMessageId = append(m.DialogMessageId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
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
					m.DialogMessageId = append(m.DialogMessageId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DialogMessageId", wireType)
			}
		case 5:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.MessageDataId = append(m.MessageDataId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.MessageDataId = append(m.MessageDataId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageDataId", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message, &mtproto.Message{})
			if err := m.Message[len(m.Message)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxUserEditMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxUserEditMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxUserEditMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerUserId", wireType)
			}
			m.PeerUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerUserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &mtproto.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxChatEditMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxChatEditMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxChatEditMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChatId", wireType)
			}
			m.PeerChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &mtproto.Message{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *DeleteMessagesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: DeleteMessagesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteMessagesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerType", wireType)
			}
			m.PeerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			m.PeerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revoke", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
			m.Revoke = bool(v != 0)
		case 5:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
				m.Id = append(m.Id, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
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
					m.Id = append(m.Id, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *DeleteHistoryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: DeleteHistoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteHistoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			m.ChannelId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerType", wireType)
			}
			m.PeerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			m.PeerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JustClear", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
			m.JustClear = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revoke", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
			m.Revoke = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxId", wireType)
			}
			m.MaxId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxDeleteMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxDeleteMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxDeleteMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Id = append(m.Id, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Id = append(m.Id, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxUserDeleteHistory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxUserDeleteHistory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxUserDeleteHistory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerUserId", wireType)
			}
			m.PeerUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerUserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JustClear", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
			m.JustClear = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxId", wireType)
			}
			m.MaxId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxChatDeleteHistory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxChatDeleteHistory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxChatDeleteHistory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChatId", wireType)
			}
			m.PeerChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxId", wireType)
			}
			m.MaxId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *ContentMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: ContentMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsMentioned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
			m.IsMentioned = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *ReadMessageContentsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: ReadMessageContentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadMessageContentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerType", wireType)
			}
			m.PeerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			m.PeerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id, &ContentMessage{})
			if err := m.Id[len(m.Id)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxUserReadMediaUnread) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxUserReadMediaUnread: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxUserReadMediaUnread: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
				m.Id = append(m.Id, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
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
					m.Id = append(m.Id, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func (m *InboxChatReadMediaUnread) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: InboxChatReadMediaUnread: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboxChatReadMediaUnread: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &Sender{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerChatId", wireType)
			}
			m.PeerChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
				m.Id = append(m.Id, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMsg
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
					return ErrInvalidLengthMsg
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMsg
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
					m.Id = append(m.Id, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMsg
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
				next, err := skipMsg(dAtA[start:])
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
	ErrInvalidLengthMsg = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("msg.proto", fileDescriptor_msg_a2fb1cd7e36d0df3) }

var fileDescriptor_msg_a2fb1cd7e36d0df3 = []byte{
	// 1204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x5f, 0xd7, 0x49, 0x36, 0x79, 0x69, 0xda, 0xe2, 0xa6, 0x6d, 0x94, 0xd2, 0xa8, 0x8d, 0x04,
	0x5a, 0x81, 0x70, 0xa4, 0xae, 0x84, 0xd4, 0x03, 0x02, 0x36, 0x45, 0xda, 0x08, 0xaa, 0x5d, 0xcc,
	0x16, 0x24, 0x0e, 0x44, 0x93, 0x78, 0xea, 0x98, 0xda, 0x9e, 0x60, 0x8f, 0xd9, 0xf6, 0x13, 0xf0,
	0x01, 0xb8, 0x20, 0xe0, 0xc0, 0x81, 0x0b, 0x7c, 0x12, 0x8e, 0x5c, 0x91, 0x38, 0x40, 0x3f, 0x01,
	0x07, 0x2e, 0x9c, 0x40, 0x33, 0x1e, 0xdb, 0xe3, 0x6c, 0x42, 0x12, 0xab, 0x7f, 0x6e, 0xc9, 0x1b,
	0xbf, 0xdf, 0xfb, 0xfd, 0xde, 0x7b, 0x7e, 0xf3, 0x0c, 0x15, 0x37, 0xb0, 0xf4, 0xb1, 0x4f, 0x28,
	0xd1, 0x8a, 0x6e, 0x60, 0x8d, 0x07, 0xcd, 0x37, 0x2c, 0x9b, 0x8e, 0xc2, 0x81, 0x3e, 0x24, 0x6e,
	0xc7, 0x22, 0x16, 0xe9, 0xf0, 0xd3, 0x41, 0x78, 0xc6, 0xff, 0xf1, 0x3f, 0xfc, 0x57, 0xe4, 0xd5,
	0x6c, 0x06, 0xc3, 0x11, 0x76, 0x91, 0x4e, 0x1d, 0x7d, 0x48, 0x7c, 0xdc, 0xa7, 0x97, 0x63, 0x1c,
	0x88, 0xb3, 0x7a, 0x7a, 0x16, 0x5c, 0x7a, 0x43, 0x61, 0x6d, 0x59, 0x84, 0x58, 0x0e, 0x4e, 0x71,
	0x9f, 0xfb, 0x68, 0x3c, 0xc6, 0xbe, 0xf0, 0x6a, 0x7f, 0x00, 0xa5, 0x8f, 0xb0, 0x67, 0x62, 0x5f,
	0x5b, 0x83, 0x15, 0xdb, 0x6c, 0x28, 0xfb, 0xca, 0x83, 0xa2, 0xb1, 0x62, 0x9b, 0x9a, 0x06, 0x05,
	0x06, 0xdf, 0x58, 0xe1, 0x16, 0xfe, 0x5b, 0x6b, 0x41, 0x15, 0x85, 0x74, 0xd4, 0x3f, 0xc7, 0x97,
	0x7d, 0xdb, 0x6c, 0xa8, 0xfb, 0xca, 0x03, 0xd5, 0xa8, 0x30, 0xd3, 0xfb, 0xf8, 0xb2, 0x67, 0xb6,
	0xbf, 0x53, 0xa0, 0x7a, 0x1a, 0x60, 0xff, 0x04, 0x07, 0x01, 0xb2, 0xb0, 0x76, 0x00, 0x85, 0x33,
	0x9f, 0xb8, 0x1c, 0xb5, 0x7a, 0x58, 0xd3, 0xb9, 0x68, 0x3d, 0x0a, 0x68, 0xf0, 0x23, 0x6d, 0x1f,
	0x56, 0xc7, 0x18, 0xfb, 0xfd, 0x30, 0xc0, 0x3e, 0xc3, 0x8c, 0xc2, 0x01, 0xb3, 0x31, 0xa4, 0x9e,
	0xa9, 0xed, 0x42, 0xc5, 0x47, 0x9e, 0x49, 0xdc, 0x34, 0x64, 0x39, 0x32, 0xf4, 0x4c, 0xed, 0x35,
	0xb8, 0xef, 0x46, 0xc1, 0x1a, 0x05, 0x1e, 0x64, 0x43, 0x77, 0x29, 0x97, 0xa6, 0x0b, 0x12, 0x46,
	0xfc, 0x00, 0x67, 0xd7, 0x1d, 0x21, 0x9a, 0x83, 0xdd, 0x70, 0x84, 0xe8, 0x04, 0x3b, 0x86, 0x74,
	0x9d, 0xec, 0x7e, 0x54, 0x60, 0xad, 0x3b, 0x42, 0x9e, 0x87, 0x9d, 0x25, 0x08, 0xbe, 0x0a, 0xeb,
	0x31, 0x41, 0xe6, 0x99, 0x72, 0xac, 0x09, 0x8e, 0xcc, 0x7a, 0x9d, 0x34, 0x7f, 0x50, 0x60, 0x83,
	0x97, 0x38, 0x74, 0xa8, 0x7d, 0x93, 0x75, 0x56, 0x67, 0x53, 0x54, 0xe7, 0x53, 0xe4, 0x75, 0xce,
	0x49, 0x71, 0xe1, 0x62, 0xe7, 0xa7, 0xf8, 0x93, 0x02, 0x9b, 0x71, 0xb1, 0x97, 0x64, 0x99, 0xb3,
	0xe2, 0xf9, 0xb9, 0xfe, 0xa6, 0x40, 0xed, 0x49, 0x48, 0x07, 0xe4, 0x22, 0x66, 0xb9, 0x07, 0xe0,
	0x91, 0xfe, 0x73, 0x3c, 0x18, 0x33, 0x00, 0xc6, 0xb5, 0x6c, 0x54, 0x3c, 0xf2, 0x49, 0x64, 0xd0,
	0x5a, 0x00, 0x03, 0x34, 0x3c, 0xb7, 0x7c, 0x12, 0x7a, 0x11, 0xb9, 0xb2, 0x21, 0x59, 0xae, 0xad,
	0x17, 0xb5, 0x77, 0xa0, 0xc6, 0x86, 0x9e, 0x19, 0x3a, 0xb8, 0x6f, 0x22, 0x8a, 0x1b, 0x45, 0xee,
	0xb1, 0xab, 0x47, 0x43, 0x4f, 0x8f, 0x87, 0x9e, 0xde, 0xf3, 0xe8, 0xc3, 0xc3, 0x8f, 0x91, 0x13,
	0x62, 0x63, 0x35, 0xf6, 0x38, 0x46, 0x14, 0xb7, 0xbf, 0x55, 0x60, 0xfd, 0x49, 0x48, 0x2d, 0x62,
	0x7b, 0xd6, 0x12, 0x35, 0xd8, 0x85, 0x0a, 0xaf, 0x81, 0x34, 0x20, 0xcb, 0xcc, 0xf0, 0x8c, 0x0d,
	0xc9, 0x1d, 0xb8, 0xcf, 0x0f, 0x85, 0xb8, 0xa2, 0x51, 0x62, 0x7f, 0x7b, 0xa6, 0xa6, 0x4f, 0x4a,
	0xab, 0x0b, 0xec, 0x4c, 0x76, 0xd3, 0xc4, 0xff, 0xac, 0x40, 0x3d, 0x21, 0xb7, 0x64, 0x97, 0xe4,
	0x63, 0x78, 0x04, 0x35, 0x97, 0x05, 0xea, 0x67, 0x9b, 0x63, 0x3a, 0xcf, 0x55, 0x57, 0xe2, 0xd4,
	0xfe, 0x5b, 0x81, 0x8d, 0x9e, 0x37, 0x20, 0x17, 0xb7, 0x3d, 0xff, 0x5f, 0x32, 0x6d, 0xe4, 0x10,
	0x2b, 0xa6, 0xcc, 0x1e, 0x2a, 0x70, 0x8c, 0xf5, 0xe8, 0x40, 0x70, 0xe9, 0x99, 0xec, 0xcd, 0x89,
	0x1f, 0x32, 0x11, 0x45, 0xec, 0xc9, 0x22, 0x87, 0xab, 0x09, 0xf3, 0x31, 0xa2, 0x28, 0xdb, 0x82,
	0xa5, 0x79, 0xe3, 0x30, 0x91, 0x7d, 0xdb, 0x17, 0xcb, 0xdd, 0xca, 0xfe, 0x47, 0x81, 0xad, 0xb4,
	0xda, 0xb7, 0x7d, 0x15, 0x4c, 0xd5, 0xae, 0x2e, 0xac, 0x5d, 0x9d, 0xa3, 0x5d, 0x5d, 0x50, 0xfb,
	0x9d, 0xdc, 0x31, 0x77, 0xab, 0xfd, 0x2b, 0x05, 0xea, 0x49, 0xdd, 0xdf, 0x33, 0x6d, 0x7a, 0xad,
	0x65, 0x5f, 0xa6, 0x03, 0x13, 0x26, 0x2c, 0x6d, 0x39, 0x99, 0xcc, 0x2e, 0xc2, 0x32, 0x4c, 0xbe,
	0x51, 0x60, 0xeb, 0x18, 0x3b, 0x98, 0x62, 0x71, 0x14, 0x18, 0xf8, 0x8b, 0x10, 0x07, 0xf4, 0xe6,
	0xe6, 0xf4, 0x36, 0x94, 0x7c, 0xfc, 0x25, 0x39, 0x8f, 0x2e, 0x92, 0xb2, 0x21, 0xfe, 0x89, 0x1d,
	0xbe, 0xc8, 0xfb, 0x61, 0xc5, 0x36, 0xdb, 0xbf, 0x2b, 0x50, 0x8f, 0xa8, 0x3d, 0xb6, 0x03, 0x4a,
	0xfc, 0xcb, 0x25, 0x98, 0xed, 0x01, 0xbc, 0xb0, 0x62, 0x54, 0x86, 0xf2, 0x7a, 0x91, 0x12, 0x57,
	0x67, 0x13, 0x2f, 0x64, 0x88, 0xef, 0x01, 0x7c, 0x1e, 0x06, 0xb4, 0x3f, 0x74, 0x30, 0xf2, 0xf9,
	0x18, 0x2a, 0x1b, 0x15, 0x66, 0xe9, 0x32, 0x83, 0xa4, 0xab, 0x94, 0xd1, 0xb5, 0x05, 0x25, 0x17,
	0x5d, 0x30, 0xb8, 0xfb, 0x1c, 0xae, 0xe8, 0xa2, 0x8b, 0x9e, 0xd9, 0x7e, 0x0c, 0x9b, 0xbc, 0x05,
	0xb2, 0xd9, 0x5f, 0x44, 0x5c, 0x94, 0xa8, 0x15, 0xfe, 0x3a, 0xb0, 0x44, 0x7d, 0xad, 0xc0, 0x76,
	0xd2, 0xd7, 0x99, 0x8c, 0x5d, 0x4f, 0x67, 0x67, 0x75, 0xab, 0x93, 0xba, 0x53, 0x7d, 0x05, 0x59,
	0x1f, 0x15, 0xa4, 0x58, 0x53, 0xe6, 0x26, 0x35, 0xbb, 0xc9, 0xd3, 0xa8, 0xaa, 0x1c, 0xb5, 0x0b,
	0x6b, 0x5d, 0xe2, 0x51, 0xec, 0x25, 0xaf, 0xd4, 0xe4, 0xa7, 0xe1, 0x01, 0xac, 0xda, 0x41, 0xdf,
	0xc5, 0x1e, 0xb5, 0x89, 0x87, 0xe3, 0x15, 0xaf, 0x6a, 0x07, 0x27, 0xb1, 0xa9, 0xfd, 0xbd, 0x02,
	0x4d, 0x03, 0x23, 0x53, 0x40, 0x08, 0xc0, 0x9b, 0x7f, 0x33, 0x5e, 0xe1, 0x54, 0xa3, 0xb5, 0x65,
	0x4b, 0xc0, 0x66, 0xd5, 0xf0, 0x7a, 0x9f, 0x40, 0x23, 0x29, 0x77, 0x44, 0xd3, 0xb4, 0xd1, 0xa9,
	0xe7, 0x63, 0x64, 0x2e, 0xde, 0x3e, 0x85, 0xe4, 0x3d, 0x23, 0x02, 0x8e, 0x25, 0x36, 0x07, 0xdc,
	0xfc, 0x52, 0x45, 0x01, 0xd5, 0x38, 0xe0, 0xe1, 0xbf, 0x65, 0x28, 0x19, 0x4f, 0xbb, 0x27, 0x81,
	0xa5, 0x1d, 0xc1, 0x3a, 0x03, 0x93, 0xd7, 0x2e, 0x4d, 0x04, 0x91, 0x6c, 0xcd, 0x74, 0x80, 0x9d,
	0x8e, 0xd9, 0xae, 0x1c, 0xb4, 0xef, 0xc5, 0xae, 0xf2, 0xea, 0x12, 0xbb, 0x4a, 0xb6, 0xa9, 0xae,
	0x6f, 0x83, 0x26, 0x5c, 0xe5, 0x0f, 0xd6, 0xad, 0xd4, 0x5b, 0x32, 0x4f, 0x05, 0xe8, 0x42, 0x3d,
	0xa1, 0x2d, 0xdf, 0xa1, 0x3b, 0x32, 0x77, 0xe9, 0xe0, 0xff, 0x40, 0x5e, 0xb8, 0x88, 0x77, 0x64,
	0x15, 0xf3, 0x40, 0x7a, 0xb0, 0x23, 0x4b, 0x91, 0x71, 0x9a, 0x13, 0x7a, 0xe6, 0x41, 0xbd, 0x09,
	0xeb, 0x4f, 0xc3, 0x60, 0x34, 0xaf, 0x16, 0xb5, 0xc4, 0xf5, 0x11, 0x21, 0x4e, 0x54, 0x08, 0x76,
	0x85, 0xe5, 0xac, 0x21, 0x73, 0xcd, 0x59, 0x43, 0xe1, 0x9a, 0xb3, 0x86, 0x47, 0x50, 0x65, 0x99,
	0x8b, 0x3d, 0xb7, 0xd3, 0xcf, 0x04, 0xf9, 0x83, 0x6a, 0x46, 0xe5, 0x36, 0xb8, 0xab, 0x9c, 0xed,
	0xdd, 0x49, 0xff, 0xf9, 0xe9, 0xae, 0xb2, 0x74, 0xcf, 0x8b, 0x3f, 0x25, 0xdd, 0x55, 0x79, 0x63,
	0x58, 0x86, 0xf7, 0x33, 0x58, 0x9b, 0xb8, 0x6d, 0x5e, 0x16, 0xde, 0x53, 0x57, 0x80, 0x66, 0x3b,
	0xc1, 0x10, 0xeb, 0x42, 0xd0, 0x7f, 0xf7, 0xec, 0x0c, 0x0f, 0x29, 0x8e, 0xd3, 0xc6, 0x50, 0x3f,
	0x84, 0x5a, 0x76, 0xbe, 0xef, 0x66, 0x40, 0xb3, 0x97, 0x77, 0xf3, 0x60, 0x36, 0xa6, 0x78, 0xb2,
	0x7d, 0x4f, 0xfb, 0x0c, 0x36, 0xa7, 0xcc, 0x5f, 0xed, 0x40, 0x00, 0xcf, 0x9e, 0xcd, 0x8b, 0x51,
	0x7e, 0xf4, 0xd6, 0x5f, 0x7f, 0xb6, 0x94, 0x5f, 0xae, 0x5a, 0xca, 0xaf, 0x57, 0x2d, 0xe5, 0x8f,
	0xab, 0x96, 0xf2, 0xe9, 0xeb, 0x1e, 0x1e, 0x84, 0x0e, 0xd2, 0xd9, 0x14, 0xeb, 0x0c, 0x9d, 0x30,
	0xa0, 0xd8, 0xef, 0xa0, 0xf1, 0xb8, 0xc3, 0x60, 0xb0, 0x67, 0x61, 0xbf, 0xe3, 0x06, 0x56, 0x87,
	0xc7, 0x1f, 0x94, 0x78, 0x84, 0x87, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x67, 0xf7, 0x98, 0x90,
	0x13, 0x15, 0x00, 0x00,
}
