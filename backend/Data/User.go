package Data

type User struct {
	Account  int    `json:"acc"`
	Password string `json:"pwd"`
}

var AllUsers []User
