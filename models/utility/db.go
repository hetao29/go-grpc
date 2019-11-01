package utility

import (
	//"database/sql"
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Conn *gorm.DB
var err error;
type Db struct{
}

func (db *Db) Init() {
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
	fmt.Println("db init")
	Conn, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3401)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (db * Db) Config(){
}
