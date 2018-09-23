package handler

import (
	"encoding/json"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"salestock-ta/app/model"
)

// handler for get all data stock out
func GetAllStockOuts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockout := []model.StockOut{}

	db.Preload("Product").Find(&stockout)
	respondWithJson(w, http.StatusOK, stockout)
}

// handler for create a single stock out
func CreateStockOuts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	log.Println("Barang Keluar")
	stockout := model.StockOut{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockout); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	validate := validator.New() // validation
	if err := validate.Struct(stockout); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// validasi product, make sure product is exist
	product := model.Product{}
	if err := db.First(&product, stockout.ProductID).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Product record not found") // print record not found
		return
	}

	stockout.Product = product

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
