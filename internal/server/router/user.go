package router

import (
	"github.com/gin-gonic/gin"
	"go-meeting/internal/server/service"
)

func userRouter(r *gin.Engine) {
	// 用户登录
	r.POST("/user/login", service.UserLogin)
}
