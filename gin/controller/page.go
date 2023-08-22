package controller

import (
	"github.com/gin-gonic/gin"

	user "myblogs/user/gin/pbfile/user"
	"myblogs/user/gin/model"

	"context"
	"errors"
)

func EditPageHandler(c *gin.Context) {
	_, err := CheckSession(c)
	if err != nil {
		return
	}
	column := c.Query("column")
	c.HTML(200, "edit_article.html", &model.ContentTitle{
		Column: column,
	})
}

func TechClassHandler(c *gin.Context) {
	username, err := CheckSession(c)
	if err != nil {
		return
	}

	articles, err := GetArticlesByColumn("tech")
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	if username != "" {
		c.HTML(200, "tech_class.html", model.UserArticles{
			Username: username,
			Articles: articles,
		})
	} else {
		c.String(400, "由于长时间没有操作，您的登录信息已过期，请重新登录")
		c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
	}
}

func CheckSession(c *gin.Context) (string, error) {
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(500, "需要登录才有权限访问此内容")
		return "", errors.New("500")
	}

	grpcConn, err := User_pool.Get()
	defer grpcConn.Close(User_pool)
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return "", errors.New("500")
	}

	grpcClient := user.NewUserServiceClient(grpcConn.Value)
	res, err := grpcClient.GetSession(context.Background(), &user.SessionReq{
		Cookie: cookie,
	})
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return "", errors.New("500")
	}

	return res.Name, nil
}

func EntertainmentClassHandler(c *gin.Context) {
	username, err := CheckSession(c)
	if err != nil {
		return
	}

	articles, err := GetArticlesByColumn("entertainment")
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	if username != "" {
		c.HTML(200, "entertainment_class.html", model.UserArticles{
			Username: username,
			Articles: articles,
		})
	} else {
		c.String(400, "由于长时间没有操作，您的登录信息已过期，请重新登录")
		c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
	}
}

func ExchangeClassHandler(c *gin.Context) {
	username, err := CheckSession(c)
	if err != nil {
		return
	}

	articles, err := GetArticlesByColumn("exchange")
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	if username != "" {
		c.HTML(200, "exchange_class.html", model.UserArticles{
			Username: username,
			Articles: articles,
		})
	} else {
		c.String(400, "由于长时间没有操作，您的登录信息已过期，请重新登录")
		c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
	}
}

func MyPageHandler(c *gin.Context) {
	username, err := CheckSession(c)
	if err != nil {
		return
	}

	articles, err := GetArticlesByAuthor(username)
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	if username != "" {
		c.HTML(200, "personal_home.html", model.UserArticles{
			Username: username,
			Articles: articles,
		})
	} else {
		c.String(400, "由于长时间没有操作，您的登录信息已过期，请重新登录")
		c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
	}
}