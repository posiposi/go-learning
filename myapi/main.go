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
		SELECT title, contents, username, nice
		FROM articles
	`

	rows, err := db.Query(dbQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 遅延実行でクエリクローズ
	defer rows.Close()

	// 構造体用スライスを用意
	articleArray := make([]models.Article, 0)
	for rows.Next() {
		// 構造体にループ処理で投入するデータ用変数を用意
		var article models.Article
		// Scanによるデータ取得
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			// スライスにデータを追加
			articleArray = append(articleArray, article)
		}
	}

	// スライスの内容を出力する
	fmt.Printf("%+v\n", articleArray)
}
