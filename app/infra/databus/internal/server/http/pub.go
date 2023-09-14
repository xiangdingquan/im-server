package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/infra/databus/internal/conf"
	"open.chat/app/infra/databus/internal/dsn"
	"open.chat/app/infra/databus/internal/server/tcp"
	"open.chat/pkg/log"
)

const (
	_nonRetriableErr = 1
	_retriableErr    = 2
)

type Record struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

type OffsetMsg struct {
	Partition int32 `json:"partition"`
	Offset    int64 `json:"offset"`
	ErrorCode int64 `json:"error_code"`
}

func pub(c *bm.Context) {
	key, secret, _ := c.Request.BasicAuth()

	dsn := &dsn.DSN{
		Role:   "pub",
		Key:    key,
		Secret: secret,
		Topic:  c.Request.Form.Get("topic"),
		Group:  c.Request.Form.Get("group"),
		Color:  c.Request.Form.Get("color"),
	}

	cfg, _, err := tcp.Auth(dsn, c.Request.Host)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		log.Error("auth failed, err(%v)", err)
		return
	}

	records, err := parseRecords(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	var d struct {
		Offsets []*OffsetMsg `json:"offsets"`
	}
	d.Offsets, err = pubRecords(dsn.Group, dsn.Topic, dsn.Color, cfg, records)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	rsp, err := json.Marshal(d)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("content-type", "application/json")
	c.Writer.Write(rsp)
}

func pubRecords(group, topic, color string, cfg *conf.Kafka, records []*Record) (offsets []*OffsetMsg, err error) {
	p, err := tcp.NewPub(nil, group, topic, color, cfg)
	if err != nil {
		return
	}

	for _, r := range records {
		var m OffsetMsg
		offsets = append(offsets, &m)
		// OffsetMsg 与 Record 一一对应
		// Publish 出错后继续循环填充
		if err != nil {
			m.ErrorCode = _nonRetriableErr
			continue
		}
		m.Partition, m.Offset, err = p.Publish([]byte(r.Key), nil, []byte(r.Value))
		if err != nil {
			m.ErrorCode = _nonRetriableErr
			p.Close(false)
		}
	}
	return
}

func parseRecords(c *bm.Context) ([]*Record, error) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	defer c.Request.Body.Close()

	var d struct {
		Records []*Record
	}

	if err = json.Unmarshal(b, &d); err != nil {
		return nil, err
	}

	return d.Records, nil
}
