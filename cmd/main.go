// Simple package that illustrates basic usage of echo
package main

import (
	"fmt"
	"simple-echo/internal/web"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func main() {
	// Pointer to echo
	e := echo.New()
	v := validator.New()
	e.Validator = &ProductValidator{validator: v}

	web.SetRoutes(e)

	// Port, handle error
	err := e.Start(":8080")
	if err != nil {
		fmt.Println(err)
	}

	// e.Logger.Print("Listening on port 8080")
	// Log and print the error
	// e.Logger.Fatal(e.Start(":8080"))
}
