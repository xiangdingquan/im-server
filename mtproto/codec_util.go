package mtproto

import (
	"time"

	"open.chat/pkg/crypto"
)

func GenerateMessageId() int64 {
	const nano = 1000 * 1000 * 1000
	unixnano := time.Now().UnixNano()

	messageId := ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
	for {
		if (messageId % 4) != 1 {
			messageId += 1
		} else {
			break
		}
	}

	return messageId
}

func generateMessageKey(msgKey, authKey []byte, incoming bool) (aesKey, aesIV []byte) {
	var x = 0
	if incoming {
		x = 8
	}

	switch MTPROTO_VERSION {
	case 2:
		t_a := make([]byte, 0, 52)
		t_a = append(t_a, msgKey[:16]...)
		t_a = append(t_a, authKey[x:x+36]...)
		sha256_a := crypto.Sha256Digest(t_a)

		t_b := make([]byte, 0, 52)
		t_b = append(t_b, authKey[40+x:40+x+36]...)
		t_b = append(t_b, msgKey[:16]...)
		sha256_b := crypto.Sha256Digest(t_b)

		aesKey = make([]byte, 0, 32)
		aesKey = append(aesKey, sha256_a[:8]...)
		aesKey = append(aesKey, sha256_b[8:8+16]...)
		aesKey = append(aesKey, sha256_a[24:24+8]...)

		aesIV = make([]byte, 0, 32)
		aesIV = append(aesIV, sha256_b[:8]...)
		aesIV = append(aesIV, sha256_a[8:8+16]...)
		aesIV = append(aesIV, sha256_b[24:24+8]...)

	default:
		aesKey = make([]byte, 0, 32)
		aesIV = make([]byte, 0, 32)
		t_a := make([]byte, 0, 48)
		t_b := make([]byte, 0, 48)
		t_c := make([]byte, 0, 48)
		t_d := make([]byte, 0, 48)

		t_a = append(t_a, msgKey...)
		t_a = append(t_a, authKey[x:x+32]...)

		t_b = append(t_b, authKey[32+x:32+x+16]...)
		t_b = append(t_b, msgKey...)
		t_b = append(t_b, authKey[48+x:48+x+16]...)

		t_c = append(t_c, authKey[64+x:64+x+32]...)
		t_c = append(t_c, msgKey...)

		t_d = append(t_d, msgKey...)
		t_d = append(t_d, authKey[96+x:96+x+32]...)

		sha1_a := crypto.Sha1Digest(t_a)
		sha1_b := crypto.Sha1Digest(t_b)
		sha1_c := crypto.Sha1Digest(t_c)
		sha1_d := crypto.Sha1Digest(t_d)

		aesKey = append(aesKey, sha1_a[0:8]...)
		aesKey = append(aesKey, sha1_b[8:8+12]...)
		aesKey = append(aesKey, sha1_c[4:4+12]...)

		aesIV = append(aesIV, sha1_a[8:8+12]...)
		aesIV = append(aesIV, sha1_b[0:8]...)
		aesIV = append(aesIV, sha1_c[16:16+4]...)
		aesIV = append(aesIV, sha1_d[0:8]...)
	}

	return
}
