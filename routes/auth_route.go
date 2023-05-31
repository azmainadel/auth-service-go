package routes

import (
	"authentication/controllers"
	"authentication/db"
	"authentication/repositories"
	"authentication/services"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(r *gin.Engine) *gin.Engine {
	postgresDb := db.GetPostgresDb()
	authRepository := repositories.GetAuthenticationRepository(postgresDb)
	authService := services.GetAuthService(authRepository)
	authController := controllers.GetAuthController(authService)
	authRoutes := r.Group("/auth")

	authRoutes.POST("/login", authController.Login)
	authRoutes.POST("/register", authController.CreateUser)

	return r
}
