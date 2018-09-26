package handler

import (
	// "log"
	// "io"
	"encoding/csv"
	"encoding/json"
	// "fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/salestock-ta/app/model"
	"gopkg.in/go-playground/validator.v9"
)

// handler for get all data stock out
func GetAllStockOuts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockout := []model.StockOut{}

	db.Preload("Product").Find(&stockout)
	respondWithJson(w, http.StatusOK, stockout)
}

// handler for get a single data stock out
func GetStockOut(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockout := model.StockOut{}

	vars := mux.Vars(r)

	stock_out_id, err := strconv.Atoi(vars["stock_out_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Preload("Product").Find(&stockout, stock_out_id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, stockout)
}

// handler for create a single stock out
func CreateStockOuts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockout := model.StockOut{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockout); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// validasi product, make sure product is exist
	product := model.Product{}
	if err := db.First(&product, stockout.ProductID).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Product record not found") // print record not found
		return
	}

	stockout.Product = product

	validate := validator.New() // validation
	if err := validate.Struct(stockout); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//====== BEGIN TRANSACTION ========//

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error Transaction")
		return
	}

	// count total_price
	stockout.TotalPrice = stockout.SellPrice * stockout.OutQty
	// create stock out
	if err := tx.Create(&stockout).Error; err != nil {
		tx.Rollback()
		return
	}

	// update product stock
	if (product.Stocks - stockout.OutQty) < 0 {
		tx.Rollback() // error stock not enough
		respondWithError(w, http.StatusInternalServerError, "Stock not enough")
		return
	}
	product.Stocks = product.Stocks - stockout.OutQty
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return
	}

	stockout.Product = product // refill stockout product with the new one

	if err := tx.Commit().Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error Transaction")
		return
	}

	//====== END OF TRANSACTION ========//

	respondWithJson(w, http.StatusCreated, stockout)
}

func DeleteStockOut(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockout := model.StockOut{}

	vars := mux.Vars(r)

	stock_out_id, err := strconv.Atoi(vars["stock_out_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Preload("Product").Find(&stockout, stock_out_id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	//====== BEGIN TRANSACTION ========//

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error Transaction")
		return
	}

	// update stock on product
	product := stockout.Product
	product.Stocks = product.Stocks + stockout.OutQty

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return
	}
	stockout.Product = product

	// deleting prosess
	if err := tx.Delete(&stockout).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit().Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error Transaction")
		return
	}

	//====== END OF TRANSACTION ========//

	respondWithJson(w, http.StatusOK, "Delete success")
}

// handler for export stock out
func ExportCsvStockOuts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockout := []model.StockOut{}
	db.Preload("Product").Find(&stockout)

	csvData, err := os.Create("csv/export_stock_outs.csv")
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	var record []string
	record = append(record, "Waktu")
	record = append(record, "SKU")
	record = append(record, "Nama Barang")
	record = append(record, "Jumlah Keluar")
	record = append(record, "Harga Jual")
	record = append(record, "Total")
	record = append(record, "Catatan")
	writer.Write(record)

	for _, worker := range stockout {
		var record []string
		record = append(record, worker.StockOutTime.Format("2006-01-02 15:04:05"))
		record = append(record, worker.Product.Sku)
		record = append(record, worker.Product.Name)
		record = append(record, strconv.Itoa(worker.OutQty))
		record = append(record, strconv.Itoa(worker.SellPrice))
		record = append(record, strconv.Itoa(worker.TotalPrice))

		switch code := worker.StatusOutCode; code {
		case 1: // terjual
			str := "Pesanan ID-" + worker.Transaction
			record = append(record, str)
		case 2:
			record = append(record, "Barang Hilang")
		case 3:
			record = append(record, "Barang Rusak")
		case 4:
			record = append(record, "Barang Rusak")
		default:
			record = append(record, "")
		}
		writer.Write(record)
	}
	writer.Flush()

	respondWithJson(w, http.StatusOK, "Export stock out to csv success check your export file at csv/export_stock_outs.csv")
}

// handler for export sales report
func ExportCsvSalesReport(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type HargaBeli struct {
		ProductID int
		Qty       int
		Total     int
	}

	// query get totoal received & total price
	query := "select stock_ins.product_id as product_id, SUM(stock_ins.order_qty) as qty, SUM(stock_ins.total_price) as total from stock_ins GROUP BY stock_ins.product_id;"

	var hargabeli []HargaBeli
	db.Raw(query).Scan(&hargabeli)
	HargaBeliBarang := make(map[int]int)
	for _, w := range hargabeli {
		HargaBeliBarang[w.ProductID] = w.Total / w.Qty
	}

	stockout := []model.StockOut{}
	db.Preload("Product").Where("status_out_code = ?", 1).Find(&stockout)

	csvData, err := os.Create("csv/export_sales_report.csv")
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	var record []string
	record = append(record, "ID Pesanan")
	record = append(record, "Waktu")
	record = append(record, "SKU")
	record = append(record, "Nama Barang")
	record = append(record, "Jumlah")
	record = append(record, "Harga Jual")
	record = append(record, "Total")
	record = append(record, "Harga Beli")
	record = append(record, "Laba")
	writer.Write(record)

	for _, worker := range stockout {
		var record []string
		record = append(record, worker.Transaction)
		record = append(record, worker.StockOutTime.Format("2006-01-02 15:04:05"))
		record = append(record, worker.Product.Sku)
		record = append(record, worker.Product.Name)
		record = append(record, strconv.Itoa(worker.OutQty))
		record = append(record, strconv.Itoa(worker.SellPrice))
		record = append(record, strconv.Itoa(worker.TotalPrice))

		if beli, ok := HargaBeliBarang[worker.ProductID]; ok {
			record = append(record, strconv.Itoa(beli)) // nilai beli rata rata barang
			// laba = total - (harga beli x jumlah)
			laba := worker.TotalPrice - (beli * worker.OutQty)
			record = append(record, strconv.Itoa(laba))
		} else {
			record = append(record, "0") // if empty stock
			record = append(record, "0")
		}

		writer.Write(record)
	}
	writer.Flush()

	respondWithJson(w, http.StatusOK, "Export sales report to csv success check your export file at csv/export_sales_report.csv")
}
