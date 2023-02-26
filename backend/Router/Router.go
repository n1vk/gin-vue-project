package Router

import (
	api "backend/LoginSystem"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	var router = gin.Default()
	router.Use(cors.Default())

	// router.NoRoute()
	router.GET("/reg", api.Register)
	router.GET("/log", api.Login)

	router.Run(":33445") // unhandled error, ignore for now
}
