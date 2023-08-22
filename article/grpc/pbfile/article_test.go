package pbfile

import (
	"testing"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCheckArticle(t *testing.T) {
	type Request struct {
		Title string
		Author string
		Column string
	}
	req := &Request{
		Title: "New dashboard",
		Author: "root",
		Column: "tech",
	}

	db, err := gorm.Open(mysql.Open("root:Chen@123@tcp(localhost:3306)/myblogs"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}

	var res int64 = 0
	data := db.Where("title = ? and class = ? and author = ?", req.Title, req.Column, req.Author).Table("articles").Count(&res)
	fmt.Println("res:", res)
	fmt.Println("data:", data)
}