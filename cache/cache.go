package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/cache"
)

const (
	CacheName   = "file"
	CacheConfig = `{"CachePath":"./static/cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`
)

var BaseCache cache.Cache

func Put(key string, data interface{}, timeout time.Duration) error {
	if BaseCache == nil {
		return errors.New("Cache is Nil")
	}

	return BaseCache.Put(key, data, timeout)
}

func Get(key string) (interface{}, error) {
	if BaseCache == nil {
		return nil, errors.New("Cache is Nil")
	}

	m := BaseCache.Get(key)
	if m == nil {
		return nil, fmt.Errorf("Cache Key:%s is Nil", key)
	}

	return m, nil
}

func Delete(key string) error {
	if BaseCache == nil {
		return errors.New("Cache is Nil")
	}

	return BaseCache.Delete(key)
}

func init() {
	BaseCache, _ = cache.NewCache(CacheName, CacheConfig)
}
