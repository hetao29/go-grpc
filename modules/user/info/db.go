package info

import (
	"encoding/json"
	"fmt"
	"modules/utility"
	"time"
)

//import "github.com/vmihailenco/msgpack/v4"

/*
CREATE TABLE `user` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(60) DEFAULT NULL,
	`password` varchar(200) DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
*/
func init() {
	//fmt.Println("world")
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
}

// User 对象
type User struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name"`
}

// MarshalBinary use msgpack
/*
func (s *User) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(s)
}
// UnmarshalBinary use msgpack
func (s *User) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, s)
}
*/
// MarshalBinary use msgpack
func (s *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalBinary use msgpack
func (s *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

// Result 结果集
type Result struct {
	Code int `json:"code"`
	//Data []Tag `json:"data"`
}

// GetByNameAndPwd 获取用户
func GetByNameAndPwd(name string, pwd string) *User {
	user := &User{}
	key := "user_" + name +"_"+pwd;
	redis := utility.GetRedis("default", "master")
	err := redis.Get(key).Scan(user)
	if err != nil {
		fmt.Printf("redis get error: %v", err)
	} else {
		fmt.Printf("redis get ok: %v", user)
		return user
	}

	db := utility.GetDb("default", "master") //.Table("user");
	if db == nil {
		return user

	}
	conn := db.Table("user")
	conn.Where("name = ? AND password = ?", name, pwd).First(&user)

	//length := redis.Len();
	r := redis.Set(key, user, 100*time.Second)
	//v:= redis.Get("name").String()
	fmt.Printf("redis error msg: %v", r.Err())
	//fmt.Println("get redis:",v);
	return user
}
