package service

import (
	"context"

	media_client "open.chat/app/service/media/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetStickers(ctx context.Context, request *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getStickers#43d4f2c - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	stickers := mtproto.MakeTLMessagesStickers(&mtproto.Messages_Stickers{
		Hash_INT32:  request.GetHash_INT32(),
		Hash_STRING: request.GetHash_STRING(),
		Stickers:    []*mtproto.Document{},
	}).To_Messages_Stickers()

	if request.Emoticon != "🎯" {
		log.Debugf("messages.getStickers#43d4f2c - reply: %s", stickers.DebugString())
		return stickers, nil
	}

	// check inputStickerSetEmpty
	set, err := s.StickerCore.GetStickerSet(ctx, mtproto.MakeTLInputStickerSetDice(&mtproto.InputStickerSet{
		Emoticon: request.Emoticon,
	}).To_InputStickerSet())
	if err != nil {
		log.Errorf("messages.getStickerSet - error: %v", err)
		return nil, err
	} else if set == nil {
		err = mtproto.ErrStickersetInvalid
		log.Errorf("messages.getStickerSet - error: %v", err)
		return nil, err
	}

	packs, idList := s.StickerCore.GetStickerPackList(ctx, set.GetId())
	set.Count = int32(len(packs))

	var (
		documents []*mtproto.Document
	)

	if len(idList) == 0 {
		documents = []*mtproto.Document{}
	} else {
		documents, err = media_client.GetDocumentByIdList(idList)
		if err != nil {
			log.Error(err.Error())
			documents = []*mtproto.Document{}
		} else {
			for i := 0; i < len(documents); i++ {
				idxImageSize := -1
				for j := 0; j < len(documents[i].Attributes); j++ {
					switch documents[i].Attributes[j].Constructor {
					case mtproto.CRC32_documentAttributeImageSize:
						documents[i].Attributes[j].PredicateName = mtproto.Predicate_documentAttributeImageSize
						idxImageSize = j
					case mtproto.CRC32_documentAttributeSticker:
						documents[i].Attributes[j].PredicateName = mtproto.Predicate_documentAttributeSticker
						documents[i].Attributes[j].Stickerset.PredicateName = mtproto.Predicate_inputStickerSetID
					case mtproto.CRC32_documentAttributeFilename:
						documents[i].Attributes[j].PredicateName = mtproto.Predicate_documentAttributeFilename
					}
				}

				if idxImageSize >= 0 && documents[i].Attributes[idxImageSize].W == 0 {
					documents[i].Attributes = append(documents[i].Attributes[:idxImageSize], documents[i].Attributes[idxImageSize+1:]...)
				}
			}
		}
	}

	stickers.Hash_INT32 = 1
	stickers.Stickers = documents

	log.Debugf("messages.getStickers#43d4f2c - reply: %s", stickers.DebugString())
	return stickers, nil
}
