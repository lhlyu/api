package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func TauriChatGPT(w http.ResponseWriter, r *http.Request) {
	log.Println("<TauriChatGPT>")
	w.Header().Set("content-type", "application/json")

	update, ok := client.Get("Tauri-ChatGPT")
	if ok {
		io.WriteString(w, update.(string))
		return
	}

	update = getTauriChatGPTLatestJosn()

	client.Set("Tauri-ChatGPT", update, time.Minute*10)

	io.WriteString(w, update.(string))
}

func getTauriChatGPTLatestJosn() string {
	resp, err := http.Get("https://api.github.com/repos/lhlyu/tauri-chatgpt/releases/latest")
	if err != nil {
		log.Println("getTauriChatGPTLatestJosn:", err)
		return "{}"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("getTauriChatGPTLatestJosn.ReadAll:", err)
		return "{}"
	}

	data := &Data{}

	if err := json.Unmarshal(body, data); err != nil {
		log.Println("getTauriChatGPTLatestJosn.Unmarshal:", err)
		return "{}"
	}

	for _, asset := range data.Assets {
		if asset.Name == "latest.json" {
			resp, err := http.Get(asset.BrowserDownloadUrl)
			if err != nil {
				log.Println("getTauriChatGPTLatestJosn.BrowserDownloadUrl:", err)
				return "{}"
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("getTauriChatGPTLatestJosn.BrowserDownloadUrl.ReadAll:", err)
				return "{}"
			}
			return string(body)
		}
	}
	return "{}"
}

type Data struct {
	TagName     string `json:"tag_name"`
	PublishedAt string `json:"published_at"`
	Assets      []struct {
		Name               string `json:"name"`
		BrowserDownloadUrl string `json:"browser_download_url"`
	} `json:"assets"`
}
