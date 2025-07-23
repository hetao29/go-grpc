package utility

import (
	//"database/sql"
	"log"
	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/config/v2"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
	//"github.com/gookit/config/v2/json"
)

// Db joj
var Db map[string]*gorm.DB

// DbConfig db
var DbConfig *config.Config

func init() {
	Db = make(map[string]*gorm.DB)
}

//var err error;

// DbConfig ....
//type Db struct {
//}

// InitDb ...
/*
func InitDb(config DbConfig) {
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
	var err error
	log.Println("db init")
	Db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3401)/test")
	if err != nil {
		log.Println(err.Error())
	}
}
*/
// InitDb it
func InitDb(cfg *config.Config) {
	DbConfig = cfg
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
	//var err error
	//log.Println("db init")
	//Db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3401)/test")
	//if err != nil {
	//	log.Println(err.Error())
	//}
}

// Shuffle s
func Shuffle(array []string, source rand.Source) {
	random := rand.New(source)
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

// GetDb 获取 DB
/**
* @param nodeName ，节点名
* @param nodeType ，节点类型，master/slave
 */
func GetDb(nodeName string, nodeType string) *gorm.DB {
	source := rand.NewSource(time.Now().UnixNano())
	master := DbConfig.Strings("db." + nodeName + "." + nodeType)
	log.Printf("\n GetDb 1:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	if len(master) == 0 {
		master = DbConfig.Strings("db." + nodeName + ".master")
		log.Printf("\n GetDb 2:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	}
	if len(master) == 0 {
		master = DbConfig.Strings("db.default." + nodeType)
		log.Printf("\n GetDb 3:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
	}
	if len(master) == 0 {
		master = DbConfig.Strings("db.default.master")
		log.Printf("\n GetDb 3:\n %#v,%d", master, len(master)) // map[string]string{"key":"val2", "key2":"val20"}
		return nil
	}
	Shuffle(master, source) // [c b a]
	uri := master[0]
	log.Printf("\n GetDb 4:\n %#v,%d, %#v", master, len(master), uri) // map[string]string{"key":"val2", "key2":"val20"}
	if Db[uri] == nil {

		//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
		var err error
		log.Println("db init")
		//Db[nodeName], err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3401)/test")
		Db[uri], err = gorm.Open("mysql", uri)
		if err != nil {
			log.Println(err.Error())
			Db[uri] = nil
			//return nil
		}
	}
	return Db[uri]
}
