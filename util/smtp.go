package util

import (
	"encoding/base64"
	"net/mail"
	"fmt"
	"net/smtp"
	"crypto/md5"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"time"
	"robin/unity/models"
)


const CHECK_CODE = "checkcode"

// Open your email SMTP server first

func Smtp(url string, _email string) {

	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	// I choose "网易"mailbox, so that is "smtp.163.com". That's different with QQ mailbox and so on. Other host of mailbox that you can check it by internet.
	host := "smtp.163.com"
	// Sender email address
	email := "***"
	// Is the authentication password for the mailbox, not the login password
	password := "***"
	toEmail := _email
	// The sender's name
	from := mail.Address{"***", email}
	// The recipient's name
	to := mail.Address{"***", toEmail}
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte("Reset password")))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	body := "Please check the url to reset your password: " + url
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)
	err := smtp.SendMail(
		host+":25",
		auth,
		email,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		panic(err)
	}
}

func GenerateActivateLink(username string,) string {
	return "http://localhost:8080/activate?username=" + username + "&" + CHECK_CODE + "=" + GenerateCheckcode(username)
}

// Generate check code, username + UUID. For security reasons, send them encrypted
func GenerateCheckcode(username string) string {
	// Here should be set a random number, but I'm going to simplify the process
	rand.Seed(time.Now().Unix())
	randnum := rand.Intn(1000000)

	data := []byte(username + ":" + string(randnum))
	logs.Debug(fmt.Sprintf("%x", md5.Sum(data)))

	user := &models.User{Username:username}
	user.RecoverCode = fmt.Sprintf("%x", md5.Sum(data))
	logs.Debug("GenerateCheckcode username:%s, code: %s", user.Username, user.RecoverCode)
	if err := user.ResetPassword("RecoverCode"); err != nil {
		logs.Debug(err)
		// Actually, here should be done a deal of error. But now, I deal it simplify.
		panic(err)
	}
	return fmt.Sprintf("%x", md5.Sum(data))
}