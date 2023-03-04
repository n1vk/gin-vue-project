package Data

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// AllUsers 数据结构应当在连接数据库后被移除
var AllUsers []User
