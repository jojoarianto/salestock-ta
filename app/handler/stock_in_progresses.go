package handler

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"salestock-ta/app/model"
)

// handler for get all progress stock in
func GetAllProgressStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter url
	i, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	progress := []model.StockInProgress{}

	db.Find(&progress, model.StockInProgress{StockInsID: i})
	respondWithJson(w, http.StatusOK, progress)
}

// handler for post data to create a progress stock in
func CreateProgressStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter

	id := vars["ID"]                          // get id
	stockin := GetStockinsOr404(db, id, w, r) // make sure stock_in_parent record is exist
	if stockin == nil {
		return // record not found
	}

	//====== BEGIN TRANSACTION

	// tx := db.Begin()

	// insert progress
	progress := model.StockInProgress{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&progress); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&progress).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//====== END TRANSACTION

	respondWithJson(w, http.StatusOK, &progress)
}
