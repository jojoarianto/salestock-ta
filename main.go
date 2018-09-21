package main

import (
	"salestock-ta/app"    // import app
	"salestock-ta/config" // import config
)

func main() {
	config := config.GetConfig() // get config app

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
