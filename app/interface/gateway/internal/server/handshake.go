package server

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

const (
	SHA_DIGEST_LENGTH = 20
)

var (
	pq      = string([]byte{0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81})
	p       = []byte{0x49, 0x4C, 0x55, 0x3B}
	q       = []byte{0x53, 0x91, 0x10, 0x73}
	dh2048P = []byte{
		0xc7, 0x1c, 0xae, 0xb9, 0xc6, 0xb1, 0xc9, 0x04, 0x8e, 0x6c, 0x52, 0x2f,
		0x70, 0xf1, 0x3f, 0x73, 0x98, 0x0d, 0x40, 0x23, 0x8e, 0x3e, 0x21, 0xc1,
		0x49, 0x34, 0xd0, 0x37, 0x56, 0x3d, 0x93, 0x0f, 0x48, 0x19, 0x8a, 0x0a,
		0xa7, 0xc1, 0x40, 0x58, 0x22, 0x94, 0x93, 0xd2, 0x25, 0x30, 0xf4, 0xdb,
		0xfa, 0x33, 0x6f, 0x6e, 0x0a, 0xc9, 0x25, 0x13, 0x95, 0x43, 0xae, 0xd4,
		0x4c, 0xce, 0x7c, 0x37, 0x20, 0xfd, 0x51, 0xf6, 0x94, 0x58, 0x70, 0x5a,
		0xc6, 0x8c, 0xd4, 0xfe, 0x6b, 0x6b, 0x13, 0xab, 0xdc, 0x97, 0x46, 0x51,
		0x29, 0x69, 0x32, 0x84, 0x54, 0xf1, 0x8f, 0xaf, 0x8c, 0x59, 0x5f, 0x64,
		0x24, 0x77, 0xfe, 0x96, 0xbb, 0x2a, 0x94, 0x1d, 0x5b, 0xcd, 0x1d, 0x4a,
		0xc8, 0xcc, 0x49, 0x88, 0x07, 0x08, 0xfa, 0x9b, 0x37, 0x8e, 0x3c, 0x4f,
		0x3a, 0x90, 0x60, 0xbe, 0xe6, 0x7c, 0xf9, 0xa4, 0xa4, 0xa6, 0x95, 0x81,
		0x10, 0x51, 0x90, 0x7e, 0x16, 0x27, 0x53, 0xb5, 0x6b, 0x0f, 0x6b, 0x41,
		0x0d, 0xba, 0x74, 0xd8, 0xa8, 0x4b, 0x2a, 0x14, 0xb3, 0x14, 0x4e, 0x0e,
		0xf1, 0x28, 0x47, 0x54, 0xfd, 0x17, 0xed, 0x95, 0x0d, 0x59, 0x65, 0xb4,
		0xb9, 0xdd, 0x46, 0x58, 0x2d, 0xb1, 0x17, 0x8d, 0x16, 0x9c, 0x6b, 0xc4,
		0x65, 0xb0, 0xd6, 0xff, 0x9c, 0xa3, 0x92, 0x8f, 0xef, 0x5b, 0x9a, 0xe4,
		0xe4, 0x18, 0xfc, 0x15, 0xe8, 0x3e, 0xbe, 0xa0, 0xf8, 0x7f, 0xa9, 0xff,
		0x5e, 0xed, 0x70, 0x05, 0x0d, 0xed, 0x28, 0x49, 0xf4, 0x7b, 0xf9, 0x59,
		0xd9, 0x56, 0x85, 0x0c, 0xe9, 0x29, 0x85, 0x1f, 0x0d, 0x81, 0x15, 0xf6,
		0x35, 0xb1, 0x05, 0xee, 0x2e, 0x4e, 0x15, 0xd0, 0x4b, 0x24, 0x54, 0xbf,
		0x6f, 0x4f, 0xad, 0xf0, 0x34, 0xb1, 0x04, 0x03, 0x11, 0x9c, 0xd8, 0xe3,
		0xb9, 0x2f, 0xcc, 0x5b,
	}

	dh2048G = []byte{0x03}
)

var (
	gBigIntDH2048P *big.Int
	gBigIntDH2048G *big.Int
)

func init() {
	gBigIntDH2048P = new(big.Int).SetBytes(dh2048P)
	gBigIntDH2048G = new(big.Int).SetBytes(dh2048G)
}

type keyCreatedF func(ctx context.Context, keyInfo *authsessionpb.AuthKeyInfo, salt *mtproto.FutureSalt) error

