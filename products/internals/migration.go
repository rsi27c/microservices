package internals

import (
	"example/productservices/pkg/models"
	"log"

	"gorm.io/gorm"
)

// Migrate the model structs to the database
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Products{}, &models.Seller{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Migrated tables successfully...")
}
