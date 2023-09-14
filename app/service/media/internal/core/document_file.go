package core

import (
	"context"
	"encoding/json"
	"time"

	"fmt"

	"open.chat/app/service/dfs/dfspb"
	"open.chat/app/service/media/internal/dal/dataobject"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

type documentData struct {
	*dataobject.DocumentsDO
}

func (m *MediaCore) DoUploadedDocumentFile2(ctx context.Context, fileMD *dfspb.DocumentFileMetadata, thumbId int64, attributes []byte) (*documentData, error) {
	data := &dataobject.DocumentsDO{
		DocumentId:       fileMD.DocumentId,
		AccessHash:       fileMD.AccessHash,
		DcId:             fileMD.DcId,
		FilePath:         fileMD.FilePath,
		FileSize:         fileMD.FileSize,
		UploadedFileName: fileMD.UploadedFileName,
		Ext:              fileMD.Ext,
		MimeType:         fileMD.MimeType,
		ThumbId:          thumbId,
		Attributes:       string(attributes),
		Version:          0,
	}
	data.Id, _, _ = m.DocumentsDAO.Insert(ctx, data)
	return &documentData{DocumentsDO: data}, nil
}

func (m *MediaCore) makeDocumentByDO(ctx context.Context, do *dataobject.DocumentsDO) *mtproto.Document {
	var (
		thumbs   []*mtproto.PhotoSize
		document *mtproto.Document = mtproto.MakeTLDocumentEmpty(nil).To_Document()
	)

	if do != nil {
		if do.ThumbId != 0 {
			thumbs = m.GetPhotoSizeList(ctx, do.ThumbId)
			log.Infof("sizeList = %#v", thumbs)
		}
		thumb := mtproto.MakeTLPhotoSizeEmpty(nil).To_PhotoSize()
		if len(thumbs) > 0 {
			thumb = thumbs[0]
			thumbs = thumbs[1:]
		}

		var attributes []*mtproto.DocumentAttribute
		err := json.Unmarshal([]byte(do.Attributes), &attributes)
		if err != nil {
			log.Error(err.Error())
			attributes = []*mtproto.DocumentAttribute{}
		}
		for i := 0; i < len(attributes); i++ {
		}
		// if do.Attributes
		document = &mtproto.Document{
			PredicateName: mtproto.Predicate_document,
			Id:            do.DocumentId,
			AccessHash:    do.AccessHash,
			Date:          int32(time.Now().Unix()),
			MimeType:      do.MimeType,
			Size2:         do.FileSize,
			Thumb:         thumb,
			DcId:          2,
			Thumbs:        thumbs,
			Version:       do.Version,
			Attributes:    attributes,
		}
	}

	return document
}

func (m *MediaCore) GetDocument(ctx context.Context, id, accessHash int64, version int32) *mtproto.Document {
	do, _ := m.DocumentsDAO.SelectByFileLocation(ctx, id, accessHash, version)
	if do == nil {
		log.Warn("")
	}
	return m.makeDocumentByDO(ctx, do)
}

func (m *MediaCore) GetDocumentList(ctx context.Context, idList []int64) []*mtproto.Document {
	doList, _ := m.DocumentsDAO.SelectByIdList(ctx, idList)
	documentList := make([]*mtproto.Document, len(doList))
	for i := 0; i < len(doList); i++ {
		documentList[i] = m.makeDocumentByDO(ctx, &doList[i])
	}
	return documentList
}

func (m *MediaCore) SaveDocument(ctx context.Context, fileName, fileExtName string, document *mtproto.Document) {
	var (
		photoId int64
		aStr    string
	)

	for i, sz := range document.GetThumbs() {
		if sz.GetLocation() == nil {
			continue
		}
		photoId = sz.GetLocation().GetVolumeId()

		photoDatasDO := &dataobject.PhotoDatasDO{
			PhotoId:    sz.Location.VolumeId,
			PhotoType:  0,
			DcId:       2,
			VolumeId:   sz.Location.VolumeId,
			LocalId:    sz.Location.LocalId,
			AccessHash: sz.Location.Secret,
			Width:      sz.W,
			Height:     sz.H,
			FileSize:   sz.Size2,
			FilePath:   fmt.Sprintf("%s/%d.%d.dat", getSizeType(i+1), sz.Location.VolumeId, sz.Location.LocalId),
			Ext:        ".jpg",
		}

		m.PhotoDatasDAO.Insert(ctx, photoDatasDO)
	}

	if document.GetAttributes() != nil {
		aBuf, _ := json.Marshal(document.GetAttributes())
		aStr = hack.String(aBuf)
	}
	data := &dataobject.DocumentsDO{
		DocumentId:       document.Id,
		AccessHash:       document.AccessHash,
		DcId:             document.DcId,
		FilePath:         fmt.Sprintf("%d.%d.dat", document.Id, document.AccessHash),
		FileSize:         document.Size2,
		UploadedFileName: fileName,
		Ext:              fileExtName,
		MimeType:         document.MimeType,
		ThumbId:          photoId,
		Attributes:       aStr,
		Version:          0,
	}
	data.Id, _, _ = m.DocumentsDAO.Insert(ctx, data)
	return
}
