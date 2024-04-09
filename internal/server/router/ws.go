package router

import (
	"github.com/gin-gonic/gin"
	"go-meeting/internal/server/service"
)

func wsRouter(r *gin.Engine) {
	// ws
	r.GET("/ws/p2p/:room_identity/:user_identity", service.WsP2PConnection)
}
