package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type Bulletin struct {
	Id string `orm:"pk"`
	Username string `json:"username"`
	Title string `json:"title"`
	Contect string `json:"contect"`
	time string `json:"time"`
}

//return table name without prefix
func (u *Bulletin) TableName() string {
	return TableName("bulletin")
}

func (n *Bulletin) ReadDB(username string) (bulletin []*Bulletin, err error) {
	var bulletins []*Bulletin
	orm.NewOrm().QueryTable("tb_bulletin").Filter("username", username).RelatedSel().All(&bulletins)
	for _, v := range bulletins {
		logs.Debug("Bulletin ID: %s", v.Id)
	}
	return bulletins, err
}