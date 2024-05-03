package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/models"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// 挿入データ生成
	article := models.Article{
		Title:    "test",
		Contents: "test contents",
		UserName: "sugiyama",
	}

	dbQuery := `
		INSERT INTO articles (title, contents, username, nice, created_at)
		VALUES (?, ?, ?, 0, now())
	`

	result, err := db.Exec(dbQuery, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
