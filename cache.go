package cache

type Cache struct {
	m map[string]CacheItem
}

type CacheItem struct {
	value int
}

func New() *Cache {
	return &Cache{
		m: make(map[string]CacheItem),
	}
}

func (c *Cache) Set(key string, v interface{}) {
	switch value := v.(type) {
	case int:
		c.m[key] = CacheItem{
			value: value,
		}
	}
}

func (c Cache) Get(key string) int {
	return c.m[key]
}

func (c *Cache) Delete(key string) {
	delete(c.m, key)
}
