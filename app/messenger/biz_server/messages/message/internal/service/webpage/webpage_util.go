package webpage

import (
	"crypto/tls"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"

	"open.chat/mtproto"
	"open.chat/pkg/http_client"
	"open.chat/pkg/log"
)

func GetWebPagePreview(rawurl string) *mtproto.WebPage {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Warnf("getWebPagePreview error - ", err)
		return mtproto.MakeTLWebPageEmpty(nil).To_WebPage()
	}

	ogContents, err := GetWebpageOgList(u.String(), []string{"image", "site_name", "title", "description", "url"})
	log.Debugf("ogContents - ", ogContents)

	if len(ogContents) == 0 {
		return mtproto.MakeTLWebPageEmpty(nil).To_WebPage()
	} else {
		var webPage = mtproto.MakeTLWebPage(&mtproto.WebPage{
			Id:         rand.Int63(),
			Url:        rawurl,
			DisplayUrl: u.String()[len(u.Scheme)+3:],
			Type:       &types.StringValue{Value: "article"},
		})
		if v, ok := ogContents["title"]; ok {
			webPage.SetTitle(&types.StringValue{Value: v})
		}
		if v, ok := ogContents["site_name"]; ok {
			webPage.SetSiteName(&types.StringValue{Value: v})
		}
		if v, ok := ogContents["description"]; ok {
			webPage.SetDescription(&types.StringValue{Value: v})
		}
		var imageBody []byte

		rawImageUrl, _ := ogContents["image"]
		if rawImageUrl != "" {
			resp, err := http.Get(rawImageUrl)
			if err != nil {
				log.Warnf("get image body error - ", err)
			} else {
				imageBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Warnf("read image body error - ", imageBody)
				}
			}
		} else {
			log.Warnf("image empty")
		}

		if len(imageBody) > 0 {
		}

		return webPage.To_WebPage()
	}
}

func GetWebpageOgList(url string, ogParams []string) (params map[string]string, err error) {
	var body string
	if strings.HasPrefix(strings.ToUpper(url), "https://") {
		body, err = http_client.Get(url).
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
			SetTimeout(30*time.Second, 10*time.Second).
			String()
	} else {
		body, err = http_client.Get(url).
			SetTimeout(30*time.Second, 10*time.Second).
			String()
	}

	if err != nil {
		log.Errorf("get url error - ", err)
		return
	}

	params = GetWebpageOgListFromContent(body, ogParams)
	return
}

func GetWebpageOgListFromContent(content string, ogParams []string) map[string]string {
	pattern := regexp.MustCompile(`<meta\s+property\s*=\s*"og:([0-9a-zA-Z-]+)"\s+content\s*=\s*"([^"]*?)"\s*/>`)
	allMatches := pattern.FindAllStringSubmatch(content, -1)
	allParams := make(map[string]string)
	for _, val := range allMatches {
		log.Infof("og val: %v = %v\n", val[1], val[2])
		k := val[1]
		v := val[2]
		allParams[k] = v
	}

	return allParams
}
