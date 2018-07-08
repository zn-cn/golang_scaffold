package view

import (
	"model"

	"github.com/labstack/echo"
)

// InitUserView 初始化 用户 view
func InitUserView(user *echo.Group) {
	user.GET("/info", model.GetUserInfo)
}
