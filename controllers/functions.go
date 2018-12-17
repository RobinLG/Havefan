package controllers

import (
	"github.com/astaxie/beego/logs"
	"robin/Havefan/models"
	"robin/Havefan/util"
	"strconv"
)

type UserController struct {
	baseController
}

func (c *UserController) AddOrder() {
	if c.Ctx.Request.Method == "POST" {
		//txhash := c.GetString("txhash")
		dishes := c.GetString("dishes")
		price := c.GetString("price")
		location := c.GetString("location")
		//address := c.GetString("address")
		//mobile := c.GetString("mobile")
		time := c.GetString("time")

		order := &models.Order{Dishes: dishes, Price: price, Location: location, Address: "111", Mobile: "222", Time:time, Flag:"0"}
		logs.Debug(order.Dishes, order.Price, order.Location, order.Address, order.Mobile, order.Time)

		err := order.Create()
		if err != nil {
			//c.History("ADD ERROR", "")
			c.Ctx.WriteString("<script>alert('ERROR');</script>")
			logs.Debug(err)
		} else {
			c.Ctx.WriteString("<script>alert('ADD SUCCESS');</script>")
		}
		c.TplName = "index.html"
	} else if c.Ctx.Request.Method == "GET" {
		////logs.Debug("Insert return: %s", id)
		c.TplName = "publish.html"
	}
}

func (c *UserController) ShowIndex() {

	orders := &models.Order{}

	var addresses []string;
	var flags []string;
	var prices []string;
	var length int;

	if messages, err := orders.ReadDB(); err != nil {
		logs.Debug(err)
		// Actually, here should be done a deal of error. But now, I deal it simplify.
		panic(err)
	} else {
		logs.Debug("showindex messages:%s", messages)
		for _, v := range messages {
			addresses = append(addresses, v.Location)
			flags = append(flags, v.Flag)
			prices = append(prices, v.Price)
		}
		length = len(addresses)
	}
	c.Data["addresses"] = addresses
	c.Data["flags"] = flags
	c.Data["prices"] = prices
	c.Data["length"] = length

	c.TplName = "index.html"
}

func (c *UserController) ShowDetail() {

	txhash_ := c.GetString("txhash")
	txhash, err := strconv.Atoi(txhash_)
	if err != nil {
		panic(err)
	}
	logs.Debug("showdetail-txhash:%s", txhash)
	order := &models.Order{}

	if message, err := order.ReadDBOne(txhash); err != nil {
		logs.Debug(err)
		// Actually, here should be done a deal of error. But now, I deal it simplify.
		panic(err)
	} else {
		logs.Debug("showindex message:%s", message[0])
		c.Data["address"] = message[0].Address
		c.Data["time"] = message[0].Time
		c.Data["location"] = message[0].Location
		c.Data["dishes"] = message[0].Dishes
		c.Data["price"] = message[0].Price
		c.Data["state"] = message[0].Flag
		c.Data["id"] = message[0].Txhash
	}
	c.TplName = "detail.html"
}

func (c *UserController) ShowCheck() {
	c.TplName = "check.html"
}

func (c *UserController) Account() {
	if c.Ctx.Request.Method == "POST" {
		wallet   := c.GetString("wallet")
		UserIdHash := c.GetString("useridhash")
		mobile   := c.GetString("mobile")
		verification := c.GetString("verification")
		logs.Debug("Account - wallet:%s, UserIdHash:%s, mobile:%s, verification:%s", wallet, UserIdHash, mobile, verification)

		// Dynamic verification code is omitted here
		if verification == "8888"{

			account := &models.Account{Wallet:wallet, Useridhash:UserIdHash, Mobile:mobile}
			logs.Debug("wallet:%s, useridhash:%s, mobile:%s", account.Wallet, account.Useridhash, account.Mobile)

			err := account.Create()
			if err != nil {
				//c.History("REGISTER ERROR", "")
				c.Ctx.WriteString("<script>alert('ERROR');</script>")
				logs.Debug(err)
			} else {
				c.Ctx.WriteString("<script>alert('REGISTER SUCCESS');</script>")
			}
			//logs.Debug("Insert return: %s", id)
			c.TplName = "index.html"
		}


	} else if c.Ctx.Request.Method == "GET" {
		c.TplName = "account.html"
	}
}

func (c *UserController) Check() {
	mobile   := c.GetString("mobile")
	logs.Debug("Check - mobile:%s", mobile)

	c.Data["mobile"] = mobile

	util.Smsmain(mobile)

	c.TplName = "account.html"
}

func (c *UserController) UpdateFlagToOne() {
	txhashstring := c.GetString("txhash")
	txhash, err := strconv.Atoi(txhashstring)

	flag := c.GetString("flag")

	logs.Debug("flag:%s", txhash)
	//txhash, err := strconv.Atoi("1")
	if err != nil {
		panic(err)
	}

	order := &models.Order{Txhash:txhash}
	order.Flag = flag
	order.UpdateFlag("flag")

	c.Redirect("index.html", 302)
}
