package service

import (
	"github.com/gin-gonic/gin"
	"go-meeting/internal/define/res"
	"go-meeting/internal/helper"
	"go-meeting/internal/models"
	"strings"
	"time"
)

func MeetingList(c *gin.Context) {
	in := new(MeetingListRequest)
	err := c.ShouldBindQuery(in)
	if err != nil {
		res.Wrong(c, -1, "参数异常")
		return
	}

	var list []*MeetingListReply
	var cnt int64
	tx := models.DB.Model(&models.RoomBasic{})
	if strings.TrimSpace(in.Keyword) != "" {
		tx.Where("name like ?", "%"+in.Keyword+"%")
	}
	err = tx.Count(&cnt).Limit(in.Size).Offset((in.Page - 1) * in.Size).
		Find(&list).Error
	if err != nil {
		res.Wrong(c, -1, "系统异常："+err.Error())
		return
	}
	res.Success(c, gin.H{
		"list":  list,
		"count": cnt,
	})
}

func MeetingCreate(c *gin.Context) {
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	in := new(MeetingCreateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		res.Wrong(c, -1, "参数异常")
		return
	}
	err = models.DB.Create(&models.RoomBasic{
		Identity: helper.GetUUID(),
		Name:     in.Name,
		BeginAt:  time.UnixMilli(in.BeginAt),
		EndAt:    time.UnixMilli(in.EndAt),
		CreateId: uc.Id,
	}).Error
	if err != nil {
		res.Wrong(c, -1, "系统异常: "+err.Error())
		return
	}
	res.Success(c, "ok")
}
