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
	gorm.Model               // add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Stock_in_time  time.Time `json:"stock_in_time"`
	ProductID      uint      `json:"product_id" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to product
	Product        Product   //`gorm:"foreignkey:ProductID"`                                                                      // belongs to product
	Order_qty      int       `json:"order_qty"`
	Received_qty   int       `json:"received_qty"`
	Purchase_price int       `json:"purchase_price"`
	Total_price    int       `json:"total_price"`
	Receipt        string    `json:"receipt"`
	Progress       []Stock_in_progress
}

// Data structure for stock_in_progress
type Stock_in_progress struct {
	gorm.Model                 // add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Progress_in_time time.Time `json:"stock_in_progress_time"`
	Stock_ins_id     uint      `json:"stock_in" sql:"type:integer REFERENCES stock_ins(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to stock in
	qty              int       `json:"qty"`
}
