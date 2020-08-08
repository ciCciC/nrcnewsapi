package util

import (
	"github.com/patrickmn/go-cache"
	"log"
	"nrcnewsapi/api/nrcnewsapi/src/model"
	"time"
)

var Cache = cache.New(1*time.Hour, 1*time.Hour)

func SetCache(key string, articles []model.ArticleItem) bool {
	log.Println("Is caching articles")
	Cache.Set(key, articles, cache.DefaultExpiration)
	return true
}

func GetCache(key string) ([]model.ArticleItem, bool) {
	var articleList []model.ArticleItem
	var found bool

	data, found := Cache.Get(key)

	if found {
		articleList = data.([]model.ArticleItem)
	}
	return articleList, found
}
