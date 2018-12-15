package main

import (
	_ "robin/unity/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"robin/unity/models"
	"github.com/astaxie/beego/logs"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	log := logs.NewLogger(10000) // Build logger, parameter is size of buffer
	log.SetLogger("console") // Output in console
	log.SetLevel(logs.LevelDebug) // Level of buffer
	log.EnableFuncCallDepth(true) // Display number of line and filename
}

func main() {
	beego.Run()
}

