package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
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

	validate := validator.New() // validation
	if err := validate.Struct(product); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

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

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	validate := validator.New() // validation
	if err := validate.Struct(product); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Save(&product).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, product)
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

func ExportCsvProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.Product{}
	db.Find(&products)

	csvData, err := os.Create("csv/export_products.csv")
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	var record []string
	record = append(record, "SKU")
	record = append(record, "Nama Item")
	record = append(record, "Jumlah Sekarang")
	writer.Write(record)

	for _, worker := range products {
		var record []string
		record = append(record, worker.Sku)
		record = append(record, worker.Name)
		record = append(record, strconv.Itoa(worker.Stocks))
		writer.Write(record)
	}
	writer.Flush()

	respondWithJson(w, http.StatusOK, "Export products to csv success check your export file at csv/export_products.csv")
}

// handler function to import products
func ImportCsvProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type ProductInsert struct {
		Sku  string
		Name string
	}
	csvData, _ := os.Open("csv/import_products.csv")

	sqlStr := ""
	reader := csv.NewReader(csvData)
	var product []ProductInsert
	for {
		line, error := reader.Read()
		if error == io.EOF { // end of line
			break
		} else if error != nil {
			log.Fatal(error)
		}

		product = append(product, ProductInsert{
			Sku:  line[0],
			Name: line[1],
		})

		sqlStr += fmt.Sprintf("INSERT INTO products (created_at, updated_at, sku, name) VALUES (datetime('now','localtime'), datetime('now','localtime'),'%s', '%s'); ", line[0], line[1])
	}

	db.Exec(sqlStr)
	peopleJson, _ := json.Marshal(product)
	respondWithJson(w, http.StatusOK, string(peopleJson))
}
