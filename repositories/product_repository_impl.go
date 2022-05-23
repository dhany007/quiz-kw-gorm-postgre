package repositories

import (
	"belajar-gorm-postgre/models"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductRepositoryImpl) CreateProduct(product *models.Product) error {
	err := repository.DB.Create(product).Error
	return err
}

func (repository *ProductRepositoryImpl) GetAllProducts() (*[]models.Product, error) {
	return nil, nil
}

func (repository *ProductRepositoryImpl) GetProductByID(id uint) (*models.Product, error) {
	return nil, nil
}

func (repository *ProductRepositoryImpl) UpdateProductByID(p *models.Product, id uint) error {
	return nil
}

func (repository *ProductRepositoryImpl) DeleteProductByID(id uint) error {
	return nil
}
