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
	bookRepository repositories.BookRepository = repositories.NewBookRepository(db)
	authService    service.AuthService         = service.NewAuthService(userRepository)
	jwtService     service.Jwtservice          = service.NewJWTService()
	userService    service.UserService         = service.NewUserService(userRepository)
	bookService    service.BookService         = service.NewBookService(bookRepository)
	userController controller.UserController   = controller.NewUserController(userService, jwtService)
	authController controller.AuthController   = controller.NewAuthController(jwtService, authService)
	bookController controller.BookController   = controller.NewBookController(bookService, jwtService)
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

	bookRoutes := r.Group("/api/book", middleware.AuthorizeJwt(jwtService))
	{
		bookRoutes.POST("/book", bookController.Insert)
		bookRoutes.GET("/book", bookController.All)
		bookRoutes.GET("/book/:id", bookController.FindById)
		bookRoutes.PUT("/book/:id", bookController.Update)
		bookRoutes.DELETE("/book/:id", bookController.Delete)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
