package core

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"open.chat/app/messenger/biz_server/messages/secretchat/dal/dataobject"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

type SecretChatData struct {
	*dataobject.SecretChatsDO
	*SecretChatCore
}

const (
	stateNone    = int8(0)
	stateCreated = int8(1)
	stateRunning = int8(2)
	stateDiscard = int8(3)
)

func (m *SecretChatCore) CreateNewSecretChatData(ctx context.Context, adminId, participantId int32, adminAuthKeyId int64, randomId int32, gA []byte) (*SecretChatData, error) {
	do := &dataobject.SecretChatsDO{
		AccessHash:     rand.Int63(),
		AdminId:        adminId,
		ParticipantId:  participantId,
		AdminAuthKeyId: adminAuthKeyId,
		RandomId:       randomId,
		GA:             hex.EncodeToString(gA),
		State:          stateCreated,
		Date:           int32(time.Now().Unix()),
	}

	id, _, err := m.SecretChatsDAO.Insert(ctx, do)
	if err != nil {
		return nil, err
	}

	do.Id = int32(id)
	if do.Id < 0 {
	}

	return &SecretChatData{SecretChatsDO: do, SecretChatCore: m}, nil
}

func (m *SecretChatCore) MakeSecretChatData(ctx context.Context, chatId int32, accessHash int64) (*SecretChatData, error) {
	do, err := m.SecretChatsDAO.Select(ctx, chatId)
	if err != nil {
		return nil, err
	}
	if do == nil {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		log.Errorf("invalid chatId %d - %v", chatId, err)
		return nil, err
	}
	return &SecretChatData{SecretChatsDO: do, SecretChatCore: m}, nil
}

func (m *SecretChatData) getGAOrGB(toUserId int32) ([]byte, error) {
	if toUserId == m.AdminId {
		return hex.DecodeString(m.GB)
	} else {
		return hex.DecodeString(m.GA)
	}
}

func (m *SecretChatData) GetSecretChatPeerId(userSelfId int32) (int32, int64) {
	if userSelfId == m.AdminId {
		return m.ParticipantId, m.ParticipantAuthKeyId
	} else {
		return m.AdminId, m.AdminAuthKeyId
	}
}

func (m *SecretChatData) DoAcceptEncryption(ctx context.Context, userId int32, authKeyId int64, gB []byte, keyFingerprint int64) error {
	if m.SecretChatsDO == nil {
		err := fmt.Errorf("invalid state, not loaded secretChat")
		log.Errorf(err.Error())
		return err
	}

	if userId != m.ParticipantId {
		err := fmt.Errorf("invalid userId and authKeyId (%d, %d)", userId, authKeyId)
		log.Errorf(err.Error())
		return err
	}

	if m.State != stateCreated {
		err := fmt.Errorf("invalid state - (chat_id: %d, state: %d)", m.Id, m.State)
		log.Errorf(err.Error())
		return err
	}

	m.ParticipantAuthKeyId = authKeyId
	m.GB = hex.EncodeToString(gB)
	m.KeyFingerprint = keyFingerprint
	m.State = stateRunning

	_, err := m.SecretChatsDAO.UpdateGB(ctx, authKeyId, m.GB, m.KeyFingerprint, m.Id)

	return err
}

func (m *SecretChatData) DoDiscardEncryption(ctx context.Context, userId int32, authKeyId int64) error {
	var err error
	if m.SecretChatsDO == nil {
		err = fmt.Errorf("invalid state, not loaded secretChat")
		log.Errorf(err.Error())
		return err
	}

	if userId == m.AdminId {
		if authKeyId != m.AdminAuthKeyId {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			log.Errorf("invalid input: %v", err)
			return err
		}
	} else if userId == m.ParticipantId {
		if authKeyId != m.ParticipantAuthKeyId {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			log.Errorf("invalid input: %v", err)
			return err
		}
	} else {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		log.Errorf("invalid input: %v", err)
		return err
	}

	if m.State != stateCreated || m.State != stateRunning {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		log.Errorf("invalid state: %v", err)
		return err
	}

	m.State = stateDiscard

	_, err = m.SecretChatsDAO.UpdateState(ctx, m.State, m.Id)

	return err
}

