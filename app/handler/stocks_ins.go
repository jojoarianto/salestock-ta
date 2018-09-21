package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "END POINT GET ALL PRODUCT")
}
