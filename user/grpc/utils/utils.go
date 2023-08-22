package utils

import (
	"net/smtp"
	"fmt"
	"time"
	math "math/rand"
	"crypto/rand"

	"myblogs/user/grpc/db"
	"myblogs/user/grpc/model"

	"github.com/jordan-wright/email"
)

func SendMail(code int, emailAddrs string) error {
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "775743977@qq.com"
	 
	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{emailAddrs}
	 
	// 设置主题
	em.Subject = "小陈博客"
	 
	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("欢迎注册小陈博客，验证码:" + fmt.Sprint(code) + "(3分钟后过期)")
	 
	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "775743977@qq.com", "arliannbrbsmbcbd", "smtp.qq.com"))
	if err != nil {
	   return err
	}
	fmt.Println("send successfully ... ")
	return nil
}

func CreateCode() int {
	math.Seed(time.Now().UnixNano())

    num := math.Intn(900000) + 100000

	return num
}

func CheckUsername(username string) bool {
	DB := db.MDB.Where("username = ?", username).Find(&model.User{})
	if DB.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

//CreateUUID 生成UUID
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		fmt.Println("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}