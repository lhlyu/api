package api

import (
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	io.WriteString(w, "hello world!")
}
