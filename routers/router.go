package routers

import (
	"robin/Havefan/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/publish", &controllers.UserController{}, "*:AddOrder")
	beego.Router("/index", &controllers.UserController{}, "get:ShowIndex")

	//beego.Router("/main", &controllers.UserController{}, "*:Perpare")
	//beego.Router("/login", &controllers.UserController{}, "*:Login")
	//beego.Router("/createPage", &controllers.UserController{}, "*:RegisterPage")
	//beego.Router("/create", &controllers.UserController{}, "post:Register")
	//beego.Router("/smsVerify", &controllers.UserController{}, "post:VerifyRegister")
	//beego.Router("/about/reset", &controllers.UserController{}, "post:Reset")
	//beego.Router("/about/link", &controllers.UserController{}, "get:LinkPage")
	//beego.Router("/about/sendLink", &controllers.UserController{}, "post:SendLink")
	//beego.Router("/activate", &controllers.UserController{}, "get:VerifyLink")
	//
	//beego.Router("/bulletin", &controllers.UserController{}, "get:WatchBulletin")
}
