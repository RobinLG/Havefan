package controllers

import (
	"github.com/astaxie/beego/logs"
	"robin/Havefan/models"
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

		order := &models.Order{Txhash: "7788", Dishes: dishes, Price: price, Location: location, Address: "111", Mobile: "222", Time:time}
		logs.Debug(order.Txhash, order.Dishes, order.Price, order.Location, order.Address, order.Mobile, order.Time)

		err := order.Create()
		if err != nil {
			//c.History("ADD ERROR", "")
			c.Ctx.WriteString("<script>alert('ADD xxx');</script>")
			logs.Debug(err)
		} else {
			c.Ctx.WriteString("<script>alert('ADD SUCCESS');</script>")
		}
	} else if c.Ctx.Request.Method == "GET" {
		////logs.Debug("Insert return: %s", id)
		c.TplName = "publish.html"
	}


}









/*

func (c *UserController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		logs.Debug("Login - username:%s, password:%s", username, password)

		user := &models.User{Username:username}
		// Look for username in "username" cols in database
		user.ReadDB()

		if user.Password == "" {
			c.History("ACCOUNT NOT EXIST", "")
		}

		if util.Md5(password) != strings.Trim(user.Password, " ") {
			c.History("PASSWORD ERROR", "")
		}

		c.SetSession("username", user.Username)
		c.Main()
	} else {
		c.TplName = "login.html"
	}
}

func (c *UserController) Register() {
	username := c.GetString("username")
	password := c.GetString("password")
	nickname := c.GetString("nickname")
	location := c.GetString("location")
	mobile := c.GetString("mobile")
	logs.Debug("Register - username:%s, password:%s, nickname:%s, location:%s, mobile:%s", username, password, nickname, location, mobile)

	c.Data["username"] = username
	c.Data["password"] = util.Md5(password)
	c.Data["nickname"] = nickname
	c.Data["location"] = location
	c.Data["mobile"] = mobile

	// SMS SERVER
	util.Smsmain()

	c.TplName = "smsVerify.html"

}

func (c *UserController) VerifyRegister() {
	username := c.GetString("username")
	password := c.GetString("password")
	nickname := c.GetString("nickname")
	location := c.GetString("location")
	mobile := c.GetString("mobile")
	verification := c.GetString("verification")
	logs.Debug("VerifyRegister - username:%s, password:%s, nickname:%s, location:%s, mobile:%s", username, password, nickname, location, mobile, verification)

	// verify the number of mobile,
	// use the AliSms:
	// The number of SMS messages sent within 1 minute shall not exceed 1;
	// The number of SMS messages sent within 1 hour shall not exceed 5;
	// The number of SMS messages sent within 1 day shall not exceed 10;


	// Dynamic verification code is omitted here
	if verification == "8888"{

		// encrypt password
		user := &models.User{Username:username, Password:password, Nickname:nickname, Location:location, Mobile:mobile, RecoverCode:""}
		logs.Debug(user.Username, user.Password, user.Nickname, user.Location, user.Mobile, user.RecoverCode)

		err := user.Create()
		if err != nil {
			c.History("REGISTER ERROR", "")
			logs.Debug(err)
		} else {
			c.Ctx.WriteString("<script>alert('REGISTER SUCCESS');</script>")
		}
		//logs.Debug("Insert return: %s", id)
		c.TplName = "login.html"
	}
}

// Send active link to email
func (c *UserController)  SendLink() {
	link := util.GenerateActivateLink(fmt.Sprintf("%s", c.GetSession("username")))
	logs.Debug(link)
	util.Smtp(link, c.GetString("email"))
	c.Ctx.WriteString("<script>alert('PLEASE CHECK YOUR EMAIL');window.history.go(-1);</script>")
	// if input msg, session will be cancel
	c.History("", "main.html")
}

func (c *UserController) VerifyLink() {
	username := c.GetString("username")
	checkCode := c.GetString("checkcode")
	logs.Debug("username: %s, code: %s", username, checkCode)

	//berifyCode := md5.Sum([]byte(username + ":" + checkCode))
	//logs.Debug("getCode: %x", berifyCode[:] )

	user := &models.User{Username:username}
	// Look for username in "username" cols in database
	user.ReadDB()

	// Compare the database verification code with the obtained verification code
	if checkCode == user.RecoverCode {
		c.Data["username"] = user.Username
		// if equals, and next step
		c.TplName = "about/reset.html"
	} else {
		c.Ctx.WriteString("<script>alert('PLEASE SEND THE EMAIL AGAIN');window.history.go(-1);</script>")
		// if input msg, session will be cancel
		c.History("", "main.html")
	}
}

func (c *UserController) Reset() {
	logs.Debug("function Reset")
	username := c.GetString("username")
	password := c.GetString("password")
	logs.Debug("username: %s, password: %s", username, password)

	user := &models.User{Username:fmt.Sprintf("%s", username)}

	logs.Debug("Reset username: %s, password: %s", user.Username, user.Password)

	if err := user.ReadDB(); err == orm.ErrNoRows {
		logs.Debug("NOT USER")
		// Actually, here should be done a deal of error. But now, I deal it simplify.
		panic(err)
	} else {
		user.Password = util.Md5(password)

		if err := user.ResetPassword("Password"); err != nil {
			logs.Debug(err)
			// Actually, here should be done a deal of error. But now, I deal it simplify.
			panic(err)
		} else {
			c.Ctx.WriteString("<script>alert('RESET SUCCESS');window.history.go(-1);</script>")
			// if input msg, session will be cancel
			c.History("", "")
		}
	}
}

func (c *UserController) LinkPage() {
	c.TplName = "about/link.html"
}

func (c *UserController) RegisterPage() {
	c.TplName = "create.html"
}

func (c *UserController) Main() {
	c.Redirect("main", 302)
}

func (c *UserController) WatchBulletin() {
	username := c.GetSession("username")
	bulletin := &models.Bulletin{Username:fmt.Sprintf("%s", username)}

	var titles []string
	var contects []string

	if bulletins, err := bulletin.ReadDB(fmt.Sprintf("%s", username)); err != nil {
		logs.Debug(err)
		// Actually, here should be done a deal of error. But now, I deal it simplify.
		panic(err)
	} else {
		logs.Debug("bullentins %s:", bulletins)
		for _, v := range bulletins {
			titles = append(titles, v.Title)
			contects = append(contects, v.Contect)
		}
	}
	c.Data["titles"] = titles
	c.Data["contects"] = contects
	logs.Debug("Bulletin title: %s", titles)

	c.TplName = "bulletin.html"
}
*/