package services

import (
	"example/productservices/repository"
	"example/productservices/pkg/models"
)

type Services interface {
	CreateProduct(product *models.Products) error
	GetProduct() (*[]models.Products, error)
	GetProductByID(id int) (*models.Products, error)
	UpdateProduct(product *models.Products, id int) error
	DeleteProduct(id int) error
}

type Service struct {
	repository.Repositories
}

func InitServices(repo repository.Repositories) Services {
	return &Service{Repositories: repo}
}

func (service *Service) CreateProduct(product *models.Products) error {
	return service.Repositories.CreateProduct(product)
}

func (service *Service) GetProduct() (*[]models.Products, error) {
	return service.Repositories.GetProduct()
}

func (service *Service) GetProductByID(id int) (*models.Products, error) {
	return service.Repositories.GetProductByID(id)
}

func (service *Service) UpdateProduct(product *models.Products, id int) error {
	return service.Repositories.UpdateProduct(product, id)
}

func (service *Service) DeleteProduct(id int) error {
	return service.Repositories.DeleteProduct(id)
}
