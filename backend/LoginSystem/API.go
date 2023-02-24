package LoginSystem

import (
	"backend/Data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func isExist(acc int) bool {
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

func isCorrect(acc int, pwd string) bool {
	for _, u := range Data.AllUsers {
		if u.Account == acc {
			if u.Password == pwd {
				return true
			}
		}
	}

	return false
}

func addUser(acc int, pwd string) {
	var usr Data.User
	usr.Account = acc
	usr.Password = pwd
	Data.AllUsers = append(Data.AllUsers, usr)
}

func Register(c *gin.Context) {
	var account, _ = strconv.Atoi(c.Request.FormValue("Acc"))
	var password = c.Request.FormValue("Pwd")

	if isExist(account) {
		c.String(http.StatusOK, "已存在用户")
	} else {
		addUser(account, password)
		c.String(http.StatusOK, "已注册")
	}
}

func Login(c *gin.Context) {
	var account, _ = strconv.Atoi(c.Request.FormValue("Acc"))
	var password = c.Request.FormValue("Pwd")

	if !isExist(account) {
		c.String(http.StatusOK, "用户名或密码错误") // 用户不存在
		return
	}

	if !isCorrect(account, password) {
		c.String(http.StatusOK, "用户名或密码错误")
		return
	}

	c.String(http.StatusOK, "登录成功")
	// Give Token
}

//func NotFound(c *gin.Context) {
//	c.JSON(http.StatusNotFound, gin.H{
//		"status": 404,
//		"error":  "Resource not found",
//	})
//}
