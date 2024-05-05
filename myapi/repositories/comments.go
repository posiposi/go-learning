package repositories

import (
	"database/sql"

	"github.com/yourname/reponame/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	// SQL構築
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	// 構造体変数定義
	var newComment models.Comment
	// httpリクエストボディから構造体のフィールドに値を代入
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message
	// クエリ実行
	result, err := db.Exec(sqlStr, newComment.ArticleID, newComment.Message)
	if err != nil {
		return models.Comment{}, err
	}
	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)
	return newComment, nil
}