func (m *SecretChatData) ToEncryptedChat(toId int32) (*mtproto.EncryptedChat, error) {
	encryptedChat := &mtproto.EncryptedChat{
		PredicateName:  mtproto.Predicate_encryptedChat,
		Id:             m.Id,
		AccessHash:     m.AccessHash,
		Date:           m.Date,
		AdminId:        m.AdminId,
		ParticipantId:  m.ParticipantId,
		KeyFingerprint: m.KeyFingerprint,
	}

	var err error
	encryptedChat.GAOrB, err = m.getGAOrGB(toId)
	return encryptedChat, err
}

func (m *SecretChatData) ToEncryptedChatWaiting() *mtproto.EncryptedChat {
	return &mtproto.EncryptedChat{
		PredicateName: mtproto.Predicate_encryptedChatWaiting,
		Id:            m.Id,
		AccessHash:    m.AccessHash,
		Date:          m.Date,
		AdminId:       m.AdminId,
		ParticipantId: m.ParticipantId,
	}
}

func (m *SecretChatData) ToEncryptedChatRequested() (*mtproto.EncryptedChat, error) {
	encryptedChat := &mtproto.EncryptedChat{
		PredicateName: mtproto.Predicate_encryptedChatRequested,
		Id:            m.Id,
		AccessHash:    m.AccessHash,
		Date:          m.Date,
		AdminId:       m.AdminId,
		ParticipantId: m.ParticipantId,
	}

	var err error
	encryptedChat.GA, err = hex.DecodeString(m.GA)
	return encryptedChat, err
}

func (m *SecretChatData) ToEncryptedChatDiscarded() *mtproto.EncryptedChat {
	return &mtproto.EncryptedChat{
		PredicateName: mtproto.Predicate_encryptedChatDiscarded,
		Id:            m.Id,
	}
}

func (m *SecretChatData) ToEncryptedChatEmpty() *mtproto.EncryptedChat {
	return &mtproto.EncryptedChat{
		PredicateName: mtproto.Predicate_encryptedChatEmpty,
		Id:            m.Id,
	}
}

func (m *SecretChatData) SendEncryptedMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, msg *mtproto.EncryptedMessage) (int32, error) {
	var err error

	if m.SecretChatsDO == nil || m.State != stateRunning {
		err = fmt.Errorf("invalid state, not loaded secretChat")
		log.Errorf(err.Error())
		return 0, err
	}

	if fromUserId == m.AdminId {
		if fromAuthKeyId != m.AdminAuthKeyId {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			log.Errorf("invalid input: %v", err)
			return 0, err
		}
	} else if fromUserId == m.ParticipantId {
		if fromAuthKeyId != m.ParticipantAuthKeyId {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			log.Errorf("invalid input: %v", err)
			return 0, err
		}
	} else {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		log.Errorf("invalid input: %v", err)
		return 0, err
	}

	var msgType int8 = 0

	if msg.Constructor != mtproto.CRC32_encryptedMessage {
		msgType = 1
	}

	peerId, peerAuthKeyId := m.GetSecretChatPeerId(fromUserId)
	qts := idgen.NextQtsId(ctx, peerAuthKeyId)

	msgData, _ := json.Marshal(msg)
	if err != nil {
		log.Errorf("json marsh error - %v", err)
		return 0, err
	}

	// 1.
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		msgDO := &dataobject.SecretChatMessagesDO{
			SenderUserId: fromUserId,
			ChatId:       m.Id,
			RandomId:     msg.RandomId,
			PeerId:       peerId,
			MessageType:  msgType,
			MessageData:  hack.String(msgData),
			Date2:        msg.Date,
		}
		msgDO.Id, _, err = m.SecretChatMessagesDAO.InsertTx(tx, msgDO)
		if err != nil {
			result.Err = err
			log.Errorf("sendEncryptedMessage error - %v", err)
			return
		}

		qtsUpdateDO := &dataobject.SecretChatQtsUpdatesDO{
			UserId:        peerId,
			AuthKeyId:     peerAuthKeyId,
			ChatId:        m.Id,
			Qts:           int32(qts),
			ChatMessageId: msgDO.Id,
			Date2:         msgDO.Date2,
		}
		_, _, err = m.SecretChatQtsUpdatesDAO.InsertTx(tx, qtsUpdateDO)
		if err != nil {
			log.Errorf("sendEncryptedMessage error - %v", err)
		}
	})

	return int32(qts), tR.Err
}
