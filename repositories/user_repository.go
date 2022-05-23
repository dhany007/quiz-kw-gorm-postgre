package repositories

import "belajar-gorm-postgre/models"

type UserRepository interface {
	CreateUser() error
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id uint) (*models.User, error)
}
