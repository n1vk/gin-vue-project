package Router

import (
	jwt "backend/JWT"
	api "backend/LoginSystem"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	var router = gin.Default()

	// Using CORS Middleware
	router.Use(cors.Default())

	// Using JWT Middleware

	// router.NoRoute()
	router.GET("/reg", api.Register)
	router.GET("/log", api.Login)

	homeRouter := router.Group("/home")
	homeRouter.Use(jwt.JWT())
	{
		homeRouter.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "hello",
			})
		})
	}

	router.Run(":33445")
}
