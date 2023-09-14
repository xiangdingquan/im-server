package dao

import (
	"context"

	"open.chat/app/infra/databus/internal/model"
	"open.chat/pkg/log"
)

const (
	_getAuth2SQL = `SELECT auth.group2,auth.operation,app.app_key,app.app_secret,auth.number,topic.topic,topic.cluster
				FROM auth LEFT JOIN app On auth.app_id=app.id LEFT JOIN topic On topic.id=auth.topic_id WHERE auth.app_id!=0 AND auth.is_delete=0`
)

func (d *Dao) Auth(c context.Context) (auths map[string]*model.Auth, err error) {
	auths = make(map[string]*model.Auth)
	rows2, err := d.db.Query(c, _getAuth2SQL)
	if err != nil {
		log.Error("getAuthStmt.Query error(%v)", err)
		return
	}
	defer rows2.Close()
	for rows2.Next() {
		a := &model.Auth{}
		if err = rows2.Scan(&a.Group, &a.Operation, &a.Key, &a.Secret, &a.Batch, &a.Topic, &a.Cluster); err != nil {
			log.Error("rows.Scan error(%v)", err)
			return
		}
		auths[a.Group] = a
	}
	return
}
