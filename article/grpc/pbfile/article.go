package pbfile

import (
	"context"
	"fmt"

	"grpc_article/db"
)

type ArticleService struct {

}

func(this *ArticleService) mustEmbedUnimplementedArticleServiceServer() {}

func(this *ArticleService) SaveArticleInfo(ctx context.Context, req *Article) (*ArticleRes, error) {
	err := db.MDB.Create(req).Error
	if err != nil {
		fmt.Println("mysql create article info error:", err)
		return &ArticleRes{}, err
	}

	return &ArticleRes{
		Id: req.ID,
	}, nil
}

func(this *ArticleService) GetArticlesInfo(ctx context.Context, req *ArticlesReq) (*ArticlesRes, error) {
	var articles []*Article
	err := db.MDB.Where("class = ?", req.Column).Find(&articles).Error
	if err != nil {
		fmt.Println("mysql find articles info error:", err)
		return &ArticlesRes{}, err
	}

	return &ArticlesRes{
		Articles: articles,
	}, nil
}

func(this *ArticleService) GetArticleInfo(ctx context.Context, req *ArticleReq) (*ArticlesRes, error) {
	var articles []*Article
	err := db.MDB.Where("title = ? AND class = ?", req.Title, req.Column).Find(&articles).Error
	if err != nil {
		fmt.Println("mysql find article info error:", err)
		return &ArticlesRes{}, err
	}

	return &ArticlesRes{
		Articles: articles,
	}, nil
}

func(this *ArticleService) CheckArtileExist(ctx context.Context, req *CheckReq) (*CheckRes, error) {
	var res int64
	db.MDB.Where("title = ? and class = ? and author = ?", req.Title, req.Column, req.Author).Table("articles").Count(&res)
	if res == 0 {
		return &CheckRes{
			Exist: false,
		}, nil
	} else {
		return &CheckRes{
			Exist: true,
		}, nil
	}
}

func(this *ArticleService) GetArtilesByAuthor(ctx context.Context, req *AuthorReq) (*ArticlesRes, error) {
	var articles []*Article
	err := db.MDB.Where("author = ?", req.Author).Find(&articles).Error
	if err != nil {
		fmt.Println("mysql find article info error:", err)
		return &ArticlesRes{}, err
	}

	return &ArticlesRes{
		Articles: articles,
	}, nil
}

func(this *ArticleService) DelArticle(ctx context.Context, req *IDReq) (*IDReq, error) {
	err := db.MDB.Where("id = ?", req.ID).Delete(&Article{}).Error
	if err != nil {
		fmt.Println("mysql delete article error:", err)
		return &IDReq{}, err
	} else {
		return &IDReq{}, nil
	}
}