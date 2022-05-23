package repositories

import "belajar-gorm-postgre/models"

type UserRepository interface {
	CreateUser(*models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUserByID(user *models.User, id uint) error
}
