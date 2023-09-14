package service

import (
	"context"

	media_client "open.chat/app/service/media/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetStickerSet(ctx context.Context, request *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getStickerSet#2619a90e - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// check stickerSet
	// check inputStickerSetEmpty
	set, err := s.StickerCore.GetStickerSet(ctx, request.GetStickerset())
	if err != nil {
		log.Errorf("messages.getStickerSet - error: %v", err)
		return nil, err
	} else if set == nil {
		err = mtproto.ErrStickersetInvalid
		log.Errorf("messages.getStickerSet - error: %v", err)
		return nil, err
	}

	packs, idList := s.StickerCore.GetStickerPackList(ctx, set.GetId())

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

				// fixed no documentAttributeImageSize data.
				if idxImageSize >= 0 && documents[i].Attributes[idxImageSize].W == 0 {
					documents[i].Attributes = append(documents[i].Attributes[:idxImageSize], documents[i].Attributes[idxImageSize+1:]...)
				}
			}
		}
	}

	set.Count = int32(len(documents))

	reply := mtproto.MakeTLMessagesStickerSet(&mtproto.Messages_StickerSet{
		Set:       set,
		Packs:     packs,
		Documents: documents,
	})

	log.Debugf("messages.getStickerSet - reply: %s", reply.DebugString())
	return reply.To_Messages_StickerSet(), nil
}
