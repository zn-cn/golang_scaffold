package model

import (
	"database/sql"
	"net/http"
	"time"
	"util"

	"github.com/labstack/echo"
)

// User 用户结构体
type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Nickname    string    `json:"nickname" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	BcryptPW    string    `json:"bcrypt_pw" validate:"required"`
	Group       string    `json:"group" validate:"required"`
	Status      int       `json:"status" validate:"required"` // status: 0 没有认证通过，10: 普通用户, 100: 管理员
	CreateDate  time.Time `json:"create_date"`
}

var userDB *sql.DB

// GetUserInfo 获取用户信息
func GetUserInfo(c echo.Context) error {
	payload := util.GetJWTInfo(c)
	userInfo := User{}
	err := userDB.QueryRow("SELECT name, nickname, email, phone_number,group FROM user WHERE name=?", payload["username"]).Scan(&userInfo.Name, &userInfo.Nickname, &userInfo.Email, &userInfo.PhoneNumber, &userInfo.Group)
	if err != nil {
		return util.RetError(http.StatusBadRequest, 400, "请求参数有问题", c)
	}
	return util.RetData(userInfo, c)
}

// 判断是否存在此用户
func verifyUserExist(username string) (bool, error) {
	var bcryptPW string
	err := userDB.QueryRow("SELECT bcrypt_pw FROM user WHERE username=?", username).Scan(&bcryptPW)
	if err != nil || bcryptPW == "" {
		return false, err
	}
	return true, nil
}
