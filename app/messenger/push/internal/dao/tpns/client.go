package tpns

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"open.chat/pkg/log"
)

const (
	//ServerURL = "https://api.tpns.tencent.com/v3/push/app"     //广州服务接入点
	ServerURL = "https://api.tpns.sh.tencent.com/v3/push/app" //上海服务接入点
	//ServerURL = "https://api.tpns.hk.tencent.com/v3/push/app"  //中国香港服务接入点
	//ServerURL = "https://api.tpns.sgp.tencent.com/v3/push/app" //新加坡服务接入点
)

// Client stores client with api key to firebase
type Client struct {
	AuthCode   string
	HTTPClient *http.Client
}

// NewClient creates a new client
func NewClient(accessID, secretKey string, timeout time.Duration) *Client {
	authCode := "Basic " + base64.StdEncoding.EncodeToString([]byte(accessID+":"+secretKey))
	return &Client{
		AuthCode:   authCode,
		HTTPClient: &http.Client{Timeout: timeout},
	}
}

// Push message to jpush
func (f *Client) Push(payload *Payload) (response *Response, err error) {
	var (
		req   *http.Request
		res   *http.Response
		t     = time.NewTimer(f.HTTPClient.Timeout)
		errCh = make(chan error, 1)
	)
	if req, err = http.NewRequest("POST", ServerURL, bytes.NewBuffer(payload.Marshal())); err != nil {
		log.Error("http.NewRequest(%s) error(%v)", ServerURL, err)
		return
	}
	req.Header.Set("Authorization", f.AuthCode)
	req.Header.Set("Content-Type", "application/json")
	go func() {
		res, err = f.HTTPClient.Do(req)
		errCh <- err
	}()
	select {
	case <-t.C:
		err = errors.New("http.Do timeout")
		return
	case err = <-errCh:
		if err != nil {
			log.Error("c.HTTPClient.Do() error(%v)", err)
			return
		}
	}
	defer res.Body.Close()
	response = &Response{StatusCode: res.StatusCode}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	} else if len(body) == 0 {
		return
	}
	log.Infof("body: %s", string(body))
	err = json.Unmarshal(body, &response)
	return
}
