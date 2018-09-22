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

	// Validasi product, make sure product is exist
	product := model.Product{}
	if err := db.First(&product, stockout.ProductID).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Product record not found") // print record not found
		return
	}
	stockout.Product = product

	// count total_price
	stockout.TotalPrice = stockout.SellPrice * stockout.OutQty

	// update product stock
	tmp := stockout.Product
	tmp.Stocks = tmp.Stocks - stockout.OutQty
	log.Println(tmp.Stocks)

	if err := db.Save(&tmp).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// create stock out
	if err := db.Save(&stockout).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, stockout)
}
