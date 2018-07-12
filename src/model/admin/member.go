package admin

import (
	"database/sql"
	"model"
	"net/http"
	"util"

	"github.com/labstack/echo"
)

var adminDB *sql.DB

func init() {
	adminDB, _ = util.GetDBSession()
}

// UpdateMemInfo 更新成员信息
func UpdateMemInfo(c echo.Context) error {

	userInfo := map[string]string{}
	if err := c.Bind(&userInfo); err != nil {
		return util.RetError(http.StatusBadRequest, 400, "参数错误", c)
	}
	if err := updateUserInfo(userInfo, userInfo["name"]); err != nil {
		return util.RetError(http.StatusBadGateway, 406, "参数错误或者服务器内部错误", c)
	}
	return util.RetData(nil, c)
}

// DeleteMem 删除成员
func DeleteMem(c echo.Context) error {
	user := map[string]string{
		"name": "",
	}
	if err := c.Bind(&user); err != nil {
		return util.RetError(http.StatusBadRequest, 400, "参数错误", c)
	}
	if err := delUser(user["name"]); err != nil {
		return util.RetError(http.StatusBadGateway, 407, "用户不存在或者服务器内部错误", c)
	}
	return util.RetData(nil, c)
}

// AddMem 增加成员
func AddMem(c echo.Context) error {
	userInfo := &model.User{}
	err := c.Bind(userInfo)
	if err != nil {
		return util.RetError(http.StatusBadRequest, 400, "参数错误", c)
	}
	// 用户状态为认证
	userInfo.Status = 10
	if err = model.DefaultRegister(userInfo); err != nil {
		return util.RetError(http.StatusBadGateway, 500, "内部错误", c)
	}
	return util.RetData(nil, c)
}

// GetMemInfoByCon 通过某种条件获取成员信息
func GetMemInfoByCon(c echo.Context) error {
	userInfo := map[string]string{}
	if err := c.Bind(&userInfo); err != nil {
		return util.RetError(http.StatusBadRequest, 400, "参数错误", c)
	}
	data, err := getMemInfoByCon(userInfo)
	if err != nil {
		return util.RetError(http.StatusBadGateway, 500, "内部错误", c)
	}
	return util.RetData(data, c)
}

func delUser(name string) error {
	_, err := adminDB.Exec("DELETE FROM user WHERE name=?", name)
	return err

}

func updateUserInfo(userInfo map[string]string, name string) error {
	// 少了一层验证
	for k, v := range userInfo {
		if _, err := adminDB.Exec("UPDATE user SET "+k+"=? WHERE name=?", v, name); err != nil {
			return err
		}
	}
	return nil
}

func getMemInfoByCon(userInfo map[string]string) (interface{}, error) {
	whereStr := ""
	for k, v := range userInfo {
		whereStr += k + "=" + v
	}
	usersInfo := []model.User{}
	rows, err := adminDB.Query("SELECT name, nickname, email, group, phone_number, status, create_date FROM user WHERE " + whereStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		userInfo := model.User{}
		err = rows.Scan(&userInfo.Name, &userInfo.Nickname, &userInfo.Email, &userInfo.Group, &userInfo.PhoneNumber, &userInfo.Status, &userInfo.CreateDate)
		if err != nil {
			return nil, err
		}
		usersInfo = append(usersInfo, userInfo)
	}
	return usersInfo, nil
}
