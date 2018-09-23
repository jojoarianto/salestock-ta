# Software Engineer (Backend) Technical Assessment: Inventory

This is Salestock backend technical assessment. Salestock give me a study case which called "**Toko Ijah**".  Main domain of this case is Inventory. 

Toko Ijah want to replace her spreadsheet by creating an application.   So, goal of this project is to provide REST API for toko ijah inventory application.

## Feature
* `Barang Masuk` Process of entering stock in (Barang masuk) can be done in stages (Progress Stock In)
* `Barang Keluar` When stock not enough, the system will give rejection

## Installation & Run

```bash
# Download this project
go get github.com/jojoarianto/salestock-ta
```

```bash
# Build and Run
cd salestock-ta
go build
./salestock-ta

# API Endpoint : http://127.0.0.1:8000
```

## Product Items Backlog
 - [X] **Mandotory:** create REST API to replace inventory spreadsheet
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
         - [ ] Update  
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
         - [ ] Delete 
 - [X] **Mandotory:** export data report in csv format
     - [X] Stock (Catatan Jumlah Barang)
     - [X] Stock in (Catatan Barang Masuk)
     - [X] Stock out (Catatan Barang Keluar)
     - [ ] Report value of product (Laporan Nilai Barang)
     - [ ] Sales report (Laporan Penjualan)
 - [X] Optional : import data from csv/spreadsheet Toko Ijah (data migration)
     - [X] Import product
     - [ ] import stock in
     - [ ] Import stock out
 - [ ] Optional : CMS UI for inventory management


## API

#### /api/products
* `GET` : Get all product
* `POST` : Create a product* 

#### /api/stock-ins
* `GET` : Get all stock in
* `POST` : Create a stock in

#### /api/stock-ins/:id
* `GET` : Get a stock in
* `DELETE` : Delete a stock in
* `PUT` : Update a stock in

#### /api/stock-ins/:id/progress
* `GET` : Get a all progress stock in by id stock in 
* `POST` : Create a stock in progress

## Usage

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

## Response

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

### Export Csv

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

## Note
README will update soon
