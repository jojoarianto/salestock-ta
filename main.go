package main

import (
	"github.com/jojoarianto/salestock-ta/app"    // import app
	"github.com/jojoarianto/salestock-ta/config" // import config
)

func main() {
	config := config.GetConfig() // get config app

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
