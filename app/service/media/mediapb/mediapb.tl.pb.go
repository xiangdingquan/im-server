package mediapb

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"open.chat/mtproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TLConstructor int32

const (
	CRC32_UNKNOWN                    TLConstructor = 0
	CRC32_photoDataRsp               TLConstructor = -471254727
	CRC32_documentId                 TLConstructor = 337093742
	CRC32_documentList               TLConstructor = -401118321
	CRC32_fileLocationSecret         TLConstructor = -1080234184
	CRC32_nbfs_uploadPhotoFile       TLConstructor = -1422618728
	CRC32_nbfs_uploadVideoFile       TLConstructor = -39771092
	CRC32_nbfs_getPhotoFileData      TLConstructor = -2063586110
	CRC32_nbfs_uploadedPhotoMedia    TLConstructor = -322886609
	CRC32_nbfs_uploadedDocumentMedia TLConstructor = -562683210
	CRC32_nbfs_getDocument           TLConstructor = -608596954
	CRC32_nbfs_getDocumentList       TLConstructor = 1710556034
	CRC32_nbfs_uploadEncryptedFile   TLConstructor = -1836531759
	CRC32_nbfs_getEncryptedFile      TLConstructor = 2089484089
	CRC32_nbfs_getFileLocationSecret TLConstructor = 514373981
)

var TLConstructor_name = map[int32]string{
	0:           "CRC32_UNKNOWN",
	-471254727:  "CRC32_photoDataRsp",
	337093742:   "CRC32_documentId",
	-401118321:  "CRC32_documentList",
	-1080234184: "CRC32_fileLocationSecret",
	-1422618728: "CRC32_nbfs_uploadPhotoFile",
	-39771092:   "CRC32_nbfs_uploadVideoFile",
	-2063586110: "CRC32_nbfs_getPhotoFileData",
	-322886609:  "CRC32_nbfs_uploadedPhotoMedia",
	-562683210:  "CRC32_nbfs_uploadedDocumentMedia",
	-608596954:  "CRC32_nbfs_getDocument",
	1710556034:  "CRC32_nbfs_getDocumentList",
	-1836531759: "CRC32_nbfs_uploadEncryptedFile",
	2089484089:  "CRC32_nbfs_getEncryptedFile",
	514373981:   "CRC32_nbfs_getFileLocationSecret",
}

var TLConstructor_value = map[string]int32{
	"CRC32_UNKNOWN":                    0,
	"CRC32_photoDataRsp":               -471254727,
	"CRC32_documentId":                 337093742,
	"CRC32_documentList":               -401118321,
	"CRC32_fileLocationSecret":         -1080234184,
	"CRC32_nbfs_uploadPhotoFile":       -1422618728,
	"CRC32_nbfs_uploadVideoFile":       -39771092,
	"CRC32_nbfs_getPhotoFileData":      -2063586110,
	"CRC32_nbfs_uploadedPhotoMedia":    -322886609,
	"CRC32_nbfs_uploadedDocumentMedia": -562683210,
	"CRC32_nbfs_getDocument":           -608596954,
	"CRC32_nbfs_getDocumentList":       1710556034,
	"CRC32_nbfs_uploadEncryptedFile":   -1836531759,
	"CRC32_nbfs_getEncryptedFile":      2089484089,
	"CRC32_nbfs_getFileLocationSecret": 514373981,
}

func (x TLConstructor) String() string {
	return proto.EnumName(TLConstructor_name, int32(x))
}

func (TLConstructor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{0}
}

// --------------------------------------------------------------------------------------------
// documentId id:long access_hash:long version:int = DocumentId;
//
// DocumentId <--
//   - TL_documentId
type DocumentId struct {
	PredicateName        string        `protobuf:"bytes,1,opt,name=predicate_name,json=predicateName,proto3" json:"predicate_name,omitempty"`
	Constructor          TLConstructor `protobuf:"varint,2,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	Id                   int64         `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	AccessHash           int64         `protobuf:"varint,4,opt,name=access_hash,json=accessHash,proto3" json:"access_hash,omitempty"`
	Version              int32         `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DocumentId) Reset()         { *m = DocumentId{} }
func (m *DocumentId) String() string { return proto.CompactTextString(m) }
func (*DocumentId) ProtoMessage()    {}
func (*DocumentId) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{0}
}
func (m *DocumentId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DocumentId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DocumentId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DocumentId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentId.Merge(dst, src)
}
func (m *DocumentId) XXX_Size() int {
	return m.Size()
}
func (m *DocumentId) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentId.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentId proto.InternalMessageInfo

func (m *DocumentId) GetPredicateName() string {
	if m != nil {
		return m.PredicateName
	}
	return ""
}

func (m *DocumentId) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *DocumentId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DocumentId) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

func (m *DocumentId) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// documentId id:long access_hash:long version:int = DocumentId;
type TLDocumentId struct {
	Data2                *DocumentId `protobuf:"bytes,1,opt,name=data2,proto3" json:"data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TLDocumentId) Reset()         { *m = TLDocumentId{} }
func (m *TLDocumentId) String() string { return proto.CompactTextString(m) }
func (*TLDocumentId) ProtoMessage()    {}
func (*TLDocumentId) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{1}
}
func (m *TLDocumentId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLDocumentId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLDocumentId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLDocumentId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLDocumentId.Merge(dst, src)
}
func (m *TLDocumentId) XXX_Size() int {
	return m.Size()
}
func (m *TLDocumentId) XXX_DiscardUnknown() {
	xxx_messageInfo_TLDocumentId.DiscardUnknown(m)
}

var xxx_messageInfo_TLDocumentId proto.InternalMessageInfo

func (m *TLDocumentId) GetData2() *DocumentId {
	if m != nil {
		return m.Data2
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// documentList documents:Vector<Document> = DocumentList;
//
// DocumentList <--
//   - TL_documentList
type DocumentList struct {
	PredicateName        string              `protobuf:"bytes,1,opt,name=predicate_name,json=predicateName,proto3" json:"predicate_name,omitempty"`
	Constructor          TLConstructor       `protobuf:"varint,2,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	Documents            []*mtproto.Document `protobuf:"bytes,3,rep,name=documents,proto3" json:"documents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DocumentList) Reset()         { *m = DocumentList{} }
func (m *DocumentList) String() string { return proto.CompactTextString(m) }
func (*DocumentList) ProtoMessage()    {}
func (*DocumentList) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{2}
}
func (m *DocumentList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DocumentList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DocumentList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DocumentList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentList.Merge(dst, src)
}
func (m *DocumentList) XXX_Size() int {
	return m.Size()
}
func (m *DocumentList) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentList.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentList proto.InternalMessageInfo

func (m *DocumentList) GetPredicateName() string {
	if m != nil {
		return m.PredicateName
	}
	return ""
}

func (m *DocumentList) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *DocumentList) GetDocuments() []*mtproto.Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

// documentList documents:Vector<Document> = DocumentList;
type TLDocumentList struct {
	Data2                *DocumentList `protobuf:"bytes,1,opt,name=data2,proto3" json:"data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLDocumentList) Reset()         { *m = TLDocumentList{} }
func (m *TLDocumentList) String() string { return proto.CompactTextString(m) }
func (*TLDocumentList) ProtoMessage()    {}
func (*TLDocumentList) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{3}
}
func (m *TLDocumentList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLDocumentList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLDocumentList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLDocumentList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLDocumentList.Merge(dst, src)
}
func (m *TLDocumentList) XXX_Size() int {
	return m.Size()
}
func (m *TLDocumentList) XXX_DiscardUnknown() {
	xxx_messageInfo_TLDocumentList.DiscardUnknown(m)
}

var xxx_messageInfo_TLDocumentList proto.InternalMessageInfo

func (m *TLDocumentList) GetData2() *DocumentList {
	if m != nil {
		return m.Data2
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// fileLocationSecret secret:long = FileLocationSecret;
//
// FileLocationSecret <--
//   - TL_fileLocationSecret
type FileLocationSecret struct {
	PredicateName        string        `protobuf:"bytes,1,opt,name=predicate_name,json=predicateName,proto3" json:"predicate_name,omitempty"`
	Constructor          TLConstructor `protobuf:"varint,2,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	Secret               int64         `protobuf:"varint,3,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FileLocationSecret) Reset()         { *m = FileLocationSecret{} }
func (m *FileLocationSecret) String() string { return proto.CompactTextString(m) }
func (*FileLocationSecret) ProtoMessage()    {}
func (*FileLocationSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{4}
}
func (m *FileLocationSecret) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileLocationSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileLocationSecret.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FileLocationSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileLocationSecret.Merge(dst, src)
}
func (m *FileLocationSecret) XXX_Size() int {
	return m.Size()
}
func (m *FileLocationSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_FileLocationSecret.DiscardUnknown(m)
}

var xxx_messageInfo_FileLocationSecret proto.InternalMessageInfo

func (m *FileLocationSecret) GetPredicateName() string {
	if m != nil {
		return m.PredicateName
	}
	return ""
}

