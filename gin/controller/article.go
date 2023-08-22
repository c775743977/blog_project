package controller

import (
	"github.com/gin-gonic/gin"

	article "myblogs/user/gin/pbfile/article"
	"myblogs/user/gin/utils"
	"myblogs/user/gin/model"

	"context"
	"fmt"
	"errors"
	"strconv"
)

func GetArticlesByAuthor(author string) ([]*article.Article, error) {
	grpcConn, err := Article_pool.Get()
	defer grpcConn.Close(Article_pool)
	if err != nil {
		return []*article.Article{}, err
	}
	grpcClient := article.NewArticleServiceClient(grpcConn.Value)
	res, err := grpcClient.GetArtilesByAuthor(context.Background(), &article.AuthorReq{
		Author: author,
	})
	if err != nil {
		return []*article.Article{}, err
	}

	return res.Articles, nil
}

func GetArticlesByColumn(column string) ([]*article.Article, error) {
	grpcConn, err := Article_pool.Get()
	defer grpcConn.Close(Article_pool)
	if err != nil {
		return []*article.Article{}, err
	}
	grpcClient := article.NewArticleServiceClient(grpcConn.Value)
	res, err := grpcClient.GetArticlesInfo(context.Background(), &article.ArticlesReq{
		Column: column,
	})
	if err != nil {
		return []*article.Article{}, err
	}

	return res.Articles, nil
}

func CheckArtileExist(c *gin.Context, column string, author string, title string) (bool, error) {
	grpcConn, err := Article_pool.Get()
	defer grpcConn.Close(Article_pool)
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return false, errors.New("500")
	}

	grpcClient := article.NewArticleServiceClient(grpcConn.Value)
	res, _ := grpcClient.CheckArtileExist(context.Background(), &article.CheckReq{
		Title: title,
		Author: author,
		Column: column,
	})
	return res.Exist, nil
}

func SaveTechArtile(c *gin.Context, art *article.Article, content string, articleID int64) {
	grpcConn, err := Article_pool.Get()
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	if articleID != 0 {
		err = utils.SaveHTMLFile(art, content, articleID)
		if err != nil {
			my_logger.Error(c.Request.RemoteAddr, err)
			c.String(500, "服务器内部发生错误，请稍等片刻")
			return
		}
	} else {
		grpcClient := article.NewArticleServiceClient(grpcConn.Value)
		res, err := grpcClient.SaveArticleInfo(context.Background(), art)
		if err != nil {
			my_logger.Error(c.Request.RemoteAddr, err)
			c.String(500, "服务器内部发生错误，请稍等片刻")
			return
		}
		err = utils.SaveHTMLFile(art, content, res.Id)
		if err != nil {
			my_logger.Error(c.Request.RemoteAddr, err)
			c.String(500, "服务器内部发生错误，请稍等片刻")
			return
		}
	}

	c.HTML(200, "publish_success.html", nil)
}


func NewArticleHandler(c *gin.Context) {
	username, err := CheckSession(c)
	if err != nil {
		return
	}
	column := c.Query("column")
	data := c.PostForm("editor-content")
	title := c.PostForm("title")
	articleID, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if articleID != 0 {
		articleReq := &article.Article{
			Column: column,
			Author: username,
			Title: title,
			Date: utils.GetTime(),
		}
		SaveTechArtile(c, articleReq, data, articleID)
		return
	}
	fmt.Println("articleID:", articleID)
	exist, err := CheckArtileExist(c, column, username, title)
	if err != nil {
		return
	}



	if exist {
		data := &model.ContentTitle{
			Error: "已在该板块发表过相关文章",
		}
		c.HTML(304, "edit_article.html", data)
	} else {
		articleReq := &article.Article{
			Column: column,
			Author: username,
			Title: title,
			Date: utils.GetTime(),
		}
		SaveTechArtile(c, articleReq, data, articleID)
	}
}

func DelArticleHandler(c *gin.Context) {
	articleID, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	grpcConn, err := Article_pool.Get()
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	grpcClient := article.NewArticleServiceClient(grpcConn.Value)
	_, err = grpcClient.DelArticle(context.Background(), &article.IDReq{
		ID: articleID,
	})
	if err != nil {
		my_logger.Error(c.Request.RemoteAddr, err)
		c.String(500, "服务器内部发生错误，请稍等片刻")
		return
	}

	MyPageHandler(c)
}

func AlterArticleHandler(c *gin.Context) {
	articleID, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	column := c.Query("column")

	title, content := utils.GetArticleContent(column, articleID)

	data := &model.ContentTitle{
		Error: "",
		Content: content,
		Title: title,
		Column: column,
		ID: articleID,
	}

	c.HTML(200, "edit_article.html", data)
}