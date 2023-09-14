package mtproto

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"

	"open.chat/pkg/log"
)

type IntermediateCodec struct {
	conn io.ReadWriteCloser
}

func NewMTProtoIntermediateCodec(conn io.ReadWriteCloser) *IntermediateCodec {
	return &IntermediateCodec{
		conn: conn,
	}
}

func (c *IntermediateCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size2 := binary.LittleEndian.Uint32(b)

	needAck := size2>>31 == 1
	_ = needAck

	size = int(size2 & 0xffffff)

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
		log.Errorf("Server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := NewMTPRawMessage(authKeyId, 0, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *IntermediateCodec) Send(msg interface{}) error {
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	b := message.Encode()
	size := len(b)

	sb := make([]byte, 4)
	binary.LittleEndian.PutUint32(sb, uint32(size))

	b = append(sb, b...)
	_, err := c.conn.Write(b)

	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *IntermediateCodec) Close() error {
	return c.conn.Close()
}
