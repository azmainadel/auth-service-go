package services

import (
	"authentication/dto"
	"authentication/interfaces"
)

type authenticationService struct {
}

func (*authenticationService) Login(input dto.LoginInput) {
	panic("unimplemented")
}

func (*authenticationService) CreateUser(input dto.CreateUserInput) {
	panic("unimplemented")
}

func GetAuthService() interfaces.AuthServiceInterface {
	return &authenticationService{}
}
