package router

import (
	"github.com/gin-gonic/gin"
	"go-meeting/internal/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 解决跨域
	r.Use(middleware.Cors())
	userRouter(r)

	// 认证
	r.Use(middleware.Auth())
	meetingRouter(r)
	return r
}
