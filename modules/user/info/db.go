package info

import (
	"encoding/json"
	"github.com/hetao29/go-grpc/modules/utility"
	"log"
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
}

// User 对象
type User struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name"`
}

// RedisVersion 1
var RedisVersion = "1"

// MarshalBinary s
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
	key := RedisVersion + "_" + "user_" + name + "_" + pwd
	redis := utility.GetRedis("default", "master")
	err := redis.Get(key).Scan(user)
	if err != nil {
		log.Printf("redis get error: %v", err)
	} else {
		log.Printf("redis get ok: %v", user)
		return user
	}

	db := utility.GetDb("default", "master") //.Table("user");
	if db == nil {
		return user

	}
	conn := db.Table("user")
	conn.Where("name = ? AND password = ?", name, pwd).First(&user)

	if user.ID > 0 {
		r := redis.Set(key, user, 100*time.Second)
		log.Printf("redis error msg: %v", r.Err())
	}
	return user
}
