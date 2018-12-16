package util

import (
	"io/ioutil"
	"encoding/json"
	"strings"
	"net/url"
	"fmt"
	"math/rand"
	"time"
	"sort"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"errors"
	"github.com/astaxie/beego/logs"
)

// SendSmsReply
type SendSmsReply struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

// especial var change
func replace(in string) string {
	rep := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	// function QueryEscape transcoding for "in", It can be used safely in URL queries
	return rep.Replace(url.QueryEscape(in))
}

// SendSms
func SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode string) (string, error) {
	paras := map[string]string{
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureNonce":   fmt.Sprintf("%d", rand.Int63()),
		"AccessKeyId":      accessKeyID,
		"SignatureVersion": "1.0",
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           "JSON",

		"Action":        "SendSms",
		"Version":       "2017-05-25",
		"RegionId":      "cn-hangzhou",
		"PhoneNumbers":  phoneNumbers,
		"SignName":      signName,
		"TemplateParam": templateParam,
		"TemplateCode":  templateCode,
	}

	var keys []string

	for k := range paras {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	logs.Debug("keys:", keys)
	var sortQueryString string

	for _, v := range keys {
		sortQueryString = fmt.Sprintf("%s&%s=%s", sortQueryString, replace(v), replace(paras[v]))
	}
	logs.Debug("sortQueryString: %s", sortQueryString)

	stringToSign := fmt.Sprintf("GET&%s&%s", replace("/"), replace(sortQueryString[1:]))
	logs.Debug("stringToSign: %s", stringToSign)

	// encrypt secret of Ali account
	mac := hmac.New(sha1.New, []byte(fmt.Sprintf("%s&", accessSecret)))
	mac.Write([]byte(stringToSign))
	sign := replace(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	str := fmt.Sprintf("http://dysmsapi.aliyuncs.com/?Signature=%s%s", sign, sortQueryString)
	logs.Debug(str)

	resp, err := http.Get(str)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ssr := &SendSmsReply{}

	// replace character
	if err := json.Unmarshal(body, ssr); err != nil {
		return "", err
	}

	// send failure, again. This function is executed the last time a message was sent.
	if ssr.Code == "SignatureNonceUsed" {
		return SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode)
	} else if ssr.Code != "OK" {
		return "", errors.New(ssr.Code)
	}
	fmt.Printf("statuCdoe:%v  body: %v \n",resp.StatusCode,string(body) )
	return ssr.Code, nil
}

func Smsmain(number string) string {
	// Your Ali account accessKeyID
	accessKeyID:="***"
	// Your Ali account accessSecret
	accessSecret:="***"
	// Registered mobile phone number
	phoneNumbers:=number
	// The sign name of Your Ali account (SMS service)
	signName:="Robin"
	// replace message template variable with json
	templateParam:=`{"code":"8888"}`
	// message template ID, check in Ali Cloud. This about where can you send the message, it's different mainland with Hong Kong/Macao/Taiwan.
	templateCode:="SMS_149101330"

	if code, err:=SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode);err!=nil{
		fmt.Println(err)
	} else {
		return code
	}
	return "ERR"
}
