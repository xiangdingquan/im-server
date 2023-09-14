package mtproto

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"

	"open.chat/pkg/crypto"
	"open.chat/pkg/net2"
)

const (
	TRANSPORT_TCP  = 1
	TRANSPORT_HTTP = 2
	TRANSPORT_UDP  = 3
)
const (
	ABRIDGED_FLAG            = 0xef
	ABRIDGED_INT32_FLAG      = 0xcdab7856
	INTERMEDIATE_FLAG        = 0xfedcba98
	PADDED_INTERMEDIATE_FLAG = 0x12345678
	UNKNOWN_FLAG             = 0x02010316
	PVRG_FLAG                = 0x47725650
	FULL_FLAG                = 0x00000000

	HTTP_HEAD_FLAG   = 0x44414548
	HTTP_POST_FLAG   = 0x54534f50
	HTTP_GET_FLAG    = 0x20544547
	HTTP_OPTION_FLAG = 0x4954504f
)

var (
	isClientType bool
)

func init() {
	net2.RegisterProtocol("mtproto", NewMTProtoTransport())
	flag.BoolVar(&isClientType, "client", false, "client conn")
}

type MTProtoTransport struct {
}

func NewMTProtoTransport() *MTProtoTransport {
	return &MTProtoTransport{}
}

func (m *MTProtoTransport) NewCodec(rw io.ReadWriter) (net2.Codec, error) {
	codec := &TransportCodec{
		codecType: TRANSPORT_TCP,
		conn:      rw.(net.Conn),
		proto:     m,
	}
	return codec, nil
}

type TransportCodec struct {
	codecType int
	conn      net.Conn
	codec     net2.Codec
	proto     *MTProtoTransport
}

func (c *TransportCodec) peekCodec() error {
	peek, _ := c.conn.(net2.PeekAble)
	firstByte, err := peek.PeekByte()
	if err != nil {
		return err
	}
	if firstByte == ABRIDGED_FLAG {
		return errors.New("")
	}
	firstInt, err := peek.PeekUint32()
	if err != nil {
		return err
	}
	if firstInt == HTTP_HEAD_FLAG ||
		firstInt == HTTP_POST_FLAG ||
		firstInt == HTTP_GET_FLAG ||
		firstInt == HTTP_OPTION_FLAG {
		return errors.New("")
	}
	if firstInt == INTERMEDIATE_FLAG {
		return errors.New("")
	}
	if firstInt == PADDED_INTERMEDIATE_FLAG {
		return errors.New("")
	}
	if firstInt == PVRG_FLAG {
		return errors.New("")
	}
	if firstInt == UNKNOWN_FLAG {
		return errors.New("")
	}
	checkFullBuf, err := peek.Peek(12)
	if err != nil {
		return err
	}
	secondInt := binary.BigEndian.Uint32(checkFullBuf[:8])
	if secondInt == FULL_FLAG {
		return errors.New("")
	}
	obfuscatedBuf, err := peek.Peek(68)
	if err != nil {
		return err
	}
	var tmp [48]byte
	for i := 0; i < 48; i++ {
		tmp[i] = obfuscatedBuf[56-i]
	}
	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		return err
	}
	d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[9:41], obfuscatedBuf[41:57])
	if err != nil {
		return err
	}
	d.Encrypt(obfuscatedBuf)
	protocolType := binary.LittleEndian.Uint32(obfuscatedBuf[57:])
	if protocolType != ABRIDGED_INT32_FLAG &&
		protocolType != INTERMEDIATE_FLAG &&
		protocolType != PADDED_INTERMEDIATE_FLAG {
		return errors.New("")
	}
	dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[61:65]))
	c.codec = NewMTProtoObfuscatedCodec(c.conn, d, e, protocolType, dcId)
	peek.Discard(68)
	return nil
}

func (c *TransportCodec) Receive() (interface{}, error) {
	if isClientType {
		if c.codec == nil {
			return nil, fmt.Errorf("")
		}
	} else {
		if c.codec == nil {
			err := c.peekCodec()
			if err != nil {
				return nil, err
			}
		}
	}
	return c.codec.Receive()
}

func (c *TransportCodec) Send(msg interface{}) error {
	if isClientType {
	} else {
		if c.codec != nil {
			return c.codec.Send(msg)
		}
	}
	return fmt.Errorf("")
}

func (c *TransportCodec) Close() error {
	if c.codec != nil {
		return c.codec.Close()
	} else {
		return nil
	}
}
