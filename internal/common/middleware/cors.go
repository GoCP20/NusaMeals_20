package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.DELETE, echo.GET, echo.PUT, echo.POST, echo.OPTIONS},
	})
}

func Logger() echo.MiddlewareFunc {
	return middleware.Logger()
}

func Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

func RequestID() echo.MiddlewareFunc {
	return middleware.RequestID()
}

func Gzip() echo.MiddlewareFunc {
	return middleware.Gzip()
}

func BodyLimit(limit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(limit)
}
