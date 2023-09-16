package middleware

import (
	"connect/util"
	"github.com/gin-gonic/gin"
)

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header
		if header.Get("x-role-id") != util.ADMIN_ROLE {
			c.AbortWithStatusJSON(403, gin.H{
				"message": "Forbidden",
			})
			return
		}
		c.Next()
	}
}
