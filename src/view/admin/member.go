package admin

import (
	adminModel "model/admin"

	"github.com/labstack/echo"
)

// InitAdminMemView 初始化后台成员管理 View
func InitAdminMemView(admin *echo.Group) {
	user := admin.Group("/user")
	initAdminUserView(user)
}

func initAdminUserView(user *echo.Group) {
	user.POST("/del", adminModel.DeleteMem)
	user.POST("/add", adminModel.AddMem)
	user.POST("/update", adminModel.UpdateMemInfo)
	// 通过条件获取成员信息
	user.GET("/info", adminModel.GetMemInfoByCon)
}
