package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"slotsevo-admin/global"

	"github.com/jordan-wright/email"
)

//@author: xiaoliang
//@function: Email
//@description: Email发送方法
//@param: subject string, body string
//@return: error
func Email(subject string, body string) error {
	to := strings.Split(global.SlotsConfig.Email.To, ",")
	return send(to, subject, body)
}

//@author: xiaoliang
//@function: ErrorToEmail
//@description: 给email中间件错误发送邮件到指定邮箱
//@param: subject string, body string
//@return: error

func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.SlotsConfig.Email.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

//@author: xiaoliang
//@function: EmailTest
//@description: Email测试方法
//@param: subject string, body string
//@return: error
func EmailTest(subject string, body string) error {
	to := []string{global.SlotsConfig.Email.From}
	return send(to, subject, body)
}

//@author: xiaoliang
//@function: send
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func send(to []string, subject string, body string) error {
	from := global.SlotsConfig.Email.From
	nickname := global.SlotsConfig.Email.Nickname
	secret := global.SlotsConfig.Email.Secret
	host := global.SlotsConfig.Email.Host
	port := global.SlotsConfig.Email.Port
	isSSL := global.SlotsConfig.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
