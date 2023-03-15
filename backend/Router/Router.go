package Router

import (
	authAPI "backend/AuthSystem"
	homeAPI "backend/HomeSystem"
	loginAPI "backend/LoginSystem"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	// 创建默认的 gin.Engine
	var router = gin.Default()

	// 使用 CORS 相关中间件处理跨域请求
	router.Use(cors.Default())

	// 处理注册和登录
	router.GET("/reg", loginAPI.Register)
	router.GET("/log", loginAPI.Login)

	// 已经登录的用户在/home路由组
	homeRouter := router.Group("/home")
	// 使用JWT中间件
	homeRouter.Use(authAPI.JWT())
	{
		homeRouter.GET("/hello", homeAPI.Hello)

		// 其他的路径路由添加在这里
	}

	router.Run("0.0.0.0:33445")
}
