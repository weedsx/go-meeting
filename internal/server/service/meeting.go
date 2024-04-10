package service

import (
	"github.com/gin-gonic/gin"
	"go-meeting/internal/define/res"
	"go-meeting/internal/helper"
	"go-meeting/internal/models"
	"strings"
	"time"
)

// MeetingList
// swagger:operation GET /meeting/list
func MeetingList(c *gin.Context) {
	in := new(MeetingListRequest)
	if err := c.ShouldBindQuery(in); err != nil {
		res.Wrong(c, -1, "参数异常")
		return
	}

	var list []*MeetingListReply
	var cnt int64
	tx := models.DB.Model(&models.RoomBasic{})
	if strings.TrimSpace(in.Keyword) != "" {
		tx.Where("name like ?", "%"+in.Keyword+"%")
	}
	if err := tx.Count(&cnt).Limit(in.Size).Offset((in.Page - 1) * in.Size).
		Find(&list).Error; err != nil {
		res.Wrong(c, -1, "系统异常："+err.Error())
		return
	}
	res.Success(c, gin.H{
		"list":  list,
		"count": cnt,
	})
}

// swagger:operation POST /meeting/create
func MeetingCreate(c *gin.Context) {
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	in := new(MeetingCreateRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		res.Wrong(c, -1, "参数异常")
		return
	}

	if err := models.DB.Create(&models.RoomBasic{
		Identity: helper.GetUUID(),
		Name:     in.Name,
		BeginAt:  time.UnixMilli(in.BeginAt),
		EndAt:    time.UnixMilli(in.EndAt),
		CreateId: uc.Id,
	}).Error; err != nil {
		res.Wrong(c, -1, "系统异常: "+err.Error())
		return
	}
	res.Success(c, "ok")
}

// swagger:operation POST /meeting/edit
func MeetingEdit(c *gin.Context) {
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	in := new(MeetingEditRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		res.Wrong(c, -1, "参数异常")
		return
	}

	if err := models.DB.Model(new(models.RoomBasic)).
		Where("identity = ? AND create_id = ?",
			in.Identity, uc.Id).Error; err != nil {
		res.Wrong(c, -1, "系统异常: "+err.Error())
		return
	}

	res.Success(c, "ok")
}

// swagger:operation DELETE /meeting/delete
func MeetingDelete(c *gin.Context) {
	identity := c.Query("identity")
	uc := c.MustGet("user_claims").(*helper.UserClaims)

	if err := models.DB.Where("identity = ? and create_id = ?",
		identity, uc.Id).Delete(&models.RoomBasic{}).Error; err != nil {
		res.Wrong(c, -1, "系统异常: "+err.Error())
		return
	}

	res.Success(c, "ok")
}
