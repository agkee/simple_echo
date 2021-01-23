// Simple package that illustrates basic usage of echo
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	// Pointer to echo
	e := echo.New()

	e.GET("/", HelloHandler)
	e.GET("/product", GetProductsHandler)
	e.GET("/product/:id", GetProductByIDHandler)
	e.POST("/product", AddProductHandler)

	// Port, handle error
	err := e.Start(":8080")
	if err != nil {
		fmt.Println(err)
	}

	// e.Logger.Print("Listening on port 8080")
	// Log and print the error
	// e.Logger.Fatal(e.Start(":8080"))
}

var products = map[int]string{
	1: "car",
	2: "computer",
	3: "phone",
}

// very simple hello handler with string
// Try with POST and see that it doesn't work
func HelloHandler(c echo.Context) error {
	return c.String(200, "HELLO")
	// return c.String(http.StatusOK, "HELLO")
}

func GetProductsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func GetProductByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	product, exists := products[id]
	if exists {
		return c.JSON(http.StatusOK, product)
	}

	return c.JSON(http.StatusNotFound, fmt.Sprintf("Product for id: %d does not exist", id))
}

func AddProductHandler(c echo.Context) error {
	type Body struct {
		Name string `json:"product_name"`
	}

	var reqBody Body
	// Binds the request body with provided type
	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	products[len(products)+1] = reqBody.Name
	return c.JSON(http.StatusOK, products)
}
