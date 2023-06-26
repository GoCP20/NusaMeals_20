package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type AuthMiddleware struct {
	jwtProvider *JWTProvider
}

func NewAuthMiddleware(jwt *JWTProvider) *AuthMiddleware {
	return &AuthMiddleware{
		jwtProvider: jwt,
	}
}

func (m *AuthMiddleware) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		payload, err := m.jwtProvider.ExtractToken(ctx)

		if err != nil {
			return ctx.JSON(http.StatusUnauthorized,
				errors.New("unauthorized").Error())
		}

		if payload.Role != "admin" {
			return ctx.JSON(http.StatusForbidden,
				errors.New("access forbidden").Error())
		}

		return next(ctx)
	}
}

func (m *AuthMiddleware) IsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		payload, err := m.jwtProvider.ExtractToken(ctx)

		if err != nil {
			return ctx.JSON(http.StatusUnauthorized,
				errors.New("unauthorized").Error())
		}

		if payload.Role != "user" {
			return ctx.JSON(http.StatusForbidden,
				errors.New("access forbidden").Error())
		}

		return next(ctx)
	}
}

func (m *AuthMiddleware) IsAuthenticated() echo.MiddlewareFunc {
	return middleware.JWT([]byte(m.jwtProvider.JWTSecret))
}

//var JWTMiddlewareConfig = middleware.JWTWithConfig(middleware.JWTConfig{
//	SigningMethod: "HS256",
//	SigningKey:    []byte(constant.SECRET_KEY),
//	TokenLookup:   "cookie:JWTCookie",
//	AuthScheme:    "user",
//})
