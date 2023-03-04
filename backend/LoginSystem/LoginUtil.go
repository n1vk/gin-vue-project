package LoginSystem

import "backend/Data"

// TODO: 这个部分的查询应当与数据库连接

func IsExist(acc string) bool {
	if len(Data.AllUsers) == 0 {
		return false
	} else {
		for _, u := range Data.AllUsers {
			if u.Account == acc {
				return true
			}
		}
	}

	return false
}

func IsCorrect(acc string, pwd string) bool {
	for _, u := range Data.AllUsers {
		if u.Account == acc {
			if u.Password == pwd {
				return true
			}
		}
	}

	return false
}

func AddUser(acc string, pwd string) {
	var usr Data.User
	usr.Account = acc
	usr.Password = pwd
	Data.AllUsers = append(Data.AllUsers, usr)
}

func DelUser(acc string, pwd string) {
	// pass
}
