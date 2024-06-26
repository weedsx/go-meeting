package models

import (
	"gorm.io/gorm"
	"time"
)

// swagger:model RoomBasic
type RoomBasic struct {
	gorm.Model
	Identity string    `gorm:"column:identity;type:varchar(36);uniqueIndex;not null" json:"identity"`
	Name     string    `gorm:"column:name;type:varchar(100);not null" json:"name"`
	BeginAt  time.Time `gorm:"column:begin_at;type:datetime;" json:"begin_at"`
	EndAt    time.Time `gorm:"column:end_at;type:datetime;" json:"end_at"`
	// CreatedId 创建者 ID
	CreateId uint `gorm:"column:create_id;type:int(11);not null" json:"create_id"`
}

func (r *RoomBasic) TableName() string {
	return "room_basic"
}
