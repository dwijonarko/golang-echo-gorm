package controllers

import (
	"fmt"
	"golang-echo/db"
	"golang-echo/entities"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func GetProduct(ctx echo.Context) error {
	db := db.DbManager()
	products := []entities.Product{}
	results := make(map[string]interface{})
	err := db.Find(&products).Error
	if err != nil {
		results["status"] = "error"
		results["data"] = nil
		results["message"] = err
	} else {
		results["status"] = "success"
		results["message"] = "Success get all products"
		results["data"] = products
	}

	return ctx.JSON(http.StatusOK, results)

}

func GetProductById(ctx echo.Context) error {
	product_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	db := db.DbManager()
	product := entities.Product{}
	err = db.First(&product, product_id).Error
	results := make(map[string]interface{})

	if err != nil {
		results["status"] = "error"
		results["data"] = nil
		results["message"] = "Product not found"
	} else {
		results["status"] = "success"
		results["message"] = fmt.Sprintf("Success get product with id %s", ctx.Param("id"))
		results["data"] = product
	}

	return ctx.JSON(http.StatusOK, results)

}

func CreateProduct(ctx echo.Context) error {
	db := db.DbManager()
	product := new(entities.Product)
	ctx.Bind(product)
	result := db.Create(&product)

	results := make(map[string]interface{})
	if result.Error != nil {
		results["status"] = "error"
		results["data"] = nil
		results["message"] = result.Error
	} else {
		results["status"] = "sucess"
		results["data"] = product
		results["message"] = "Product save succesfully"
	}
	return ctx.JSON(http.StatusOK, results)
}

func UpdateProduct(ctx echo.Context) error {
	product_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	db := db.DbManager()
	product := entities.Product{}
	err = db.First(&product, product_id).Error
	results := make(map[string]interface{})
	ctx.Bind(&product)
	if err != nil {
		results["status"] = "error"
		results["data"] = nil
		results["message"] = "Product not found"
	} else {
		db.Updates(product)
		results["status"] = "success"
		results["message"] = fmt.Sprintf("Success update product with id %s", ctx.Param("id"))
		results["data"] = product
	}

	return ctx.JSON(http.StatusOK, results)
}

func DeleteProduct(ctx echo.Context) error {
	product_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	db := db.DbManager()
	product := entities.Product{}
	err = db.First(&product, product_id).Error
	results := make(map[string]interface{})
	if err != nil {
		results["status"] = "error"
		results["data"] = nil
		results["message"] = "Product not found"
	} else {
		db.Delete(&product)
		results["status"] = "success"
		results["message"] = fmt.Sprintf("Success delete product with id %s", ctx.Param("id"))
		results["data"] = product
	}

	return ctx.JSON(http.StatusOK, results)
}