func (m *FileLocationSecret) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *FileLocationSecret) GetSecret() int64 {
	if m != nil {
		return m.Secret
	}
	return 0
}

// fileLocationSecret secret:long = FileLocationSecret;
type TLFileLocationSecret struct {
	Data2                *FileLocationSecret `protobuf:"bytes,1,opt,name=data2,proto3" json:"data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TLFileLocationSecret) Reset()         { *m = TLFileLocationSecret{} }
func (m *TLFileLocationSecret) String() string { return proto.CompactTextString(m) }
func (*TLFileLocationSecret) ProtoMessage()    {}
func (*TLFileLocationSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{5}
}
func (m *TLFileLocationSecret) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLFileLocationSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLFileLocationSecret.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLFileLocationSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLFileLocationSecret.Merge(dst, src)
}
func (m *TLFileLocationSecret) XXX_Size() int {
	return m.Size()
}
func (m *TLFileLocationSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TLFileLocationSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TLFileLocationSecret proto.InternalMessageInfo

func (m *TLFileLocationSecret) GetData2() *FileLocationSecret {
	if m != nil {
		return m.Data2
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// photoDataRsp photo_id:long access_hash:long date:int size_list:Vector<PhotoSize> = PhotoDataRsp;
//
// PhotoDataRsp <--
//   - TL_photoDataRsp
type PhotoDataRsp struct {
	PredicateName        string               `protobuf:"bytes,1,opt,name=predicate_name,json=predicateName,proto3" json:"predicate_name,omitempty"`
	Constructor          TLConstructor        `protobuf:"varint,2,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	PhotoId              int64                `protobuf:"varint,3,opt,name=photo_id,json=photoId,proto3" json:"photo_id,omitempty"`
	AccessHash           int64                `protobuf:"varint,4,opt,name=access_hash,json=accessHash,proto3" json:"access_hash,omitempty"`
	Date                 int32                `protobuf:"varint,5,opt,name=date,proto3" json:"date,omitempty"`
	SizeList             []*mtproto.PhotoSize `protobuf:"bytes,6,rep,name=size_list,json=sizeList,proto3" json:"size_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PhotoDataRsp) Reset()         { *m = PhotoDataRsp{} }
func (m *PhotoDataRsp) String() string { return proto.CompactTextString(m) }
func (*PhotoDataRsp) ProtoMessage()    {}
func (*PhotoDataRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{6}
}
func (m *PhotoDataRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PhotoDataRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PhotoDataRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PhotoDataRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhotoDataRsp.Merge(dst, src)
}
func (m *PhotoDataRsp) XXX_Size() int {
	return m.Size()
}
func (m *PhotoDataRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PhotoDataRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PhotoDataRsp proto.InternalMessageInfo

func (m *PhotoDataRsp) GetPredicateName() string {
	if m != nil {
		return m.PredicateName
	}
	return ""
}

func (m *PhotoDataRsp) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *PhotoDataRsp) GetPhotoId() int64 {
	if m != nil {
		return m.PhotoId
	}
	return 0
}

func (m *PhotoDataRsp) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

func (m *PhotoDataRsp) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *PhotoDataRsp) GetSizeList() []*mtproto.PhotoSize {
	if m != nil {
		return m.SizeList
	}
	return nil
}

// photoDataRsp photo_id:long access_hash:long date:int size_list:Vector<PhotoSize> = PhotoDataRsp;
type TLPhotoDataRsp struct {
	Data2                *PhotoDataRsp `protobuf:"bytes,1,opt,name=data2,proto3" json:"data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLPhotoDataRsp) Reset()         { *m = TLPhotoDataRsp{} }
func (m *TLPhotoDataRsp) String() string { return proto.CompactTextString(m) }
func (*TLPhotoDataRsp) ProtoMessage()    {}
func (*TLPhotoDataRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{7}
}
func (m *TLPhotoDataRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLPhotoDataRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLPhotoDataRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLPhotoDataRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLPhotoDataRsp.Merge(dst, src)
}
func (m *TLPhotoDataRsp) XXX_Size() int {
	return m.Size()
}
func (m *TLPhotoDataRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_TLPhotoDataRsp.DiscardUnknown(m)
}

var xxx_messageInfo_TLPhotoDataRsp proto.InternalMessageInfo

func (m *TLPhotoDataRsp) GetData2() *PhotoDataRsp {
	if m != nil {
		return m.Data2
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// nbfs.uploadPhotoFile owner_id:long file:InputFile is_profile:bool = PhotoDataRsp;
type TLNbfsUploadPhotoFile struct {
	Constructor          TLConstructor      `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	OwnerId              int64              `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	File                 *mtproto.InputFile `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	IsProfile            bool               `protobuf:"varint,4,opt,name=is_profile,json=isProfile,proto3" json:"is_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TLNbfsUploadPhotoFile) Reset()         { *m = TLNbfsUploadPhotoFile{} }
func (m *TLNbfsUploadPhotoFile) String() string { return proto.CompactTextString(m) }
func (*TLNbfsUploadPhotoFile) ProtoMessage()    {}
func (*TLNbfsUploadPhotoFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{8}
}
func (m *TLNbfsUploadPhotoFile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsUploadPhotoFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsUploadPhotoFile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsUploadPhotoFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsUploadPhotoFile.Merge(dst, src)
}
func (m *TLNbfsUploadPhotoFile) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsUploadPhotoFile) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsUploadPhotoFile.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsUploadPhotoFile proto.InternalMessageInfo

func (m *TLNbfsUploadPhotoFile) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsUploadPhotoFile) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *TLNbfsUploadPhotoFile) GetFile() *mtproto.InputFile {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *TLNbfsUploadPhotoFile) GetIsProfile() bool {
	if m != nil {
		return m.IsProfile
	}
	return false
}

// --------------------------------------------------------------------------------------------
// nbfs.uploadVideoFile owner_id:long file:InputFile = Document;
type TLNbfsUploadVideoFile struct {
	Constructor          TLConstructor      `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	OwnerId              int64              `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	File                 *mtproto.InputFile `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TLNbfsUploadVideoFile) Reset()         { *m = TLNbfsUploadVideoFile{} }
func (m *TLNbfsUploadVideoFile) String() string { return proto.CompactTextString(m) }
func (*TLNbfsUploadVideoFile) ProtoMessage()    {}
func (*TLNbfsUploadVideoFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{9}
}
func (m *TLNbfsUploadVideoFile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsUploadVideoFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsUploadVideoFile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsUploadVideoFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsUploadVideoFile.Merge(dst, src)
}
func (m *TLNbfsUploadVideoFile) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsUploadVideoFile) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsUploadVideoFile.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsUploadVideoFile proto.InternalMessageInfo

func (m *TLNbfsUploadVideoFile) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsUploadVideoFile) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *TLNbfsUploadVideoFile) GetFile() *mtproto.InputFile {
	if m != nil {
		return m.File
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// nbfs.getPhotoFileData photo_id:long = PhotoDataRsp;
type TLNbfsGetPhotoFileData struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	PhotoId              int64         `protobuf:"varint,2,opt,name=photo_id,json=photoId,proto3" json:"photo_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLNbfsGetPhotoFileData) Reset()         { *m = TLNbfsGetPhotoFileData{} }
func (m *TLNbfsGetPhotoFileData) String() string { return proto.CompactTextString(m) }
func (*TLNbfsGetPhotoFileData) ProtoMessage()    {}
func (*TLNbfsGetPhotoFileData) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{10}
}
func (m *TLNbfsGetPhotoFileData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsGetPhotoFileData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsGetPhotoFileData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsGetPhotoFileData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsGetPhotoFileData.Merge(dst, src)
}
func (m *TLNbfsGetPhotoFileData) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsGetPhotoFileData) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsGetPhotoFileData.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsGetPhotoFileData proto.InternalMessageInfo

func (m *TLNbfsGetPhotoFileData) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsGetPhotoFileData) GetPhotoId() int64 {
	if m != nil {
		return m.PhotoId
	}
	return 0
}

