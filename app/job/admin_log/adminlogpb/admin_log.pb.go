package adminlogpb

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
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

type ChannelAdminLogEventData struct {
	LogUserId            int32                               `protobuf:"varint,1,opt,name=log_user_id,json=logUserId,proto3" json:"log_user_id,omitempty"`
	ChannelId            int32                               `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Event                *mtproto.ChannelAdminLogEventAction `protobuf:"bytes,3,opt,name=event" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ChannelAdminLogEventData) Reset()         { *m = ChannelAdminLogEventData{} }
func (m *ChannelAdminLogEventData) String() string { return proto.CompactTextString(m) }
func (*ChannelAdminLogEventData) ProtoMessage()    {}
func (*ChannelAdminLogEventData) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_log_0bd7cea4b8a694ab, []int{0}
}
func (m *ChannelAdminLogEventData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChannelAdminLogEventData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChannelAdminLogEventData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ChannelAdminLogEventData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelAdminLogEventData.Merge(dst, src)
}
func (m *ChannelAdminLogEventData) XXX_Size() int {
	return m.Size()
}
func (m *ChannelAdminLogEventData) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelAdminLogEventData.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelAdminLogEventData proto.InternalMessageInfo

func (m *ChannelAdminLogEventData) GetLogUserId() int32 {
	if m != nil {
		return m.LogUserId
	}
	return 0
}

func (m *ChannelAdminLogEventData) GetChannelId() int32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *ChannelAdminLogEventData) GetEvent() *mtproto.ChannelAdminLogEventAction {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelAdminLogEventData)(nil), "pushpb.ChannelAdminLogEventData")
}
func (this *ChannelAdminLogEventData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&adminlogpb.ChannelAdminLogEventData{")
	s = append(s, "LogUserId: "+fmt.Sprintf("%#v", this.LogUserId)+",")
	s = append(s, "ChannelId: "+fmt.Sprintf("%#v", this.ChannelId)+",")
	if this.Event != nil {
		s = append(s, "Event: "+fmt.Sprintf("%#v", this.Event)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAdminLog(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ChannelAdminLogEventData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelAdminLogEventData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LogUserId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAdminLog(dAtA, i, uint64(m.LogUserId))
	}
	if m.ChannelId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAdminLog(dAtA, i, uint64(m.ChannelId))
	}
	if m.Event != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAdminLog(dAtA, i, uint64(m.Event.Size()))
		n1, err := m.Event.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAdminLog(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ChannelAdminLogEventData) Size() (n int) {
	var l int
	_ = l
	if m.LogUserId != 0 {
		n += 1 + sovAdminLog(uint64(m.LogUserId))
	}
	if m.ChannelId != 0 {
		n += 1 + sovAdminLog(uint64(m.ChannelId))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovAdminLog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAdminLog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAdminLog(x uint64) (n int) {
	return sovAdminLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChannelAdminLogEventData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdminLog
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
			return fmt.Errorf("proto: ChannelAdminLogEventData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChannelAdminLogEventData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogUserId", wireType)
			}
			m.LogUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdminLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogUserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			m.ChannelId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdminLog
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdminLog
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
				return ErrInvalidLengthAdminLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &mtproto.ChannelAdminLogEventAction{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdminLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdminLog
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
func skipAdminLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdminLog
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
					return 0, ErrIntOverflowAdminLog
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
					return 0, ErrIntOverflowAdminLog
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
				return 0, ErrInvalidLengthAdminLog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAdminLog
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
				next, err := skipAdminLog(dAtA[start:])
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
	ErrInvalidLengthAdminLog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdminLog   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("admin_log.proto", fileDescriptor_admin_log_0bd7cea4b8a694ab) }

var fileDescriptor_admin_log_0bd7cea4b8a694ab = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4a, 0x34, 0x31,
	0x18, 0xc6, 0xc9, 0xf7, 0xb1, 0x0b, 0x66, 0x0b, 0x61, 0xb0, 0x18, 0x16, 0x0c, 0x8b, 0x36, 0xdb,
	0x98, 0x88, 0x56, 0x96, 0xbb, 0x6a, 0xb1, 0x60, 0xb5, 0x60, 0x63, 0x33, 0x24, 0x99, 0x98, 0x8c,
	0x64, 0xf2, 0x86, 0x49, 0x22, 0x78, 0x0f, 0x0f, 0x65, 0xe9, 0x11, 0x74, 0x4e, 0xe0, 0x11, 0x64,
	0x12, 0xb1, 0xb2, 0x7b, 0xfe, 0xf0, 0xbc, 0x3f, 0x78, 0xf1, 0x21, 0x6f, 0xfb, 0xce, 0x35, 0x16,
	0x34, 0xf5, 0x03, 0x44, 0xa8, 0xe6, 0x3e, 0x05, 0xe3, 0xc5, 0xf2, 0x4c, 0x77, 0xd1, 0x24, 0x41,
	0x25, 0xf4, 0x4c, 0x83, 0x06, 0x96, 0x6b, 0x91, 0x1e, 0xb3, 0xcb, 0x26, 0xab, 0x32, 0x5b, 0x1e,
	0x05, 0x69, 0x54, 0xcf, 0x69, 0xb4, 0x34, 0xbc, 0x38, 0x59, 0xd2, 0x93, 0x57, 0x84, 0xeb, 0x6b,
	0xc3, 0x9d, 0x53, 0x76, 0x33, 0x71, 0xee, 0x40, 0xdf, 0x3e, 0x2b, 0x17, 0x6f, 0x78, 0xe4, 0x15,
	0xc1, 0x0b, 0x0b, 0xba, 0x49, 0x41, 0x0d, 0x4d, 0xd7, 0xd6, 0x68, 0x85, 0xd6, 0xb3, 0xfd, 0x81,
	0x05, 0x7d, 0x1f, 0xd4, 0xb0, 0x6b, 0xab, 0x63, 0x8c, 0x65, 0xd9, 0x4e, 0xf5, 0xbf, 0x52, 0xff,
	0x24, 0xbb, 0xb6, 0xba, 0xc2, 0x33, 0x35, 0xdd, 0xaa, 0xff, 0xaf, 0xd0, 0x7a, 0x71, 0x71, 0x4a,
	0xfb, 0x98, 0xa1, 0xf4, 0x2f, 0xe0, 0x46, 0xc6, 0x0e, 0xdc, 0xbe, 0x2c, 0xb6, 0xdb, 0xaf, 0x4f,
	0x82, 0xde, 0x46, 0x82, 0xde, 0x47, 0x82, 0x3e, 0x46, 0x82, 0x1e, 0xce, 0x9d, 0x12, 0xc9, 0x72,
	0x2a, 0x0d, 0x8f, 0x4c, 0xda, 0x14, 0xa2, 0x1a, 0x18, 0xf7, 0x9e, 0x3d, 0x81, 0x60, 0xbf, 0x2f,
	0x2a, 0xca, 0x82, 0xf6, 0x42, 0xcc, 0x33, 0xec, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x81,
	0x62, 0xd8, 0x41, 0x01, 0x00, 0x00,
}
