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

// swagger:operation POST /user/login
func UserLogin(c *gin.Context) {
	in := new(UserLoginRequest)
	if err := c.ShouldBindJSON(in); err != nil {
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

	if err := models.DB.Where("username = ? and password = ?",
		in.Username, in.Password).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.Wrong(c, -1, "用户名或密码错误")
			return
		}
		res.Wrong(c, -1, "Get UserBasic Error:"+err.Error())
		return
	}

	if token, err := helper.GenerateToken(data.ID, data.Username); err != nil {
		res.Wrong(c, -1, "GenerateToken Error:"+err.Error())
		return
	} else {
		res.Success(c, gin.H{
			"token": token,
		})
	}
}
