package handler

import (
	// "fmt"
	// "log"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"salestock-ta/app/model"
)

// get all product
func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.Product{} // array of product
	db.Find(&products)
	respondWithJson(w, http.StatusOK, products)
}

// create a product
func CreateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, product)
}
