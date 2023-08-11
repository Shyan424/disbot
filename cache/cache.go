package cache

import (
	"sync"
)

type Cache[T any] struct {
	lock  *sync.Mutex
	cache map[string]map[string]T
}

func NewCache[T any]() Cache[T] {
	return Cache[T]{&sync.Mutex{}, make(map[string]map[string]T)}
}

func (c *Cache[T]) Get(guild string, key string) (T, bool) {
	value, ok := c.cache[guild][key]

	return value, ok
}

func (c *Cache[T]) GetAllByGuild(guild string) map[string]T {
	return c.cache[guild]
}

func (c *Cache[T]) GetAll() map[string]map[string]T {
	return c.cache
}

func (c *Cache[T]) Set(guild string, key string, t T) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[guild][key] = t
}
