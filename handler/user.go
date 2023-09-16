package handler

import (
	"connect/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	service service.IUser
}

func NewUser(service service.IUser) *User {
	return &User{
		service: service,
	}
}

func (h *User) GetUsers(c *gin.Context) {
	user, err := h.service.GetUsers(c)
	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, user)
}
