package interfaces

import "github.com/gin-gonic/gin"

type AuthControllerInterface interface {
	Login(c *gin.Context)
	CreateUser(c *gin.Context)
}
