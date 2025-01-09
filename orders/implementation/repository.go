package implementation

import (
	"example/orderservices/pkg/models"

	"gorm.io/gorm"
)

type Repositories interface {
	CreateOrder(order *models.Order) error
	GetOrder() (*[]models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	UpdateOrder(order *models.Order, id int) error
	DeleteOrder(id int) error
}

type Repository struct {
	DB *gorm.DB
}

func InitRepository(db *gorm.DB) Repositories {
	return &Repository{DB: db}
}

func (repo *Repository) CreateOrder(order *models.Order) error {
	data := repo.DB.Create(order)
	if data.Error != nil {
		return data.Error
	}

	return nil
}

func (repo *Repository) GetOrder() (*[]models.Order, error) {
	var orders []models.Order

	data := repo.DB.Model(orders).Find(&orders)
	if data.Error != nil {
		return nil, data.Error
	}

	return &orders, nil
}

func (repo *Repository) GetOrderByID(id int) (*models.Order, error) {
	var order models.Order

	data := repo.DB.Where("order_id=?", id).First(&order)
	if data.Error != nil {
		return nil, data.Error
	}

	return &order, nil
}

func (repo *Repository) UpdateOrder(order *models.Order, id int) error {
	data := repo.DB.Where("order_id=?", id).Updates(&order)
	if data.Error != nil {
		return data.Error
	}

	return nil
}

func (repo *Repository) DeleteOrder(id int) error {
	data := repo.DB.Where("order_id=?", id).Delete(&models.Order{})
	if data.Error != nil {
		return data.Error
	}

	return nil
}
