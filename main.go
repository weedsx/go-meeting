package main

import (
	"go-meeting/server"
	"go-meeting/server/models"
	"go-meeting/server/router"
	"log"
)

// Package classification Your_API_Name.
//
// swagger:meta
func main() {
	internal.ConfigInit()
	models.InitDB()
	engine := router.Router()
	err := engine.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		log.Fatal("run error: ", err)
		return
	}
}
