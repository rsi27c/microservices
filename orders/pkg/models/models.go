package models

import "time"

type Order struct {
	ID         int      `json:"id,omitempty" gorm:"primary_key"`
	CustomerID int      `json:"customer_id"`
	Status     string      `json:"status"`
	CreatedOn  time.Time   `json:"created_on,omitempty" gorm:"autoCreateTime"`
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type OrderItem struct {
	ProductCode string `json:"product_code"`
	Name        string `json:"name"`
	UnitPrice   int    `json:"unit_price"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}
