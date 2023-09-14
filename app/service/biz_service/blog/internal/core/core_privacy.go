package core

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/pkg/errors"
	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

const (
	privacyKeyBegin        int8 = 1
	privacyKeyHideTo            = 1
	privacyKeySkipFrom          = 2
	privacyKeyShowInDays        = 3
	privacyKeyShowInCounts      = 4
	privacyKeyEnd               = 5

	privacyRuleAllowAll      = 1
	privacyRuleRestrictUsers = 2
	privacyRuleAllowDays     = 3
	privacyRuleAllowCounts   = 4
)

type PrivacyRule struct {
	Rule   int8    `json:"rule"`
	Users  []int32 `json:"users"`
	Counts int32   `json:"counts"`
	Days   int32   `json:"days"`
}

func (m *BlogCore) SetPrivacy(ctx context.Context, userId int32, key int8, rules string) (err error) {
	if key != privacyKeyShowInDays && key != privacyKeyShowInCounts {
		err = errors.New("invalid arguments")
		return
	}

	err = m.touchPrivacy(ctx, userId)
	if err != nil {
		return
	}

	err = m.setPrivacy(ctx, userId, key, rules)
	return
}

func (m *BlogCore) setPrivacy(ctx context.Context, userId int32, key int8, rules string) (err error) {
	log.Debugf("[blogsPrivacy] setPrivacy(%d, %d, %s)", userId, key, rules)

	do := &dataobject.BlogUserPrivaciesDO{
		UserId:  userId,
		KeyType: key,
		Rules:   rules,
	}
	_, _, err = m.BlogUserPrivaciesDAO.InsertOrUpdate(ctx, do)
	if err != nil {
		return
	}

	err = m.Redis.SetPrivacy(ctx, userId, key, rules)
	return
}

func (m *BlogCore) ModifyPrivacyUsers(ctx context.Context, userId int32, key int8, uidList []int32, isAdding bool) (err error) {
	if key != privacyKeyHideTo && key != privacyKeySkipFrom {
		err = errors.New("invalid arguments")
		return
	}

	rules, err := m.getPrivacy(ctx, userId, key)
	if err != nil {
		return
	}

	log.Debugf("[blogsReply] modifyPrivacyUsers, oldRules: %v", rules)

	var newList []int32
	for _, r := range rules {
		if r.Rule == privacyRuleAllowAll {
			if isAdding {
				newList = uidList
			}
			break
		}
		if r.Rule == privacyRuleRestrictUsers {
			if isAdding {
				newList = uidList
				for _, uid := range r.Users {
					contains, _ := util.Contains(uid, newList)
					if !contains {
						newList = append(newList, uid)
					}
				}
			} else {
				newList = util.Int32Subtract(r.Users, uidList)
			}
			break
		}
	}

	var newRule PrivacyRule
	if len(newList) == 0 {
		newRule = PrivacyRule{Rule: privacyRuleAllowAll}
	} else {
		newRule = PrivacyRule{
			Rule:  privacyRuleRestrictUsers,
			Users: newList,
		}
	}

	s, err := json.Marshal([]PrivacyRule{newRule})
	if err != nil {
		return
	}

	err = m.setPrivacy(ctx, userId, key, string(s))
	return
}

func (m *BlogCore) GetUserPrivacy(ctx context.Context, userId int32) (out map[int8]string, err error) {
	err = m.touchPrivacy(ctx, userId)
	if err != nil {
		return
	}

	out, err = m.Redis.GetUserPrivacy(ctx, userId)
	return
}

func (m *BlogCore) CheckPrivacyHideTo(ctx context.Context, settingOwnerUID, uidToCheck int32) (canSee bool, err error) {
	canSee = false

	rules, err := m.getPrivacy(ctx, settingOwnerUID, privacyKeyHideTo)
	if err != nil {
		return
	}

	for _, rule := range rules {
		if rule.Rule == privacyRuleAllowAll {
			canSee = true
			log.Debugf("[blogsPrivacy] CheckPrivacyHideTo, can see, settingOwnerUID:%d, uidToCheck:%d", settingOwnerUID, uidToCheck)
			return
		} else if rule.Rule == privacyRuleRestrictUsers {
			for _, u := range rule.Users {
				if u == uidToCheck {
					canSee = false
					log.Debugf("[blogsPrivacy] CheckPrivacyHideTo, can not see, settingOwnerUID:%d, uidToCheck:%d", settingOwnerUID, uidToCheck)
					return
				}
			}
			canSee = true
			log.Debugf("[blogsPrivacy] CheckPrivacyHideTo, can see, settingOwnerUID:%d, uidToCheck:%d", settingOwnerUID, uidToCheck)
			return
		}
	}

	log.Errorf("BlogCore.CheckPrivacyHideTo(%d, %d), invalid rules", settingOwnerUID, uidToCheck)
	err = ecode.Error(-1, "invalid rules")
	return
}

func (m *BlogCore) CheckPrivacySkipFrom(ctx context.Context, settingOwnerUID, uidToCheck int32) (canSee bool, err error) {
	canSee = false

	rules, err := m.getPrivacy(ctx, settingOwnerUID, privacyKeySkipFrom)
	if err != nil {
		return
	}

	for _, rule := range rules {
		if rule.Rule == privacyRuleAllowAll {
			canSee = true
			log.Debugf("[blogsPrivacy] CheckPrivacySkipFrom, can see, settingOwnerUID:%d, uidToCheck:%d", settingOwnerUID, uidToCheck)
			return
		} else if rule.Rule == privacyRuleRestrictUsers {
			for _, u := range rule.Users {
				if u == uidToCheck {
					canSee = false
					log.Debugf("[blogsPrivacy] CheckPrivacySkipFrom, can not see, settingOwnerUID:%d, uidToCheck:%d", settingOwnerUID, uidToCheck)
					return
				}
			}
			canSee = true
			log.Debugf("[blogsPrivacy] CheckPrivacySkipFrom, can see, settingOwnerUID:%d, uidToCheck:%d", settingOwnerUID, uidToCheck)
			return
		}
	}

	log.Errorf("BlogCore.CheckPrivacySkipFrom(%d, %d), invalid rules", settingOwnerUID, uidToCheck)
	err = ecode.Error(-1, "invalid rules")
	return
}

