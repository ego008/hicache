package hicache

import "sync"

type Cache struct {
	hash map[string]interface{}
	lock sync.RWMutex
}

func New() *Cache {
	var c Cache
	c.hash = make(map[string]interface{})
	return &c
}

func (c *Cache) Set(k string, v interface{}) {
	c.lock.Lock()
	c.hash[k] = v
	c.lock.Unlock()
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.lock.RLock()
	v, ok := c.hash[k]
	c.lock.RUnlock()
	return v, ok
}

func (c *Cache) Count() int {
	return len(c.hash)
}

func (c *Cache) Del(k string) {
	c.lock.Lock()
	delete(c.hash, k)
	c.lock.Unlock()
}

func (c *Cache) Flush() {
	c.lock.Lock()
	c.hash = make(map[string]interface{})
	c.lock.Unlock()
}

func (c *Cache) Incr(k string, n int) int {
	v, ok := c.Get(k)
	new_v := n
	if ok {
		switch v.(type) {
		case int:
			new_v += v.(int)
		}
	}
	c.Set(k, new_v)
	return new_v
}
