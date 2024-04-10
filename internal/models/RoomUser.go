package models

import "gorm.io/gorm"

// swagger:model RoomUser
type RoomUser struct {
	gorm.Model
	Rid uint `gorm:"column:room_id;type:int(11);not null" json:"rid"`
	Uid uint `gorm:"column:user_id;type:int(11);not null" json:"uid"`
}

func (ru *RoomUser) TableName() string {
	return "room_user"
}
