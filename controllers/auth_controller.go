package controllers

import (
	"authentication/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authenticationController struct {
}

// CreateUser implements interfaces.AuthControllerInterface.
func (*authenticationController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

// Login implements interfaces.AuthControllerInterface.
func (*authenticationController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in ",
	})
}

func GetAuthController() interfaces.AuthControllerInterface {
	return &authenticationController{}
}
