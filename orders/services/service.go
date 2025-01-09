package services

import (
	"example/orderservices/implementation"
	"example/orderservices/pkg/models"
)

type Services interface {
	CreateOrder(order *models.Order) error
	GetOrder() (*[]models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	UpdateOrder(order *models.Order, id int) error
	DeleteOrder(id int) error
}

type Service struct {
	implementation.Repositories
}

func InitServices(repo implementation.Repositories) Services {
	return &Service{Repositories: repo}
}

func (service *Service) CreateOrder(order *models.Order) error {
	return service.Repositories.CreateOrder(order)
}

func (service *Service) GetOrder() (*[]models.Order, error) {
	return service.Repositories.GetOrder()
}

func (service *Service) GetOrderByID(id int) (*models.Order, error) {
	return service.Repositories.GetOrderByID(id)
}

func (service *Service) UpdateOrder(order *models.Order, id int) error {
	return service.Repositories.UpdateOrder(order, id)
}

func (service *Service) DeleteOrder(id int) error {
	return service.Repositories.DeleteOrder(id)
}
