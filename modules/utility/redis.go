package utility

import (
	//"database/sql"
	//"github.com/jinzhu/gorm"
	"github.com/go-redis/redis/v7"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
)

/**
* Redis Ring
* https://godoc.org/github.com/go-redis/redis#Ring
*/
var Redis *redis.Ring
//var err error;
type RedisConfig struct{
}

func InitRedis(RedisConfig) {
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
	fmt.Println("redis init")
	//Conn, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3401)/test")
	Redis := redis.NewRing(&redis.RingOptions{
Addrs: map[string]string{
"shard1": ":7000",
"shard2": ":7001",
"shard3": ":7002",
},
})
Redis.Ping()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
}
