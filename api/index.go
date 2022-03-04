package api

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
	"time"
)

var client *cache.Cache

func init() {
	// 5分钟过期，10分钟清理一次
	client = cache.New(5*time.Minute, 10*time.Minute)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	io.WriteString(w, "hello world!")
}

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
