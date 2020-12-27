package conf

import (
	"github.com/gin-gonic/gin"
	"go-gin-web-demo/app/river-api/internal/handle"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handle.AddUser)
		//userRouter.POST("/login", handler.UserLogin)
	}
	return router
}
