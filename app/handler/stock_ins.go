package handler

import (
	// "fmt"
	// "log"
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"salestock-ta/app/model"
)

// handler for get all data stockin
func GetAllStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockin := []model.StockIn{}

	db.Preload("Progress").Preload("Product").Find(&stockin)
	respondWithJson(w, http.StatusOK, stockin)
}

// handler for get single data stockin
func GetStockIn(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stock_in_id, err := strconv.Atoi(vars["stock_in_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	stockin := model.StockIn{}
	if err := db.Preload("Progress").Preload("Product").Find(&stockin, stock_in_id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, stockin)
}

// handler for create a single stock in
func CreateStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockin := model.StockIn{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockin); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	product_id := stockin.ProductID

	product := model.Product{}
	if err := db.First(&product, product_id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error()) // print record not found
		return
	}
	stockin.Product = product

	validate := validator.New() // validation
	if err := validate.Struct(stockin); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// count total price
	stockin.TotalPrice = stockin.PurchasePrice * stockin.OrderQty

	if err := db.Save(&stockin).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, stockin)
}

// handler for delete data stockins by id
func DeleteStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter url
	stock_in_id, err := strconv.Atoi(vars["stock_in_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	stockin := GetStockinsOr404(db, stock_in_id, w, r) // make sure record is exist
	if stockin == nil {
		return // record not found
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
	product := stockin.Product
	if (product.Stocks - stockin.ReceivedQty) < 0 {
		tx.Rollback() // error stock not enough
		respondWithError(w, http.StatusInternalServerError, "Stock not enough")
		return
	}
	product.Stocks = product.Stocks - stockin.ReceivedQty

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return
	}
	stockin.Product = product

	// deleting prosess
	if err := tx.Delete(&stockin).Error; err != nil {
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

func UpdateStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter url
	stock_in_id, err := strconv.Atoi(vars["stock_in_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	stockin := GetStockinsOr404(db, stock_in_id, w, r) // make sure record is exist
	if stockin == nil {
		return // record not found
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockin); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&stockin).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, stockin)
}

// getStockinsOr404 gets a stockin instance if exists, or respond the 404 error otherwise
func GetStockinsOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.StockIn {
	stockin := model.StockIn{}

	if err := db.Preload("Product").First(&stockin, id).Error; err != nil { // Get record with primary key (only works for integer primary key)
		respondWithError(w, http.StatusNotFound, "Record stock in not found") // print record not found
		return nil
	}
	return &stockin
}

func ExportCsvStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type Stock_in_report struct {
		Waktu time.Time
		Sku   string
		Nama  string
	}

	var stockin []Stock_in_report
	db.Raw("SELECT stock_ins.stock_in_time, products.sku, products.name FROM stock_ins INNER JOIN products ON stock_ins.product_id=products.id").Scan(&stockin)

	csvData, err := os.Create("csv/export_stock_ins.csv")
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	var record []string
	record = append(record, "Waktu")
	record = append(record, "SKU")
	record = append(record, "Nama Barang")
	// record = append(record, "Jumlah Pemesanan")
	// record = append(record, "Jumlah Diterima")
	writer.Write(record)

	for _, worker := range stockin {
		var record []string
		record = append(record, worker.Waktu.String())
		record = append(record, worker.Sku)
		record = append(record, worker.Nama)
		// record = append(record, strconv.Itoa(worker.OrderQty))
		// record = append(record, strconv.Itoa(worker.ReceivedQty))

		writer.Write(record)
	}
	writer.Flush()

	respondWithJson(w, http.StatusOK, "Export stock ins to csv success")
}
