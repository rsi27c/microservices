package repository

import (
	"example/productservices/pkg/models"

	"gorm.io/gorm"
)

type Repositories interface {
	CreateProduct(product *models.Products) error
	GetProduct() (*[]models.Products, error)
	GetProductByID(id int) (*models.Products, error)
	UpdateProduct(product *models.Products, id int) error
	DeleteProduct(id int) error
}

type Repository struct {
	DB *gorm.DB
}

func InitRepository(db *gorm.DB) Repositories {
	return &Repository{DB: db}
}

func (repo *Repository) CreateProduct(product *models.Products) error {
	data := repo.DB.Create(product)
	if data.Error != nil {
		return data.Error
	}

	return nil
}

func (repo *Repository) GetProduct() (*[]models.Products, error) {
	var product []models.Products

	data := repo.DB.Model(product).Find(&product)
	if data.Error != nil {
		return nil, data.Error
	}

	return &product, nil
}

func (repo *Repository) GetProductByID(id int) (*models.Products, error) {
	var product models.Products

	data := repo.DB.Where("product_id=?", id).First(&product)
	if data.Error != nil {
		return nil, data.Error
	}

	return &product, nil
}

func (repo *Repository) UpdateProduct(product *models.Products, id int) error {
	data := repo.DB.Where("product_id=?", id).Updates(&product)
	if data.Error != nil {
		return data.Error
	}

	return nil
}

func (repo *Repository) DeleteProduct(id int) error {
	data := repo.DB.Where("product_id=?", id).Delete(&models.Products{})
	if data.Error != nil {
		return data.Error
	}

	return nil
}
