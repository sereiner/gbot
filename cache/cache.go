package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sereiner/library/types"
)

type Cache struct {
	prefix string
	redis  *redis.Client
}

func NewCache(addr, password string) *Cache {
	return &Cache{
		prefix: "gbot",
		redis: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		}),
	}
}

func (c *Cache) Get(key string) (string, error) {
	data, err := c.redis.Get(c.prefix + key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return data, nil
		}
		return "", nil
	}

	return data, nil
}

func (c *Cache) Many(keys []string) (map[string]string, error) {
	data, err := c.redis.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	r := make(map[string]string)
	for k, v := range data {
		r[keys[k]] = types.GetString(v)
	}

	return r, nil
}

func (c *Cache) Put(key, value string, ttl int64) error {
	panic("implement me")
}

func (c *Cache) PutMany(v map[string]string, ttl int64) error {
	panic("implement me")
}

func (c *Cache) Increment(key string, delta int64) (n int64, err error) {
	panic("implement me")
}

func (c *Cache) Decrement(key string, delta int64) (n int64, err error) {
	panic("implement me")
}

func (c *Cache) Forever(key, value string) error {
	panic("implement me")
}

func (c *Cache) Forget(key string) error {
	panic("implement me")
}

func (c *Cache) Exists(key string) bool {
	panic("implement me")
}

func (c *Cache) Flush() {
	panic("implement me")
}

func (c *Cache) GetPrefix() string {
	panic("implement me")
}
