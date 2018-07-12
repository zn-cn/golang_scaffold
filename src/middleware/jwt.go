package middleware

import (
	"config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// DefaultJWTConfig jwt配置
var DefaultJWTConfig = middleware.JWTConfig{
	SigningKey:  []byte(config.Conf.SecretKey),
	TokenLookup: "header:" + echo.HeaderAuthorization,
	AuthScheme:  "Bearer",
	Claims:      jwt.MapClaims{},
	ContextKey:  "user",
	Skipper:     Skipper,
}

// Skipper 过滤
func Skipper(c echo.Context) bool {
	if c.Path() == "/api/v1/login" || c.Path() == "/api/v1/signup" {
		return true
	}
	return false
}
