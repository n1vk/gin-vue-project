package JWT

import "github.com/golang-jwt/jwt"

// Payload is a customized JWT Claim
type Payload struct {
	jwt.StandardClaims
	Account int `json:"acc"`
}

func (c Payload) Valid() error {
	return nil
}
