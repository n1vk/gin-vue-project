package JWT

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

// Random secret used for generating token
var serverSecret = []byte("gin-vue-project")

// Middleware logic
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var status string = ""
		var data interface{}

		token := c.Query("token")
		if token == "" {
			status = "no_token"
		} else {
			payload, err := ParseToken(token)
			if err != nil {
				status = "token_invalid"
			} else if time.Now().Unix() > payload.ExpiresAt {
				status = "token_expired"
			}
		}

		if status != "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"stauts": status,
				"data":   data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GenToken(acc int) (string, error) {

	customClaim := Payload{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "goServer",
		}, acc,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)

	return token.SignedString(serverSecret)
}

func ParseToken(tokenString string) (*Payload, error) {
	acquiredClaim := new(Payload)

	token, err := jwt.ParseWithClaims(tokenString, acquiredClaim, func(token *jwt.Token) (interface{}, error) {
		return serverSecret, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if token.Valid {
		return acquiredClaim, nil
	}

	return nil, err
}
