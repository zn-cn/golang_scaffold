package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetJWTInfo 获取 payload
func GetJWTInfo(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
