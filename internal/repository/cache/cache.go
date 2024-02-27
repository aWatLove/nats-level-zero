package cache

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"sync"
)

type Cache struct {
	Mutex sync.Mutex
	Data  map[string]model.Order
}

func NewCache() *Cache {
	var cache Cache
	cache.Data = make(map[string]model.Order)
	return &cache
}
