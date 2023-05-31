package controllers

import (
	"authentication/dto"
	"authentication/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type authenticationController struct {
	authenticationService interfaces.AuthServiceInterface
}

func (authController *authenticationController) CreateUser(c *gin.Context) {
	var createUserInput dto.CreateUserInput
	err := c.ShouldBindWith(&createUserInput, binding.JSON)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid format",
		})
		return
	}

	authController.authenticationService.CreateUser(createUserInput)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

func (authController *authenticationController) Login(c *gin.Context) {
	var loginInput dto.LoginInput
	err := c.ShouldBindWith(&loginInput, binding.JSON)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid format",
		})
		return
	}

	authController.authenticationService.Login(loginInput)
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in ",
	})
}

func GetAuthController(service interfaces.AuthServiceInterface) interfaces.AuthControllerInterface {
	return &authenticationController{
		authenticationService: service,
	}
}
