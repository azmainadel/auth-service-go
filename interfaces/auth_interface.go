package interfaces

import (
	"authentication/dto"

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
