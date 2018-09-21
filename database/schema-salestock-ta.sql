-- create products tabel
-- CREATE TABLE IF NOT EXISTS products (
-- 	id INTEGER PRIMARY KEY, 
-- 	sku TEXT, 
-- 	name TEXT,
-- 	stocks INTEGER
-- );

CREATE TABLE IF NOT EXISTS "products" (
	"id" integer primary key autoincrement,
	"created_at" datetime,
	"updated_at" datetime,
	"deleted_at" datetime,
	"sku" varchar(255),
	"name" varchar(255),
	"stocks" integer 
);
CREATE TABLE sqlite_sequence(name,seq);
CREATE INDEX idx_products_deleted_at ON "products"(deleted_at) ;

-- create stock_ins tabel
CREATE TABLE IF NOT EXISTS stock_ins (
	id INTEGER PRIMARY KEY, 
	transaction_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	product_id INTEGER,
	order_qty INTEGER,
	received_qty INTEGER,
	purchase_price INTEGER,
	total_price INTEGER,
	receipt TEXT,
	FOREIGN KEY (product_id) REFERENCES products(id)
);

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
