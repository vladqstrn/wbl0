package wbl0

import (
	"github.com/jackc/pgx/v4"
)

type CacheManager struct {
	cache  map[string]interface{}
	dbConn *pgx.Conn
}

func CreateCacheManager(dbConn *pgx.Conn) *CacheManager {
	gc := GetCache(dbConn)
	c := make(map[string]interface{})
	c = gc
	return &CacheManager{
		c,
		dbConn,
	}
}

func (c *CacheManager) Set(key string, value interface{}) {
	c.cache[key] = value
	PushCache(c.dbConn, key, value)
}
func (c *CacheManager) Get(key string) (interface{}, bool) {
	val, ok := c.cache[key]
	return val, ok
}
