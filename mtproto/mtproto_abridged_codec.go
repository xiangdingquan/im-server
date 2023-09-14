package mtproto

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"

	"open.chat/pkg/log"
)

type AbridgedCodec struct {
	conn io.ReadWriteCloser
}

func NewMTProtoAbridgedCodec(conn io.ReadWriteCloser) *AbridgedCodec {
	return &AbridgedCodec{
		conn: conn,
	}
}

func (c *AbridgedCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 1)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	needAck := b[0]>>7 == 1
	_ = needAck

	b[0] = b[0] & 0x7f

	if b[0] < 0x7f {
		size = int(b[0]) << 2
		if size == 0 {
			return nil, nil
		}
	} else {
		b2 := make([]byte, 3)
		n, err = io.ReadFull(c.conn, b2)
		if err != nil {
			return nil, err
		}
		size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2
	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			log.Errorf("readFull2 error: %v", err)
			return nil, err
		}
		left -= n
	}
	if size > 4096 {
		log.Debugf("readFull2: %s", hex.EncodeToString(buf[:256]))
	}

	if size == 4 {
		log.Errorf("server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := NewMTPRawMessage(authKeyId, 0, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *AbridgedCodec) Send(msg interface{}) error {
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	b := message.Encode()

	sb := make([]byte, 4)
	size := len(b) / 4

	if size < 127 {
		sb = []byte{byte(size)}
	} else {
		binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
	}

	b = append(sb, b...)
	_, err := c.conn.Write(b)

	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *AbridgedCodec) Close() error {
	return c.conn.Close()
}
