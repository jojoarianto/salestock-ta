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

// Note : Gorm not work foreignkey with sqlite (https://github.com/jinzhu/gorm/issues/635)
// Data structure for stock_in
type Stock_ins struct {
	gorm.Model                 // add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Transaction_time time.Time `json:"transaction_time"`
	ProductID        uint      `json:"product_id" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to product
	Product          Product   //`gorm:"foreignkey:ProductID"`                                                                      // belongs to product
	Order_qty        int       `json:"order_qty"`
	Received_qty     int       `json:"received_qty"`
	Purchase_price   int       `json:"purchase_price"`
	Total_price      int       `json:"total_price"`
	Receipt          string    `json:"receipt"`
}