type handshake struct {
	rsa            *crypto.RSACryptor
	keyFingerprint uint64
	dh2048p        []byte
	dh2048g        []byte
	keyCreatedF
}

func newHandshake(keyFile string, keyFingerprint uint64, createdCB keyCreatedF) (*handshake, error) {
	rsa, err := crypto.NewRSACryptor(keyFile)
	if err != nil {
		return nil, err
	}
	return &handshake{
		rsa:            rsa,
		keyFingerprint: keyFingerprint,
		dh2048p:        dh2048P,
		dh2048g:        dh2048G,
		keyCreatedF:    createdCB,
	}, nil
}
func (s *handshake) onReqPq(request *mtproto.TLReqPq) (*mtproto.TLResPQ, error) {
	log.Debugf("req_pq#60469778 - {\"request\":%s", request.DebugString())
	if request.GetNonce() == nil || len(request.GetNonce()) != 16 {
		err := fmt.Errorf("onReqPq - invalid nonce: %v", request)
		log.Error(err.Error())
		return nil, err
	}

	resPQ := &mtproto.TLResPQ{Data2: &mtproto.ResPQ{
		Nonce:                       request.Nonce,
		ServerNonce:                 crypto.GenerateNonce(16),
		Pq:                          pq,
		ServerPublicKeyFingerprints: []int64{int64(s.keyFingerprint)},
	}}
	log.Debugf("req_pq#60469778 reply - {\"resPQ\":%s", resPQ.DebugString())
	return resPQ, nil
}

func (s *handshake) onReqPqMulti(request *mtproto.TLReqPqMulti) (*mtproto.TLResPQ, error) {
	log.Debugf("req_pq_multi#be7e8ef1 - {\"request\":%s", request.DebugString())
	if request.GetNonce() == nil || len(request.GetNonce()) != 16 {
		err := fmt.Errorf("onReqPq - invalid nonce: %v", request)
		log.Error(err.Error())
		return nil, err
	}

	resPQ := &mtproto.TLResPQ{Data2: &mtproto.ResPQ{
		Nonce:                       request.Nonce,
		ServerNonce:                 crypto.GenerateNonce(16),
		Pq:                          pq,
		ServerPublicKeyFingerprints: []int64{int64(s.keyFingerprint)},
	}}
	log.Debugf("req_pq_multi#be7e8ef1 - reply: %s", resPQ.DebugString())
	return resPQ, nil
}

