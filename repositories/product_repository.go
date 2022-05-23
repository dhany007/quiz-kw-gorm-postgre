package repositories

import "belajar-gorm-postgre/models"

type ProductRepository interface {
	CreateProduct(*models.Product) error
	GetAllProducts() (*[]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProductByID(p *models.Product, id uint) error
	DeleteProductByID(id uint) error
}
