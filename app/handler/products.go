package handler

import (
	// "fmt"
	// "log"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// update a product
func UpdateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	product_id, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	product := GetProductOr404(db, product_id, w, r)
	if product == nil {
		return
	}
}

// delete a product
func DeleteProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	product_id, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	product := GetProductOr404(db, product_id, w, r)
	if product == nil {
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, "Delete product success")
}

// GetProductOr404 gets a product instance if exists, or respond the 404 error otherwise
func GetProductOr404(db *gorm.DB, product_id int, w http.ResponseWriter, r *http.Request) *model.Product {
	product := model.Product{}
	if err := db.First(&product, product_id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}
	return &product
}
