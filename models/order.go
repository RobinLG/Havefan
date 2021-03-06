package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type Order struct {
	Txhash   int `orm:"pk"`
	Dishes   string `json:"dishes"`
	Price    string `json:"price"`
	Location string `json:"location"`
	Address  string `json:"address"`
	Mobile   string `json:"mobile"`
	Time     string `json:"time"`
	Flag     string `json:"flag"`
}

//return table name without prefix
func (r *Order) TableName() string {
	return TableName("order")
}

func (r *Order) ReadDB() (message []*Order, err error) {
	var messages []*Order
	orm.NewOrm().QueryTable("tb_order").RelatedSel().All(&messages)
	logs.Debug(messages)
	return messages, err
}

func (r *Order) ReadDBOne(txhash int) (messageone []*Order, err error) {
	var message []*Order
	orm.NewOrm().QueryTable("tb_order").Filter("txhash", txhash).RelatedSel().All(&message)
	logs.Debug(message)
	return message, err
}

func (r *Order) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(r)
	return err
}

func (r *Order) UpdateFlag(flag string) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(r, flag)
	return err
}