func (m *BlogCore) CheckPrivacySkipFromList(ctx context.Context, settingOwnerUID int32, uidListToCheck []int32) (canSeeList []int32, err error) {
	log.Debugf("[blogsPrivacy] CheckPrivacySkipFromList, settingOwnerUID:%d, uidListToCheck:%v", settingOwnerUID, uidListToCheck)
	rules, err := m.getPrivacy(ctx, settingOwnerUID, privacyKeySkipFrom)
	if err != nil {
		return
	}

	for _, rule := range rules {
		if rule.Rule == privacyRuleAllowAll {
			canSeeList = uidListToCheck
			log.Debugf("[blogsPrivacy] CheckPrivacySkipFromList, settingOwnerUID:%d, canSeeList:%v", settingOwnerUID, canSeeList)
			return
		} else if rule.Rule == privacyRuleRestrictUsers {
			m := make(map[int32]bool)
			for _, u := range rule.Users {
				m[u] = true
			}
			canSeeList = make([]int32, 0)
			for _, u := range uidListToCheck {
				if _, ok := m[u]; !ok {
					canSeeList = append(canSeeList, u)
				}
			}
			log.Debugf("[blogsPrivacy] CheckPrivacySkipFromList, settingOwnerUID:%d, canSeeList:%v", settingOwnerUID, canSeeList)
			return
		}
	}

	log.Errorf("BlogCore.CheckPrivacySkipFromList(%d, _), invalid rules", settingOwnerUID)
	err = ecode.Error(-1, "invalid rules")
	return
}

func (m *BlogCore) GetPrivacyShowInCounts(ctx context.Context, settingOwnerUID int32) (int32, error) {
	rules, err := m.getPrivacy(ctx, settingOwnerUID, privacyKeyShowInCounts)
	if err != nil {
		return 0, err
	}

	for _, rule := range rules {
		if rule.Rule == privacyRuleAllowAll {
			log.Debugf("[blogsPrivacy] GetPrivacyShowInCounts, settingOwnerUID:%d, counts:%d", settingOwnerUID, -1)
			return -1, nil
		} else if rule.Rule == privacyRuleAllowCounts {
			log.Debugf("[blogsPrivacy] GetPrivacyShowInCounts, settingOwnerUID:%d, counts:%d", settingOwnerUID, rule.Counts)
			return rule.Counts, nil
		}
	}

	log.Errorf("BlogCore.GetPrivacyShowInCounts(%d), invalid rules", settingOwnerUID)
	return 0, ecode.Error(-1, "invalid rules")
}

func (m *BlogCore) GetPrivacyShowInDays(ctx context.Context, settingOwnerUID int32) (int32, error) {
	rules, err := m.getPrivacy(ctx, settingOwnerUID, privacyKeyShowInDays)
	if err != nil {
		return 0, err
	}

	for _, rule := range rules {
		if rule.Rule == privacyRuleAllowAll {
			log.Debugf("[blogsPrivacy] GetPrivacyShowInDays, settingOwnerUID:%d, days:%d", settingOwnerUID, -1)
			return -1, nil
		} else if rule.Rule == privacyRuleAllowDays {
			log.Debugf("[blogsPrivacy] GetPrivacyShowInDays, settingOwnerUID:%d, days:%d", settingOwnerUID, rule.Days)
			return rule.Days, nil
		}
	}

	log.Errorf("BlogCore.GetPrivacyShowInDays(%d), invalid rules", settingOwnerUID)
	return 0, ecode.Error(-1, "invalid rules")
}

func (m *BlogCore) getPrivacy(ctx context.Context, userId int32, key int8) ([]PrivacyRule, error) {
	err := m.touchPrivacy(ctx, userId)
	if err != nil {
		return nil, err
	}

	s, err := m.Redis.GetPrivacy(ctx, userId, key)
	if err != nil {
		return nil, err
	}

	rules := make([]PrivacyRule, 0)
	err = json.Unmarshal([]byte(s), &rules)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (m *BlogCore) touchPrivacy(ctx context.Context, userId int32) (err error) {
	isExists, err := m.Redis.IsPrivacyExists(ctx, userId)
	if err != nil {
		return err
	}

	if isExists {
		return nil
	}

	l, err := m.BlogUserPrivaciesDAO.SelectUserPrivacy(ctx, userId)
	if err != nil {
		return err
	}

	mapping := make(map[int8]string, 0)
	for _, do := range l {
		mapping[do.KeyType] = do.Rules
	}

	for i := privacyKeyBegin; i < privacyKeyEnd; i++ {
		_, ok := mapping[i]
		if !ok {
			r := struct {
				Rule int8 `json:"rule"`
			}{
				Rule: privacyRuleAllowAll,
			}
			var b []byte
			b, err = json.Marshal(r)
			if err != nil {
				log.Errorf("marshal in touchPrivacy, rule: %v, error: %v", r, err)
				return
			}
			mapping[i] = "[" + string(b) + "]"
		}
	}

	err = m.Redis.SetPrivacyList(ctx, userId, mapping)
	if err != nil {
		return
	}

	err = m.Redis.ExpirePrivacy(ctx, userId, 86400)

	return
}
