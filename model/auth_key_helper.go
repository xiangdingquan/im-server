package model

import "fmt"

const (
	AuthKeyTypeUnknown   = -1
	AuthKeyTypePerm      = 0
	AuthKeyTypeTemp      = 1
	AuthKeyTypeMediaTemp = 2
)

type BoundAuthKeyIdInfo struct {
	PermAuthKeyId      int64
	TempAuthKeyId      int64
	MediaTempAuthKeyId int64
}

type AuthKeyData struct {
	AuthKeyType        int    `json:"auth_key_type,omitempty"`
	AuthKeyId          int64  `json:"auth_key_id,omitempty"`
	AuthKey            []byte `json:"auth_key,omitempty"`
	PermAuthKeyId      int64  `json:"perm_auth_key_id,omitempty"`
	TempAuthKeyId      int64  `json:"temp_auth_key_id,omitempty"`
	MediaTempAuthKeyId int64  `json:"media_temp_auth_key_id,omitempty"`
}

func NewAuthKeyInfo(keyId int64, key []byte, keyType int) *AuthKeyData {
	keyData := &AuthKeyData{
		AuthKeyId:          keyId,
		AuthKey:            key,
		AuthKeyType:        keyType,
		PermAuthKeyId:      0,
		TempAuthKeyId:      0,
		MediaTempAuthKeyId: 0,
	}

	switch keyType {
	case AuthKeyTypePerm:
		keyData.PermAuthKeyId = keyId
	case AuthKeyTypeTemp:
		keyData.TempAuthKeyId = keyId
	case AuthKeyTypeMediaTemp:
		keyData.MediaTempAuthKeyId = keyId
	}

	return keyData
}

func (m *AuthKeyData) DebugString() string {
	return fmt.Sprintf(`{"auth_key_type":%d,"auth_key_id":%d,"perm_auth_key_id"=%d,"temp_auth_key_id"=%d,"media_temp_auth_key_id"=%d}`,
		m.AuthKeyType,
		m.AuthKeyId,
		m.PermAuthKeyId,
		m.TempAuthKeyId,
		m.MediaTempAuthKeyId)
}

// Impl cache.Value interface
func (m *AuthKeyData) Size() int {
	return 1
}
