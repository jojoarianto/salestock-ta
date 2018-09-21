package config

// object config for database connection
type Config struct {
	DB *DBConfig
}

// database gorm setup
type DBConfig struct {
	Dialeg string
	DBUri  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialeg: "sqlite3",                       // use sqlite database
			DBUri:  "database/salestock-ta.sqlite3", // db uri for sqlite3 which is url file
		},
	}
}
