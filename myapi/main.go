package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	// article_id = 1記事のいいねを取得
	selectQuery := `
		SELECT nice
		FROM articles
		WHERE article_id = ?
		`
	targetArticleId := 3
	row := tx.QueryRow(selectQuery, targetArticleId)
	// 記事が取得できないパターンはロールバックする
	if row.Err() != nil {
		fmt.Println(row.Err())
		tx.Rollback()
		return
	}

	// いいね数読み込み
	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// いいね数をインクリメントし、更新実行
	updateQuery := `
		UPDATE articles
		SET nice = ?
		WHERE article_id = ?
		`
	_, err = tx.Exec(updateQuery, niceNum+1, targetArticleId)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	// コミット実行
	tx.Commit()
	fmt.Printf("nice num: %d\n", niceNum+1)
}
