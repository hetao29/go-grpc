package user

import (
	"fmt"
	"net/http"
	"encoding/json"
	"database/sql"
	"models/utility"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE TABLE `user` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(60) DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
*/
func init() {
	fmt.Println("world")
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
}
type User struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Result struct {
	Code int   `json:"code"`
	Data []Tag `json:"data"`
}
func Get()(User){
	user := &User{};
	utility.Conn.First(&user)
	return user;
}

