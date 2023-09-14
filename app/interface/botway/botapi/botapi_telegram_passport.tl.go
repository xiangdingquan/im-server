package botapi

import (
	"encoding/json"
)

type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data,omitempty"`
	PhoneNumber string          `json:"phone_number,omitempty"`
	Email       string          `json:"email,omitempty"`
	Files       []*PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile   `json:"front_side,omitempty"`
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"`
	Selfie      *PassportFile   `json:"selfie,omitempty"`
	Translation []*PassportFile `json:"translation,omitempty"`
	Hash        string          `json:"hash"`
}

func (m *EncryptedPassportElement) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *EncryptedPassportElement) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

func (m *EncryptedCredentials) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *EncryptedCredentials) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type PassportElementError struct {
	Source      string   `json:"source"`
	Type        string   `json:"type"`
	FieldName   string   `json:"field_name"`
	DataHash    string   `json:"data_hash"`
	Message     string   `json:"message"`
	FileHash    string   `json:"file_hash"`
	FileHashes  []string `json:"file_hashes"`
	ElementHash string   `json:"element_hash"`
}

func (m *PassportElementError) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *PassportElementError) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials []*EncryptedCredentials     `json:"credentials"`
}

func (m *PassportData) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *PassportData) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int32  `json:"file_size"`
	FileDate     int32  `json:"file_date"`
}

func (m *PassportFile) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *PassportFile) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}
