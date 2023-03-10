package LoginSystem

import (
	"backend/AuthSystem"
	"backend/Data"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = InitDatabase()

func Register(c *gin.Context) {
	var account = c.Request.FormValue("account")
	var password = c.Request.FormValue("password")

	if IsExist(db, account) {
		c.JSON(http.StatusOK, gin.H{
			"code":    Data.ERROR_EXISTED_USER,
			"message": "已存在用户",
			"account": account,
		})
	} else {
		AddUser(db, account, password)
		c.JSON(http.StatusOK, gin.H{
			"code":    Data.REGISTER_SUCCESS,
			"message": "注册成功",
			"account": account,
		})
	}
}

func Login(c *gin.Context) {
	var account = c.Request.FormValue("account")
	var password = c.Request.FormValue("password")

	if !IsExist(db, account) {

		c.JSON(http.StatusOK, gin.H{
			"code":    Data.ERROR_NONEXISTENT_USER,
			"message": "用户名或密码错误",
			"account": account,
		})
		return
	}

	if !IsCorrect(db, account, password) {
		// c.String(http.StatusOK, "用户名或密码错误")
		c.JSON(http.StatusOK, gin.H{
			"code":    Data.ERROR_WRONG_PASSWORD,
			"message": "用户名或密码错误",
			"account": account,
		})
		return
	}

	// data := make(map[string]interface{})
	token, err := AuthSystem.GenToken(account, password)
	if err != nil {

	} else {
		// c.String(http.StatusOK, "登录成功")
		c.JSON(http.StatusOK, gin.H{
			"code":    Data.LOGIN_SUCCESS,
			"message": "登录成功",
			"account": account,
			"token":   token,
		})
	}
}
