package AuthSystem

import (
	"backend/Data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

// 任意规定的secret用于生成token
var serverSecret = []byte("gin-vue-project")

// JWT 是中间件函数
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		var code = Data.VALIDATE_SUCCESS

		// 从收到的数据包中获得token
		// 前端所有对/home的请求应该有token=xxx
		token := c.Query("token")
		if token == "" {
			code = Data.INVALID_PARAMS
		} else {
			payload, err := ParseToken(token)
			if err != nil {
				code = Data.ERROR_TOKEN_CHECK_FAIL
			} else if time.Now().Unix() > payload.ExpiresAt {
				code = Data.ERROR_TOKEN_EXPIRED
			}
		}

		if code != Data.VALIDATE_SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  code,
				"token": token,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// GenToken 使用用户的信息生成密钥，返回字符串类型的token和错误
func GenToken(acc string, pwd string) (string, error) {

	customClaim := Data.CustomClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "goServer",
		}, acc, pwd,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)

	return token.SignedString(serverSecret)
}

// ParseToken 解析字符串密钥
func ParseToken(tokenString string) (*Data.CustomClaim, error) {
	acquiredClaim := new(Data.CustomClaim)

	// 解析需要密钥，自定义Claim和自定义secret
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
