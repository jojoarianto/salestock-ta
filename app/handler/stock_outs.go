package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
	stockout := model.StockOut{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockout); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Search the product first
	vars := mux.Vars(r)
	product_id := vars["product_id"]
	product := model.Product{}

	if err := db.First(&product, product_id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error()) // print record not found
		return
	}
	stockout.Product = product

	if err := db.Save(&stockout).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, stockout)
}
