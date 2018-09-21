-- create products tabel
-------------------------------

CREATE TABLE IF NOT EXISTS "products" (
	"id" integer primary key autoincrement,
	"sku" varchar(255),
	"name" varchar(255),
	"stocks" integer,
	"created_at" datetime, -- auto from gorm.Model
	"updated_at" datetime, -- auto from gorm.Model
	"deleted_at" datetime  -- auto from gorm.Model
);
CREATE TABLE sqlite_sequence(name,seq);
CREATE INDEX idx_products_deleted_at ON "products"(deleted_at) ;

-- end products tabel
-------------------------------

CREATE TABLE IF NOT EXISTS "stock_ins" (
	"id" integer primary key autoincrement,
	"transaction_time" datetime,
	"product_id" integer,
	"order_qty" integer,
	"received_qty" integer,
	"purchase_price" integer,
	"total_price" integer,
	"receipt" varchar(255),
	"created_at" datetime, -- auto from gorm.Model
	"updated_at" datetime, -- auto from gorm.Model
	"deleted_at" datetime, -- auto from gorm.Model 
);
CREATE INDEX idx_stock_ins_deleted_at ON "stock_ins"(deleted_at) ;


-- create stock_ins_progress tabel
CREATE TABLE IF NOT EXISTS stock_ins_progress (
	id INTEGER PRIMARY KEY, 
	transaction_progress_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	stock_ins_id INTEGER,
	in_qty INTEGER
);

-- create stock_outs tabel
CREATE TABLE IF NOT EXISTS stock_outs (
	id INTEGER PRIMARY KEY, 
	transaction_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	product_id INTEGER,
	out_qty INTEGER,
	sell_price INTEGER,
	total_price INTEGER,
	note TEXT
);
