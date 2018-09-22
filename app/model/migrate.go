package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{}, &StockIn{}, StockInProgress{})
	log.Println("Migration has been processed")

	// Note : Gorm not work foreignkey with sqlite (https://github.com/jinzhu/gorm/issues/635)
	// db.Model(&StockIns{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")

	return db
}
