package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"strconv"
)

type AesCTR128KeySizeError int

func (k AesCTR128KeySizeError) Error() string {
	return "AesCTR128KeySizeError: invalid key size " + strconv.Itoa(int(k))
}

type AesCTR128Encrypt struct {
	stream cipher.Stream
}

func NewAesCTR128Encrypt(key []byte, iv []byte) (*AesCTR128Encrypt, error) {
	block2, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(iv) != 16 {
		return nil, AesCTR128KeySizeError(len(iv))
	}

	stream2 := cipher.NewCTR(block2, iv)

	return &AesCTR128Encrypt{
		stream: stream2,
	}, nil
}

func (e *AesCTR128Encrypt) Encrypt(plaintext []byte) []byte {
	e.stream.XORKeyStream(plaintext, plaintext)
	return plaintext
}
