package routes

import (
	"authentication/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(r *gin.Engine) *gin.Engine {
	authController := controllers.GetAuthController()
	authRoutes := r.Group("/auth")

	authRoutes.POST("/login", authController.Login)
	authRoutes.POST("/register", authController.CreateUser)

	return r
}
