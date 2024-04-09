package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-meeting/internal/define/res"
	"log"
	"net/http"
	"sync"
)

var wsP2PConnMap = sync.Map{}

func WsP2PConnection(c *gin.Context) {
	// 0. 获取房间和用户的信息
	// 1. 升级协议
	// 2. 存储当前的连接信息
	// 3. 监听发过来的消息

	// 获取房间和用户的信息
	in := new(WsP2PConnectionRequest)
	err := c.ShouldBindUri(in)
	if err != nil {
		res.Wrong(c, -1, "地址参数错误"+err.Error())
		return
	}

	// 升级协议
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// 存储当前的连接信息
	userConnMap := new(sync.Map)
	value, ok := wsP2PConnMap.Load(in.RoomIdentity)
	if ok {
		userConnMap = value.(*sync.Map)
	}
	userConnMap.Store(in.UserIdentity, conn)
	wsP2PConnMap.Store(in.RoomIdentity, userConnMap)

	// 监听发过来的消息
	for {
		_, data, e := conn.ReadMessage()
		if e != nil {
			return
		}
		v, okk := wsP2PConnMap.Load(in.RoomIdentity)
		if okk {
			v.(*sync.Map).Range(func(key, value any) bool {
				err2 := value.(*websocket.Conn).WriteMessage(websocket.TextMessage, data)
				if err2 != nil {
					log.Println("WriteMessage err.", err2)
				}
				return true
			})
		}
	}
}
