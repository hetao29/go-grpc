package user

import (
	"fmt"
	//"net/http"
	//"encoding/json"
	//"database/sql"
	"modules/utility"
	//"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE TABLE `user` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(60) DEFAULT NULL,
	`password` varchar(200) DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
*/
func init() {
	fmt.Println("world")
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
}
type User struct {
	ID        uint `gorm:"primary_key"`
	Name string `json:"name"`
}
type Result struct {
	Code int   `json:"code"`
	//Data []Tag `json:"data"`
}
func GetByNameAndPwd(name string,pwd string)(*User){
	conn := utility.Conn.Table("user");
	user := &User{};
	conn.Where("name = ? AND password = ?",name,pwd).First(&user)
	return user;
}

