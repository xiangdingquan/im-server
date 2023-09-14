package mtproto

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"open.chat/pkg/crypto"
	"open.chat/pkg/log"
)

type PaddedIntermediateCodec struct {
	conn io.ReadWriteCloser
}

func NewMTProtoPaddedIntermediateCodec(conn io.ReadWriteCloser) *PaddedIntermediateCodec {
	return &PaddedIntermediateCodec{
		conn: conn,
	}
}

func (c *PaddedIntermediateCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b))
	log.Debugf("size1: %d", size)

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
		log.Infof("ReadFull2: %s", hex.EncodeToString(buf[:256]))
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

func (c *PaddedIntermediateCodec) Send(msg interface{}) error {
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	b := message.Encode()

	sb := make([]byte, 4)
	size := len(b)

	binary.LittleEndian.PutUint32(sb, uint32(size))
	b = append(sb, b...)
	b = append(b, crypto.GenerateNonce(int(rand.Uint32()%16))...)

	_, err := c.conn.Write(b)

	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *PaddedIntermediateCodec) Close() error {
	return c.conn.Close()
}
