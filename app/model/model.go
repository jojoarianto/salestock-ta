package model

import (
	"github.com/jinzhu/gorm" // for object relational mapping
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Data structure for product
// Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type Product struct {
	gorm.Model
	Sku    string `json:"sku"`
	Name   string `json:"name"`
	Stocks uint   `json:"stocks"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{})
	return db
}
