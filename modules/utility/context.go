package utility

import (
	//"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	//"net/http"
)
type Context struct{
	gorm.Model
//		Writer http.ResponseWriter;
//		Request *http.Request;
		Db *gorm.DB
}
func (con *Context) GetDb() (*gorm.DB){
	return con.Db;
}
