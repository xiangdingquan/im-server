package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"open.chat/mtproto"

	"open.chat/app/admin/api_server/api"
)

var pkcs1PemPrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAvKLEOWTzt9Hn3/9Kdp/RdHcEhzmd8xXeLSpHIIzaXTLJDw8B
hJy1jR/iqeG8Je5yrtVabqMSkA6ltIpgylH///FojMsX1BHu4EPYOXQgB0qOi6kr
08iXZIH9/iOPQOWDsL+Lt8gDG0xBy+sPe/2ZHdzKMjX6O9B4sOsxjFrk5qDoWDri
oJorAJ7eFAfPpOBf2w73ohXudSrJE0lbQ8pCWNpMY8cB9i8r+WBitcvouLDAvmtn
TX7akhoDzmKgpJBYliAY4qA73v7u5UIepE8QgV0jCOhxJCPubP8dg+/PlLLVKyxU
5CdiQtZj2EMy4s9xlNKzX8XezE0MHEa6bQpnFwIDAQABAoIBACd+SGjfyursZoiO
MW/ejAK/PFJ3bKtNI8P++v9Enh8vF8swUBgMmzIdv93jZfnnD1mtT46kU6mXd3fy
FMunGVrjlwkLKET9MC8B5U46Es6T/H4fAA8KCzA+ywefOEnVA5pIsB7dIFFhyNDB
uO8zrBAFfsu+Y1KMlggsZaZGDXB/WVyUJDbEOMZstVx4uNhpcEgKYp28YQMP/yvv
dp4UgnTxXXXpDghzO5iqi5tUWY0p1lH2ii2OZBxEdqdDl7TirorhUDYIivyoe3B5
H30RNBRok/6w7W0WPyY2lSIcjd3cLPte6vx0QfBXVo2A6N9LTKAtAw3iWBp0x9NZ
N5p8OeECgYEA8QywXlM8nH5M7Sg2sMUYBOHA22O26ZPio7rJzcb8dlkV5gVHm+Kl
aDP61Uy8KoYABQ5kFdem/IQAUPepLxmJmiqfbwOIjfajOD3uVAQunFnDCHBWm4Uk
onbpdA5NlT/OUoSjIBemiBR/4CpDK1cEby/sg+EvQaGxqtedEe4xFmcCgYEAyFXe
MyAAOLpzmnCs9NYTTvMPofW8y+kLDodfbskl7M8q6l20VMo/E+g1gQ+65Aah901Z
/LKGi6HpzmHi5q9O2OJtqyI6FVwjXa07M5ueDbHcVKJw4hC9W0oHpMg8hqumPAWF
+MoN/Toy77p5LzoR30WUdhPvOAJPEL1p2a6r29ECgYEAiXfCEVkI5PqGZm2bmv4b
75TLhpJ8WwMSqms48Vi828V8Xpy+NOFxkVargv9rBBk9Y6TMYUSGH9Yr1AEZhBnd
RoVuPUJXmxaACPAQvetQpavvNR3T1od82AZWpvANQMONp7Oqz/+M4mhGcRHJEqti
hQJgsOk4KQbMqvChy/r6FZsCgYEAwyaqgkD9FkXC0UJLqWFUg8bQhqPcGwLUC34h
n8kAUbPpiU5omWQ+mATPAf8xvmkbo81NCJVb7W93U90U7ET/2NSRonCABkiwBtP2
ZKqGB68oA6YNspo960ytL38DPui80aFLxXQGtpPYBKEw5al6uXWNTozSrkvJe3QY
Rb4amdECgYBpGk7zPcK1TbJ++W5fkiory4qOdf0L1Zf0NbML4fY6dIww+dwMVUpq
FbsgCLqimqOFaaECU+LQEFUHHM7zrk7NBf7GzBvQ+qJx8zhJ66sFVox+IirBUyR9
Vh0+z5tIbFbKmYkO06NbeMlq87JexSlocPZtA3HMhEga5/0fHNHsNw==
-----END RSA PRIVATE KEY-----
`)

func (s *Service) Resolve(ctx context.Context, i *api.ResolveRequest) (r *api.ResolveResponse, err error) {
	t2 := api.GetDnsType(i.Type)
	if t2 == api.DnsTypeUnknown {
		err = fmt.Errorf("invalid type")
	}

	r = &api.ResolveResponse{
		Status: 0,
		TC:     false,
		RD:     false,
		RA:     true,
		AD:     true,
		CD:     false,
		Question: []api.Question{
			{
				Name: i.Name,
				Type: t2,
			},
		},
		Answer:  []api.Answer{},
		Comment: "",
	}

	switch t2 {
	case api.DnsTypeA:
		r.Answer = []api.Answer{
			{
				Name: i.Name,
				Type: t2,
				TTL:  559,
				Data: "192.168.1.150",
			},
		}
	case api.DnsTypeTXT:
		var data []byte
		r.Answer = []api.Answer{
			{
				Name: i.Name,
				Type: t2,
				TTL:  59,
				Data: base64.StdEncoding.EncodeToString(data),
			},
		}
	}

	return
}

func makeSimpleConfig() *mtproto.Help_ConfigSimple {
	secret, _ := base64.StdEncoding.DecodeString("7v3aJUx42fogKsU2B56IuAh3d3cuZ29vZ2xlLmNvbQ==")
	config := mtproto.MakeTLHelpConfigSimple(&mtproto.Help_ConfigSimple{
		Date:    1562949126,
		Expires: 1565541126,
		Rules: []*mtproto.AccessPointRule{
			mtproto.MakeTLAccessPointRule(&mtproto.AccessPointRule{
				PhonePrefixRules: "",
				DcId:             2,
				Ips: []*mtproto.IpPort{
					mtproto.MakeTLIpPortSecret(&mtproto.IpPort{
						Ipv4:   -1959013790,
						Port:   14544,
						Secret: secret,
					}).To_IpPort(),
				},
			}).To_AccessPointRule(),
		},
	}).To_Help_ConfigSimple()

	return config
}
