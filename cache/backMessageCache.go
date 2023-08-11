package cache

import (
	"discordbot/model/vo"
	"sync"
)

type BackMessageCache struct {
	lock  *sync.Mutex
	cache map[string]map[string][]vo.BackMessageVo
}

func NewBackMessageCache() BackMessageCache {
	return BackMessageCache{&sync.Mutex{}, make(map[string]map[string][]vo.BackMessageVo)}
}

func (c *BackMessageCache) Get(guild string, key string) ([]vo.BackMessageVo, bool) {
	value, ok := c.cache[guild][key]

	return value, ok
}

func (c *BackMessageCache) GetAllByGuild(guild string) map[string][]vo.BackMessageVo {
	return c.cache[guild]
}

func (c *BackMessageCache) GetAll() map[string]map[string][]vo.BackMessageVo {
	return c.cache
}

func (c *BackMessageCache) Set(guild string, vo vo.BackMessageVo) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[guild][vo.Key] = append(c.cache[guild][vo.Key], vo)
}

func (c *BackMessageCache) SetAll(guild string, voSlice []vo.BackMessageVo) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for i := range voSlice {
		value := voSlice[i]
		c.cache[guild][value.Key] = append(c.cache[guild][value.Key], value)
	}
}
