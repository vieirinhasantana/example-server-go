package main

import (
	"fmt"

	"github.com/vieirinhasantana/example-server-go/pkg/domain"
)

func main() {
	// e := echo.New()
	// e.GET("/products", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })
	// e.GET("/product/:id", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })
	// e.POST("/product/:id", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })
	// e.Logger.Fatal(e.Start(":8080"))

	product := domain.NewProduct()
	resp, _ := product.Create("produto teste", "client123")
	fmt.Println(resp)
}
