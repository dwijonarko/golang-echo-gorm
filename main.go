package main

import (
	"golang-echo/db"
	"golang-echo/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.ProductHandler(e)
	db.Connect()
	e.Logger.Fatal(e.Start(":8080"))
}
