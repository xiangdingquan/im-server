package relaypb

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
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

type CallConnections struct {
	Id                     int64                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerTag                []byte                     `protobuf:"bytes,2,opt,name=peer_tag,json=peerTag,proto3" json:"peer_tag,omitempty"`
	Connection             *mtproto.PhoneConnection   `protobuf:"bytes,3,opt,name=connection" json:"connection,omitempty"`
	AlternativeConnections []*mtproto.PhoneConnection `protobuf:"bytes,4,rep,name=alternative_connections,json=alternativeConnections" json:"alternative_connections,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                   `json:"-"`
	XXX_unrecognized       []byte                     `json:"-"`
	XXX_sizecache          int32                      `json:"-"`
}

func (m *CallConnections) Reset()         { *m = CallConnections{} }
func (m *CallConnections) String() string { return proto.CompactTextString(m) }
func (*CallConnections) ProtoMessage()    {}
func (*CallConnections) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_fb7168e89d12c41a, []int{0}
}
func (m *CallConnections) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CallConnections) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CallConnections.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CallConnections) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallConnections.Merge(dst, src)
}
func (m *CallConnections) XXX_Size() int {
	return m.Size()
}
func (m *CallConnections) XXX_DiscardUnknown() {
	xxx_messageInfo_CallConnections.DiscardUnknown(m)
}

var xxx_messageInfo_CallConnections proto.InternalMessageInfo

func (m *CallConnections) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CallConnections) GetPeerTag() []byte {
	if m != nil {
		return m.PeerTag
	}
	return nil
}

func (m *CallConnections) GetConnection() *mtproto.PhoneConnection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *CallConnections) GetAlternativeConnections() []*mtproto.PhoneConnection {
	if m != nil {
		return m.AlternativeConnections
	}
	return nil
}

type RelayCreateCallRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AdminId              int32    `protobuf:"varint,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	ParticipantId        int32    `protobuf:"varint,3,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelayCreateCallRequest) Reset()         { *m = RelayCreateCallRequest{} }
func (m *RelayCreateCallRequest) String() string { return proto.CompactTextString(m) }
func (*RelayCreateCallRequest) ProtoMessage()    {}
func (*RelayCreateCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_fb7168e89d12c41a, []int{1}
}
func (m *RelayCreateCallRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RelayCreateCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RelayCreateCallRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RelayCreateCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayCreateCallRequest.Merge(dst, src)
}
func (m *RelayCreateCallRequest) XXX_Size() int {
	return m.Size()
}
func (m *RelayCreateCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayCreateCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RelayCreateCallRequest proto.InternalMessageInfo

func (m *RelayCreateCallRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RelayCreateCallRequest) GetAdminId() int32 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

func (m *RelayCreateCallRequest) GetParticipantId() int32 {
	if m != nil {
		return m.ParticipantId
	}
	return 0
}

type RelaydiscardCallRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelaydiscardCallRequest) Reset()         { *m = RelaydiscardCallRequest{} }
func (m *RelaydiscardCallRequest) String() string { return proto.CompactTextString(m) }
func (*RelaydiscardCallRequest) ProtoMessage()    {}
func (*RelaydiscardCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_fb7168e89d12c41a, []int{2}
}
func (m *RelaydiscardCallRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RelaydiscardCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RelaydiscardCallRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RelaydiscardCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelaydiscardCallRequest.Merge(dst, src)
}
func (m *RelaydiscardCallRequest) XXX_Size() int {
	return m.Size()
}
func (m *RelaydiscardCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RelaydiscardCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RelaydiscardCallRequest proto.InternalMessageInfo

func (m *RelaydiscardCallRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*CallConnections)(nil), "relaypb.CallConnections")
	proto.RegisterType((*RelayCreateCallRequest)(nil), "relaypb.RelayCreateCallRequest")
	proto.RegisterType((*RelaydiscardCallRequest)(nil), "relaypb.RelaydiscardCallRequest")
}
func (this *CallConnections) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&relaypb.CallConnections{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	s = append(s, "PeerTag: "+fmt.Sprintf("%#v", this.PeerTag)+",")
	if this.Connection != nil {
		s = append(s, "Connection: "+fmt.Sprintf("%#v", this.Connection)+",")
	}
	if this.AlternativeConnections != nil {
		s = append(s, "AlternativeConnections: "+fmt.Sprintf("%#v", this.AlternativeConnections)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RelayCreateCallRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&relaypb.RelayCreateCallRequest{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	s = append(s, "AdminId: "+fmt.Sprintf("%#v", this.AdminId)+",")
	s = append(s, "ParticipantId: "+fmt.Sprintf("%#v", this.ParticipantId)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RelaydiscardCallRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&relaypb.RelaydiscardCallRequest{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRelay(v interface{}, typ string) string {
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

// RPCRelayClient is the client API for RPCRelay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCRelayClient interface {
	RelayCreateCall(ctx context.Context, in *RelayCreateCallRequest, opts ...grpc.CallOption) (*CallConnections, error)
	RelayDiscardCall(ctx context.Context, in *RelaydiscardCallRequest, opts ...grpc.CallOption) (*mtproto.Bool, error)
}

type rPCRelayClient struct {
	cc *grpc.ClientConn
}

func NewRPCRelayClient(cc *grpc.ClientConn) RPCRelayClient {
	return &rPCRelayClient{cc}
}

func (c *rPCRelayClient) RelayCreateCall(ctx context.Context, in *RelayCreateCallRequest, opts ...grpc.CallOption) (*CallConnections, error) {
	out := new(CallConnections)
	err := c.cc.Invoke(ctx, "/relaypb.RPCRelay/relayCreateCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCRelayClient) RelayDiscardCall(ctx context.Context, in *RelaydiscardCallRequest, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, "/relaypb.RPCRelay/relayDiscardCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCRelay service

type RPCRelayServer interface {
	RelayCreateCall(context.Context, *RelayCreateCallRequest) (*CallConnections, error)
	RelayDiscardCall(context.Context, *RelaydiscardCallRequest) (*mtproto.Bool, error)
}

func RegisterRPCRelayServer(s *grpc.Server, srv RPCRelayServer) {
	s.RegisterService(&_RPCRelay_serviceDesc, srv)
}

func _RPCRelay_RelayCreateCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelayCreateCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCRelayServer).RelayCreateCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relaypb.RPCRelay/RelayCreateCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCRelayServer).RelayCreateCall(ctx, req.(*RelayCreateCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCRelay_RelayDiscardCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelaydiscardCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCRelayServer).RelayDiscardCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relaypb.RPCRelay/RelayDiscardCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCRelayServer).RelayDiscardCall(ctx, req.(*RelaydiscardCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCRelay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "relaypb.RPCRelay",
	HandlerType: (*RPCRelayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "relayCreateCall",
			Handler:    _RPCRelay_RelayCreateCall_Handler,
		},
		{
			MethodName: "relayDiscardCall",
			Handler:    _RPCRelay_RelayDiscardCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relay.proto",
}

func (m *CallConnections) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallConnections) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRelay(dAtA, i, uint64(m.Id))
	}
	if len(m.PeerTag) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRelay(dAtA, i, uint64(len(m.PeerTag)))
		i += copy(dAtA[i:], m.PeerTag)
	}
	if m.Connection != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRelay(dAtA, i, uint64(m.Connection.Size()))
		n1, err := m.Connection.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.AlternativeConnections) > 0 {
		for _, msg := range m.AlternativeConnections {
			dAtA[i] = 0x22
			i++
			i = encodeVarintRelay(dAtA, i, uint64(msg.Size()))
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

func (m *RelayCreateCallRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RelayCreateCallRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRelay(dAtA, i, uint64(m.Id))
	}
	if m.AdminId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRelay(dAtA, i, uint64(m.AdminId))
	}
	if m.ParticipantId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRelay(dAtA, i, uint64(m.ParticipantId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RelaydiscardCallRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RelaydiscardCallRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRelay(dAtA, i, uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRelay(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CallConnections) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRelay(uint64(m.Id))
	}
	l = len(m.PeerTag)
	if l > 0 {
		n += 1 + l + sovRelay(uint64(l))
	}
	if m.Connection != nil {
		l = m.Connection.Size()
		n += 1 + l + sovRelay(uint64(l))
	}
	if len(m.AlternativeConnections) > 0 {
		for _, e := range m.AlternativeConnections {
			l = e.Size()
			n += 1 + l + sovRelay(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RelayCreateCallRequest) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRelay(uint64(m.Id))
	}
	if m.AdminId != 0 {
		n += 1 + sovRelay(uint64(m.AdminId))
	}
	if m.ParticipantId != 0 {
		n += 1 + sovRelay(uint64(m.ParticipantId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RelaydiscardCallRequest) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRelay(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRelay(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRelay(x uint64) (n int) {
	return sovRelay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CallConnections) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelay
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
			return fmt.Errorf("proto: CallConnections: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallConnections: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerTag", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
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
				return ErrInvalidLengthRelay
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerTag = append(m.PeerTag[:0], dAtA[iNdEx:postIndex]...)
			if m.PeerTag == nil {
				m.PeerTag = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Connection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
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
				return ErrInvalidLengthRelay
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Connection == nil {
				m.Connection = &mtproto.PhoneConnection{}
			}
			if err := m.Connection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlternativeConnections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
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
				return ErrInvalidLengthRelay
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AlternativeConnections = append(m.AlternativeConnections, &mtproto.PhoneConnection{})
			if err := m.AlternativeConnections[len(m.AlternativeConnections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRelay
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
func (m *RelayCreateCallRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelay
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
			return fmt.Errorf("proto: RelayCreateCallRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RelayCreateCallRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminId", wireType)
			}
			m.AdminId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipantId", wireType)
			}
			m.ParticipantId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipantId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRelay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRelay
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
func (m *RelaydiscardCallRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelay
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
			return fmt.Errorf("proto: RelaydiscardCallRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RelaydiscardCallRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRelay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRelay
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
func skipRelay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRelay
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
					return 0, ErrIntOverflowRelay
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
					return 0, ErrIntOverflowRelay
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
				return 0, ErrInvalidLengthRelay
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRelay
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
				next, err := skipRelay(dAtA[start:])
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
	ErrInvalidLengthRelay = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRelay   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("relay.proto", fileDescriptor_relay_fb7168e89d12c41a) }

var fileDescriptor_relay_fb7168e89d12c41a = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0xad, 0x37, 0xc0, 0x56, 0x2e, 0x6d, 0x91, 0x85, 0xda, 0x90, 0x43, 0x88, 0x22, 0x90, 0x96,
	0x03, 0x8e, 0x54, 0x24, 0xc4, 0xb9, 0x81, 0x43, 0x2f, 0xa8, 0x44, 0x9c, 0xb8, 0x44, 0x8e, 0x3d,
	0x24, 0x46, 0x89, 0x1d, 0x6c, 0x07, 0x69, 0x3f, 0x87, 0xbf, 0x81, 0x1b, 0x9f, 0x00, 0xfb, 0x05,
	0x7c, 0x02, 0x8a, 0xbb, 0xda, 0xa4, 0xb4, 0xea, 0xcd, 0xcf, 0xef, 0xcd, 0x9b, 0x37, 0x33, 0xf8,
	0xc0, 0x40, 0xcb, 0xd6, 0xb4, 0x37, 0xda, 0x69, 0xb2, 0xf4, 0xa0, 0xaf, 0xa2, 0x97, 0xb5, 0x74,
	0xcd, 0x50, 0x51, 0xae, 0xbb, 0xac, 0xd6, 0xb5, 0xce, 0x3c, 0x5f, 0x0d, 0x9f, 0x3d, 0xf2, 0xc0,
	0xbf, 0xae, 0xea, 0xa2, 0xc8, 0xf2, 0x06, 0x3a, 0x46, 0x5d, 0x4b, 0xb9, 0x36, 0x50, 0xba, 0x75,
	0x0f, 0x76, 0xcb, 0x3d, 0x9e, 0x38, 0xbb, 0x56, 0xfc, 0xea, 0x37, 0xfd, 0x89, 0xf0, 0x71, 0xce,
	0xda, 0x36, 0xd7, 0x4a, 0x01, 0x77, 0x52, 0x2b, 0x4b, 0x8e, 0xf0, 0x42, 0x8a, 0x10, 0x25, 0x68,
	0x15, 0x14, 0x0b, 0x29, 0xc8, 0x13, 0xbc, 0xdf, 0x03, 0x98, 0xd2, 0xb1, 0x3a, 0x5c, 0x24, 0x68,
	0xf5, 0xb0, 0x58, 0x8e, 0xf8, 0x23, 0xab, 0xc9, 0x1b, 0x8c, 0xf9, 0xae, 0x32, 0x0c, 0x12, 0xb4,
	0x3a, 0x38, 0x0b, 0x69, 0xe7, 0xbc, 0x39, 0xbd, 0x6c, 0xb4, 0x82, 0xc9, 0xb9, 0x98, 0x69, 0xc9,
	0x07, 0x7c, 0xca, 0x5a, 0x07, 0x46, 0x31, 0x27, 0xbf, 0x41, 0x39, 0x31, 0x36, 0xbc, 0x97, 0x04,
	0x77, 0xda, 0x9c, 0xcc, 0x0a, 0x67, 0xb9, 0xd3, 0x2f, 0xf8, 0xa4, 0x18, 0xf7, 0x96, 0x1b, 0x60,
	0x0e, 0xc6, 0xa9, 0x0a, 0xf8, 0x3a, 0x80, 0x75, 0xb7, 0x4d, 0xc4, 0x44, 0x27, 0x55, 0x29, 0x85,
	0x9f, 0xe8, 0x7e, 0xb1, 0xf4, 0xf8, 0x42, 0x90, 0xe7, 0xf8, 0xa8, 0x67, 0xc6, 0x49, 0x2e, 0x7b,
	0xa6, 0xdc, 0x28, 0x08, 0xbc, 0xe0, 0x70, 0xf6, 0x7b, 0x21, 0xd2, 0x17, 0xf8, 0xd4, 0xf7, 0x12,
	0xd2, 0x72, 0x66, 0xc4, 0x1d, 0xcd, 0xce, 0xbe, 0x23, 0xbc, 0x5f, 0x5c, 0xe6, 0x5e, 0x4e, 0xde,
	0xe3, 0x63, 0x73, 0x3d, 0x23, 0x79, 0x4a, 0xb7, 0xd7, 0xa6, 0xb7, 0xa7, 0x8f, 0xc2, 0x9d, 0xe0,
	0xbf, 0x4b, 0xa5, 0x7b, 0xe4, 0x1d, 0x7e, 0xe4, 0xc9, 0xb7, 0x53, 0x0e, 0x92, 0x5c, 0x37, 0xbc,
	0x19, 0x31, 0x3a, 0xdc, 0xed, 0xf6, 0x5c, 0xeb, 0x36, 0xdd, 0x3b, 0x7f, 0xfd, 0xf7, 0x4f, 0x8c,
	0x7e, 0x6c, 0x62, 0xf4, 0x6b, 0x13, 0xa3, 0xdf, 0x9b, 0x18, 0x7d, 0x7a, 0xa6, 0xa0, 0x1a, 0x5a,
	0x46, 0x79, 0xc3, 0x5c, 0x06, 0xca, 0x81, 0xe9, 0x8d, 0xb4, 0x90, 0x79, 0xeb, 0x6c, 0xdb, 0xa0,
	0x7a, 0xe0, 0x5d, 0x5e, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x47, 0x39, 0xb6, 0xbe, 0x02,
	0x00, 0x00,
}
