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

	if err := db.Save(&stockin).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, stockin)
}

// handler for delete data stockinby id
func DeleteStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter

	id := vars["ID"]
	stockin := getStockinsOr404(db, id, w, r)
	if stockin == nil {
		return // record not found
	}

	if err := db.Delete(&stockin).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, nil)
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
