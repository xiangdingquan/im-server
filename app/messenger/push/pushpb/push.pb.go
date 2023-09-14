// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: push.proto

package pushpb // import "open.chat/app/messenger/push/pushpb"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import mtproto "open.chat/mtproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PushUpdatesIfNot struct {
	UserId               int32            `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Excludes             []int64          `protobuf:"varint,2,rep,packed,name=excludes" json:"excludes,omitempty"`
	Updates              *mtproto.Updates `protobuf:"bytes,3,opt,name=updates" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PushUpdatesIfNot) Reset()         { *m = PushUpdatesIfNot{} }
func (m *PushUpdatesIfNot) String() string { return proto.CompactTextString(m) }
func (*PushUpdatesIfNot) ProtoMessage()    {}
func (*PushUpdatesIfNot) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_d8af02f530495561, []int{0}
}
func (m *PushUpdatesIfNot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushUpdatesIfNot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushUpdatesIfNot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PushUpdatesIfNot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushUpdatesIfNot.Merge(dst, src)
}
func (m *PushUpdatesIfNot) XXX_Size() int {
	return m.Size()
}
func (m *PushUpdatesIfNot) XXX_DiscardUnknown() {
	xxx_messageInfo_PushUpdatesIfNot.DiscardUnknown(m)
}

var xxx_messageInfo_PushUpdatesIfNot proto.InternalMessageInfo

func (m *PushUpdatesIfNot) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PushUpdatesIfNot) GetExcludes() []int64 {
	if m != nil {
		return m.Excludes
	}
	return nil
}

func (m *PushUpdatesIfNot) GetUpdates() *mtproto.Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

type PushUpdates struct {
	AuthKeyId            int64            `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	Updates              *mtproto.Updates `protobuf:"bytes,2,opt,name=updates" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PushUpdates) Reset()         { *m = PushUpdates{} }
func (m *PushUpdates) String() string { return proto.CompactTextString(m) }
func (*PushUpdates) ProtoMessage()    {}
func (*PushUpdates) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_d8af02f530495561, []int{1}
}
func (m *PushUpdates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushUpdates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushUpdates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PushUpdates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushUpdates.Merge(dst, src)
}
func (m *PushUpdates) XXX_Size() int {
	return m.Size()
}
func (m *PushUpdates) XXX_DiscardUnknown() {
	xxx_messageInfo_PushUpdates.DiscardUnknown(m)
}

var xxx_messageInfo_PushUpdates proto.InternalMessageInfo

func (m *PushUpdates) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *PushUpdates) GetUpdates() *mtproto.Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

