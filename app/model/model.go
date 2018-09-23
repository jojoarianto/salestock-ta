package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Data structure for product
type Product struct { // BARANG
	gorm.Model
	Sku    string `gorm:"type:varchar(100);unique_index" json:"sku"`
	Name   string `json:"name"`
	Stocks int    `json:"stocks"`
}

// Note : Gorm not work foreignkey with sqlite3 (https://github.com/jinzhu/gorm/issues/635)
// Struct for stock_in (data barang masuk)
type StockIn struct { // BARANG MASUK
	gorm.Model
	StockInTime   time.Time         `json:"stock_in_time"`
	ProductID     int               `json:"product_id" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to product
	Product       Product           `json:"product"`
	OrderQty      int               `json:"order_qty"`
	ReceivedQty   int               `json:"received_qty"`
	PurchasePrice int               `json:"purchase_price"`
	TotalPrice    int               `json:"total_price"`
	Receipt       string            `json:"receipt"`
	Progress      []StockInProgress `gorm:"ForeignKey:StockInsID" json:"progress"`
	StausInCode   int               `json:"status_in_code"` // 0. waiting, 1 completed
}

// Struct for stock_in_progress (progress data barang masuk)
type StockInProgress struct { // PROGRESS BARANG MASUK
	gorm.Model
	ProgressInTime time.Time `validate:"required" json:"stock_in_progress_time"`
	StockInsID     int       `validate:"required,numeric,min=1" json:"stock_ins_id" sql:"type:integer REFERENCES stock_ins(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to stock in
	Qty            int       `validate:"required,numeric,min=1" json:"qty"`
}

// Struct for stock_out
type StockOut struct { // BARANG KELUAR
	gorm.Model
	StockOutTime time.Time `validate:"required" json:"stock_out_time"`
	ProductID    int       `validate:"required,numeric,min=1" json:"product_id" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"` // belongs to product
	Product      Product   `json:"product"`
	OutQty       int       `validate:"required,numeric" json:"out_qty"`
	SellPrice    int       `validate:"omitempty,numeric" json:"sell_price"`
	TotalPrice   int       `json:"total_price"`
	Transaction  string    `json:"transaction_id"`                                    // transaction null if barang tidak terjual
	StausOutCode int       `validate:"required,numeric,min=1" json:"status_out_code"` // 1. Terjual, 2. Barang Hilang, 3. Rusak, 4 Barang Sample
}
