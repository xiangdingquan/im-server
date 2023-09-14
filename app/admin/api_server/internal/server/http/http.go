package http

import (
	"math"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"

	"open.chat/app/admin/api_server/api"
	"open.chat/app/admin/api_server/internal/service"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

var (
	svc *service.Service
)

func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/api")
	{
		g2 := g.Group("/push")
		{
			g2.POST("/serviceNotifications", pushServiceNotifications)
		}

		g2 = g.Group("/user")
		{
			g2.GET("/createPredefinedUser", createPredefinedUser)
			g2.POST("/createPredefinedUser", createPredefinedUser)
			g2.GET("/createPredefinedUser2", createPredefinedUser2)
			g2.POST("/createPredefinedUser2", createPredefinedUser2)

			g2.GET("/updatePredefinedUsername", updatePredefinedUsername)
			g2.POST("/updatePredefinedUsername", updatePredefinedUsername)

			g2.GET("/updatePredefinedProfile", updatePredefinedProfile)
			g2.POST("/updatePredefinedProfile", updatePredefinedProfile)

			g2.GET("/updatePredefinedVerified", updatePredefinedVerified)
			g2.POST("/updatePredefinedVerified", updatePredefinedVerified)

			g2.GET("/updatePredefinedCode", updatePredefinedCode)
			g2.POST("/updatePredefinedCode", updatePredefinedCode)

			g2.GET("/updatePredefinedProfilePhoto", updatePredefinedProfilePhoto)
			g2.POST("/updatePredefinedProfilePhoto", updatePredefinedProfilePhoto)

			g2.GET("/getPredefinedUsers", getPredefinedUsers)
			g2.POST("/getPredefinedUsers", getPredefinedUsers)
		}

		g2 = g.Group("/ban")
		{
			g2.GET("/toggleBan", toggleBan)
			g2.POST("/toggleBan", toggleBan)
			g2.GET("/ban", ban)
			g2.POST("/ban", ban)
			g2.GET("/unBan", unBan)
			g2.POST("/unBan", unBan)
		}

		g2 = g.Group("/sms")
		{
			g2.GET("/sendCode", sendVerifyCode)
			g2.POST("/sendCode", sendVerifyCode)
			g2.GET("/verifyCode", verifyCode)
			g2.POST("/verifyCode", verifyCode)
		}

		g2 = g.Group("/dns")
		{
			g2.GET("/resolve", resolve)
			g2.POST("/resolve", resolve)
		}

		g2 = g.Group("/contact")
		{
			g2.GET("/addContact", addContact)
			g2.POST("/addContact", addContact)
		}

		g2 = g.Group("/chat")
		{
			g2.GET("/getAllChats", getAllChats)
			g2.POST("/getAllChats", getAllChats)
		}
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func DoHandlerHelper(c *bm.Context, req interface{}, cb func(c *bm.Context) (interface{}, error)) {
	var (
		b             binding.Binding
		contentType   = c.Request.Header.Get("Content-Type")
		authorization = c.Request.Header.Get("Authorization")
	)

	if !svc.Auth(authorization) {
		err := mtproto.ErrAccessTokenInvalid
		log.Error("auth (%s) error(%v)", authorization, err)
		c.JSON(nil, err)
		return
	}

	if c.Request.Method == "GET" {
		b = binding.Form
	} else {
		var stripContentTypeParam = func(contentType string) string {
			i := strings.Index(contentType, ";")
			if i != -1 {
				contentType = contentType[:i]
			}
			return contentType
		}

		contentType = stripContentTypeParam(contentType)
		switch contentType {
		case binding.MIMEJSON:
			b = binding.JSON
		case binding.MIMEXML, binding.MIMEXML2:
			b = binding.XML
		case binding.MIMEMultipartPOSTForm:
			b = binding.FormMultipart
		case binding.MIMEPOSTForm:
			b = binding.FormPost
		default:
			b = binding.Form
		}
	}

	if err := c.BindWith(req, b); err != nil {
		log.Errorf("bind form error: %v", err)
		c.JSON(nil, err)
		return
	}

	r, err := cb(c)
	c.JSON(r, err)
}

func pushServiceNotifications(c *bm.Context) {
}

func createPredefinedUser(c *bm.Context) {
	req := new(api.CreatePredefinedUser)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.CreatePredefinedUser(c, req)
	})
}

func createPredefinedUser2(c *bm.Context) {
	req := new(api.CreatePredefinedUser)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.CreatePredefinedUser(c, req)
	})
}

func updatePredefinedUsername(c *bm.Context) {
	req := new(api.UpdatePredefinedUsername)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.UpdatePredefinedUsername(c, req)
	})
}

func updatePredefinedProfile(c *bm.Context) {
	req := new(api.UpdatePredefinedProfile)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.UpdatePredefinedProfile(c, req)
	})
}

func updatePredefinedProfilePhoto(c *bm.Context) {
	req := new(api.UpdatePredefinedProfilePhoto)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.UpdatePredefinedProfilePhoto(c, req)
	})
}

func updatePredefinedVerified(c *bm.Context) {
	req := new(api.UpdatePredefinedVerified)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.UpdatePredefinedVerified(c, req)
	})
}

func updatePredefinedCode(c *bm.Context) {
	req := new(api.UpdatePredefinedCode)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.UpdatePredefinedCode(c, req)
	})
}

func getPredefinedUsers(c *bm.Context) {
	req := new(api.GetPredefinedUsers)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.GetPredefinedUsers(c, req)
	})
}

func toggleBan(c *bm.Context) {
	req := new(api.ToggleBan)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.ToggleBan(c, req)
	})
}

func ban(c *bm.Context) {
	req := new(api.ToggleBan)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		if req.Expires == 0 {
			req.Expires = int32(math.MaxInt32)
		}
		return svc.ToggleBan(c, req)
	})
}

func unBan(c *bm.Context) {
	req := new(api.ToggleBan)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		if req.Expires != 0 {
			req.Expires = 0
		}
		return svc.ToggleBan(c, req)
	})
}

func sendVerifyCode(c *bm.Context) {
	req := new(api.SendVerifyCode)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.SendVerifyCode(c, req)
	})
}

func verifyCode(c *bm.Context) {
	req := new(api.VerifyCode)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.VerifyCode(c, req)
	})
}

func addContact(c *bm.Context) {
	req := new(api.AddContact)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.AddContact(c, req)
	})
}

func getAllChats(c *bm.Context) {
	req := new(api.GetAllChats)
	DoHandlerHelper(c, req, func(c *bm.Context) (i interface{}, err error) {
		return svc.GetAllChats(c, req)
	})
}
