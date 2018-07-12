package util

import (
	"net/http"

	"github.com/labstack/echo"
)

// RetError 错误返回
func RetError(code, status int, errMsg string, c echo.Context) error {
	return c.JSON(code, map[string]interface{}{
		"status":  status,
		"err_msg": errMsg,
	})
}

// RetData 数据返回
func RetData(data interface{}, c echo.Context) error {
	if data != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
			"data":   data,
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
		})
	}

}
