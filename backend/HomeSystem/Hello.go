package HomeSystem

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Hello 是占位用处理函数，暂时仅发送一个验证通过的请求
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Accepted",
	})
}
