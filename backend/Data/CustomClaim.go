package Data

import "github.com/golang-jwt/jwt"

// CustomClaim 是在 JWT 标准请求后添加自定义字段的结构
type CustomClaim struct {
	jwt.StandardClaims
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (c CustomClaim) Valid() error {
	return nil
}
