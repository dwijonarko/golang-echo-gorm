package router

import (
	"golang-echo/controllers"

	"github.com/labstack/echo/v4"
)

func ProductHandler(e *echo.Echo) {
	e.GET("/", controllers.Hello)
	e.GET("/products", controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
	e.GET("/products/:id", controllers.GetProductById)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)
}
