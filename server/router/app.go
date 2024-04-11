package router

import (
	"github.com/gin-gonic/gin"
	"go-meeting/server/middleware"
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
	r.Static("/web", "./web")

	userRouter(r)
	wsRouter(r)
	// 认证
	r.Use(middleware.Auth())
	meetingRouter(r)
	return r
}
