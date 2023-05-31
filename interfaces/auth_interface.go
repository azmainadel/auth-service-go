package interfaces

import (
	"authentication/dto"
	"authentication/models"

	"github.com/gin-gonic/gin"
)

type AuthControllerInterface interface {
	Login(c *gin.Context)
	CreateUser(c *gin.Context)
}

type AuthServiceInterface interface {
	Login(input dto.LoginInput)
	CreateUser(input dto.CreateUserInput)
}

type AuthRepositoryInterface interface {
	Create(user models.User) *models.User
}