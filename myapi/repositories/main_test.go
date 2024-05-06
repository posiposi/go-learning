package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// 前後処理で共通仕様のDB変数を定義
// 変数dbに代入するとtearDown()での使用時にコンパイルエラーとなるため
var testDB *sql.DB

// テストの前処理
func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// テストの後処理
func tearDown() {
	testDB.Close()
}

// テスト実行関数定義
func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		// testing.MはFatal系メソッドが存在しないので、エラー時はExitで終了させる
		os.Exit(1)
	}

	// テスト実行
	m.Run()

	// 後処理実行
	tearDown()
}
