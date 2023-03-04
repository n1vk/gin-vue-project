package LoginSystem

import (
	"backend/Data"
	"backend/JWT"
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
		// c.String(http.StatusOK, "已存在用户")
		c.JSON(http.StatusOK, gin.H{
			"status":  "existent",
			"message": "已存在用户",
			"account": account,
		})
	} else {
		addUser(account, password)
		// c.String(http.StatusOK, "已注册")
		c.JSON(http.StatusOK, gin.H{
			"status":  "registered",
			"message": "注册成功",
			"account": account,
		})
	}
}

func Login(c *gin.Context) {
	var account, _ = strconv.Atoi(c.Request.FormValue("Acc"))
	var password = c.Request.FormValue("Pwd")

	if !isExist(account) {
		// c.String(http.StatusOK, "用户名或密码错误") // 用户不存在
		c.JSON(http.StatusOK, gin.H{
			"status":  "nonexistent",
			"message": "用户名或密码错误",
			"account": account,
		})
		return
	}

	if !isCorrect(account, password) {
		// c.String(http.StatusOK, "用户名或密码错误")
		c.JSON(http.StatusOK, gin.H{
			"status":  "wrong-password",
			"message": "用户名或密码错误",
			"account": account,
		})
		return
	}

	data := make(map[string]interface{})
	token, err := JWT.GenToken(account)
	if err != nil {

	} else {
		data["token"] = token
	}

	// c.String(http.StatusOK, "登录成功")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "登录成功",
		"account": account,
		"data":    data,
	})
	// Give Token
}

//func NotFound(c *gin.Context) {
//	c.JSON(http.StatusNotFound, gin.H{
//		"status": 404,
//		"error":  "Resource not found",
//	})
//}
