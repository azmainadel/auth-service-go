package dto

type CreateUserInput struct {
	Username string
	Name     string
	Password string
}

type LoginInput struct {
	Username string
	Password string
}
