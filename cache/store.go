package cache

type Store interface {
	Get(key string) (string, error)
	Many(keys []string) (map[string]string, error)
	Put(key, value string, ttl int64) error
	PutMany(v map[string]string, ttl int64) error
	Increment(key string, delta int64) (n int64, err error)
	Decrement(key string, delta int64) (n int64, err error)
	Forever(key, value string) error
	Forget(key string) error
	Exists(key string) bool
	Flush()
	GetPrefix() string
}
