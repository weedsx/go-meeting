package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-meeting/internal/define/res"
	"go-meeting/internal/helper"
	"go-meeting/internal/models"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func UserLogin(c *gin.Context) {
	in := new(UserLoginRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		res.Wrong(c, -1, "request error")
		return
	}

	if strings.TrimSpace(in.Username) == "" ||
		strings.TrimSpace(in.Password) == "" {
		res.Wrong(c, http.StatusOK, "用户或密码请勿为空")
		return
	}

	in.Password = helper.GetMd5(in.Password)
	data := new(models.UserBasic)
	err = models.DB.Where("username = ? and password = ?",
		in.Username, in.Password).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.Wrong(c, -1, "用户名或密码错误")
			return
		}
		res.Wrong(c, -1, "Get UserBasic Error:"+err.Error())
		return
	}

	token, err := helper.GenerateToken(data.ID, data.Username)
	if err != nil {
		res.Wrong(c, -1, "GenerateToken Error:"+err.Error())
		return
	}
	res.Success(c, gin.H{
		"token": token,
	})
}
