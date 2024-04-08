package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:llh112358@tcp(127.0.0.1:3306)/go_meeting?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&UserBasic{}, &RoomBasic{}, &RoomUser{})
	if err != nil {
		log.Fatal("db.AutoMigrate failed: ", err)
		return
	}
	DB = db
}
