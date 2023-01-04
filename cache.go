package cache

type Cache struct {
	m map[string]int
}

func New() *Cache {
	return &Cache{
		m: make(map[string]int),
	}
}

func (c *Cache) Set(key string, v interface{}) {
	switch value := v.(type) {
	case int:
		c.m[key] = value
	}
}

func (c *Cache) Get(key string) int {
	return c.m[key]
}

func (c *Cache) Delete(key string) {
	delete(c.m, key)
}
