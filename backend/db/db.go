package db

import (
	"database/sql"
	"fmt"
	"log"

	// "os"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

// DB 各repositoryで利用するDB接続情報
var DB *sql.DB

func init() {
	/* ===== データベースへ接続する. ===== */
	// ユーザ
	user := "user"
	// パスワード
	password := "mysql-yuki9015"
	//// 接続先ホスト
	//host := "localhost"
	//host := "vocabulary-book_mysql"
	//// 接続先ポート
	//port := "3306"
	// 接続先データベース
	database := "vocabulary_book_api"
	// user:password@tcp(host:port)/database
	var err error
	DB, err = sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, password, database))
	if err != nil {
		log.Fatal(err)
	}
}