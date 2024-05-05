package repositories

import (
	"database/sql"
	"fmt"

	"github.com/yourname/reponame/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	var sqlStr = `
	INSET INTO articles (title, contents, user_name, nice_num, created_at)
	VALUES (?, ?, ?, ?, ?)
	`

	// 構造体変数を定義
	var newArticle models.Article
	// httpリクエストボディから構造体のフィールドに値を代入
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	// 構造体のフィールドからデータベースに値を挿入
	result, err := db.Exec(sqlStr, newArticle.Title, newArticle.Contents, newArticle.UserName)
	if err != nil {
		return models.Article{}, fmt.Errorf("failed to insert article: %w", err)
	}

	// データ挿入後のIDを取得し、構造体のIDフィールドに代入
	// ここまで到達しているということはDBへのデータ投入は成功しているので、errorは変数に存在し得ない
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)

	return newArticle, nil
}
