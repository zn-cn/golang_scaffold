package model

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"util"

	"github.com/labstack/echo"
)

// User 用户结构体
type User struct {
	ID          int       `json:"id" query:"id" example:"1" format:"int64"`
	Name        string    `json:"name" query:"name" validate:"required" example:"我的名字"`
	Nickname    string    `json:"nickname" query:"nickname" validate:"required" example:"nickname"`
	Email       string    `json:"email" query:"email" validate:"required,email" example:"example@qq.com"`
	PhoneNumber string    `json:"phone_number" query:"phone_number" validate:"required" example:"15837891235"`
	BcryptPW    string    `json:"bcrypt_pw" query:"bcrypt_pw" validate:"required" example:"123456789"`
	Group       string    `json:"group" query:"group" validate:"required" example:"程序组"`
	Status      int       `json:"status" query:"status" example:"0"` // status: 0 没有认证通过，10: 普通用户, 100: 管理员
	CreateDate  time.Time `json:"create_date" query:"create_date" example:"2018-07-13"`
}

var userDB *sql.DB

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags user info
// @Accept  json
// @Produce  json
// @Success 200 {object} util.DataRes
// @Failure 400 {object} util.ErrorRes
// @Failure 404 {object} util.ErrorRes
// @Failure 500 {object} util.ErrorRes
// @Security ApiKeyAuth
// @Router /api/v1/login [post]
func GetUserInfo(c echo.Context) error {
	payload := util.GetJWTInfo(c)
	userInfo := User{}
	err := userDB.QueryRow("SELECT name, nickname, email, phone_number, group FROM user WHERE name=?", payload["username"]).Scan(&userInfo.Name, &userInfo.Nickname, &userInfo.Email, &userInfo.PhoneNumber, &userInfo.Group)
	if err != nil {
		return util.RetError(http.StatusBadRequest, 400, "请求参数有问题", c)
	}
	return util.RetData(userInfo, c)
}

// 判断是否存在此用户
func verifyUserExist(username string) (bool, error) {
	var bcryptPW string
	err := userDB.QueryRow("SELECT bcrypt_pw FROM user WHERE name=?", username).Scan(&bcryptPW)
	if err != nil || bcryptPW == "" {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
