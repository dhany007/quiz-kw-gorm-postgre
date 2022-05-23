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
	products := []models.Product{}
	err := repository.DB.Find(&products).Error
	return &products, err
}

func (repository *ProductRepositoryImpl) GetProductByID(id uint) (*models.Product, error) {
	product := models.Product{}
	err := repository.DB.First(&product, "id=?", id).Error

	return &product, err
}

func (repository *ProductRepositoryImpl) UpdateProductByID(p *models.Product, id uint) error {
	product := models.Product{}
	err := repository.DB.Model(&product).Where("id = ?", id).Updates(
		models.Product{
			Name:   p.Name,
			Brand:  p.Brand,
			UserID: p.UserID,
		},
	).Error

	return err
}

func (repository *ProductRepositoryImpl) DeleteProductByID(id uint) error {
	product := models.Product{}
	err := repository.DB.Where("id=?", id).Delete(&product, id).Error

	return err
}
