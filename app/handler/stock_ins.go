package handler

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"salestock-ta/app/model"
)

// handler for get all data stockin
func GetStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockin := []model.Stock_ins{} // array of stock_in

	db.Find(&stockin) // get all stock in
	respondWithJson(w, http.StatusOK, stockin)
}

func CreateStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockin := model.Stock_ins{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockin); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// search the product first
	vars := mux.Vars(r) // get parameter
	product_id := vars["product_id"]
	product := model.Product{}
	if err := db.First(&product, product_id).Error; err != nil { // Get record with primary key (only works for integer primary key)
		respondWithError(w, http.StatusNotFound, err.Error()) // print record not found
		return
	}
	// product := getProductsOr404(db, product_id, w, r) // make sure record is exist

	stockin.Product = product

	if err := db.Save(&stockin).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, stockin)
}

// handler for delete data stockins by id
func DeleteStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter

	id := vars["ID"]                          // get id
	stockin := getStockinsOr404(db, id, w, r) // make sure record is exist
	if stockin == nil {
		return // record not found
	}

	if err := db.Delete(&stockin).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, nil)
}

func UpdateStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter

	id := vars["ID"]                          // get id
	stockin := getStockinsOr404(db, id, w, r) // make sure record is exist
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
func getStockinsOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Stock_ins {
	stockin := model.Stock_ins{}

	if err := db.First(&stockin, id).Error; err != nil { // Get record with primary key (only works for integer primary key)
		respondWithError(w, http.StatusNotFound, err.Error()) // print record not found
		return nil
	}
	return &stockin
}
