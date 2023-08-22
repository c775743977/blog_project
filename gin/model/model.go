package model

import (
	article "myblogs/user/gin/pbfile/article"
)

type UserError struct {
	Username string
	Error string
}

type UserArticles struct {
	Username string
	Articles []*article.Article
}

type ColumnError struct {
	Column string
	Error string
	Title string
}

type ContentTitle struct {
	Error string
	Title string
	Content string
	Column string
	ID int64
}