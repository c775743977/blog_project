package controller

import (
	"github.com/gin-gonic/gin"

	"myblogs/user/gin/logger"
	user "myblogs/user/gin/pbfile/user"
	"myblogs/user/gin/pool"

	"context"
	"fmt"
	_ "errors"
)

var (
	my_logger = logger.NewLogger()
	User_pool = pool.NewPool("localhost:8000", 10)
	Article_pool = pool.NewPool("localhost:8001", 10)
	verifyCodeError = "rpc error: code = Unknown desc = email-code verification failed!"
	loginFailError = "rpc error: code = Unknown desc = 用户名或密码错误"
	userExistError = "rpc error: code = Unknown desc = 该用户名已被注册"
)

func LoginHandler(c *gin.Context) {
	cookie, _ := c.Cookie("uuid")
	if cookie != "" {
		IndexHandler(c)
	}
	c.HTML(200, "login.html", nil)
}

func IndexHandler(c *gin.Context) {
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		fmt.Println("cookie:", cookie)
		c.HTML(200, "index.html", nil)
		return
	}
	grpcConn, err := User_pool.Get()
	defer grpcConn.Close(User_pool)
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	grpcClient := user.NewUserServiceClient(grpcConn.Value)
	res, _ := grpcClient.GetSession(context.Background(), &user.SessionReq{
		Cookie: cookie,
	})
	
	if res.Name != "" {
		c.HTML(200, "index.html", res.Name)
	} else {
		c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
		c.String(400, "由于长时间没有操作，您的登录信息已过期，请重新登录")
	}
}

func LoginResultHandler(c *gin.Context) {
	cookie, _ := c.Cookie("uuid")
	fmt.Println("cookie:", cookie)
	if cookie != "" {
		IndexHandler(c)
		return
	}


	grpcConn, err := User_pool.Get()
	defer grpcConn.Close(User_pool)
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	grpcClient := user.NewUserServiceClient(grpcConn.Value)
	res, err := grpcClient.Login(context.Background(), &user.LoginReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		my_logger.Warn(c.Request.RemoteAddr, err)
		if err.Error() == loginFailError {
			c.HTML(400, "login.html", "用户名或密码错误")
		} else {
			c.String(500, "服务器内部发生错误，请稍等片刻")
		}
		return
	}
	c.SetCookie("uuid", res.Cookie, 0, "/", "localhost", false, true)
	c.HTML(200, "index.html", username)
}

func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func SendCodeHandler(c *gin.Context) {
	email := c.Query("email")

	grpcConn, err := User_pool.Get()
	defer grpcConn.Close(User_pool)
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		return
	}

	grpcClient := user.NewUserServiceClient(grpcConn.Value)
	_, err = grpcClient.SendCode(context.Background(), &user.SendCodeReq{
		Email: email,
	})
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		return
	}
}

func VerifyCodeHandler(c *gin.Context) {
	repwd := c.PostForm("repasswd")
	data := &user.VerifyCodeReq{}
	c.Bind(data)
	if repwd != data.Password {
		c.HTML(400, "register.html", "两次输入的密码不一致")
		return
	}

	grpcConn, err := User_pool.Get()
	defer grpcConn.Close(User_pool)
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}
	
	grpcClient := user.NewUserServiceClient(grpcConn.Value)
	_, err = grpcClient.VerifyCode(context.Background(), data)
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		if err.Error() == verifyCodeError {
			c.HTML(400, "register.html", "email-code verification failed!")
		} else if err.Error() == userExistError {
			c.HTML(400, "register.html", "该用户名已存在")
		} else {
			c.String(500, "服务器内部发生错误，请稍等片刻")
		}
		return
	}
	c.HTML(200, "register_success.html", nil)
}

func LogoutHandler(c *gin.Context) {
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		IndexHandler(c)
		return
	}

	c.SetCookie("uuid", "", -1 , "/", "localhost", false, true)

	grpcConn, err := User_pool.Get()
	defer grpcConn.Close(User_pool)
	if err != nil {
		fmt.Println("get conn from pool error:", err)
		my_logger.Error(c.Request.RemoteAddr, err)
	}

	grpcClient := user.NewUserServiceClient(grpcConn.Value)
	_, err = grpcClient.Logout(context.Background(), &user.LogoutReq{
		Cookie: cookie,
	})
	if err != nil {
		IndexHandler(c)
		return
	} else {
		c.String(500, "服务器内部发生错误，请稍等片刻")
		my_logger.Error(c.Request.RemoteAddr, err)
	}
}