package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"salestock-ta/app/model"
)

func GetStockIns(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "END POINT GET ALL Stockin")
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
