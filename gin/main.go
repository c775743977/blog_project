package main

import (
	"github.com/gin-gonic/gin"

	_"fmt"

	"myblogs/user/gin/controller"
)


func main() {
	// 关闭grpc连接池
	defer controller.User_pool.Close()
	defer controller.Article_pool.Close()

	router := gin.Default()
	router.LoadHTMLGlob("pages/*")
	// router.LoadHTMLGlob("articles/**/*")

	// 这里表示访问/pages/***的静态资源时去./pages下的匹配。 比如我有图片123.png放在./img下，那么第二个参数就应该是"./img"
	// 然后在前端代码中，读取123.png的href是"/images/123.png"，那么第一个参数就应该是"/images"
	router.Static("/pages", "./pages")

	router.Static("/articles", "./articles")

	router.GET("/login", controller.LoginHandler)
	router.POST("/login", controller.LoginResultHandler)

	router.GET("/register", controller.Register)

	router.GET("/sendCode", controller.SendCodeHandler)

	router.POST("/verifyCode", controller.VerifyCodeHandler)

	router.GET("/index", controller.IndexHandler)

	router.GET("/logout", controller.LogoutHandler)

	router.GET("/tech", controller.TechClassHandler)

	router.GET("/entertainment", controller.EntertainmentClassHandler)

	router.GET("exchange", controller.ExchangeClassHandler)

	router.GET("/tech/newArticle", controller.EditPageHandler)

	router.GET("/entertainment/newArticle", controller.EditPageHandler)

	router.GET("/exchange/newArticle", controller.EditPageHandler)

	router.POST("/edit_article", controller.NewArticleHandler)

	router.GET("/myPage", controller.MyPageHandler)

	router.GET("/delete_article", controller.DelArticleHandler)

	router.GET("/alter_article", controller.AlterArticleHandler)

	router.Run(":8080")
}