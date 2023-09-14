package mtproto

const (
	MTPROTO_VERSION = 2
)

type TLObject interface {
	Encode(layer int32) []byte
	Decode(dBuf *DecodeBuf) error
	String() string
	//DebugString() string
}
