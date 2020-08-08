package util

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

var Cache = cache.New(1*time.Hour, 1*time.Hour)

func SetCache(key string, objectToCache interface{}) bool {
	log.Println("Is caching..")
	Cache.Set(key, objectToCache, cache.DefaultExpiration)
	return true
}

func GetCache(key string) (interface{}, bool) {
	var cachedObject interface{}
	var found bool

	data, found := Cache.Get(key)

	if found {
		cachedObject = data.(interface{})
	}
	return cachedObject, found
}
