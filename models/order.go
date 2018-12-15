package models

import (
	"github.com/astaxie/beego/orm"
)

type Order struct {
	Txhash   string `orm:"pk"`
	Dishes   string `json:"dishes"`
	Price    string `json:"price"`
	Location string `json:"location"`
	Address  string `json:"address"`
	Mobile   string `json:"mobile"`
	Time     string `json:"time"`
}

//return table name without prefix
func (r *Order) TableName() string {
	return TableName("order")
}

func (r *Order) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(r)
	return err
}

func (r *Order) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(r)
	return err
}