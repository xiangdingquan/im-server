package helper

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

type (
	// MethodHandler .
	MethodHandler func(context.Context, *grpc_util.RpcMetadata, *DataJSON) (*ResultJSON, error)
	// DataJSON .
	DataJSON struct {
		*mtproto.DataJSON
	}

	// ResultJSON .
	ResultJSON struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data,omitempty"`
	}
)

func (m *DataJSON) JsonCall(data interface{}, fn func(interface{}) *ResultJSON) (*ResultJSON, error) {
	if data != nil {
		err := m.GetJSONData(data)
		if err != nil {
			return nil, errors.New("json data is wrong")
		}
	}
	return fn(data), nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetJSONData .
func (m *DataJSON) GetJSONData(data interface{}) error {
	//fmt.Printf(m.GetData() + "\n")
	return json.Unmarshal([]byte(m.GetData()), data)
}

// GetJSONData .
func (m *ResultJSON) GetJSONData(data interface{}) error {
	strJSON, err := json.Marshal(m.Data)
	if err != nil {
		strJSON = []byte(`{}`)
	}
	return json.Unmarshal([]byte(strJSON), data)
}

// ToDataJSON .
func (r *ResultJSON) ToDataJSON() *DataJSON {
	if r == nil {
		return nil
	}
	strJSON, err := json.Marshal(r)
	if err != nil {
		strJSON = []byte(`{"code":"-1","msg":"error","data":{}}`)
	}
	return &DataJSON{mtproto.MakeTLDataJSON(&mtproto.DataJSON{
		Data: string(strJSON),
	}).To_DataJSON()}
}

// Render render it to http response writer.
func (r *ResultJSON) Render(w http.ResponseWriter) (err error) {
	var jsonBytes []byte
	r.WriteContentType(w)
	if jsonBytes, err = json.Marshal(r); err != nil {
		err = errors.WithStack(err)
		return
	}
	if _, err = w.Write(jsonBytes); err != nil {
		err = errors.WithStack(err)
	}
	return
}

// WriteContentType write content-type to http response writer.
func (r *ResultJSON) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	header["Content-Type"] = []string{"application/json; charset=utf-8"}
}