func (s *handshake) onReqDHParams(ctx *HandshakeStateCtx, request *mtproto.TLReq_DHParams) (*mtproto.Server_DH_Params, error) {
	log.Debugf("req_DH_params#d712e4be - state: {%s}, request: %s", ctx.DebugString(), request.DebugString())

	var (
		err error
	)
	if !bytes.Equal(request.Nonce, ctx.Nonce) {
		err = fmt.Errorf("onReq_DHParams - Invalid Nonce, req: %s, back: %s",
			util.HexDump(request.Nonce),
			util.HexDump(ctx.Nonce))
		log.Error(err.Error())
		return nil, err
	}

	if !bytes.Equal(request.ServerNonce, ctx.ServerNonce) {
		err = fmt.Errorf("onReq_DHParams - Wrong ServerNonce, req: %s, back: %s",
			util.HexDump(request.ServerNonce),
			util.HexDump(ctx.ServerNonce))
		log.Error(err.Error())
		return nil, err
	}

	if !bytes.Equal([]byte(request.P), p) {
		err = fmt.Errorf("onReq_DHParams - Invalid p valuee")
		log.Error(err.Error())
		return nil, err
	}

	if !bytes.Equal([]byte(request.Q), q) {
		err = fmt.Errorf("onReq_DHParams - Invalid q value")
		log.Error(err.Error())
		return nil, err
	}

	if request.PublicKeyFingerprint != int64(s.keyFingerprint) {
		err = fmt.Errorf("onReq_DHParams - Invalid PublicKeyFingerprint value")
		log.Error(err.Error())
		return nil, err
	}

	encryptedPQInnerData := s.rsa.Decrypt([]byte(request.EncryptedData))
	if len(encryptedPQInnerData) != 255 {
		log.Error("need len(encryptedPQInnerData) = 255")
		return nil, fmt.Errorf("process Req_DHParams - len(encryptedPQInnerData) != 255")
	}

	if !checkSha1(encryptedPQInnerData, 255-SHA_DIGEST_LENGTH) {
		log.Error("process Req_DHParams - sha1Check error")
		return nil, fmt.Errorf("process Req_DHParams - sha1Check error")
	}

	dbuf := mtproto.NewDecodeBuf(encryptedPQInnerData[SHA_DIGEST_LENGTH:])
	o := dbuf.Object()
	if dbuf.GetError() != nil {
		err = fmt.Errorf("onReq_DHParams - decode P_Q_inner_data error")
		log.Error(err.Error())
		return nil, err
	}

	var pqInnerData *mtproto.P_QInnerData
	switch innerData := o.(type) {
	case *mtproto.TLPQInnerData:
		ctx.handshakeType = model.AuthKeyTypePerm
		pqInnerData = innerData.To_P_QInnerData()
	case *mtproto.TLPQInnerDataDc:
		ctx.handshakeType = model.AuthKeyTypePerm
		pqInnerData = innerData.To_P_QInnerData()
	case *mtproto.TLPQInnerDataTemp:
		ctx.handshakeType = model.AuthKeyTypeTemp
		ctx.ExpiresIn = innerData.GetExpiresIn()
		pqInnerData = innerData.To_P_QInnerData()
	case *mtproto.TLPQInnerDataTempDc:
		if innerData.GetDc() < 0 {
			ctx.handshakeType = model.AuthKeyTypeMediaTemp
		} else {
			ctx.handshakeType = model.AuthKeyTypeTemp
		}
		ctx.ExpiresIn = innerData.GetExpiresIn()
		pqInnerData = innerData.To_P_QInnerData()
	default:
		err = fmt.Errorf("onReq_DHParams - decode P_Q_inner_data error")
		log.Error(err.Error())
		return nil, err
	}

	if !bytes.Equal([]byte(pqInnerData.GetPq()), []byte(pq)) {
		log.Error("process Req_DHParams - Invalid p_q_inner_data.pq value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid p_q_inner_data.pq value")
	}

	if !bytes.Equal([]byte(pqInnerData.GetP()), p) {
		log.Error("process Req_DHParams - Invalid p_q_inner_data.p value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid p_q_inner_data.p value")
	}

	if !bytes.Equal([]byte(pqInnerData.GetQ()), q) {
		log.Error("process Req_DHParams - Invalid p_q_inner_data.q value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid p_q_inner_data.q value")
	}

	if !bytes.Equal(pqInnerData.GetNonce(), ctx.Nonce) {
		log.Error("process Req_DHParams - Invalid Nonce")
		return nil, fmt.Errorf("process Req_DHParams - InvalidNonce")
	}

	if !bytes.Equal(pqInnerData.GetServerNonce(), ctx.ServerNonce) {
		log.Error("process Req_DHParams - Wrong ServerNonce")
		return nil, fmt.Errorf("process Req_DHParams - Wrong ServerNonce")
	}
	ctx.NewNonce = pqInnerData.GetNewNonce()
	A := crypto.GenerateNonce(256)
	ctx.A = A
	ctx.P = s.dh2048p

	bigIntA := new(big.Int).SetBytes(A)

	gA := new(big.Int).Exp(gBigIntDH2048G, bigIntA, gBigIntDH2048P)

	serverDHInnerData := &mtproto.TLServer_DHInnerData{Data2: &mtproto.Server_DHInnerData{
		Nonce:       ctx.Nonce,
		ServerNonce: ctx.ServerNonce,
		G:           int32(s.dh2048g[0]),
		GA:          string(gA.Bytes()),
		DhPrime:     string(s.dh2048p),
		ServerTime:  int32(time.Now().Unix()),
	}}

	serverDHInnerDataBuf := serverDHInnerData.Encode(0)

	tmpAesKeyAndIV := make([]byte, 64)
	sha1A := sha1.Sum(append(ctx.NewNonce, ctx.ServerNonce...))
	sha1B := sha1.Sum(append(ctx.ServerNonce, ctx.NewNonce...))
	sha1C := sha1.Sum(append(ctx.NewNonce, ctx.NewNonce...))
	copy(tmpAesKeyAndIV, sha1A[:])
	copy(tmpAesKeyAndIV[20:], sha1B[:])
	copy(tmpAesKeyAndIV[40:], sha1C[:])
	copy(tmpAesKeyAndIV[60:], ctx.NewNonce[:4])

	tmpLen := 20 + len(serverDHInnerDataBuf)
	if tmpLen%16 > 0 {
		tmpLen = (tmpLen/16 + 1) * 16
	} else {
		tmpLen = 20 + len(serverDHInnerDataBuf)
	}

	tmpEncryptedAnswer := make([]byte, tmpLen)
	sha1Tmp := sha1.Sum(serverDHInnerDataBuf)
	copy(tmpEncryptedAnswer, sha1Tmp[:])
	copy(tmpEncryptedAnswer[20:], serverDHInnerDataBuf)

	e := crypto.NewAES256IGECryptor(tmpAesKeyAndIV[:32], tmpAesKeyAndIV[32:64])
	tmpEncryptedAnswer, _ = e.Encrypt(tmpEncryptedAnswer)

	serverDHParamsOk := &mtproto.TLServer_DHParamsOk{Data2: &mtproto.Server_DH_Params{
		Nonce:           ctx.Nonce,
		ServerNonce:     ctx.ServerNonce,
		EncryptedAnswer: hack.String(tmpEncryptedAnswer),
	}}

	log.Debugf("onReq_DHParams - state: {%s}, reply: %s", ctx.DebugString(), serverDHParamsOk.DebugString())

	return serverDHParamsOk.To_Server_DH_Params(), nil
}

