package services

import (
	"authentication/dto"
	"authentication/interfaces"
	"authentication/models"
)

type authenticationService struct {
	repository interfaces.AuthRepositoryInterface
}

func (authService *authenticationService) Login(input dto.LoginInput) {
	panic("unimplemented")
}

func (authService *authenticationService) CreateUser(input dto.CreateUserInput) {
	user := authService.repository.Create(models.User{
		Username: input.Username,
		Name:     input.Name,
		Password: input.Password,
	})

	return user
}

func GetAuthService(repo interfaces.AuthRepositoryInterface) interfaces.AuthServiceInterface {
	return &authenticationService{
		repository: repo,
	}
}
