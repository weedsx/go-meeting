package main

import (
	"go-meeting/internal/models"
	"go-meeting/internal/server/router"
	"log"
)

func main() {
	models.InitDB()
	engine := router.Router()
	err := engine.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		log.Fatal("run error: ", err)
		return
	}
}