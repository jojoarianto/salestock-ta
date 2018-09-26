# Software Engineer (Backend) Technical Assessment: Inventory

This is Salestock backend technical assessment. Salestock give me a study case which called "**Toko Ijah**".  Main domain of this case is Inventory. 

Toko Ijah want to replace her spreadsheet by creating an application.   So, goal of this project is to provide REST API for toko ijah inventory application.

## Installation & Run

```bash
# Download this project
go get github.com/jojoarianto/salestock-ta

# It's take several minute to download project
```

Make sure you have set up $GOPATH

```bash
# Build and Run
cd $GOPATH/src/github.com/jojoarianto/salestock-ta
go build
./salestock-ta

# API Endpoint : http://127.0.0.1:8000
```
## Library & Dependency
Project library : 
* Github.com/gorilla/mux
* Github.com/jinzhu/gorm
* Github.com/jinzhu/gorm/dialects/sqlite
* Gopkg.in/go-playground/validator.v9

## Model Data
Products (Barang)
```go 
type Product struct {
    gorm.Model
    Sku    string
    Name   string 
    Stocks int
}
```

Stock Ins (Barang Masuk)
```go
type StockIn struct {
    gorm.Model
    StockInTime   time.Time
    ProductID     int
    Product       Product // belongs to
    OrderQty      int 
    ReceivedQty   int 
    PurchasePrice int 
    TotalPrice    int 
    Receipt       string 
    Progress      []StockInProgress // has many progress
    StausInCode   int // 0. waiting, 1 completed
}
```
Stock In Progress (Progress Barang Masuk)
```go
type StockInProgress struct { 
    gorm.Model
    ProgressInTime time.Time
    StockInsID     int // belongs to
    Qty            int 
}
```
Stock Outs (Barang Keluar)
```go
type StockOut struct { 
    gorm.Model
    StockOutTime  time.Time
    ProductID     int  // belongs to
    Product       Product 
    OutQty        int
    SellPrice     int 
    TotalPrice    int 
    Transaction   string // transaction null jika barang tidak terjual
    StatusOutCode int // 1. Terjual, 2. Barang Hilang, 3. Barang Rusak, 4 Barang Sample
}
```


## Features
* `Barang Masuk` Process of entering stock in (Barang masuk) can be done in stages (Progress Stock In)
* `Barang Keluar` When stock not enough, the system will give rejection and rollback the insertion
* `Import Barang` by using csv file


## Product Items Backlog
 - [X] **Mandatory:** create REST API to replace inventory spreadsheet
     - [X] Product (Barang)
         - [X] Get all
         - [X] Get by id 
         - [X] Create 
         - [X] Update 
         - [X] Delete
     - [X] Stock In (Barang Masuk)
         - [X] Get all
         - [X] Get by id
         - [X] Create
         - [X] Update  
         - [X] Delete
     - [X] Stock In Progress (Tahapan Barang Masuk)
         - [X] Get all progress by stock_in_id
         - [X] Create 
         - [ ] Update 
         - [ ] Delete  
     - [X] Stock Out 
         - [X] Get All
         - [X] Get by id 
         - [X] Create
         - [ ] Update
         - [X] Delete 
 - [X] **Mandatory:** export data report in csv format
     - [X] Stock (Catatan Jumlah Barang)
     - [X] Stock in (Catatan Barang Masuk)
     - [X] Stock out (Catatan Barang Keluar)
     - [ ] Report value of product (Laporan Nilai Barang)
     - [X] Sales report (Laporan Penjualan)
 - [X] Optional : import data from csv/spreadsheet Toko Ijah (data migration)
     - [X] Import product
     - [ ] import stock in
     - [ ] Import stock out
 - [ ] Optional : CMS UI for inventory management


## API ENDPOINT

#### /api/products
* `GET` : Get all product
* `POST` : Create a product* 

#### /api/products/:product_id
* `GET` : Get a product by id
* `PUT` : Update a product
* `DELETE` : Delete a product* 

#### /api/stock-ins
* `GET` : Get all stock in
* `POST` : Create a stock in

#### /api/stock-ins/:stock_in_id
* `GET` : Get a stock in
* `DELETE` : Delete a stock in
* `PUT` : Update a stock in

#### /api/stock-ins/:stock_in_id/progress
* `GET` : Get all stock in progress by id stock in 
* `POST` : Create a stock in progress

#### /api/stock-outs
* `GET` : Get all stock out
* `POST` : Create a stock outs

#### /api/stock-outs/:stock_out_id
* `GET` : Get a stock out

## Export & Import URL
#### /export/products
* `GET` : Export to get all product
* Out file directory : `csv/export_products.csv` 

#### /export/stock_ins
* `GET` : Export to get all stock in transaction 
* Out file directory : `csv/export_stock_ins.csv` 

#### /export/stock_Outs
* `GET` : Export to get all stock out transaction 
 * Out file directory : `csv/export_stock_outs.csv` 

#### /export/sales
* `GET` : Export to get all sales report transaction 
* Out file directory : `csv/export_sales_report.csv` 

#### /import/products
* `GET` : Import products from csv file
* File directory : `csv/import_products.csv`


### Usage Examples

POST `/api/products` with json
```json
{
    "sku": "SSI-D00791015-LL-BWH",
    "name": "Zalekia Plain Casual Blouse (L,Broken White)"
}
```

POST `/api/stock-ins` with json
```json
{
    "stock_in_time":"2018-09-21T14:42:49.77869956+07:00",
    "product_id":1,
    "order_qty":54,
    "purchase_price":77000,
    "receipt":"20170823-75140"
}
```

POST `/api/stock-ins/1/progress` with json
```json
{
    "stock_in_progress_time":"2017-08-26T14:42:49.77869956+07:00",
    "qty":54
}
```

