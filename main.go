package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/iamgoangle/go-api-101/controllers"
)

type Company struct {
	Name  string
	City  string
	Other string
}

func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// METHOD: GET
	// ROUTE: /products
	app.Handle("GET", "/products", controllers.GetProducts)

	// METHOD: POST
	// ROUTE product
	// JSON raw body
	app.Handle("POST", "product", controllers.PostProduct)

	app.Run(iris.Addr("localhost:3000"), iris.WithoutServerError(iris.ErrServerClosed))
}
