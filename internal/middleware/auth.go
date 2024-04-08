package middleware

import (
	"github.com/gin-gonic/gin"
	"go-meeting/internal/define/res"
	"go-meeting/internal/helper"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaims, err := helper.ParseToken(auth)
		if err != nil {
			c.Abort()
			res.Error(c, http.StatusForbidden, "Unauthorized Authorization")
			return
		}
		if userClaims == nil {
			c.Abort()
			res.Error(c, http.StatusForbidden, "Unauthorized Admin")
			return
		}
		c.Set("user_id", userClaims)
		c.Next()
	}
}
