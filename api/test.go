package api

import (
	"io"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	kind := r.URL.Query().Get("kind")
	switch kind {
	case "int":
		io.WriteString(w, NewOkResp(10086).JsonString())
	case "float":
		io.WriteString(w, NewOkResp(3.1415).JsonString())
	case "bool":
		io.WriteString(w, NewOkResp(true).JsonString())
	case "string":
		io.WriteString(w, NewOkResp("this is test").JsonString())
	case "array":
		io.WriteString(w, NewOkResp([]int{1, 1, 2, 3, 5, 8, 13}).JsonString())
	case "json":
		io.WriteString(w, NewOkResp(struct {
			Name   string  `json:"name"`
			Age    int     `json:"age"`
			Height float64 `json:"height"`
			IsVip  bool    `json:"is_vip"`
		}{
			Name:   "Tom",
			Age:    18,
			Height: 178.9,
			IsVip:  true,
		}).JsonString())
	default:
		io.WriteString(w, NewOkResp([]string{"int", "float", "bool", "string", "array", "json"}).JsonString())
	}

}
