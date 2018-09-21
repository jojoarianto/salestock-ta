package model

import (
	"time"

	"github.com/jinzhu/gorm" // for object relational mapping
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Data structure for product
type Product struct {
	gorm.Model        // add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Sku        string `gorm:"type:varchar(100);unique_index" json:"sku"`
	Name       string `json:"name"`
	Stocks     uint   `json:"stocks"`
}

// Data structure for stock_in
type Stock_ins struct {
	gorm.Model                 // add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Transaction_time time.Time `json:"transaction_time"`
	Product          Product   // belongs to
	ProductID        int       `json:"product_id"` // belongs to
	Order_qty        int       `json:"order_qty"`
	Received_qty     int       `json:"received_qty"`
	Purchase_price   int       `json:"purchase_price"`
	Total_price      int       `json:"total_price"`
	Receipt          string    `json:"receipt"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Stock_ins{})
	return db
}
