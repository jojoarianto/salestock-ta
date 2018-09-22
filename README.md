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
* `POST` : Create a product

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

GET `/api/stock-in` with this json
```json
[
    {
        "ID": 1,
        "CreatedAt": "2018-09-22T15:36:03.697961927+07:00",
        "UpdatedAt": "2018-09-22T17:05:34.153181848+07:00",
        "DeletedAt": null,
        "stock_in_time": "0001-01-01T00:00:00Z",
        "product_id": 1,
        "Product": {
            "ID": 1,
            "CreatedAt": "2018-09-22T15:35:37.861424368+07:00",
            "UpdatedAt": "2018-09-22T17:05:34.152808305+07:00",
            "DeletedAt": null,
            "sku": "SSI-D00791015-LL-BWH",
            "name": "Zalekia Plain Casual Blouse (L,Broken White)",
            "stocks": 0
        },
        "order_qty": 50,
        "received_qty": 0,
        "purchase_price": 1000,
        "total_price": 100000,
        "receipt": "IRIANTO-99-NEW-99",
        "Progress": [
            {
                "ID": 1,
                "CreatedAt": "2018-09-22T16:43:30.793632901+07:00",
                "UpdatedAt": "2018-09-22T17:05:34.153397007+07:00",
                "DeletedAt": null,
                "stock_in_progress_time": "2018-09-21T14:42:49.77869956+07:00",
                "stock_ins_id": 1,
                "qty": 3
            },
            {
                "ID": 2,
                "CreatedAt": "2018-09-22T16:50:20.024516655+07:00",
                "UpdatedAt": "2018-09-22T17:05:34.153575862+07:00",
                "DeletedAt": null,
                "stock_in_progress_time": "2018-09-21T14:42:49.77869956+07:00",
                "stock_ins_id": 1,
                "qty": 5
            }
        ]
    },
    ...
]
```

Post `/api/products` with this json
```json
{
	"sku": "SSI-D00791015-LL-BWH",
	"name": "Zalekia Plain Casual Blouse (L,Broken White)"
}
```

Post `/api/stock-ins` with this json
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

Post `/api/stock-ins` Response
```json
{
    "ID": 1,
    "CreatedAt": "2018-09-22T03:40:31.544120826+07:00",
    "UpdatedAt": "2018-09-22T03:40:31.544120826+07:00",
    "DeletedAt": null,
    "transaction_time": "2018-09-21T14:42:49.77869956+07:00",
    "product_id": 1,
    "Product": {
        "ID": 1,
        "CreatedAt": "2018-09-22T03:36:48.762505285+07:00",
        "UpdatedAt": "2018-09-22T03:40:31.536253105+07:00",
        "DeletedAt": null,
        "sku": "SSI-D00791015-LL-BWH",
        "name": "Zalekia Plain Casual Blouse (L,Broken White)",
        "stocks": 0
    },
    "order_qty": 100,
    "received_qty": 0,
    "purchase_price": 1000,
    "total_price": 100000,
    "receipt": "IRIANTO-99-NEW-99"
}
```


## Note
README will update soon
