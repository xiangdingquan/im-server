package dfspb

type DocumentFileMetadata struct {
	FileId           int64
	DocumentId       int64
	AccessHash       int64
	DcId             int32
	FileSize         int32
	FilePath         string
	UploadedFileName string
	Ext              string
	Md5Hash          string
	MimeType         string
}

type PhotoFileMetadata struct {
	FileId    int64
	PhotoId   int64
	PhotoType int8
	SizeType  string
	DcId      int32
	VolumeId  int64
	LocalId   int32
	SecretId  int64
	Width     int32
	Height    int32
	FileSize  int32
	FilePath  string
	Ext       string
}

type EncryptedFileMetadata struct {
	FileId          int64
	EncryptedFileId int64
	AccessHash      int64
	DcId            int32
	FileSize        int32
	FilePath        string
	Md5Hash         string
}
