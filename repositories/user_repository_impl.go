package repositories

import (
	"belajar-gorm-postgre/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) CreateUser() error {
	return nil
}

func (repository *UserRepositoryImpl) GetAllUsers() (*[]models.User, error) {
	return nil, nil
}

func (repository *UserRepositoryImpl) GetUserByID(id uint) (*models.User, error) {
	return nil, nil
}
