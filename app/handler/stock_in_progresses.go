package handler

import (
	// "fmt"
	"encoding/json"
	// "log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"salestock-ta/app/model"
)

// handler for get all progress stock in
func GetAllProgressStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get parameter url
	i, err := strconv.Atoi(vars["stock_in_id"])
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
	vars := mux.Vars(r) // get parameter url
	stock_in_id, err := strconv.Atoi(vars["stock_in_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	stockin := GetStockinsOr404(db, stock_in_id, w, r) // make sure stock_in_parent record is exist
	if stockin == nil {
		return // record not found
	}

	//====== BEGIN TRANSACTION
	// 1. Create progress_stock_in
	// 2. Update stock_in quantity received
	// 3. Update stockin stock in status code (When order_qty == purchase_price)
	// 4. Update stock on product

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

	progress := model.StockInProgress{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&progress); err != nil {
		tx.Rollback()
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	progress.StockInsID = stock_in_id // set stock_in_id

	validate := validator.New() // validation
	if err := validate.Struct(progress); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// 1. Create progress_stock_in
	if err := db.Create(&progress).Error; err != nil {
		tx.Rollback()
		return
	}

	// 2. Update stock_in quantity received
	stockin.ReceivedQty = stockin.ReceivedQty + progress.Qty
	// 3. Update stockin stock in status code (When order_qty == purchase_price)
	if stockin.ReceivedQty >= stockin.OrderQty {
		stockin.StausInCode = 1
	}
	if err := db.Save(&stockin).Error; err != nil {
		tx.Rollback()
		return
	}

	// 4. Update stock on product
	product := stockin.Product
	product.Stocks = product.Stocks + progress.Qty
	if err := db.Save(&product).Error; err != nil {
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error Transaction")
		return
	}
	//====== END TRANSACTION

	respondWithJson(w, http.StatusOK, &stockin)
}
