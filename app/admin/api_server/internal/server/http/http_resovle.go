package http

import (
	"encoding/json"
	"net/http"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"

	"open.chat/app/admin/api_server/api"
	"open.chat/pkg/log"
)

func resolve(c *bm.Context) {
	req := new(api.ResolveRequest)
	if err := c.BindWith(req, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		log.Errorf("resolve error: %v", err)
		return
	}

	var (
		r     *api.ResolveResponse
		err   error
		bytes []byte
	)

	r, err = svc.Resolve(c, req)
	if err != nil {
		log.Errorf("resolve error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	bytes, err = json.Marshal(r)
	if err != nil {
		log.Errorf("resolve error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Bytes(http.StatusOK, "application/json; charset=utf-8", bytes)
}
