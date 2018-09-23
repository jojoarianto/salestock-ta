package app

import (
	"log"
	"net/http"

	"salestock-ta/app/handler"
	"salestock-ta/app/model"
	"salestock-ta/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm" // Go ORM (Object Relational Mapping) for sql
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// app has router using mux and db instance
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// init app with predefined app configuration
func (a *App) Initialize(config *config.Config) {

	db, err := gorm.Open(config.DB.Dialeg, config.DB.DBUri)
	if err != nil {
		log.Fatal("Fail to connect database %s", err.Error())
	}

	log.Printf("Database connected")

	a.DB = model.DBMigrate(db) // migration
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling products
	a.Router.HandleFunc("/api/products", a.GetAllProducts).Methods("GET")
	a.Router.HandleFunc("/api/products", a.CreateProduct).Methods("POST")
	a.Router.HandleFunc("/api/products/{product_id}", a.UpdateProduct).Methods("PUT")
	a.Router.HandleFunc("/api/products/{product_id}", a.DeleteProduct).Methods("DELETE")

	// Routing for handling stock_in
	a.Router.HandleFunc("/api/stock-ins", a.GetAllStockIns).Methods("GET")
	a.Router.HandleFunc("/api/stock-ins/{id}", a.GetStockIn).Methods("GET")
	a.Router.HandleFunc("/api/stock-ins", a.CreateStockIns).Methods("POST")
	a.Router.HandleFunc("/api/stock-ins/{id}", a.UpdateStockIns).Methods("PUT")
	a.Router.HandleFunc("/api/stock-ins/{id}", a.DeleteStockIns).Methods("DELETE")

	// Routing for handling stock_in_progress
	a.Router.HandleFunc("/api/stock-ins/{stock_in_id}/progress", a.GetAllProgressStockIns).Methods("GET")
	a.Router.HandleFunc("/api/stock-ins/{stock_in_id}/progress", a.CreateProgressStockIns).Methods("POST")

	// Routing for handling stock_out
	a.Router.HandleFunc("/api/stock-outs", a.GetAllStockOuts).Methods("GET")
	a.Router.HandleFunc("/api/stock-outs", a.CreateStockOuts).Methods("POST")

	// Routing for export csv handling
	a.Router.HandleFunc("/export/products", a.ExportCsvProducts).Methods("GET")
	a.Router.HandleFunc("/export/stock-ins", a.ExportCsvStockIns).Methods("GET")
	a.Router.HandleFunc("/export/stock-outs", a.ExportCsvStockOuts).Methods("GET")

	// Routing for migration
	a.Router.HandleFunc("/import/products", a.ImportCsvProducts).Methods("GET")
}

// =============================
//    PRODUCT
// =============================

// handler for get all product
func (a *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProducts(a.DB, w, r)
}

// handler for create a product
func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(a.DB, w, r)
}

// handler for update a product
func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	handler.UpdateProduct(a.DB, w, r)
}

// handler for update a product
func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	handler.DeleteProduct(a.DB, w, r)
}

// =============================
//    STOCK IN / BARANG MASUK
// =============================

// handler for get all stock in
func (a *App) GetAllStockIns(w http.ResponseWriter, r *http.Request) {
	handler.GetAllStockIns(a.DB, w, r)
}

// handler for get single stock in
func (a *App) GetStockIn(w http.ResponseWriter, r *http.Request) {
	handler.GetStockIn(a.DB, w, r)
}

// handler for create stock_in
func (a *App) CreateStockIns(w http.ResponseWriter, r *http.Request) {
	handler.CreateStockIns(a.DB, w, r)
}

// handler for create stock_in
func (a *App) UpdateStockIns(w http.ResponseWriter, r *http.Request) {
	handler.UpdateStockIns(a.DB, w, r)
}

// handler for delete stock_in
func (a *App) DeleteStockIns(w http.ResponseWriter, r *http.Request) {
	handler.DeleteStockIns(a.DB, w, r)
}

// =============================
//    PROGRESS STOCK IN / PROGRESS BARANG MASUK
// =============================

// handler for get all [rpgress] stock in
func (a *App) GetAllProgressStockIns(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProgressStockIns(a.DB, w, r)
}

// handler for create a single progress stock in
func (a *App) CreateProgressStockIns(w http.ResponseWriter, r *http.Request) {
	handler.CreateProgressStockIns(a.DB, w, r)
}

// =============================
//    STOCK OUT / BARANG KELUAR
// =============================

// handler for get all stock out
func (a *App) GetAllStockOuts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllStockOuts(a.DB, w, r)
}

// handler for create stock_out
func (a *App) CreateStockOuts(w http.ResponseWriter, r *http.Request) {
	handler.CreateStockOuts(a.DB, w, r)
}

// =============================
//    EXPORT CSV
// =============================

func (a *App) ExportCsvProducts(w http.ResponseWriter, r *http.Request) {
	handler.ExportCsvProducts(a.DB, w, r)
}

func (a *App) ExportCsvStockIns(w http.ResponseWriter, r *http.Request) {
	handler.ExportCsvStockIns(a.DB, w, r)
}

func (a *App) ExportCsvStockOuts(w http.ResponseWriter, r *http.Request) {
	handler.ExportCsvStockOuts(a.DB, w, r)
}

// =============================
//    IMPORT CSV
// =============================

//  handler for migrate data product
func (a *App) ImportCsvProducts(w http.ResponseWriter, r *http.Request) {
	handler.ImportCsvProducts(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Printf("Start a server on port %s", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
