package model

import (
	"config"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"
	"util"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var indexDB *sql.DB

func init() {
	var err error
	indexDB, _ = util.GetDBSession()
	userDB, err = util.GetDBSession()
	if err != nil {
		fmt.Println(err)
	}
}

// Login index model
// @Summary Login
// @Description login by name and pw
// @Tags user login
// @Accept  json
// @Produce  json
// @Param name body string true "your unique name"
// @Param bcrypt_pw body string true "your pw"
// @Success 200 {object} util.DataRes
// @Failure 400 {object} util.ErrorRes
// @Failure 404 {object} util.ErrorRes
// @Failure 500 {object} util.ErrorRes
// @Router /api/v1/login [post]
func Login(c echo.Context) error {
	user := map[string]string{
		"name":      "",
		"bcrypt_pw": "",
	}
	err := c.Bind(&user)
	if err != nil {
		return util.RetError(http.StatusBadRequest, 400, "参数错误", c)
	}

	var bcryptPW string
	var status int
	var id int
	err = indexDB.QueryRow("SELECT id, bcrypt_pw, status FROM user WHERE name=?", user["name"]).Scan(&id, &bcryptPW, &status)
	if err != nil {
		fmt.Println(err)
		return util.RetError(http.StatusBadRequest, 400, "密码验证错误", c)
	}
	if ok := util.CheckPasswordHash(user["bcrypt_pw"], bcryptPW); !ok {
		fmt.Println("密码验证错误")
		return util.RetError(http.StatusBadRequest, 400, "密码验证错误", c)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user["name"]
	// 判断是否通过认证
	if status < 10 {
		fmt.Println("还没有认证")
		return util.RetError(http.StatusBadRequest, 400, "还没有认证", c)
	}
	// 判断是否为管理员
	if status >= 100 {
		claims["admin"] = true
	} else {
		claims["admin"] = false
	}
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Generate encoded token and send it as response.
	t, _ := token.SignedString([]byte(config.Conf.SecretKey))

	// 生成 csrf token
	csrfToken := util.GetCSRFToken()
	return util.RetData(map[string]interface{}{
		"token": t,
		"csrf":  csrfToken,
	}, c)
}

// Register register model
// @Summary Register
// @Description register your account
// @Tags user register
// @Accept  json
// @Produce  json
// @Param name body string true "your unique name"
// @Param bcrypt_pw body string true "your pw"
// @Param nickname body string true "your nickname"
// @Param email body string true "your email"
// @Param phone_number body string true "your phone_number"
// @Param group body string true "your group"
// @Success 200 {object} util.DataRes
// @Failure 400 {object} util.ErrorRes
// @Failure 404 {object} util.ErrorRes
// @Failure 500 {object} util.ErrorRes
// @Router /api/v1/signup [post]
func Register(c echo.Context) error {
	userInfo := &User{}
	if err := c.Bind(userInfo); err != nil {
		return util.RetError(http.StatusBadRequest, 400, "参数错误", c)
	}
	// 用户状态为认证
	userInfo.Status = 0
	// 参数验证
	if err := c.Validate(userInfo); err != nil {
		return util.RetError(http.StatusBadRequest, 400, "参数错误", c)
	}
	// 注册

	if err := DefaultRegister(userInfo); err != nil {
		fmt.Println(err)
		return util.RetError(http.StatusBadGateway, 500, "内部错误", c)
	}
	return util.RetData(nil, c)
}

// DefaultRegister 默认注册器
func DefaultRegister(user *User) error {
	if ok, _ := verifyUserExist(user.Name); ok {
		return errors.New("此用户已经存在")
	}
	// insert
	stmt, err := indexDB.Prepare("INSERT INTO `user`(name, nickname, email, phone_number, bcrypt_pw, `group`, status, create_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	bcryptPW, _ := util.HashPassword(user.BcryptPW)
	_, err = stmt.Exec(user.Name, user.Nickname, user.Email, user.PhoneNumber, bcryptPW, user.Group, user.Status, time.Now())
	if err != nil {
		return err
	}
	return nil
}
