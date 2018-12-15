package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

// Add database connection to orm of go
func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	logs.Debug("Database URL: ", dsn)
	orm.RegisterDataBase("default", "mysql", dsn)
	// add User to orm
	orm.RegisterModel(new(Order))
	orm.RegisterModel(new(Bulletin))

}

// Return table name with prefix
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
