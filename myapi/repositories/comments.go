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

// 指定IDの記事についたコメント一覧を取得する関数
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
