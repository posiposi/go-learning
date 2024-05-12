package services

import (
	"log"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// articleIDを元に、記事データを取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 記事に紐づくコメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 記事データにコメント一覧をセット
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// 指定ページの記事一覧を取得
	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

func PostNiceService(articleID int) (models.Article, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// いいね数+1処理実行
	err = repositories.UpdateNiceNum(db, articleID)
	if err != nil {
		log.Print("いいね数の更新に失敗しました")
		log.Println(err)
		return models.Article{}, err
	}
	//　いいね数を増分した記事データをIDで取得する
	updatedArticle, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		log.Print("記事データの取得に失敗しました")
		log.Println(err)
		return models.Article{}, err
	}

	return models.Article{
		ID:        updatedArticle.ID,
		Title:     updatedArticle.Title,
		Contents:  updatedArticle.Contents,
		UserName:  updatedArticle.UserName,
		NiceNum:   updatedArticle.NiceNum,
		CreatedAt: updatedArticle.CreatedAt,
	}, nil
}
