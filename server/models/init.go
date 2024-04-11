package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	// "username:password@tcp(host:port)/go_meeting?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/go_meeting?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"))
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
