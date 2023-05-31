package services

import (
	"authentication/config"
	"authentication/dto"
	"authentication/interfaces"
	"authentication/models"

	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	repository interfaces.AuthRepositoryInterface
}

func (authService *authenticationService) Login(input dto.LoginInput) {
	panic("unimplemented")
}

func (authService *authenticationService) CreateUser(input dto.CreateUserInput) *models.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), config.GetBcryptConfig().DEFAULT_COST)
	if err != nil {
		panic("Password hashing failed")
	}

	user := authService.repository.Create(models.User{
		Username: input.Username,
		Name:     input.Name,
		Password: string(hashedPassword),
	})

	return user
}

func GetAuthService(repo interfaces.AuthRepositoryInterface) interfaces.AuthServiceInterface {
	return &authenticationService{
		repository: repo,
	}
}