POST `/api/stock-outs` with json
```json
{
    "stock_out_time":"2018-01-01T14:42:49.77869956+07:00",
    "product_id":1,
    "out_qty":1,
    "sell_price":130000,
    "transaction_id":"20180101-023993",
    "status_out_code":1
}
```

### Response

GET all stock in `/api/stock-ins`
```json
[
    {
        "ID": 1,
        "CreatedAt": "2018-09-23T06:10:28.568665081+07:00",
        "UpdatedAt": "2018-09-23T06:18:20.26672724+07:00",
        "DeletedAt": null,
        "stock_in_time": "2018-09-21T14:42:49.77869956+07:00",
        "product_id": 1,
        "product": {
            "ID": 1,
            "CreatedAt": "2018-09-23T06:00:18.659528013+07:00",
            "UpdatedAt": "2018-09-23T07:09:45.76298692+07:00",
            "DeletedAt": null,
            "sku": "SSI-D00791015-LL-BWH",
            "name": "Zalekia Plain Casual Blouse (L,Broken White)",
            "stocks": 54
        },
        "order_qty": 54,
        "received_qty": 54,
        "purchase_price": 77000,
        "total_price": 4158000,
        "receipt": "20170823-75140",
        "progress": [
            {
                "ID": 1,
                "CreatedAt": "2018-09-23T06:16:09.830367981+07:00",
                "UpdatedAt": "2018-09-23T06:16:09.830367981+07:00",
                "DeletedAt": null,
                "stock_in_progress_time": "2017-08-26T14:42:49.77869956+07:00",
                "stock_ins_id": 1,
                "qty": 1
            },
            {
                "ID": 2,
                "CreatedAt": "2018-09-23T06:18:20.211968571+07:00",
                "UpdatedAt": "2018-09-23T06:18:20.211968571+07:00",
                "DeletedAt": null,
                "stock_in_progress_time": "2017-08-26T14:42:49.77869956+07:00",
                "stock_ins_id": 1,
                "qty": 53
            }
        ],
        "status_in_code": 1
    },
    ...
]
```

GET all stock outs `/api/stock-outs`
```json
[
    {
        "ID": 1,
        "CreatedAt": "2018-09-23T06:23:08.338073577+07:00",
        "UpdatedAt": "2018-09-23T06:23:08.338073577+07:00",
        "DeletedAt": null,
        "stock_out_time": "2018-01-01T14:42:49.77869956+07:00",
        "product_id": 1,
        "product": {
            "ID": 1,
            "CreatedAt": "2018-09-23T12:21:57Z",
            "UpdatedAt": "2018-09-23T18:29:31.81898317+07:00",
            "DeletedAt": null,
            "sku": "SSI-D00791015-LL-BWH",
            "name": "Zalekia Plain Casual Blouse (L,Broken White)",
            "stocks": 89
        },
        "out_qty": 1,
        "sell_price": 130,
        "total_price": 130,
        "transaction_id": "20180101-023993",
        "status_out_code": 1
    },
    ...
]
```

### Examples of Export Csv Result

Export csv stock ins `GET` `/export/stock-ins`
```csv
SKU,Nama Item,Jumlah Sekarang
SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",83
SSI-D00791077-MM-BWH,"Zalekia Plain Casual Blouse (M,Broken White)",0
SSI-D00791091-XL-BWH,"Zalekia Plain Casual Blouse (XL,Broken White)",0
SSI-D00864612-LL-NAV,"Deklia Plain Casual Blouse (L,Navy)",0
SSI-D00864652-SS-NAV,"Deklia Plain Casual Blouse (S,Navy)",8

```

Export csv stock ins `GET` `/export/stock-ins`
```csv
Waktu,SKU,Nama Barang,Jumlah Pemesanan,Harga Diterima,Harga Beli,Total,Nomer Kwitansi,Catatan
2018/09/21 14:42,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",54,129,77000,4158000,20170823-75140,2017/08/26 terima 1; 2017/08/26 terima 53; 2017/08/26 terima 50; 2017/08/26 terima 25; 
2018/09/21 14:42,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",100,20,77000,7700000,20170823-75140,2017/08/26 terima 20; Masih menunggu
2018/09/21 14:42,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",10,0,77000,770000,20170823-75140,Masih menunggu
2018/09/21 14:42,SSI-D00864652-SS-NAV,"Deklia Plain Casual Blouse (S,Navy)",10,0,77000,770000,20170823-75140,Masih menunggu

```

Export csv stock outs `GET` `/export/stock-outs`
```csv
Waktu,SKU,Nama Barang,Jumlah Keluar,Harga Jual,Total,Catatan
2018/01/01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",1,130,130,Pesanan ID-20180101-023993
2018/01/01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",1,130,130,Pesanan ID-20180101-023993
2018/01/01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",10,130,1300,Pesanan ID-20180101-023993
2018/01/01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",10,130,1300,Pesanan ID-20180101-023993
2018/01/01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",10,130,1300,Pesanan ID-20180101-023993
2018/01/01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",100,130,13000,Pesanan ID-20180101-023993
2018/01/01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",100,130,13000,Pesanan ID-20180101-023993
```

Export csv sales report `GET` `/export/sales`
```csv
ID Pesanan,Waktu,SKU,Nama Barang,Jumlah,Harga Jual,Total,Harga Beli,Laba
20180109-853724,2018-01-01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",1,100000,100000,77000,23000
20180109-853724,2018-01-01 14:42:49,SSI-D00864652-SS-NAV,"Deklia Plain Casual Blouse (S,Navy)",2,125000,250000,77000,96000
20180109-853724,2018-01-01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",1,97000,97000,77000,20000
20180109-853724,2018-01-01 14:42:49,SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",2,100000,200000,77000,46000
```
