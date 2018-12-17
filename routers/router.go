package routers

import (
	"robin/Havefan/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/publish", &controllers.UserController{}, "*:AddOrder")
	beego.Router("/index", &controllers.UserController{}, "get:ShowIndex")
	beego.Router("/detail", &controllers.UserController{}, "get:ShowDetail")
	beego.Router("/account", &controllers.UserController{}, "*:Account")
	beego.Router("/check", &controllers.UserController{}, "post:Check")
	beego.Router("/verify", &controllers.UserController{}, "get:ShowCheck")
	beego.Router("/changeflagone", &controllers.UserController{}, "post:UpdateFlagToOne")

}
