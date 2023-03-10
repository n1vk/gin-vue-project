package LoginSystem

import "database/sql"

// TODO: 这个部分的查询应当与数据库连接

func IsExist(db *sql.DB, acc string) bool {
	//if len(Data.AllUsers) == 0 {
	//	return false
	//} else {
	//	for _, u := range Data.AllUsers {
	//		if u.Account == acc {
	//			return true
	//		}
	//	}
	//}

	rows, err := db.Query("SELECT * FROM users WHERE account=?", acc)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

func IsCorrect(db *sql.DB, acc string, pwd string) bool {

	rows, err := db.Query("SELECT * FROM users WHERE account=? AND password= ?", acc, pwd)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}

	//for _, u := range Data.AllUsers {
	//	if u.Account == acc {
	//		if u.Password == pwd {
	//			return true
	//		}
	//	}
	//}

	return false
}

func AddUser(db *sql.DB, acc string, pwd string) {
	//var usr Data.User
	//usr.Account = acc
	//usr.Password = pwd
	//Data.AllUsers = append(Data.AllUsers, usr)

	statement, err := db.Prepare("INSERT INTO users (account, password) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()

	_, err = statement.Exec(acc, pwd)
	if err != nil {
		panic(err.Error())
	}

	return
}

func DelUser(acc string, pwd string) {
	// pass
}
