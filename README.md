# Software Engineer (Backend) Technical Assessment: Inventory

**case study** : Toko Ijah

## Todo
 - [ ] **Mandotory:** create REST API to replace inventory spreadsheet
 - [ ] **Mandotory:** export data report in csv format
 - [ ] Optional : import data from csv/spreadsheet Toko Ijah (data migration)
 - [ ] Optional : CMS UI for inventory management

## API

#### /api/products
* `GET` : Get all products

#### /api/stock-ins
* `GET` : Get all stock in
* `POST` : Create a stock in

## Usage

Post `/stock-ins` with this json
```json
{
	"transaction_time":"2018-09-21T14:42:49.77869956+07:00",
	"product_id":1,
	"order_qty":100,
	"received_qty":0,
	"purchase_price":1000,
	"total_price":1000000,
	"receipt":"ASP"
}
```

## Note
README will update soon
