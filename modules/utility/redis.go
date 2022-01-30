package utility

import (
	"github.com/go-redis/redis/v7"
	"github.com/gookit/config/v2"
)

// Redis r
/*
* Redis Ring
* https://godoc.org/github.com/go-redis/redis#Ring
 */
var Redis map[string]*redis.Ring

// RedisConfig r
var RedisConfig *config.Config

//var err error;
func init() {
	Redis = make(map[string]*redis.Ring)
}

// InitRedis i
func InitRedis(cfg *config.Config) {
	RedisConfig = cfg
}

// GetRedis g
func GetRedis(nodeName string, nodeType string) *redis.Ring {
	uri := nodeName + "." + nodeType
	master := RedisConfig.Strings("redis." + nodeName + "." + nodeType)
	//log.Printf("\n GetDb 1:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	if len(master) == 0 {
		master = RedisConfig.Strings("redis." + nodeName + ".master")
	//	log.Printf("\n GetDb 2:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	}
	if len(master) == 0 {
		master = RedisConfig.Strings("redis.default." + nodeType)
	//	log.Printf("\n GetDb 3:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	}
	if len(master) == 0 {
		master = RedisConfig.Strings("redis.default.master")
	//	log.Printf("\n GetDb 3:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
		return nil
	}
	addrs := make(map[string]string)
	for _, ele := range master {
		addrs[ele] = ele
	}
	if Redis[uri] == nil {
		Redis[uri] = redis.NewRing(&redis.RingOptions{
			Addrs: addrs,
		})
	}
	return Redis[uri]
}
