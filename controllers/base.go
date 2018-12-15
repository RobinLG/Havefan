package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type baseController struct {
	beego.Controller
	o orm.Ormer
}

// Check if the user is logged in
func (p *baseController) Perpare() {
	logs.Debug("function prepare")
	p.o = orm.NewOrm()
	if p.GetSession("username") == nil {
		p.History("Not Logged In", "login")
	} else {
		p.Data["Username"] = p.GetSession("username")
		p.Data["url"] = "about/link.html"
		p.Data["bulletin"] = "bulletin.html"
		p.TplName = "main.html"
	}
}

// Use this function to redirect
func (p *baseController) History(msg string, url string) {
	if url == "" {
		// Return and refresh page
		p.Ctx.WriteString("<script>alert('"+ msg +"');window.history.go(-1);</script>")
		logs.Debug(msg)
		p.StopRun()
	} else {
		p.Redirect(url, 302)
	}
}
