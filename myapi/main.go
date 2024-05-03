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

	dbQuery := `
		SELECT *
		FROM articles
		WHERE article_id = ?
	`

	articleId := 3
	row := db.QueryRow(dbQuery, articleId)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// 構造体用スライスを用意
	articleArray := make([]models.Article, 0)
	// 構造体にループ処理で投入するデータ用変数を用意
	var article models.Article
	// NULL許容カラム定義
	var createdTime sql.NullTime
	// Scanによるデータ取得
	err = row.Scan(
		&article.ID,
		&article.Title,
		&article.Contents,
		&article.UserName,
		&article.NiceNum,
		&createdTime)

	// データ生成日がNULLであるかの確認(NULLでなければtrue)
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	if err != nil {
		fmt.Println(err)
		return
	} else {
		// スライスにデータを追加
		articleArray = append(articleArray, article)
	}

	// スライスの内容を出力する
	fmt.Printf("%+v\n", articleArray)
}
