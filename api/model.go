package api

import "encoding/json"

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewOkResp(data interface{}) *Resp {
	return &Resp{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}

func NewFailResp(code int, msg string) *Resp {
	return &Resp{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (r *Resp) JsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}
