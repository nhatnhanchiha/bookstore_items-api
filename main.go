package main

import (
	"github.com/nhatnhanchiha/bookstore_items-api/app"
	"os"
)

func main() {
	_ = os.Setenv("LOG_LEVEL", "info")
	app.StartApplication()
}
