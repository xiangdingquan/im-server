package core

import (
	"context"
	"encoding/json"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

func makePrivacyRulesByDO(do *dataobject.UserPrivaciesDO) []*mtproto.PrivacyRule {
	var (
		rules = make([]*mtproto.PrivacyRule, 0)
	)

	if err := json.Unmarshal([]byte(do.Rules), &rules); err != nil {
		log.Errorf("getPrivacy - Unmarshal PrivacyRulesData(%d)error: %v", do.Id, err)
	}

	return rules
}

func (m *UserCore) GetPrivacy(ctx context.Context, userId int32, keyType int) (rules []*mtproto.PrivacyRule, err error) {
	var (
		do *dataobject.UserPrivaciesDO
	)

	rules = make([]*mtproto.PrivacyRule, 0)
	if do, err = m.UserPrivaciesDAO.SelectPrivacy(ctx, userId, int8(keyType)); err != nil {
		return
	} else if do == nil {
		rules = append(rules, mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule())
		return
	} else {
		if err = json.Unmarshal([]byte(do.Rules), &rules); err != nil {
			log.Errorf("getPrivacy - Unmarshal PrivacyRulesData(%d)error: %v", do.Id, err)
			return
		}
	}
	return
}

func (m *UserCore) refreshPrivacyToCache(ctx context.Context, userId int32) error {
	doList, err := m.UserPrivaciesDAO.SelectPrivacyAll(ctx, userId)
	if err != nil {
		log.Errorf("refreshToCache err: %v", err)
		return err
	}
	privacyList := make(map[int][]*mtproto.PrivacyRule)
	for i := 0; i < len(doList); i++ {
		rules := make([]*mtproto.PrivacyRule, 0)
		if err = json.Unmarshal([]byte(doList[i].Rules), &rules); err != nil {
			log.Errorf("getPrivacy - Unmarshal PrivacyRulesData(%d)error: %v", doList[i].Id, err)
		} else {
			rules = []*mtproto.PrivacyRule{
				mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule(),
			}
		}
		privacyList[int(doList[i].KeyType)] = rules
	}
	return m.Redis.SetPrivacyList(ctx, userId, privacyList)
}

func (m *UserCore) CheckPrivacy(ctx context.Context, keyType int, selfId, peerId int32, isContact bool) bool {
	rules, err := m.GetPrivacy(ctx, selfId, keyType)
	if err != nil {
		return false
	} else if len(rules) == 0 {
		return true
	}

	return model.CheckPrivacyIsAllow(selfId, rules, peerId, func(id, checkId int32) bool {
		return isContact
	}, func(checkId int32, idList []int32) bool {
		return true
	})
}

func (m *UserCore) SetPrivacy(ctx context.Context, userId int32, keyType int, rules []*mtproto.PrivacyRule) (err error) {
	bData, _ := json.Marshal(rules)
	do := &dataobject.UserPrivaciesDO{
		UserId:  userId,
		KeyType: int8(keyType),
		Rules:   hack.String(bData),
	}
	_, _, err = m.UserPrivaciesDAO.InsertOrUpdate(ctx, do)

	m.Redis.SetPrivacy(ctx, userId, keyType, rules)
	return err
}

func (m *UserCore) CreateAllPrivacy(ctx context.Context, userId int32) (err error) {
	defaultValues := []*mtproto.PrivacyRule{
		mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule(),
	}

	doList := make([]*dataobject.UserPrivaciesDO, model.MAX_KEY_TYPE)
	bData, _ := json.Marshal(defaultValues)
	for i := model.STATUS_TIMESTAMP; i <= model.MAX_KEY_TYPE; i++ {
		doList = append(doList, &dataobject.UserPrivaciesDO{
			UserId:  userId,
			KeyType: int8(i),
			Rules:   hack.String(bData),
		})
	}
	_, _, err = m.UserPrivaciesDAO.InsertBulk(ctx, doList)
	if err != nil {
		return
	}

	// cache
	privacyList := make(map[int][]*mtproto.PrivacyRule)
	for i := model.STATUS_TIMESTAMP; i <= model.MAX_KEY_TYPE; i++ {
		privacyList[i] = defaultValues
	}
	m.Redis.SetPrivacyList(ctx, userId, privacyList)
	return err
}
