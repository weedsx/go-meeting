package models

import "gorm.io/gorm"

// swagger:model UserBasic
type UserBasic struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(100);uniqueIndex;not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(36);not null" json:"password"`
	// Sdp 用于建立 p2p 通信的通道
	Sdp string `gorm:"column:sdp;type:text" json:"sdp"`
}

func (u *UserBasic) TableName() string {
	return "user_basic"
}
