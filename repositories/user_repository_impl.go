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

func (repository *UserRepositoryImpl) CreateUser(user *models.User) error {
	err := repository.DB.Create(user).Error

	return err
}

func (repository *UserRepositoryImpl) GetAllUsers() (*[]models.User, error) {
	users := []models.User{}
	err := repository.DB.Find(&users).Error

	return &users, err
}

func (repository *UserRepositoryImpl) GetUserByID(id uint) (*models.User, error) {
	user := models.User{}
	err := repository.DB.Find(&user, "id=?", id).Error

	return &user, err
}
