package services

import (
	"log"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		log.Print("DB接続に失敗しました")
		log.Println(err)
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		log.Print("コメント投稿に失敗しました")
		log.Println(err)
		return models.Comment{}, err
	}
	return newComment, nil
}
