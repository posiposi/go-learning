package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

// InsertComment関数のテスト
func TestInsertComment(t *testing.T) {
	newComment := models.Comment{
		ArticleID: 1,
		Message:   "This is a comment",
	}
	// CommentIDの現行最大値は2
	expectedCommentID := 3
	result, err := repositories.InsertComment(testDB, newComment)
	// DB接続自体のエラー
	if err != nil {
		t.Error(err)
	}
	// アサーション実行
	if result.ArticleID != newComment.ArticleID {
		t.Errorf("want %d but got %d\n", newComment.ArticleID, result.ArticleID)
	}
	if result.Message != newComment.Message {
		t.Errorf("want %s but got %s\n", newComment.Message, result.Message)
	}
	if result.CommentID != expectedCommentID {
		t.Errorf("want %d but got %d\n", expectedCommentID, result.CommentID)
	}
	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, newComment.Message)
	})
}
