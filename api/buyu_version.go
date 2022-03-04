package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	buyu_key  = "buyu"
	buyu_url  = "https://api.github.com/repos/lhlyu/buyu/releases/latest"
	buyu_name = "app-release.apk"
)

type BuyuInfo struct {
	TagName    string      `json:"tag_name"`
	Draft      bool        `json:"draft"`
	Prerelease bool        `json:"prerelease"`
	Assets     []BuyuAsset `json:"assets"`
	Body       string      `json:"body"`
}

type BuyuAsset struct {
	Name               string `json:"name"`
	Label              string `json:"label"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}

type RespBuyuVersion struct {
	Version    string `json:"version"`
	Content    string `json:"content"`
	Download   string `json:"download"`
	Prerelease bool   `json:"prerelease"`
}

func BuyuVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	val, ok := client.Get(buyu_key)
	if ok {
		io.WriteString(w, NewOkResp(val).JsonString())
		return
	}

	resp, err := http.Get(buyu_url)
	if err != nil {
		log.Println("BuyuVersion - 请求异常:", err)
		io.WriteString(w, NewFailResp(1, err.Error()).JsonString())
		return
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	result := &BuyuInfo{}
	if err := json.Unmarshal(b, result); err != nil {
		log.Println("BuyuVersion - 反序列化异常:", err)
		io.WriteString(w, NewFailResp(2, err.Error()).JsonString())
		return
	}
	if result.TagName == "" {
		log.Println("BuyuVersion - 数据为空")
		io.WriteString(w, NewFailResp(3, "数据为空").JsonString())
		return
	}

	if len(result.Assets) == 0 {
		log.Println("BuyuVersion - 没有资源可下载")
		io.WriteString(w, NewFailResp(4, "没有资源可下载").JsonString())
		return
	}
	data := &RespBuyuVersion{
		Version:    result.TagName,
		Content:    result.Body,
		Download:   "",
		Prerelease: result.Prerelease,
	}
	for _, asset := range result.Assets {
		if asset.Name == buyu_name {
			data.Download = asset.BrowserDownloadUrl
			break
		}
	}
	if data.Download == "" {
		log.Println("BuyuVersion - 下载链接为空")
		io.WriteString(w, NewFailResp(5, "下载链接为空").JsonString())
		return
	}
	client.Set(buyu_key, data, time.Minute*5)
	io.WriteString(w, NewOkResp(data).JsonString())
	return
}
