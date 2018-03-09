package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type Product struct {
	Title       string  `json:"Title"`
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
}

type Products []Product

func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// METHOD: GET
	// ROUTE: /
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Hello world")
	})

	// METHOD: GET
	// ROUTE: /products
	app.Handle("GET", "/products", func(ctx iris.Context) {
		var Products Products
		var IPhoneX = Product{
			Title:       "iPhone X",
			Description: "This is a new generation of iphone",
			Price:       30000.00,
		}

		var SamsungS9 = Product{
			Title:       "Samsung Galaxy S9",
			Description: "Samsung Galaxy Phone",
			Price:       29000.00,
		}

		Products = append(Products, IPhoneX)
		Products = append(Products, SamsungS9)

		ctx.JSON(Products)
	})

	app.Run(iris.Addr("localhost:3000"), iris.WithoutServerError(iris.ErrServerClosed))
}
