package repositories

import (
	"authentication/interfaces"
	"authentication/models"

	"gorm.io/gorm"
)

type authenticationRepository struct {
	database *gorm.DB
}

func (authRepository *authenticationRepository) Create(user models.User) *models.User {
	authRepository.database.Save(&user)
	return &user
}

func GetAuthenticationRepository(db *gorm.DB) interfaces.AuthRepositoryInterface {
	return &authenticationRepository{
		database: db,
	}
}