// --------------------------------------------------------------------------------------------
// nbfs.uploadedPhotoMedia owner_id:long media:InputMedia is_profile:bool = MessageMedia;
type TLNbfsUploadedPhotoMedia struct {
	Constructor          TLConstructor       `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	OwnerId              int64               `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Media                *mtproto.InputMedia `protobuf:"bytes,3,opt,name=media,proto3" json:"media,omitempty"`
	IsProfile            bool                `protobuf:"varint,4,opt,name=is_profile,json=isProfile,proto3" json:"is_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TLNbfsUploadedPhotoMedia) Reset()         { *m = TLNbfsUploadedPhotoMedia{} }
func (m *TLNbfsUploadedPhotoMedia) String() string { return proto.CompactTextString(m) }
func (*TLNbfsUploadedPhotoMedia) ProtoMessage()    {}
func (*TLNbfsUploadedPhotoMedia) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{11}
}
func (m *TLNbfsUploadedPhotoMedia) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsUploadedPhotoMedia) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsUploadedPhotoMedia.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsUploadedPhotoMedia) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsUploadedPhotoMedia.Merge(dst, src)
}
func (m *TLNbfsUploadedPhotoMedia) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsUploadedPhotoMedia) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsUploadedPhotoMedia.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsUploadedPhotoMedia proto.InternalMessageInfo

func (m *TLNbfsUploadedPhotoMedia) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsUploadedPhotoMedia) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *TLNbfsUploadedPhotoMedia) GetMedia() *mtproto.InputMedia {
	if m != nil {
		return m.Media
	}
	return nil
}

func (m *TLNbfsUploadedPhotoMedia) GetIsProfile() bool {
	if m != nil {
		return m.IsProfile
	}
	return false
}

// --------------------------------------------------------------------------------------------
// nbfs.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
type TLNbfsUploadedDocumentMedia struct {
	Constructor          TLConstructor       `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	OwnerId              int64               `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Media                *mtproto.InputMedia `protobuf:"bytes,3,opt,name=media,proto3" json:"media,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TLNbfsUploadedDocumentMedia) Reset()         { *m = TLNbfsUploadedDocumentMedia{} }
func (m *TLNbfsUploadedDocumentMedia) String() string { return proto.CompactTextString(m) }
func (*TLNbfsUploadedDocumentMedia) ProtoMessage()    {}
func (*TLNbfsUploadedDocumentMedia) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{12}
}
func (m *TLNbfsUploadedDocumentMedia) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsUploadedDocumentMedia) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsUploadedDocumentMedia.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsUploadedDocumentMedia) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsUploadedDocumentMedia.Merge(dst, src)
}
func (m *TLNbfsUploadedDocumentMedia) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsUploadedDocumentMedia) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsUploadedDocumentMedia.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsUploadedDocumentMedia proto.InternalMessageInfo

func (m *TLNbfsUploadedDocumentMedia) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsUploadedDocumentMedia) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *TLNbfsUploadedDocumentMedia) GetMedia() *mtproto.InputMedia {
	if m != nil {
		return m.Media
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// nbfs.getDocument document_id:DocumentId = Document;
type TLNbfsGetDocument struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	DocumentId           *DocumentId   `protobuf:"bytes,2,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLNbfsGetDocument) Reset()         { *m = TLNbfsGetDocument{} }
func (m *TLNbfsGetDocument) String() string { return proto.CompactTextString(m) }
func (*TLNbfsGetDocument) ProtoMessage()    {}
func (*TLNbfsGetDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{13}
}
func (m *TLNbfsGetDocument) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsGetDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsGetDocument.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsGetDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsGetDocument.Merge(dst, src)
}
func (m *TLNbfsGetDocument) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsGetDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsGetDocument.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsGetDocument proto.InternalMessageInfo

func (m *TLNbfsGetDocument) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsGetDocument) GetDocumentId() *DocumentId {
	if m != nil {
		return m.DocumentId
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// nbfs.getDocumentList id_list:Vector<DocumentId> = DocumentList;
type TLNbfsGetDocumentList struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	IdList               []*DocumentId `protobuf:"bytes,2,rep,name=id_list,json=idList,proto3" json:"id_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLNbfsGetDocumentList) Reset()         { *m = TLNbfsGetDocumentList{} }
func (m *TLNbfsGetDocumentList) String() string { return proto.CompactTextString(m) }
func (*TLNbfsGetDocumentList) ProtoMessage()    {}
func (*TLNbfsGetDocumentList) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{14}
}
func (m *TLNbfsGetDocumentList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsGetDocumentList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsGetDocumentList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsGetDocumentList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsGetDocumentList.Merge(dst, src)
}
func (m *TLNbfsGetDocumentList) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsGetDocumentList) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsGetDocumentList.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsGetDocumentList proto.InternalMessageInfo

func (m *TLNbfsGetDocumentList) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsGetDocumentList) GetIdList() []*DocumentId {
	if m != nil {
		return m.IdList
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// nbfs.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
type TLNbfsUploadEncryptedFile struct {
	Constructor          TLConstructor               `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	OwnerId              int64                       `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	File                 *mtproto.InputEncryptedFile `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TLNbfsUploadEncryptedFile) Reset()         { *m = TLNbfsUploadEncryptedFile{} }
func (m *TLNbfsUploadEncryptedFile) String() string { return proto.CompactTextString(m) }
func (*TLNbfsUploadEncryptedFile) ProtoMessage()    {}
func (*TLNbfsUploadEncryptedFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{15}
}
func (m *TLNbfsUploadEncryptedFile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsUploadEncryptedFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsUploadEncryptedFile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsUploadEncryptedFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsUploadEncryptedFile.Merge(dst, src)
}
func (m *TLNbfsUploadEncryptedFile) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsUploadEncryptedFile) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsUploadEncryptedFile.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsUploadEncryptedFile proto.InternalMessageInfo

func (m *TLNbfsUploadEncryptedFile) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsUploadEncryptedFile) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *TLNbfsUploadEncryptedFile) GetFile() *mtproto.InputEncryptedFile {
	if m != nil {
		return m.File
	}
	return nil
}

// --------------------------------------------------------------------------------------------
// nbfs.getEncryptedFile id:long access_hash:long = EncryptedFile;
type TLNbfsGetEncryptedFile struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	Id                   int64         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	AccessHash           int64         `protobuf:"varint,3,opt,name=access_hash,json=accessHash,proto3" json:"access_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLNbfsGetEncryptedFile) Reset()         { *m = TLNbfsGetEncryptedFile{} }
func (m *TLNbfsGetEncryptedFile) String() string { return proto.CompactTextString(m) }
func (*TLNbfsGetEncryptedFile) ProtoMessage()    {}
func (*TLNbfsGetEncryptedFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{16}
}
func (m *TLNbfsGetEncryptedFile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsGetEncryptedFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsGetEncryptedFile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsGetEncryptedFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsGetEncryptedFile.Merge(dst, src)
}
func (m *TLNbfsGetEncryptedFile) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsGetEncryptedFile) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsGetEncryptedFile.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsGetEncryptedFile proto.InternalMessageInfo

func (m *TLNbfsGetEncryptedFile) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsGetEncryptedFile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TLNbfsGetEncryptedFile) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

// --------------------------------------------------------------------------------------------
// nbfs.getFileLocationSecret volume_id:long local_id:int = FileLocationSecret;
type TLNbfsGetFileLocationSecret struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=mediapb.TLConstructor" json:"constructor,omitempty"`
	VolumeId             int64         `protobuf:"varint,2,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
	LocalId              int32         `protobuf:"varint,3,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLNbfsGetFileLocationSecret) Reset()         { *m = TLNbfsGetFileLocationSecret{} }
func (m *TLNbfsGetFileLocationSecret) String() string { return proto.CompactTextString(m) }
func (*TLNbfsGetFileLocationSecret) ProtoMessage()    {}
func (*TLNbfsGetFileLocationSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_mediapb_tl_0451dcac72627873, []int{17}
}
func (m *TLNbfsGetFileLocationSecret) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLNbfsGetFileLocationSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLNbfsGetFileLocationSecret.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TLNbfsGetFileLocationSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLNbfsGetFileLocationSecret.Merge(dst, src)
}
func (m *TLNbfsGetFileLocationSecret) XXX_Size() int {
	return m.Size()
}
func (m *TLNbfsGetFileLocationSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TLNbfsGetFileLocationSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TLNbfsGetFileLocationSecret proto.InternalMessageInfo

func (m *TLNbfsGetFileLocationSecret) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLNbfsGetFileLocationSecret) GetVolumeId() int64 {
	if m != nil {
		return m.VolumeId
	}
	return 0
}

func (m *TLNbfsGetFileLocationSecret) GetLocalId() int32 {
	if m != nil {
		return m.LocalId
	}
	return 0
}

