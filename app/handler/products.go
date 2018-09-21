package handler

import (
	// "fmt"
	// "log"
	"net/http"

	"github.com/jinzhu/gorm"
	"salestock-ta/app/model"
)

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.Product{}
	db.Find(&products)
	respondWithJson(w, http.StatusOK, products)
}
