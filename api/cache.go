package api

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var client *cache.Cache

func init() {
	// 5分钟过期，10分钟清理一次
	client = cache.New(5*time.Minute, 10*time.Minute)
}