func init() {
	proto.RegisterType((*DocumentId)(nil), "mediapb.DocumentId")
	proto.RegisterType((*TLDocumentId)(nil), "mediapb.TL_documentId")
	proto.RegisterType((*DocumentList)(nil), "mediapb.DocumentList")
	proto.RegisterType((*TLDocumentList)(nil), "mediapb.TL_documentList")
	proto.RegisterType((*FileLocationSecret)(nil), "mediapb.FileLocationSecret")
	proto.RegisterType((*TLFileLocationSecret)(nil), "mediapb.TL_fileLocationSecret")
	proto.RegisterType((*PhotoDataRsp)(nil), "mediapb.PhotoDataRsp")
	proto.RegisterType((*TLPhotoDataRsp)(nil), "mediapb.TL_photoDataRsp")
	proto.RegisterType((*TLNbfsUploadPhotoFile)(nil), "mediapb.TL_nbfs_uploadPhotoFile")
	proto.RegisterType((*TLNbfsUploadVideoFile)(nil), "mediapb.TL_nbfs_uploadVideoFile")
	proto.RegisterType((*TLNbfsGetPhotoFileData)(nil), "mediapb.TL_nbfs_getPhotoFileData")
	proto.RegisterType((*TLNbfsUploadedPhotoMedia)(nil), "mediapb.TL_nbfs_uploadedPhotoMedia")
	proto.RegisterType((*TLNbfsUploadedDocumentMedia)(nil), "mediapb.TL_nbfs_uploadedDocumentMedia")
	proto.RegisterType((*TLNbfsGetDocument)(nil), "mediapb.TL_nbfs_getDocument")
	proto.RegisterType((*TLNbfsGetDocumentList)(nil), "mediapb.TL_nbfs_getDocumentList")
	proto.RegisterType((*TLNbfsUploadEncryptedFile)(nil), "mediapb.TL_nbfs_uploadEncryptedFile")
	proto.RegisterType((*TLNbfsGetEncryptedFile)(nil), "mediapb.TL_nbfs_getEncryptedFile")
	proto.RegisterType((*TLNbfsGetFileLocationSecret)(nil), "mediapb.TL_nbfs_getFileLocationSecret")
	proto.RegisterEnum("mediapb.TLConstructor", TLConstructor_name, TLConstructor_value)
}
func (this *DocumentId) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&mediapb.DocumentId{")
	s = append(s, "PredicateName: "+fmt.Sprintf("%#v", this.PredicateName)+",\n")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "AccessHash: "+fmt.Sprintf("%#v", this.AccessHash)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLDocumentId) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&mediapb.TLDocumentId{")
	if this.Data2 != nil {
		s = append(s, "Data2: "+fmt.Sprintf("%#v", this.Data2)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DocumentList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&mediapb.DocumentList{")
	s = append(s, "PredicateName: "+fmt.Sprintf("%#v", this.PredicateName)+",\n")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	if this.Documents != nil {
		s = append(s, "Documents: "+fmt.Sprintf("%#v", this.Documents)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLDocumentList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&mediapb.TLDocumentList{")
	if this.Data2 != nil {
		s = append(s, "Data2: "+fmt.Sprintf("%#v", this.Data2)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *FileLocationSecret) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&mediapb.FileLocationSecret{")
	s = append(s, "PredicateName: "+fmt.Sprintf("%#v", this.PredicateName)+",\n")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "Secret: "+fmt.Sprintf("%#v", this.Secret)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLFileLocationSecret) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&mediapb.TLFileLocationSecret{")
	if this.Data2 != nil {
		s = append(s, "Data2: "+fmt.Sprintf("%#v", this.Data2)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PhotoDataRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&mediapb.PhotoDataRsp{")
	s = append(s, "PredicateName: "+fmt.Sprintf("%#v", this.PredicateName)+",\n")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "PhotoId: "+fmt.Sprintf("%#v", this.PhotoId)+",\n")
	s = append(s, "AccessHash: "+fmt.Sprintf("%#v", this.AccessHash)+",\n")
	s = append(s, "Date: "+fmt.Sprintf("%#v", this.Date)+",\n")
	if this.SizeList != nil {
		s = append(s, "SizeList: "+fmt.Sprintf("%#v", this.SizeList)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLPhotoDataRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&mediapb.TLPhotoDataRsp{")
	if this.Data2 != nil {
		s = append(s, "Data2: "+fmt.Sprintf("%#v", this.Data2)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsUploadPhotoFile) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&mediapb.TLNbfsUploadPhotoFile{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "OwnerId: "+fmt.Sprintf("%#v", this.OwnerId)+",\n")
	if this.File != nil {
		s = append(s, "File: "+fmt.Sprintf("%#v", this.File)+",\n")
	}
	s = append(s, "IsProfile: "+fmt.Sprintf("%#v", this.IsProfile)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsUploadVideoFile) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&mediapb.TLNbfsUploadVideoFile{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "OwnerId: "+fmt.Sprintf("%#v", this.OwnerId)+",\n")
	if this.File != nil {
		s = append(s, "File: "+fmt.Sprintf("%#v", this.File)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsGetPhotoFileData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&mediapb.TLNbfsGetPhotoFileData{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "PhotoId: "+fmt.Sprintf("%#v", this.PhotoId)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsUploadedPhotoMedia) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&mediapb.TLNbfsUploadedPhotoMedia{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "OwnerId: "+fmt.Sprintf("%#v", this.OwnerId)+",\n")
	if this.Media != nil {
		s = append(s, "Media: "+fmt.Sprintf("%#v", this.Media)+",\n")
	}
	s = append(s, "IsProfile: "+fmt.Sprintf("%#v", this.IsProfile)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsUploadedDocumentMedia) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&mediapb.TLNbfsUploadedDocumentMedia{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "OwnerId: "+fmt.Sprintf("%#v", this.OwnerId)+",\n")
	if this.Media != nil {
		s = append(s, "Media: "+fmt.Sprintf("%#v", this.Media)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsGetDocument) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&mediapb.TLNbfsGetDocument{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	if this.DocumentId != nil {
		s = append(s, "DocumentId: "+fmt.Sprintf("%#v", this.DocumentId)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsGetDocumentList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&mediapb.TLNbfsGetDocumentList{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	if this.IdList != nil {
		s = append(s, "IdList: "+fmt.Sprintf("%#v", this.IdList)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsUploadEncryptedFile) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&mediapb.TLNbfsUploadEncryptedFile{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "OwnerId: "+fmt.Sprintf("%#v", this.OwnerId)+",\n")
	if this.File != nil {
		s = append(s, "File: "+fmt.Sprintf("%#v", this.File)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsGetEncryptedFile) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&mediapb.TLNbfsGetEncryptedFile{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "AccessHash: "+fmt.Sprintf("%#v", this.AccessHash)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLNbfsGetFileLocationSecret) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&mediapb.TLNbfsGetFileLocationSecret{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "VolumeId: "+fmt.Sprintf("%#v", this.VolumeId)+",\n")
	s = append(s, "LocalId: "+fmt.Sprintf("%#v", this.LocalId)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMediapbTl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCNbfsClient is the client API for RPCNbfs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCNbfsClient interface {
	// nbfs.uploadPhotoFile owner_id:long file:InputFile is_profile:bool = PhotoDataRsp;
	NbfsUploadPhotoFile(ctx context.Context, in *TLNbfsUploadPhotoFile, opts ...grpc.CallOption) (*PhotoDataRsp, error)
	// nbfs.uploadVideoFile owner_id:long file:InputFile = Document;
	NbfsUploadVideoFile(ctx context.Context, in *TLNbfsUploadVideoFile, opts ...grpc.CallOption) (*mtproto.Document, error)
	// nbfs.getPhotoFileData photo_id:long = PhotoDataRsp;
	NbfsGetPhotoFileData(ctx context.Context, in *TLNbfsGetPhotoFileData, opts ...grpc.CallOption) (*PhotoDataRsp, error)
	// nbfs.uploadedPhotoMedia owner_id:long media:InputMedia is_profile:bool = MessageMedia;
	NbfsUploadedPhotoMedia(ctx context.Context, in *TLNbfsUploadedPhotoMedia, opts ...grpc.CallOption) (*mtproto.MessageMedia, error)
	// nbfs.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
	NbfsUploadedDocumentMedia(ctx context.Context, in *TLNbfsUploadedDocumentMedia, opts ...grpc.CallOption) (*mtproto.MessageMedia, error)
	// nbfs.getDocument document_id:DocumentId = Document;
	NbfsGetDocument(ctx context.Context, in *TLNbfsGetDocument, opts ...grpc.CallOption) (*mtproto.Document, error)
	// nbfs.getDocumentList id_list:Vector<DocumentId> = DocumentList;
	NbfsGetDocumentList(ctx context.Context, in *TLNbfsGetDocumentList, opts ...grpc.CallOption) (*DocumentList, error)
	// nbfs.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
	NbfsUploadEncryptedFile(ctx context.Context, in *TLNbfsUploadEncryptedFile, opts ...grpc.CallOption) (*mtproto.EncryptedFile, error)
	// nbfs.getEncryptedFile id:long access_hash:long = EncryptedFile;
	NbfsGetEncryptedFile(ctx context.Context, in *TLNbfsGetEncryptedFile, opts ...grpc.CallOption) (*mtproto.EncryptedFile, error)
	// nbfs.getFileLocationSecret volume_id:long local_id:int = FileLocationSecret;
	NbfsGetFileLocationSecret(ctx context.Context, in *TLNbfsGetFileLocationSecret, opts ...grpc.CallOption) (*FileLocationSecret, error)
}

type rPCNbfsClient struct {
	cc *grpc.ClientConn
}

func NewRPCNbfsClient(cc *grpc.ClientConn) RPCNbfsClient {
	return &rPCNbfsClient{cc}
}

func (c *rPCNbfsClient) NbfsUploadPhotoFile(ctx context.Context, in *TLNbfsUploadPhotoFile, opts ...grpc.CallOption) (*PhotoDataRsp, error) {
	out := new(PhotoDataRsp)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_uploadPhotoFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsUploadVideoFile(ctx context.Context, in *TLNbfsUploadVideoFile, opts ...grpc.CallOption) (*mtproto.Document, error) {
	out := new(mtproto.Document)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_uploadVideoFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsGetPhotoFileData(ctx context.Context, in *TLNbfsGetPhotoFileData, opts ...grpc.CallOption) (*PhotoDataRsp, error) {
	out := new(PhotoDataRsp)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_getPhotoFileData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsUploadedPhotoMedia(ctx context.Context, in *TLNbfsUploadedPhotoMedia, opts ...grpc.CallOption) (*mtproto.MessageMedia, error) {
	out := new(mtproto.MessageMedia)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_uploadedPhotoMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsUploadedDocumentMedia(ctx context.Context, in *TLNbfsUploadedDocumentMedia, opts ...grpc.CallOption) (*mtproto.MessageMedia, error) {
	out := new(mtproto.MessageMedia)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_uploadedDocumentMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsGetDocument(ctx context.Context, in *TLNbfsGetDocument, opts ...grpc.CallOption) (*mtproto.Document, error) {
	out := new(mtproto.Document)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_getDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsGetDocumentList(ctx context.Context, in *TLNbfsGetDocumentList, opts ...grpc.CallOption) (*DocumentList, error) {
	out := new(DocumentList)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_getDocumentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsUploadEncryptedFile(ctx context.Context, in *TLNbfsUploadEncryptedFile, opts ...grpc.CallOption) (*mtproto.EncryptedFile, error) {
	out := new(mtproto.EncryptedFile)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_uploadEncryptedFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsGetEncryptedFile(ctx context.Context, in *TLNbfsGetEncryptedFile, opts ...grpc.CallOption) (*mtproto.EncryptedFile, error) {
	out := new(mtproto.EncryptedFile)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_getEncryptedFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCNbfsClient) NbfsGetFileLocationSecret(ctx context.Context, in *TLNbfsGetFileLocationSecret, opts ...grpc.CallOption) (*FileLocationSecret, error) {
	out := new(FileLocationSecret)
	err := c.cc.Invoke(ctx, "/mediapb.RPCNbfs/nbfs_getFileLocationSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCNbfsServer is the server API for RPCNbfs service.
type RPCNbfsServer interface {
	// nbfs.uploadPhotoFile owner_id:long file:InputFile is_profile:bool = PhotoDataRsp;
	NbfsUploadPhotoFile(context.Context, *TLNbfsUploadPhotoFile) (*PhotoDataRsp, error)
	// nbfs.uploadVideoFile owner_id:long file:InputFile = Document;
	NbfsUploadVideoFile(context.Context, *TLNbfsUploadVideoFile) (*mtproto.Document, error)
	// nbfs.getPhotoFileData photo_id:long = PhotoDataRsp;
	NbfsGetPhotoFileData(context.Context, *TLNbfsGetPhotoFileData) (*PhotoDataRsp, error)
	// nbfs.uploadedPhotoMedia owner_id:long media:InputMedia is_profile:bool = MessageMedia;
	NbfsUploadedPhotoMedia(context.Context, *TLNbfsUploadedPhotoMedia) (*mtproto.MessageMedia, error)
	// nbfs.uploadedDocumentMedia owner_id:long media:InputMedia = MessageMedia;
	NbfsUploadedDocumentMedia(context.Context, *TLNbfsUploadedDocumentMedia) (*mtproto.MessageMedia, error)
	// nbfs.getDocument document_id:DocumentId = Document;
	NbfsGetDocument(context.Context, *TLNbfsGetDocument) (*mtproto.Document, error)
	// nbfs.getDocumentList id_list:Vector<DocumentId> = DocumentList;
	NbfsGetDocumentList(context.Context, *TLNbfsGetDocumentList) (*DocumentList, error)
	// nbfs.uploadEncryptedFile owner_id:long file:InputEncryptedFile = EncryptedFile;
	NbfsUploadEncryptedFile(context.Context, *TLNbfsUploadEncryptedFile) (*mtproto.EncryptedFile, error)
	// nbfs.getEncryptedFile id:long access_hash:long = EncryptedFile;
	NbfsGetEncryptedFile(context.Context, *TLNbfsGetEncryptedFile) (*mtproto.EncryptedFile, error)
	// nbfs.getFileLocationSecret volume_id:long local_id:int = FileLocationSecret;
	NbfsGetFileLocationSecret(context.Context, *TLNbfsGetFileLocationSecret) (*FileLocationSecret, error)
}

func RegisterRPCNbfsServer(s *grpc.Server, srv RPCNbfsServer) {
	s.RegisterService(&_RPCNbfs_serviceDesc, srv)
}

func _RPCNbfs_NbfsUploadPhotoFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsUploadPhotoFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsUploadPhotoFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsUploadPhotoFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsUploadPhotoFile(ctx, req.(*TLNbfsUploadPhotoFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsUploadVideoFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsUploadVideoFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsUploadVideoFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsUploadVideoFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsUploadVideoFile(ctx, req.(*TLNbfsUploadVideoFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsGetPhotoFileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsGetPhotoFileData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsGetPhotoFileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsGetPhotoFileData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsGetPhotoFileData(ctx, req.(*TLNbfsGetPhotoFileData))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsUploadedPhotoMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsUploadedPhotoMedia)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsUploadedPhotoMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsUploadedPhotoMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsUploadedPhotoMedia(ctx, req.(*TLNbfsUploadedPhotoMedia))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsUploadedDocumentMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsUploadedDocumentMedia)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsUploadedDocumentMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsUploadedDocumentMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsUploadedDocumentMedia(ctx, req.(*TLNbfsUploadedDocumentMedia))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsGetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsGetDocument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsGetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsGetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsGetDocument(ctx, req.(*TLNbfsGetDocument))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsGetDocumentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsGetDocumentList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsGetDocumentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsGetDocumentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsGetDocumentList(ctx, req.(*TLNbfsGetDocumentList))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsUploadEncryptedFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsUploadEncryptedFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsUploadEncryptedFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsUploadEncryptedFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsUploadEncryptedFile(ctx, req.(*TLNbfsUploadEncryptedFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsGetEncryptedFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsGetEncryptedFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsGetEncryptedFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsGetEncryptedFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsGetEncryptedFile(ctx, req.(*TLNbfsGetEncryptedFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCNbfs_NbfsGetFileLocationSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLNbfsGetFileLocationSecret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCNbfsServer).NbfsGetFileLocationSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCNbfs/NbfsGetFileLocationSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCNbfsServer).NbfsGetFileLocationSecret(ctx, req.(*TLNbfsGetFileLocationSecret))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCNbfs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mediapb.RPCNbfs",
	HandlerType: (*RPCNbfsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "nbfs_uploadPhotoFile",
			Handler:    _RPCNbfs_NbfsUploadPhotoFile_Handler,
		},
		{
			MethodName: "nbfs_uploadVideoFile",
			Handler:    _RPCNbfs_NbfsUploadVideoFile_Handler,
		},
		{
			MethodName: "nbfs_getPhotoFileData",
			Handler:    _RPCNbfs_NbfsGetPhotoFileData_Handler,
		},
		{
			MethodName: "nbfs_uploadedPhotoMedia",
			Handler:    _RPCNbfs_NbfsUploadedPhotoMedia_Handler,
		},
		{
			MethodName: "nbfs_uploadedDocumentMedia",
			Handler:    _RPCNbfs_NbfsUploadedDocumentMedia_Handler,
		},
		{
			MethodName: "nbfs_getDocument",
			Handler:    _RPCNbfs_NbfsGetDocument_Handler,
		},
		{
			MethodName: "nbfs_getDocumentList",
			Handler:    _RPCNbfs_NbfsGetDocumentList_Handler,
		},
		{
			MethodName: "nbfs_uploadEncryptedFile",
			Handler:    _RPCNbfs_NbfsUploadEncryptedFile_Handler,
		},
		{
			MethodName: "nbfs_getEncryptedFile",
			Handler:    _RPCNbfs_NbfsGetEncryptedFile_Handler,
		},
		{
			MethodName: "nbfs_getFileLocationSecret",
			Handler:    _RPCNbfs_NbfsGetFileLocationSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mediapb.tl.proto",
}

func (m *DocumentId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DocumentId) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PredicateName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(len(m.PredicateName)))
		i += copy(dAtA[i:], m.PredicateName)
	}
	if m.Constructor != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.Id != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Id))
	}
	if m.AccessHash != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.AccessHash))
	}
	if m.Version != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLDocumentId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLDocumentId) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data2 != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Data2.Size()))
		n1, err := m.Data2.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DocumentList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DocumentList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PredicateName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(len(m.PredicateName)))
		i += copy(dAtA[i:], m.PredicateName)
	}
	if m.Constructor != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if len(m.Documents) > 0 {
		for _, msg := range m.Documents {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintMediapbTl(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLDocumentList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLDocumentList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data2 != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Data2.Size()))
		n2, err := m.Data2.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *FileLocationSecret) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileLocationSecret) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PredicateName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(len(m.PredicateName)))
		i += copy(dAtA[i:], m.PredicateName)
	}
	if m.Constructor != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.Secret != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Secret))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLFileLocationSecret) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLFileLocationSecret) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data2 != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Data2.Size()))
		n3, err := m.Data2.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PhotoDataRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PhotoDataRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PredicateName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(len(m.PredicateName)))
		i += copy(dAtA[i:], m.PredicateName)
	}
	if m.Constructor != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.PhotoId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.PhotoId))
	}
	if m.AccessHash != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.AccessHash))
	}
	if m.Date != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Date))
	}
	if len(m.SizeList) > 0 {
		for _, msg := range m.SizeList {
			dAtA[i] = 0x32
			i++
			i = encodeVarintMediapbTl(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLPhotoDataRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLPhotoDataRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data2 != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Data2.Size()))
		n4, err := m.Data2.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsUploadPhotoFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsUploadPhotoFile) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.OwnerId))
	}
	if m.File != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.File.Size()))
		n5, err := m.File.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.IsProfile {
		dAtA[i] = 0x20
		i++
		if m.IsProfile {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsUploadVideoFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsUploadVideoFile) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.OwnerId))
	}
	if m.File != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.File.Size()))
		n6, err := m.File.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsGetPhotoFileData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsGetPhotoFileData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.PhotoId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.PhotoId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsUploadedPhotoMedia) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsUploadedPhotoMedia) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.OwnerId))
	}
	if m.Media != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Media.Size()))
		n7, err := m.Media.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.IsProfile {
		dAtA[i] = 0x20
		i++
		if m.IsProfile {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsUploadedDocumentMedia) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsUploadedDocumentMedia) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.OwnerId))
	}
	if m.Media != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Media.Size()))
		n8, err := m.Media.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsGetDocument) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsGetDocument) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.DocumentId != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.DocumentId.Size()))
		n9, err := m.DocumentId.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsGetDocumentList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsGetDocumentList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if len(m.IdList) > 0 {
		for _, msg := range m.IdList {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMediapbTl(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsUploadEncryptedFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsUploadEncryptedFile) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.OwnerId))
	}
	if m.File != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.File.Size()))
		n10, err := m.File.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsGetEncryptedFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsGetEncryptedFile) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.Id != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Id))
	}
	if m.AccessHash != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.AccessHash))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLNbfsGetFileLocationSecret) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLNbfsGetFileLocationSecret) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Constructor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.Constructor))
	}
	if m.VolumeId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.VolumeId))
	}
	if m.LocalId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMediapbTl(dAtA, i, uint64(m.LocalId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMediapbTl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DocumentId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PredicateName)
	if l > 0 {
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.Id != 0 {
		n += 1 + sovMediapbTl(uint64(m.Id))
	}
	if m.AccessHash != 0 {
		n += 1 + sovMediapbTl(uint64(m.AccessHash))
	}
	if m.Version != 0 {
		n += 1 + sovMediapbTl(uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLDocumentId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data2 != nil {
		l = m.Data2.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DocumentList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PredicateName)
	if l > 0 {
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if len(m.Documents) > 0 {
		for _, e := range m.Documents {
			l = e.Size()
			n += 1 + l + sovMediapbTl(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLDocumentList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data2 != nil {
		l = m.Data2.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FileLocationSecret) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PredicateName)
	if l > 0 {
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.Secret != 0 {
		n += 1 + sovMediapbTl(uint64(m.Secret))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLFileLocationSecret) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data2 != nil {
		l = m.Data2.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PhotoDataRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PredicateName)
	if l > 0 {
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.PhotoId != 0 {
		n += 1 + sovMediapbTl(uint64(m.PhotoId))
	}
	if m.AccessHash != 0 {
		n += 1 + sovMediapbTl(uint64(m.AccessHash))
	}
	if m.Date != 0 {
		n += 1 + sovMediapbTl(uint64(m.Date))
	}
	if len(m.SizeList) > 0 {
		for _, e := range m.SizeList {
			l = e.Size()
			n += 1 + l + sovMediapbTl(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLPhotoDataRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data2 != nil {
		l = m.Data2.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsUploadPhotoFile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		n += 1 + sovMediapbTl(uint64(m.OwnerId))
	}
	if m.File != nil {
		l = m.File.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.IsProfile {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsUploadVideoFile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		n += 1 + sovMediapbTl(uint64(m.OwnerId))
	}
	if m.File != nil {
		l = m.File.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsGetPhotoFileData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.PhotoId != 0 {
		n += 1 + sovMediapbTl(uint64(m.PhotoId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsUploadedPhotoMedia) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		n += 1 + sovMediapbTl(uint64(m.OwnerId))
	}
	if m.Media != nil {
		l = m.Media.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.IsProfile {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsUploadedDocumentMedia) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		n += 1 + sovMediapbTl(uint64(m.OwnerId))
	}
	if m.Media != nil {
		l = m.Media.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsGetDocument) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.DocumentId != nil {
		l = m.DocumentId.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsGetDocumentList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if len(m.IdList) > 0 {
		for _, e := range m.IdList {
			l = e.Size()
			n += 1 + l + sovMediapbTl(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsUploadEncryptedFile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.OwnerId != 0 {
		n += 1 + sovMediapbTl(uint64(m.OwnerId))
	}
	if m.File != nil {
		l = m.File.Size()
		n += 1 + l + sovMediapbTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsGetEncryptedFile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.Id != 0 {
		n += 1 + sovMediapbTl(uint64(m.Id))
	}
	if m.AccessHash != 0 {
		n += 1 + sovMediapbTl(uint64(m.AccessHash))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLNbfsGetFileLocationSecret) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovMediapbTl(uint64(m.Constructor))
	}
	if m.VolumeId != 0 {
		n += 1 + sovMediapbTl(uint64(m.VolumeId))
	}
	if m.LocalId != 0 {
		n += 1 + sovMediapbTl(uint64(m.LocalId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMediapbTl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMediapbTl(x uint64) (n int) {
	return sovMediapbTl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DocumentId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DocumentId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DocumentId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PredicateName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PredicateName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessHash", wireType)
			}
			m.AccessHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessHash |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLDocumentId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_documentId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_documentId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data2 == nil {
				m.Data2 = &DocumentId{}
			}
			if err := m.Data2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DocumentList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DocumentList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DocumentList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PredicateName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PredicateName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Documents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Documents = append(m.Documents, &mtproto.Document{})
			if err := m.Documents[len(m.Documents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLDocumentList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_documentList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_documentList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data2 == nil {
				m.Data2 = &DocumentList{}
			}
			if err := m.Data2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FileLocationSecret) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileLocationSecret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileLocationSecret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PredicateName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PredicateName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			m.Secret = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Secret |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLFileLocationSecret) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_fileLocationSecret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_fileLocationSecret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data2 == nil {
				m.Data2 = &FileLocationSecret{}
			}
			if err := m.Data2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PhotoDataRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PhotoDataRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PhotoDataRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PredicateName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PredicateName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhotoId", wireType)
			}
			m.PhotoId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PhotoId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessHash", wireType)
			}
			m.AccessHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessHash |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			m.Date = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Date |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SizeList = append(m.SizeList, &mtproto.PhotoSize{})
			if err := m.SizeList[len(m.SizeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLPhotoDataRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_photoDataRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_photoDataRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data2 == nil {
				m.Data2 = &PhotoDataRsp{}
			}
			if err := m.Data2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsUploadPhotoFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_uploadPhotoFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_uploadPhotoFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			m.OwnerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.File == nil {
				m.File = &mtproto.InputFile{}
			}
			if err := m.File.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsProfile", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsProfile = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsUploadVideoFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_uploadVideoFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_uploadVideoFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			m.OwnerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.File == nil {
				m.File = &mtproto.InputFile{}
			}
			if err := m.File.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsGetPhotoFileData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_getPhotoFileData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_getPhotoFileData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhotoId", wireType)
			}
			m.PhotoId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PhotoId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsUploadedPhotoMedia) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_uploadedPhotoMedia: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_uploadedPhotoMedia: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			m.OwnerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Media", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Media == nil {
				m.Media = &mtproto.InputMedia{}
			}
			if err := m.Media.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsProfile", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsProfile = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsUploadedDocumentMedia) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_uploadedDocumentMedia: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_uploadedDocumentMedia: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			m.OwnerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Media", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Media == nil {
				m.Media = &mtproto.InputMedia{}
			}
			if err := m.Media.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsGetDocument) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_getDocument: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_getDocument: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DocumentId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DocumentId == nil {
				m.DocumentId = &DocumentId{}
			}
			if err := m.DocumentId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsGetDocumentList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_getDocumentList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_getDocumentList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdList = append(m.IdList, &DocumentId{})
			if err := m.IdList[len(m.IdList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsUploadEncryptedFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_uploadEncryptedFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_uploadEncryptedFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			m.OwnerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediapbTl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.File == nil {
				m.File = &mtproto.InputEncryptedFile{}
			}
			if err := m.File.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsGetEncryptedFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_getEncryptedFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_getEncryptedFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessHash", wireType)
			}
			m.AccessHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessHash |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLNbfsGetFileLocationSecret) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_nbfs_getFileLocationSecret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_nbfs_getFileLocationSecret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= (TLConstructor(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeId", wireType)
			}
			m.VolumeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VolumeId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalId", wireType)
			}
			m.LocalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LocalId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMediapbTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediapbTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMediapbTl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMediapbTl
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMediapbTl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMediapbTl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMediapbTl
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMediapbTl(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMediapbTl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMediapbTl   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mediapb.tl.proto", fileDescriptor_mediapb_tl_0451dcac72627873) }

var fileDescriptor_mediapb_tl_0451dcac72627873 = []byte{
	// 1231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x5f, 0x68, 0x1c, 0xd5,
	0x17, 0xde, 0xbb, 0x49, 0x76, 0x93, 0x93, 0x36, 0xbf, 0xe9, 0x4d, 0x93, 0xce, 0x6f, 0x63, 0xb6,
	0xeb, 0xd4, 0xc6, 0x6d, 0x63, 0xb3, 0xb8, 0xf5, 0x41, 0x7c, 0x10, 0x34, 0xad, 0xb8, 0xba, 0x89,
	0x71, 0x12, 0x2d, 0x14, 0x64, 0x99, 0x9d, 0xb9, 0xd9, 0x1d, 0xd8, 0xdd, 0x19, 0xe6, 0xce, 0x46,
	0xda, 0x27, 0x95, 0x16, 0x7d, 0x11, 0x11, 0x44, 0xf2, 0x96, 0xd2, 0x8a, 0x08, 0x0a, 0x7d, 0x52,
	0x5a, 0x7c, 0x11, 0x41, 0xcc, 0xa3, 0x8a, 0xa0, 0x7d, 0x48, 0xd1, 0x28, 0x05, 0x1f, 0x94, 0xd2,
	0x22, 0xf8, 0xb7, 0x2b, 0x73, 0xe7, 0xcf, 0xce, 0xec, 0xcc, 0xc4, 0xda, 0x1a, 0xf5, 0x6d, 0xef,
	0xb9, 0xdf, 0x9e, 0xfb, 0x9d, 0xef, 0x9c, 0x7b, 0xce, 0x1d, 0xe0, 0x9a, 0x44, 0x51, 0x25, 0xbd,
	0x3a, 0x63, 0x36, 0x66, 0x74, 0x43, 0x33, 0x35, 0x9c, 0x76, 0x2c, 0x99, 0x43, 0x35, 0xd5, 0xac,
	0xb7, 0xab, 0x33, 0xb2, 0xd6, 0x2c, 0xd4, 0xb4, 0x9a, 0x56, 0x60, 0xfb, 0xd5, 0xf6, 0x32, 0x5b,
	0xb1, 0x05, 0xfb, 0x65, 0xff, 0x2f, 0xb3, 0x9b, 0xca, 0x75, 0xd2, 0x94, 0x2c, 0x47, 0xf4, 0x44,
	0x4b, 0xb6, 0xad, 0xc2, 0x05, 0x04, 0x70, 0x44, 0x93, 0xdb, 0x4d, 0xd2, 0x32, 0x4b, 0x0a, 0xde,
	0x0f, 0x23, 0xba, 0x41, 0x14, 0x55, 0x96, 0x4c, 0x52, 0x69, 0x49, 0x4d, 0xc2, 0xa3, 0x1c, 0xca,
	0x0f, 0x89, 0x3b, 0x3d, 0xeb, 0xbc, 0xd4, 0x24, 0xf8, 0x7e, 0x18, 0x96, 0xb5, 0x16, 0x35, 0x8d,
	0xb6, 0x6c, 0x6a, 0x06, 0x9f, 0xcc, 0xa1, 0xfc, 0x48, 0x71, 0x7c, 0xc6, 0xe5, 0xba, 0x54, 0x9e,
	0xed, 0xee, 0x8a, 0x7e, 0x28, 0x1e, 0x81, 0xa4, 0xaa, 0xf0, 0x7d, 0x39, 0x94, 0xef, 0x13, 0x93,
	0xaa, 0x82, 0xf7, 0xc2, 0xb0, 0x24, 0xcb, 0x84, 0xd2, 0x4a, 0x5d, 0xa2, 0x75, 0xbe, 0x9f, 0x6d,
	0x80, 0x6d, 0x7a, 0x54, 0xa2, 0x75, 0xcc, 0x43, 0x7a, 0x85, 0x18, 0x54, 0xd5, 0x5a, 0xfc, 0x40,
	0x0e, 0xe5, 0x07, 0x44, 0x77, 0x29, 0x3c, 0x00, 0x3b, 0x97, 0xca, 0x15, 0xa5, 0x4b, 0xfe, 0x00,
	0x0c, 0x28, 0x92, 0x29, 0x15, 0x19, 0xe7, 0xe1, 0xe2, 0xa8, 0xc7, 0xa7, 0x1b, 0xa0, 0x68, 0x23,
	0x84, 0x33, 0x08, 0x76, 0xb8, 0xd6, 0xb2, 0x4a, 0xcd, 0xed, 0x0f, 0xbc, 0x00, 0x43, 0x2e, 0x55,
	0xca, 0xf7, 0xe5, 0xfa, 0xf2, 0xc3, 0xc5, 0x5d, 0x33, 0x4d, 0x93, 0x65, 0xc1, 0x23, 0x28, 0x76,
	0x31, 0xc2, 0x83, 0xf0, 0x3f, 0x5f, 0x78, 0x8c, 0xe4, 0x74, 0x30, 0xc0, 0xb1, 0x50, 0x80, 0x16,
	0xca, 0x0d, 0xf1, 0x65, 0x04, 0xf8, 0x11, 0xb5, 0x41, 0xca, 0x9a, 0x2c, 0x99, 0xaa, 0xd6, 0x5a,
	0x24, 0xb2, 0x41, 0xfe, 0x81, 0x40, 0xc7, 0x21, 0x45, 0xd9, 0x51, 0x4e, 0x96, 0x9d, 0x95, 0xf0,
	0x18, 0x8c, 0x2d, 0x95, 0x2b, 0xcb, 0x61, 0x46, 0xf7, 0x06, 0xa3, 0x9a, 0xf0, 0x0e, 0x09, 0xb3,
	0x77, 0x63, 0xbb, 0x8a, 0x60, 0xc7, 0x42, 0x5d, 0x33, 0xb5, 0x23, 0x92, 0x29, 0x89, 0x54, 0xdf,
	0xfe, 0xa8, 0xfe, 0x0f, 0x83, 0xba, 0x75, 0x60, 0xc5, 0xab, 0xde, 0x34, 0x5b, 0x97, 0x6e, 0xa2,
	0x84, 0x31, 0xf4, 0x2b, 0x92, 0x49, 0x9c, 0xfa, 0x65, 0xbf, 0xad, 0x72, 0xa0, 0xea, 0x49, 0x52,
	0x69, 0xa8, 0xd4, 0xe4, 0x53, 0xac, 0x1c, 0xb0, 0x57, 0x0e, 0x2c, 0xb4, 0x45, 0xf5, 0x24, 0x11,
	0x07, 0x2d, 0x90, 0x95, 0x55, 0xa7, 0x1c, 0x74, 0x7f, 0xd0, 0xb1, 0xe5, 0xe0, 0x97, 0xc6, 0x95,
	0xec, 0x5d, 0x04, 0x7b, 0x96, 0xca, 0x95, 0x56, 0x75, 0x99, 0x56, 0xda, 0x7a, 0x43, 0x93, 0x14,
	0x86, 0xb2, 0x34, 0xee, 0x95, 0x05, 0xfd, 0x25, 0x59, 0xb4, 0x67, 0x5b, 0xc4, 0xb0, 0x64, 0x49,
	0xda, 0xb2, 0xb0, 0x75, 0x49, 0xc1, 0x53, 0xd0, 0x6f, 0x25, 0x9b, 0xa9, 0xe5, 0x0f, 0xae, 0xd4,
	0xd2, 0xdb, 0xa6, 0x75, 0xac, 0xc8, 0xf6, 0xf1, 0x24, 0x80, 0x4a, 0x2b, 0xba, 0xa1, 0x31, 0xb4,
	0xa5, 0xde, 0xa0, 0x38, 0xa4, 0xd2, 0x05, 0xdb, 0x20, 0xbc, 0x1e, 0xe2, 0xfd, 0xb4, 0xaa, 0x90,
	0x7f, 0x9f, 0xb7, 0xa0, 0x01, 0xef, 0xf2, 0xaa, 0x11, 0xd3, 0x13, 0xd3, 0x92, 0xfd, 0xf6, 0x88,
	0x79, 0x75, 0x96, 0x0c, 0xd4, 0x99, 0xf0, 0x3e, 0x82, 0x4c, 0x50, 0x09, 0x62, 0xe7, 0x70, 0xce,
	0x72, 0xbb, 0x3d, 0x62, 0x1c, 0x80, 0x01, 0xe6, 0xc0, 0x51, 0x63, 0x34, 0xa8, 0x06, 0x3b, 0x58,
	0xb4, 0x11, 0x7f, 0x96, 0xc7, 0x35, 0x04, 0x93, 0xbd, 0xec, 0xdd, 0xb6, 0xf5, 0x9f, 0x08, 0x40,
	0x38, 0x8d, 0x60, 0xd4, 0x97, 0x51, 0x97, 0xdc, 0x6d, 0xf0, 0xba, 0x0f, 0x86, 0xdd, 0xfe, 0xed,
	0x52, 0x8b, 0x19, 0x4b, 0xd0, 0x1d, 0x63, 0xc2, 0xf3, 0xbe, 0x8a, 0xf7, 0xf1, 0x60, 0x13, 0xe0,
	0xd6, 0xb9, 0xdc, 0x03, 0x69, 0x55, 0xb1, 0xdb, 0x4d, 0x92, 0xb5, 0x9b, 0x48, 0x1e, 0x29, 0x55,
	0x61, 0xdd, 0xe6, 0x1c, 0x82, 0x89, 0x60, 0xb6, 0x8e, 0xb6, 0x64, 0xe3, 0x84, 0x6e, 0x12, 0x65,
	0xfb, 0x6e, 0x5e, 0x21, 0x70, 0xf3, 0x26, 0x82, 0xa9, 0x0a, 0x9c, 0xef, 0x5c, 0xc1, 0xd3, 0x28,
	0x70, 0x07, 0xff, 0x2e, 0x8a, 0xf6, 0x1b, 0x25, 0x19, 0xf7, 0x46, 0xe9, 0xeb, 0x6d, 0xf0, 0xc2,
	0xab, 0xbe, 0xda, 0xae, 0x11, 0x33, 0x62, 0xea, 0xde, 0x3a, 0x99, 0x09, 0x18, 0x5a, 0xd1, 0x1a,
	0xed, 0x26, 0xe9, 0x0a, 0x36, 0x68, 0x1b, 0x4a, 0x8a, 0x25, 0x66, 0x43, 0x93, 0xa5, 0x86, 0x3b,
	0x95, 0x06, 0xc4, 0x34, 0x5b, 0x97, 0x94, 0x83, 0x1f, 0xf7, 0x5b, 0xcf, 0x23, 0x9f, 0x5b, 0xbc,
	0x0b, 0x76, 0xce, 0x8a, 0xb3, 0x87, 0x8b, 0x95, 0xa7, 0xe6, 0x1f, 0x9f, 0x7f, 0xe2, 0xd8, 0x3c,
	0x97, 0xc0, 0x7b, 0x01, 0xdb, 0x26, 0xff, 0x5c, 0xe1, 0x2e, 0x5e, 0x7b, 0x63, 0xed, 0x46, 0xa7,
	0xd3, 0xe9, 0x20, 0xcc, 0x03, 0x67, 0x03, 0xba, 0xf5, 0xc9, 0xfd, 0xb0, 0x7e, 0xf9, 0x0c, 0xea,
	0xfe, 0xd5, 0xff, 0x42, 0xe1, 0x5e, 0xf9, 0x62, 0xe3, 0x03, 0xe7, 0xaf, 0xfb, 0x81, 0xb7, 0x01,
	0xe1, 0x91, 0xcf, 0x5d, 0xb8, 0x7c, 0xfd, 0xb7, 0x5f, 0x6d, 0xd8, 0xdd, 0x90, 0xb1, 0x61, 0x51,
	0x93, 0x89, 0x5b, 0x5d, 0xfb, 0xec, 0xd2, 0x2f, 0xf1, 0x40, 0x6f, 0x14, 0x70, 0xef, 0xac, 0x9f,
	0xfa, 0xbe, 0x63, 0x03, 0xf3, 0x30, 0xe1, 0x03, 0xf6, 0xf6, 0x66, 0xee, 0xc3, 0x4b, 0xcf, 0xbd,
	0xf5, 0x93, 0x8d, 0x3c, 0x08, 0x93, 0x21, 0x97, 0xfe, 0x9e, 0xca, 0x9d, 0x5f, 0x3f, 0x75, 0xc5,
	0x09, 0xe7, 0x10, 0xe4, 0x22, 0xb0, 0x81, 0x0e, 0xc6, 0xbd, 0xf7, 0xd1, 0x97, 0xd7, 0x7f, 0xb7,
	0xe1, 0xfb, 0x60, 0x3c, 0x48, 0xc2, 0x45, 0x72, 0x6f, 0xae, 0x5e, 0xd9, 0x70, 0x40, 0x42, 0x20,
	0xa4, 0x9e, 0xbb, 0xce, 0xbd, 0xf0, 0xe2, 0xe7, 0xe7, 0x53, 0x78, 0x1a, 0xb2, 0xa1, 0x73, 0x03,
	0x85, 0xce, 0x7d, 0xda, 0x39, 0xfb, 0xda, 0xcf, 0xee, 0xa9, 0x3d, 0xa1, 0x07, 0x91, 0x17, 0x6f,
	0xbc, 0xfd, 0x6d, 0x1a, 0xe7, 0x03, 0x91, 0x44, 0xd6, 0x2b, 0xb7, 0x71, 0xed, 0xec, 0x8f, 0x28,
	0xd3, 0xff, 0xd2, 0xb9, 0x6c, 0xa2, 0xf8, 0x5d, 0x0a, 0xd2, 0xe2, 0xc2, 0xec, 0x7c, 0x75, 0x99,
	0xe2, 0x27, 0x61, 0x77, 0xe4, 0x0b, 0x22, 0xe7, 0x2b, 0xe5, 0xc8, 0x4c, 0x66, 0xa2, 0x5f, 0x27,
	0x42, 0x02, 0xcf, 0x05, 0x5c, 0x76, 0x87, 0x7b, 0x9c, 0x4b, 0x0f, 0x91, 0x09, 0xbf, 0x9f, 0x85,
	0x04, 0x5e, 0x84, 0xb1, 0xe8, 0x99, 0x7c, 0x67, 0xc8, 0x5f, 0x2f, 0x24, 0x9e, 0xe3, 0x31, 0xd8,
	0x13, 0x37, 0x76, 0xf7, 0xc5, 0xd0, 0xf4, 0x83, 0x2c, 0xc7, 0x0e, 0xd3, 0x39, 0x42, 0xa9, 0x54,
	0x23, 0xcc, 0x2c, 0x24, 0xf0, 0x33, 0x90, 0xd9, 0x62, 0x22, 0x4e, 0xc5, 0xfa, 0x0e, 0xe0, 0xe2,
	0xdd, 0x1f, 0x05, 0x2e, 0x34, 0xce, 0xee, 0x88, 0xd2, 0xc1, 0xdd, 0x8d, 0xd6, 0xd4, 0xcd, 0x7a,
	0xef, 0x34, 0xca, 0x6d, 0xe5, 0xca, 0x42, 0x64, 0xa2, 0x3f, 0x51, 0x84, 0x04, 0x3e, 0x0e, 0x7c,
	0xec, 0x70, 0xb9, 0x2b, 0x26, 0xec, 0x00, 0x2a, 0x33, 0xee, 0x31, 0x0d, 0xd8, 0x85, 0x04, 0x5e,
	0xea, 0x96, 0x40, 0xd0, 0x71, 0x64, 0x09, 0xdc, 0xac, 0x57, 0xd9, 0x49, 0x55, 0x74, 0x83, 0x9f,
	0x8a, 0x72, 0x1d, 0xc6, 0x65, 0xb6, 0xfa, 0xba, 0x11, 0x12, 0x0f, 0x3f, 0x74, 0xf5, 0xeb, 0x2c,
	0x5a, 0xdf, 0xcc, 0xa2, 0x4f, 0x36, 0xb3, 0xe8, 0xab, 0xcd, 0x2c, 0x5a, 0xfd, 0x26, 0x9b, 0x38,
	0x3e, 0xcd, 0x3e, 0xe3, 0xe5, 0xba, 0x64, 0x16, 0xe4, 0x46, 0x9b, 0x9a, 0xc4, 0x28, 0x48, 0xba,
	0x5e, 0xa0, 0xc4, 0x58, 0x51, 0x65, 0x52, 0x60, 0x1e, 0x0b, 0x8e, 0xdf, 0x6a, 0x8a, 0xd1, 0x3f,
	0xfc, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x61, 0xeb, 0xfe, 0x3a, 0x10, 0x00, 0x00,
}