func init() {
	proto.RegisterType((*PushUpdatesIfNot)(nil), "pushpb.PushUpdatesIfNot")
	proto.RegisterType((*PushUpdates)(nil), "pushpb.PushUpdates")
}
func (this *PushUpdatesIfNot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pushpb.PushUpdatesIfNot{")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",")
	s = append(s, "Excludes: "+fmt.Sprintf("%#v", this.Excludes)+",")
	if this.Updates != nil {
		s = append(s, "Updates: "+fmt.Sprintf("%#v", this.Updates)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PushUpdates) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pushpb.PushUpdates{")
	s = append(s, "AuthKeyId: "+fmt.Sprintf("%#v", this.AuthKeyId)+",")
	if this.Updates != nil {
		s = append(s, "Updates: "+fmt.Sprintf("%#v", this.Updates)+",")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPush(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PushUpdatesIfNot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushUpdatesIfNot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPush(dAtA, i, uint64(m.UserId))
	}
	if len(m.Excludes) > 0 {
		dAtA2 := make([]byte, len(m.Excludes)*10)
		var j1 int
		for _, num1 := range m.Excludes {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintPush(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.Updates != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPush(dAtA, i, uint64(m.Updates.Size()))
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

func (m *PushUpdates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushUpdates) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AuthKeyId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPush(dAtA, i, uint64(m.AuthKeyId))
	}
	if m.Updates != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPush(dAtA, i, uint64(m.Updates.Size()))
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

func encodeVarintPush(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PushUpdatesIfNot) Size() (n int) {
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovPush(uint64(m.UserId))
	}
	if len(m.Excludes) > 0 {
		l = 0
		for _, e := range m.Excludes {
			l += sovPush(uint64(e))
		}
		n += 1 + sovPush(uint64(l)) + l
	}
	if m.Updates != nil {
		l = m.Updates.Size()
		n += 1 + l + sovPush(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PushUpdates) Size() (n int) {
	var l int
	_ = l
	if m.AuthKeyId != 0 {
		n += 1 + sovPush(uint64(m.AuthKeyId))
	}
	if m.Updates != nil {
		l = m.Updates.Size()
		n += 1 + l + sovPush(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPush(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPush(x uint64) (n int) {
	return sovPush(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PushUpdatesIfNot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPush
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
			return fmt.Errorf("proto: PushUpdatesIfNot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushUpdatesIfNot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
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
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPush
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
				m.Excludes = append(m.Excludes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPush
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
					return ErrInvalidLengthPush
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPush
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
					m.Excludes = append(m.Excludes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Excludes", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
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
				return ErrInvalidLengthPush
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
			skippy, err := skipPush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPush
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
func (m *PushUpdates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPush
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
			return fmt.Errorf("proto: PushUpdates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushUpdates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthKeyId", wireType)
			}
			m.AuthKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
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
				return ErrInvalidLengthPush
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
			skippy, err := skipPush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPush
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
func skipPush(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPush
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
					return 0, ErrIntOverflowPush
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
					return 0, ErrIntOverflowPush
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
				return 0, ErrInvalidLengthPush
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPush
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
				next, err := skipPush(dAtA[start:])
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
	ErrInvalidLengthPush = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPush   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("push.proto", fileDescriptor_push_d8af02f530495561) }

var fileDescriptor_push_d8af02f530495561 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0xc9, 0x14, 0x3b, 0x9a, 0x6e, 0x86, 0x22, 0x58, 0xba, 0x28, 0x65, 0x56, 0x45, 0x98,
	0x04, 0xf4, 0x00, 0x82, 0xbb, 0x22, 0x88, 0x14, 0x5c, 0xe8, 0x66, 0x48, 0xd3, 0x37, 0xcd, 0x60,
	0xdb, 0x84, 0x26, 0x01, 0x7b, 0x43, 0x97, 0x1e, 0x41, 0x7b, 0x02, 0x8f, 0x20, 0x4d, 0x1d, 0x99,
	0xdd, 0xec, 0xde, 0xf7, 0x92, 0xf7, 0xfd, 0xf0, 0x63, 0xac, 0xac, 0x16, 0x44, 0xf5, 0xd2, 0xc8,
	0xd0, 0x9f, 0x66, 0x55, 0xc6, 0x9b, 0x7a, 0x6f, 0x84, 0x2d, 0x09, 0x97, 0x2d, 0xad, 0x65, 0x2d,
	0xa9, 0x7b, 0x2e, 0xed, 0xce, 0x91, 0x03, 0x37, 0xcd, 0x67, 0xf1, 0xa5, 0xe6, 0x02, 0x5a, 0x46,
	0x4c, 0x43, 0xf4, 0xd0, 0xf1, 0x79, 0xbb, 0xd6, 0x78, 0xf5, 0x64, 0xb5, 0x78, 0x56, 0x15, 0x33,
	0xa0, 0xf3, 0xdd, 0xa3, 0x34, 0xe1, 0x15, 0x5e, 0x5a, 0x0d, 0xfd, 0x76, 0x5f, 0x45, 0x28, 0x45,
	0xd9, 0x59, 0xe1, 0x4f, 0x98, 0x57, 0x61, 0x8c, 0xcf, 0xe1, 0x9d, 0x37, 0xb6, 0x02, 0x1d, 0x2d,
	0x52, 0x2f, 0xf3, 0x8a, 0x7f, 0x0e, 0xaf, 0xf1, 0xd2, 0xce, 0x92, 0xc8, 0x4b, 0x51, 0x16, 0xdc,
	0xac, 0x48, 0x6b, 0x5c, 0x06, 0xf9, 0x93, 0x17, 0x87, 0x0f, 0xeb, 0x17, 0x1c, 0x1c, 0x85, 0x86,
	0x09, 0x0e, 0x98, 0x35, 0x62, 0xfb, 0x06, 0xc3, 0x21, 0xd3, 0x2b, 0x2e, 0xa6, 0xd5, 0x03, 0x0c,
	0x79, 0x75, 0xac, 0x5e, 0x9c, 0x50, 0xdf, 0xdf, 0xfd, 0x7c, 0x27, 0xe8, 0x63, 0x4c, 0xd0, 0xe7,
	0x98, 0xa0, 0xaf, 0x31, 0x41, 0xaf, 0x9b, 0x0e, 0x4a, 0xdb, 0x30, 0xc2, 0x05, 0x33, 0x94, 0x37,
	0x56, 0x1b, 0xe8, 0x29, 0x53, 0x8a, 0xb6, 0xa0, 0x35, 0x74, 0x35, 0xf4, 0x74, 0x2a, 0x94, 0xce,
	0xad, 0x96, 0xbe, 0x13, 0xdf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x85, 0x0a, 0x61, 0x20, 0x72,
	0x01, 0x00, 0x00,
}