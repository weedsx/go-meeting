package router

import (
	"github.com/gin-gonic/gin"
	"go-meeting/internal/server/service"
)

func meetingRouter(r *gin.Engine) {
	// 会议列表
	r.GET("/meeting/list", service.MeetingList)
	// 创建会议
	r.POST("/meeting/create", service.MeetingCreate)
	// 修改会议
	r.POST("/meeting/edit", service.MeetingEdit)
	// 删除会议
	r.DELETE("/meeting/delete", service.MeetingDelete)
}
