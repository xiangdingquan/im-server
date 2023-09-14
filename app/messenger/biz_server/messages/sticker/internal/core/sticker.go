package core

import (
	"context"
	"math/rand"

	"github.com/gogo/protobuf/types"
	"open.chat/app/messenger/biz_server/messages/sticker/internal/dal/dataobject"
	"open.chat/mtproto"
	"open.chat/pkg/util"
)

func makeStickerSet(do *dataobject.StickerSetsDO) *mtproto.StickerSet {
	stickers := mtproto.MakeTLStickerSet(&mtproto.StickerSet{
		Archived:      util.Int8ToBool(do.Archived),
		Official:      util.Int8ToBool(do.Official),
		Masks:         util.Int8ToBool(do.Masks),
		Animated:      util.Int8ToBool(do.Animated),
		InstalledDate: nil, // &types.Int32Value{Value: 1544590989},
		Id:            do.StickerSetId,
		AccessHash:    do.AccessHash,
		Title:         do.Title,
		ShortName:     do.ShortName,
		Thumb:         nil,
		ThumbDcId:     nil,
		Count:         do.Count,
		Hash:          rand.Int31(),
		Installed:     false,
	}).To_StickerSet()

	if do.InstalledDate != 0 {
		stickers.InstalledDate = &types.Int32Value{Value: do.InstalledDate}
		stickers.Installed = true
	}
	return stickers
}

func (m *StickerCore) GetStickerSetList(ctx context.Context, hash int32) []*mtproto.StickerSet {
	//
	doList, _ := m.StickerSetsDAO.SelectAll(ctx)
	stickers := make([]*mtproto.StickerSet, len(doList))
	for i := 0; i < len(doList); i++ {
		stickers[i] = makeStickerSet(&doList[i])
	}
	return stickers
}

func (m *StickerCore) GetStickerSet(ctx context.Context, stickerSet *mtproto.InputStickerSet) (*mtproto.StickerSet, error) {
	var (
		inputSet = stickerSet
		set      *mtproto.StickerSet
	)

	switch stickerSet.PredicateName {
	case mtproto.Predicate_inputStickerSetEmpty:
		return nil, mtproto.ErrStickersetInvalid
	case mtproto.Predicate_inputStickerSetID:
		do, _ := m.StickerSetsDAO.SelectByID(ctx, inputSet.GetId(), inputSet.GetAccessHash())
		if do != nil {
			set = makeStickerSet(do)
		}
	case mtproto.Predicate_inputStickerSetShortName:
		do, _ := m.StickerSetsDAO.SelectByShortName(ctx, inputSet.GetShortName())
		if do != nil {
			set = makeStickerSet(do)
		}
	case mtproto.Predicate_inputStickerSetAnimatedEmoji:
		do, _ := m.StickerSetsDAO.SelectByShortName(ctx, "AnimatedEmojies")
		if do != nil {
			set = makeStickerSet(do)
		}
	case mtproto.Predicate_inputStickerSetDice:
		if stickerSet.Emoticon == "🎲" { // 🎲
			do, _ := m.StickerSetsDAO.SelectByShortName(ctx, "AnimatedDice2")
			if do != nil {
				set = makeStickerSet(do)
			}
		} else if stickerSet.Emoticon == "🎯" { // 🎯
			do, _ := m.StickerSetsDAO.SelectByShortName(ctx, "AnimatedDart")
			if do != nil {
				set = makeStickerSet(do)
			}
		} else if stickerSet.Emoticon == "🏀" { // 🏀
			do, _ := m.StickerSetsDAO.SelectByShortName(ctx, "AnimatedBasketball")
			if do != nil {
				set = makeStickerSet(do)
			}
		} else {
			do, _ := m.StickerSetsDAO.SelectByShortName(ctx, "AnimatedDices")
			if do != nil {
				set = makeStickerSet(do)
			}
		}
	default:
		return nil, mtproto.ErrStickersetInvalid
	}

	return set, nil
}

func (m *StickerCore) GetStickerPackList(ctx context.Context, setId int64) ([]*mtproto.StickerPack, []int64) {
	doList, _ := m.StickerPacksDAO.SelectBySetID(ctx, setId)
	idList := make([]int64, 0, len(doList))

	packs := make(map[string][]int64)

	for i := 0; i < len(doList); i++ {
		idList = append(idList, doList[i].DocumentId)

		if _, ok := packs[doList[i].Emoticon]; !ok {
			packs[doList[i].Emoticon] = []int64{doList[i].DocumentId}
		} else {
			packs[doList[i].Emoticon] = append(packs[doList[i].Emoticon], doList[i].DocumentId)
		}
	}

	packs2 := make([]*mtproto.StickerPack, 0, len(packs))
	for k, v := range packs {
		packs2 = append(packs2, &mtproto.StickerPack{
			PredicateName: mtproto.Predicate_stickerPack,
			Constructor:   mtproto.CRC32_stickerPack,
			Emoticon:      k,
			Documents:     v,
		})
	}

	return packs2, idList
}
