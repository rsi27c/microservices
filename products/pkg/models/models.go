package models

type Products struct {
	ProductID   int      `json:"product_id" gorm:"primary_key;"`
	ProductName string   `json:"product_name"`
	Price       int      `json:"price"`
	Sellers     []Seller `json:"orders,omitempty" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Seller struct {
	SellerID    int    `json:"seller_id" gorm:"primary_key"`
	SellerName  string `json:"seller_name"`
	ProductID   int    `json:"product_id"`
	PhoneNumber int    `json:"phone_number"`
}
