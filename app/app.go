package app

import (
	"log"
	"net/http"

	"salestock-ta/app/handler"
	// "salestock-ta/app/model"
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

	a.DB = db // set db

	// a.DB = model.DBMigrate(db) // open for migration purpose only
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the products
	a.Router.HandleFunc("/api/products", a.GetAllProducts).Methods("GET")

	// Routing for handling stock_in
	a.Router.HandleFunc("/api/stock-ins", a.GetStockIns).Methods("GET")
	a.Router.HandleFunc("/api/stock-ins", a.CreateStockIns).Methods("POST")
	a.Router.HandleFunc("/api/stock-ins", a.DeleteStockIns).Methods("DELETE")
}

// handler for get all product
func (a *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProducts(a.DB, w, r)
}

// handler for get all stock in
func (a *App) GetStockIns(w http.ResponseWriter, r *http.Request) {
	handler.GetStockIns(a.DB, w, r)
}

// handler for create stock_in
func (a *App) CreateStockIns(w http.ResponseWriter, r *http.Request) {
	handler.CreateStockIns(a.DB, w, r)
}

// handler for delete stock_in
func (a *App) DeleteStockIns(w http.ResponseWriter, r *http.Request) {
	handler.DeleteStockIns(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Printf("Start a server on port %s", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
