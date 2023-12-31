package databus

import (
	encoding_json "encoding/json"
	"fmt"
	"io"
	"math"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
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

type Header struct {
	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata" json:"metadata" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptorDatabus, []int{0} }

func (m *Header) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type MessagePB struct {
	Key       string                   `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	Value     encoding_json.RawMessage `protobuf:"bytes,2,opt,name=value,proto3,casttype=encoding/json.RawMessage" json:"value"`
	Topic     string                   `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic"`
	Partition int32                    `protobuf:"varint,4,opt,name=partition,proto3" json:"partition"`
	Offset    int64                    `protobuf:"varint,5,opt,name=offset,proto3" json:"offset"`
	Timestamp int64                    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp"`
	Metadata  map[string]string        `protobuf:"bytes,7,rep,name=metadata" json:"metadata" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MessagePB) Reset()                    { *m = MessagePB{} }
func (m *MessagePB) String() string            { return proto.CompactTextString(m) }
func (*MessagePB) ProtoMessage()               {}
func (*MessagePB) Descriptor() ([]byte, []int) { return fileDescriptorDatabus, []int{1} }

func (m *MessagePB) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MessagePB) GetValue() encoding_json.RawMessage {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MessagePB) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *MessagePB) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *MessagePB) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *MessagePB) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessagePB) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "infra.databus.Header")
	proto.RegisterType((*MessagePB)(nil), "infra.databus.MessagePB")
}
func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			dAtA[i] = 0xa
			i++
			v := m.Metadata[k]
			mapSize := 1 + len(k) + sovDatabus(uint64(len(k))) + 1 + len(v) + sovDatabus(uint64(len(v)))
			i = encodeVarintDatabus(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDatabus(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDatabus(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *MessagePB) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessagePB) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDatabus(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDatabus(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if len(m.Topic) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDatabus(dAtA, i, uint64(len(m.Topic)))
		i += copy(dAtA[i:], m.Topic)
	}
	if m.Partition != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDatabus(dAtA, i, uint64(m.Partition))
	}
	if m.Offset != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDatabus(dAtA, i, uint64(m.Offset))
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintDatabus(dAtA, i, uint64(m.Timestamp))
	}
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			dAtA[i] = 0x3a
			i++
			v := m.Metadata[k]
			mapSize := 1 + len(k) + sovDatabus(uint64(len(k))) + 1 + len(v) + sovDatabus(uint64(len(v)))
			i = encodeVarintDatabus(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDatabus(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDatabus(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintDatabus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Header) Size() (n int) {
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDatabus(uint64(len(k))) + 1 + len(v) + sovDatabus(uint64(len(v)))
			n += mapEntrySize + 1 + sovDatabus(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *MessagePB) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovDatabus(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovDatabus(uint64(l))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovDatabus(uint64(l))
	}
	if m.Partition != 0 {
		n += 1 + sovDatabus(uint64(m.Partition))
	}
	if m.Offset != 0 {
		n += 1 + sovDatabus(uint64(m.Offset))
	}
	if m.Timestamp != 0 {
		n += 1 + sovDatabus(uint64(m.Timestamp))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDatabus(uint64(len(k))) + 1 + len(v) + sovDatabus(uint64(len(v)))
			n += mapEntrySize + 1 + sovDatabus(uint64(mapEntrySize))
		}
	}
	return n
}

func sovDatabus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDatabus(x uint64) (n int) {
	return sovDatabus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatabus
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabus
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
				return ErrInvalidLengthDatabus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDatabus
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDatabus
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDatabus
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDatabus
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthDatabus
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDatabus(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthDatabus
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatabus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatabus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MessagePB) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatabus
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
			return fmt.Errorf("proto: MessagePB: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessagePB: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabus
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
				return ErrInvalidLengthDatabus
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabus
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
				return ErrInvalidLengthDatabus
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabus
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
				return ErrInvalidLengthDatabus
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partition", wireType)
			}
			m.Partition = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Partition |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabus
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
				return ErrInvalidLengthDatabus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDatabus
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDatabus
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDatabus
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDatabus
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthDatabus
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDatabus(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthDatabus
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatabus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatabus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDatabus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDatabus
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
					return 0, ErrIntOverflowDatabus
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
					return 0, ErrIntOverflowDatabus
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
				return 0, ErrInvalidLengthDatabus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDatabus
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
				next, err := skipDatabus(dAtA[start:])
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
	ErrInvalidLengthDatabus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDatabus   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("library/queue/databus/databus.proto", fileDescriptorDatabus) }

var fileDescriptorDatabus = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x41, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0x99, 0xc6, 0xa4, 0x66, 0x6c, 0x41, 0x06, 0x85, 0xb1, 0x8b, 0x26, 0xb4, 0x28, 0x01,
	0x31, 0x01, 0xdd, 0x88, 0xdd, 0x05, 0x04, 0x37, 0x85, 0x92, 0xa5, 0xbb, 0x49, 0x3b, 0x89, 0xa3,
	0x4d, 0x26, 0x26, 0x13, 0xa5, 0xc7, 0x10, 0x3c, 0x94, 0x4b, 0x4f, 0x10, 0xa4, 0xcb, 0x1c, 0xc1,
	0x95, 0x64, 0x12, 0xd3, 0xd6, 0x03, 0xb8, 0xfa, 0xdf, 0xff, 0x78, 0xff, 0x97, 0x90, 0x17, 0x38,
	0x5e, 0x32, 0x3f, 0x25, 0xe9, 0xca, 0x79, 0xce, 0x69, 0x4e, 0x9d, 0x05, 0x11, 0xc4, 0xcf, 0xb3,
	0xdf, 0x69, 0x27, 0x29, 0x17, 0x1c, 0xf5, 0x59, 0x1c, 0xa4, 0xc4, 0x6e, 0xcc, 0xc1, 0x45, 0xc8,
	0xc4, 0x43, 0xee, 0xdb, 0x73, 0x1e, 0x39, 0x21, 0x0f, 0xb9, 0x23, 0x53, 0x7e, 0x1e, 0xc8, 0x4d,
	0x2e, 0x52, 0xd5, 0xd7, 0xa3, 0x77, 0x00, 0xb5, 0x3b, 0x4a, 0x16, 0x34, 0x45, 0x53, 0xb8, 0x1f,
	0x51, 0x41, 0x2a, 0x10, 0x06, 0xa6, 0x62, 0x1d, 0x5c, 0x8e, 0xed, 0x1d, 0xb6, 0x5d, 0x07, 0xed,
	0x69, 0x93, 0xba, 0x8d, 0x45, 0xba, 0x72, 0x7b, 0x65, 0x61, 0xb4, 0x87, 0x5e, 0xab, 0x06, 0x13,
	0xd8, 0xdf, 0x09, 0xa2, 0x43, 0xa8, 0x3c, 0xd1, 0x15, 0x06, 0x26, 0xb0, 0x74, 0xaf, 0x92, 0xe8,
	0x08, 0xaa, 0x2f, 0x64, 0x99, 0x53, 0xdc, 0x91, 0x5e, 0xbd, 0xdc, 0x74, 0xae, 0xc1, 0xe8, 0x4d,
	0x81, 0xfa, 0x94, 0x66, 0x19, 0x09, 0xe9, 0xcc, 0x45, 0x27, 0x5b, 0x97, 0x6e, 0xb7, 0x2c, 0x8c,
	0x6a, 0xad, 0x11, 0x93, 0x6d, 0x44, 0xcf, 0x3d, 0x2d, 0x0b, 0xa3, 0x36, 0xbe, 0x0b, 0x03, 0xd3,
	0x78, 0xce, 0x17, 0x2c, 0x0e, 0x9d, 0xc7, 0x8c, 0xc7, 0xb6, 0x47, 0x5e, 0x1b, 0x64, 0xf3, 0x24,
	0x64, 0x40, 0x55, 0xf0, 0x84, 0xcd, 0xb1, 0x22, 0xc9, 0x7a, 0x75, 0x2c, 0x0d, 0xaf, 0x1e, 0xe8,
	0x1c, 0xea, 0x09, 0x49, 0x05, 0x13, 0x8c, 0xc7, 0x78, 0xcf, 0x04, 0x96, 0xea, 0xf6, 0xcb, 0xc2,
	0xd8, 0x98, 0xde, 0x46, 0xa2, 0x11, 0xd4, 0x78, 0x10, 0x64, 0x54, 0x60, 0xd5, 0x04, 0x96, 0xe2,
	0xc2, 0xb2, 0x30, 0x1a, 0xc7, 0x6b, 0x66, 0x05, 0x14, 0x2c, 0xa2, 0x99, 0x20, 0x51, 0x82, 0x35,
	0x19, 0x93, 0xc0, 0xd6, 0xf4, 0x36, 0x12, 0xcd, 0xb6, 0x0a, 0xe9, 0xca, 0x42, 0xce, 0xfe, 0x14,
	0xd2, 0x7e, 0xa2, 0x7f, 0xe8, 0xc4, 0x3d, 0xfe, 0x58, 0x0f, 0xc1, 0xe7, 0x7a, 0x08, 0xbe, 0xd6,
	0x43, 0x70, 0xdf, 0x6d, 0xde, 0xc1, 0xd7, 0xe4, 0x8f, 0x74, 0xf5, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xa5, 0x22, 0x22, 0x61, 0xad, 0x02, 0x00, 0x00,
}
