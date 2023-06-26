package main

import (
	"reglog/internal/common/constant"

	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"

	"reglog/internal/common/config"
	"reglog/internal/common/middleware"
	"reglog/internal/common/util"
	"reglog/internal/route"
)

func main() {
	// Setup JWT Secret Key
	jwtProvider := middleware.NewJWTProvider(constant.JWT_SECRET_KEY)

	// ECHO HTTP SERVER
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Custom Middleware
	e.Use(middleware.CORS())
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}

	// Initialize database connection
	config.InitDB()
	config.InitialMigration()

	// Get Routes
	route := route.Config{
		Echo:        e,
		DBConn:      config.DB,
		JwtProvider: jwtProvider,
	}
	route.New()

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
