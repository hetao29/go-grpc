package utility

import (
	//"database/sql"
	//"github.com/jinzhu/gorm"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/gookit/config/v2"
	//"math/rand"
	//"time"
	//_ "github.com/go-sql-driver/mysql"
)

/**
* Redis Ring
* https://godoc.org/github.com/go-redis/redis#Ring
 */
var Redis map[string]*redis.Ring
var RedisConfig *config.Config

//var err error;
func init() {
	Redis = make(map[string]*redis.Ring)
}

func InitRedis(cfg *config.Config) {
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
	RedisConfig = cfg
	//Conn, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3401)/test")
	//Redis := redis.NewRing(&redis.RingOptions{
	//Addrs: map[string]string{
	//"shard1": ":7000",
	//"shard2": ":7001",
	//"shard3": ":7002",
	//},
	//})
	//Redis.Ping()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
}
func GetRedis(nodeName string, nodeType string) *redis.Ring {
	//source := rand.NewSource(time.Now().UnixNano())
	uri := nodeName + "." + nodeType
	master := RedisConfig.Strings("redis." + nodeName + "." + nodeType)
	fmt.Printf("\n GetDb 1:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	if len(master) == 0 {
		master = RedisConfig.Strings("redis." + nodeName + ".master")
		fmt.Printf("\n GetDb 2:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	}
	if len(master) == 0 {
		master = RedisConfig.Strings("redis.default." + nodeType)
		fmt.Printf("\n GetDb 3:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	}
	if len(master) == 0 {
		master = RedisConfig.Strings("redis.default.master")
		fmt.Printf("\n GetDb 3:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
		return nil
	}
	addrs := make(map[string]string)
	for _, ele := range master {
		addrs[ele] = ele
	}
	//Shuffle(master, source) // [c b a]
	//uri := master[0]
	//fmt.Printf("\n GetDb 4:\n %#v,%d, %#v", master, len(master), uri) // map[string]string{"key":"val2", "key2":"val20"}
	if Redis[uri] == nil {
		Redis[uri] = redis.NewRing(&redis.RingOptions{
			Addrs: addrs,
		})

		//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
		//var err error
		//fmt.Println("redis init")
		////Db[nodeName], err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3401)/test")
		//Db[uri], err = gorm.Open("mysql", uri)
		//if err != nil {
		//	fmt.Println(err.Error())
		//	Db[uri] = nil
		//	//return nil
		//}
	}
	return Redis[uri]
}