func (s *handshake) onSetClientDHParams(ctx *HandshakeStateCtx, request *mtproto.TLSetClient_DHParams) (*mtproto.SetClient_DHParamsAnswer, error) {
	log.Debugf("set_client_DH_params#f5045f1f - state: {%s}, request: %s", ctx.DebugString(), request.DebugString())
	if !bytes.Equal(request.Nonce, ctx.Nonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong Nonce")
		log.Error(err.Error())
		return nil, err
	}

	if !bytes.Equal(request.ServerNonce, ctx.ServerNonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong ServerNonce")
		log.Error(err.Error())
		return nil, err
	}

	bEncryptedData := []byte(request.EncryptedData)

	tmpAesKeyAndIv := make([]byte, 64)
	sha1A := sha1.Sum(append(ctx.NewNonce, ctx.ServerNonce...))
	sha1B := sha1.Sum(append(ctx.ServerNonce, ctx.NewNonce...))
	sha1C := sha1.Sum(append(ctx.NewNonce, ctx.NewNonce...))
	copy(tmpAesKeyAndIv, sha1A[:])
	copy(tmpAesKeyAndIv[20:], sha1B[:])
	copy(tmpAesKeyAndIv[40:], sha1C[:])
	copy(tmpAesKeyAndIv[60:], ctx.NewNonce[:4])

	d := crypto.NewAES256IGECryptor(tmpAesKeyAndIv[:32], tmpAesKeyAndIv[32:64])
	decryptedData, err := d.Decrypt(bEncryptedData)
	if err != nil {
		err := fmt.Errorf("process SetClient_DHParams - AES256IGECryptor descrypt error")
		log.Error(err.Error())
		return nil, err
	}
	dBuf := mtproto.NewDecodeBuf(decryptedData[20:])
	clientDHInnerData := mtproto.MakeTLClient_DHInnerData(nil)
	clientDHInnerData.Data2.Constructor = mtproto.TLConstructor(dBuf.Int())
	err = clientDHInnerData.Decode(dBuf)
	if err != nil {
		log.Errorf("processSetClient_DHParams - TLClient_DHInnerData decode error: %s", err)
		return nil, err
	}

	log.Debugf("processSetClient_DHParams - client_DHInnerData: %#v", clientDHInnerData.String())
	if !bytes.Equal(clientDHInnerData.GetNonce(), ctx.Nonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong client_DHInnerData's Nonce")
		log.Error(err.Error())
		return nil, err
	}

	if !bytes.Equal(clientDHInnerData.GetServerNonce(), ctx.ServerNonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong client_DHInnerData's ServerNonce")
		log.Error(err.Error())
		return nil, err
	}

	bigIntA := new(big.Int).SetBytes(ctx.A)
	authKeyNum := new(big.Int)
	authKeyNum.Exp(new(big.Int).SetBytes([]byte(clientDHInnerData.GetGB())), bigIntA, gBigIntDH2048P)

	authKey := make([]byte, 256)

	copy(authKey[256-len(authKeyNum.Bytes()):], authKeyNum.Bytes())

	authKeyAuxHash := make([]byte, len(ctx.NewNonce))
	copy(authKeyAuxHash, ctx.NewNonce)
	authKeyAuxHash = append(authKeyAuxHash, byte(0x01))
	sha1D := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1D[:]...)
	sha1E := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1E[:]...)

	authKeyId := int64(binary.LittleEndian.Uint64(authKeyAuxHash[len(ctx.NewNonce)+1+12 : len(ctx.NewNonce)+1+12+8]))
	if s.saveAuthKeyInfo(ctx, model.NewAuthKeyInfo(authKeyId, authKey, ctx.handshakeType)) {
		dhGenOk := &mtproto.TLDhGenOk{Data2: &mtproto.SetClient_DHParamsAnswer{
			Nonce:         ctx.Nonce,
			ServerNonce:   ctx.ServerNonce,
			NewNonceHash1: calcNewNonceHash(ctx.NewNonce, authKey, 0x01),
		}}

		log.Debugf("onSetClient_DHParams - ctx: {%s}, reply: %s", ctx.DebugString(), dhGenOk.DebugString())
		return dhGenOk.To_SetClient_DHParamsAnswer(), nil
	} else {
		dhGenRetry := &mtproto.TLDhGenRetry{Data2: &mtproto.SetClient_DHParamsAnswer{
			Nonce:         ctx.Nonce,
			ServerNonce:   ctx.ServerNonce,
			NewNonceHash2: calcNewNonceHash(ctx.NewNonce, authKey, 0x02),
		}}

		log.Debugf("onSetClient_DHParams - ctx: {%v}, reply: %s", ctx.DebugString(), dhGenRetry.DebugString())
		return dhGenRetry.To_SetClient_DHParamsAnswer(), nil
	}
}

