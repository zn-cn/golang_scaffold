package view

import (
	"model"

	"github.com/labstack/echo"
)

// InitIndexView 初始化 index View
func InitIndexView(index *echo.Group) {
	index.POST("/login", model.Login)
	index.POST("/signup", model.Register)
}
