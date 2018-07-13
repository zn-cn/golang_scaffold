package util

import (
	"net/http"

	"github.com/labstack/echo"
)

// ErrorRes 错误返回
type ErrorRes struct {
	status int
	errMsg string
}

// DataRes 数据返回
type DataRes struct {
	status int
	data   interface{}
}

// RetError 错误返回
func RetError(code, status int, errMsg string, c echo.Context) error {
	// return c.JSON(code, map[string]interface{}{
	// 	"status":  status,
	// 	"err_msg": errMsg,
	// })
	return c.JSON(code, ErrorRes{
		status: status,
		errMsg: errMsg,
	})
}

// RetData 数据返回
func RetData(data interface{}, c echo.Context) error {
	return c.JSON(http.StatusOK, DataRes{
		status: 200,
		data:   data,
	})
}