func (s *handshake) onMsgsAck(state *HandshakeStateCtx, request *mtproto.TLMsgsAck) error {
	log.Debugf("msgs_ack#62d6b459 - state: {%s}, request: %s", state.DebugString(), request.DebugString())

	switch state.State {
	case STATE_pq_res:
		state.State = STATE_pq_ack
	case STATE_DH_params_res:
		state.State = STATE_DH_params_ack
	case STATE_dh_gen_res:
		state.State = STATE_dh_gen_ack
	default:
		return fmt.Errorf("invalid state: %v", state)
	}

	return nil
}

func (s *handshake) saveAuthKeyInfo(ctx *HandshakeStateCtx, key *model.AuthKeyData) bool {
	var (
		salt       = int64(0)
		serverSalt *mtproto.TLFutureSalt
		now        = int32(time.Now().Unix())
	)

	for a := 7; a >= 0; a-- {
		salt <<= 8
		salt |= int64(ctx.NewNonce[a] ^ ctx.ServerNonce[a])
	}

	serverSalt = &mtproto.TLFutureSalt{Data2: &mtproto.FutureSalt{
		ValidSince: now,
		ValidUntil: now + 30*60,
		Salt:       salt,
	}}

	keyInfo := &authsessionpb.AuthKeyInfo{
		AuthKeyId:          key.AuthKeyId,
		AuthKey:            key.AuthKey,
		AuthKeyType:        int32(key.AuthKeyType),
		PermAuthKeyId:      key.PermAuthKeyId,
		TempAuthKeyId:      key.TempAuthKeyId,
		MediaTempAuthKeyId: key.MediaTempAuthKeyId,
	}

	err := s.keyCreatedF(context.Background(), keyInfo, serverSalt.To_FutureSalt())
	if err != nil {
		log.Errorf("saveAuthKeyInfo not successful - auth_key_id:%d, err:%v", key.AuthKeyId, err)
		return false
	}
	return true
}

func calcNewNonceHash(newNonce, authKey []byte, b byte) []byte {
	authKeyAuxHash := make([]byte, len(newNonce))
	copy(authKeyAuxHash, newNonce)
	authKeyAuxHash = append(authKeyAuxHash, b)
	sha1D := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1D[:]...)
	sha1E := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1E[:]...)
	return authKeyAuxHash[len(authKeyAuxHash)-16:]
}

func checkSha1(data []byte, maxPaddingLen int) bool {
	var (
		dataLen  = len(data)
		sha1Data = data[:SHA_DIGEST_LENGTH]
	)

	for paddingLen := 0; paddingLen < maxPaddingLen; paddingLen++ {
		sha1Check := sha1.Sum(data[SHA_DIGEST_LENGTH : dataLen-paddingLen])
		if bytes.Equal(sha1Check[:], sha1Data) {
			return true
		}
	}
	return false
}
