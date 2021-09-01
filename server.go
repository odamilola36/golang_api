package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odamilola36/golang_api/config"
	"github.com/odamilola36/golang_api/controller"
	"github.com/odamilola36/golang_api/middleware"
	"github.com/odamilola36/golang_api/repositories"
	"github.com/odamilola36/golang_api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = config.SetupDatabaseConnection()
	userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	authService    service.AuthService         = service.NewAuthService(userRepository)
	jwtService     service.Jwtservice          = service.NewJWTService()
	userService    service.UserService         = service.NewUserService(userRepository)
	userController controller.UserController   = controller.NewUserController(userService, jwtService)
	authController controller.AuthController   = controller.NewAuthController(jwtService, authService)
)

func main() {
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("/api/user", middleware.AuthorizeJwt(jwtService))
	{
		userRoutes.PUT("/profile", userController.UpdateUser)
		userRoutes.GET("/profile", userController.Profile)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
