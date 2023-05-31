package routes

import (
	"authentication/controllers"
	"authentication/services"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(r *gin.Engine) *gin.Engine {
	authService := services.GetAuthService()
	authController := controllers.GetAuthController(authService)
	authRoutes := r.Group("/auth")

	authRoutes.POST("/login", authController.Login)
	authRoutes.POST("/register", authController.CreateUser)

	return r
}
