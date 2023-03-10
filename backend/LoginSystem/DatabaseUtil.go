package LoginSystem

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	// 数据库配置 MySQL 8.0.32
	// 账户    : "root"
	// 密码    : "mysqlPassword"
	// 端口默认 : "3306"
	// 数据库   : "gin"
	// URL     : jdbc:mysql://localhost:3306/gin

	db, err := sql.Open("mysql", "root:mysqlPassword@tcp(localhost:3306)/gin")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("SELECT 1 FROM users LIMIT 1")
	if err != nil {
		// 没有users表
		createTable := `CREATE TABLE users (
    			id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    			account VARCHAR(32) NOT NULL,
    			password VARCHAR(256) NOT NULL
            )`

		_, err = db.Exec(createTable)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Created table 'users' ")

	} else {
		fmt.Println("'users' table exists")
	}

	return db
}
