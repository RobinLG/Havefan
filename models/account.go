package models

import (
	"github.com/astaxie/beego/orm"
)

type Account struct {
	Id         int    `orm:"pk"`
	Wallet     string `json:"wallet"`
	Useridhash string `json:"useridhash"`
	Mobile     string `json:"mobile"`
}

//return table name without prefix
func (a *Account) TableName() string {
	return TableName("account")
}

func (a *Account) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(a)
	return err
}

func (a *Account) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(a)
	return err
}