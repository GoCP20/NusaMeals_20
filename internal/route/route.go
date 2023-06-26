package route

import (
	"reglog/internal/common/middleware"
	"reglog/internal/controller"
	"reglog/internal/repository"
	"reglog/internal/usecase"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Config struct {
	// ECHO TOP LEVEL INSTANCE
	Echo        *echo.Echo
	DBConn      *gorm.DB
	JwtProvider *middleware.JWTProvider
}

func (cfg *Config) New() {
	// Get Auth middleware to filter authorization user/admin
	authMiddleware := middleware.NewAuthMiddleware(cfg.JwtProvider)

	// dependency injection
	userRepository := repository.NewUserRepository(cfg.DBConn)
	categoryRepository := repository.NewCategoryRepository(cfg.DBConn)
	menuRepository := repository.NewMenuRepository(cfg.DBConn)
	//paymentRepository := repository.NewPaymentRepository(cfg.DBConn)
	//orderRepository := repository.NewOrderRepository(cfg.DBConn)
	//tableRepository := repository.NewTableRepository(cfg.DBConn)
	//orderDetailRepository := repository.NewOrderDetailRepository(cfg.DBConn)

	userUseCase := usecase.NewUserUseCase(userRepository, cfg.JwtProvider)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository, menuRepository)
	menuUseCase := usecase.NewMenuUseCase(menuRepository)
	//paymentUseCase := usecase.NewPaymentUseCase(paymentRepository, orderRepository, userRepository)
	//orderUseCase := usecase.NewOrderUseCase(orderRepository, userRepository, menuRepository)
	//tableUseCase := usecase.NewTableUseCase(tableRepository)
	//orderDetailUseCase := usecase.NewOrderDetailUseCase(orderDetailRepository)

	// Routes

	// AUTH
	authController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/login", authController.LoginController)
	//register ADMIN
	adminController := controller.NewAuthController(userUseCase)
	cfg.Echo.POST("/admin/register", adminController.RegisterAdminController)
	//register USER
	cfg.Echo.POST("/register", authController.RegisterUserController)

	// USER
	userController := controller.NewUserController(userUseCase)
	user := cfg.Echo.Group("/users", authMiddleware.IsAuthenticated())
	user.GET("/user", userController.GetUserByID, authMiddleware.IsUser)
	user.GET("/username", userController.GetUserByUsername, authMiddleware.IsUser)
	user.GET("/email", userController.GetUserByEmail, authMiddleware.IsUser)
	user.PUT("/:id", userController.UpdateUser, authMiddleware.IsUser)

	// ADMIN
	admin := cfg.Echo.Group("/admin", authMiddleware.IsAdmin)
	admin.GET("/users", userController.GetAllUser)
	admin.GET("/users/id", userController.GetUserByID)
	admin.GET("/users/username", userController.GetUserByUsername)
	admin.GET("/users/email", userController.GetUserByEmail)
	admin.PUT("/users/:id", userController.UpdateUser)

}
