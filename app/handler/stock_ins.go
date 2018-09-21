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
func GetAllStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	stockin := []model.Stock_ins{}

	db.Preload("Product").Find(&stockin)
	respondWithJson(w, http.StatusOK, stockin)
}

// handler for get single ddata stockin
func GetStockIn(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter url
	id := vars["id"]

	stockin := model.Stock_ins{}
	if err := db.Preload("Product").Find(&stockin, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
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
	vars := mux.Vars(r)
	product_id := vars["product_id"]
	product := model.Product{}

	// Search product
	// Get record with primary key (only works for integer primary key)
	if err := db.First(&product, product_id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error()) // print record not found
		return
	}
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
