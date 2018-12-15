package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Username string `orm:"pk"`
	Nickname string `json:"nickname"`
	Location string `json:"location"`
	Password string `json:"password"`
	Mobile string `json:"mobile"`
	// When user need to reset password, recoverCode will br put in database. If CHECK_CODE of activate link equals recoverCode, user reset password successful.
	RecoverCode string `json:"recover_code"`
}

//return table name without prefix
func (u *User) TableName() string {
	return TableName("user")
}

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return err
}

func (u *User) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(u)
	return err
}

func (u *User) ResetPassword(password string) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(u, password)
	return err
}