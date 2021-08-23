package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odamilola36/golang_api/config"
	"github.com/odamilola36/golang_api/controller"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main (){
	r := gin.Default()
	
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}