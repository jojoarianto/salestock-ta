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
type StockIn struct { // BARANG MASUK
	gorm.Model                      // add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	StockInTime   time.Time         `json:"stock_in_time"`
	ProductID     int               `json:"product_id" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to product
	Product       Product           `json:"product"`
	OrderQty      int               `json:"order_qty"`
	ReceivedQty   int               `json:"received_qty"`
	PurchasePrice int               `json:"purchase_price"`
	TotalPrice    int               `json:"total_price"`
	Receipt       string            `json:"receipt"`
	Progress      []StockInProgress `gorm:"ForeignKey:StockInsID" json:"progress"`
	StausOutCode  int               `json:"status_in_code"` // 1. waiting, 2 completed
}

// Data structure for stock_in_progress
type StockInProgress struct {
	gorm.Model               // add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	ProgressInTime time.Time `json:"stock_in_progress_time"`
	StockInsID     int       `json:"stock_ins_id" sql:"type:integer REFERENCES stock_ins(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to stock in
	Qty            int       `json:"qty"`
}

type StockOut struct {
	gorm.Model
	StockOutTime time.Time `json:"stock_out_time"`
	ProductID    int       `json:"product_id" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to product
	Product      Product   `json:"product"`
	OutQty       int       `json:"out_qty"`
	SellPrice    int       `json:"sell_price"`
	TotalPrice   int       `json:"total_price"`
	Transaction  string    `json:"transaction_id"`  // catatan null if barang tidak terjual
	StausOutCode int       `json:"status_out_code"` // 1. Terjual, 2. Barang Hlang, 3. Rusak, 4 Barang Sample
}
