package internals

import (
	"example/orderservices/pkg/models"
	"log"

	"gorm.io/gorm"
)

// Migrate the model structs to the database
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Order{}, &models.OrderItem{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Migrated tables successfully...")
}
