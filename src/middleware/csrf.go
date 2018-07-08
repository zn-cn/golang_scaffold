package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// DefaultCSRFConfig 默认CSRFConfig
var DefaultCSRFConfig = middleware.CSRFConfig{
	Skipper:      middleware.DefaultSkipper,
	TokenLength:  32,
	TokenLookup:  "header:" + echo.HeaderXCSRFToken,
	ContextKey:   "csrf",
	CookieName:   "_csrf",
	CookieMaxAge: 86400,
}
