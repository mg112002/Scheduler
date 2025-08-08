package caching

import (
	"log"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache
var once sync.Once

func InitCache() {
	once.Do(func() {
		Cache = cache.New(5*time.Minute, 10*time.Minute)
	})
	log.Println("Cache initialized")
}
