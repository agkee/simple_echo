package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	// Pointer to echo
	e := echo.New()

	e.GET("/", HelloHandler)

	// Port, handle error
	err := e.Start(":8080")
	if err != nil {
		fmt.Println(err)
	}

	// e.Logger.Print("Listening on port 8080")
	// Log and print the error
	// e.Logger.Fatal(e.Start(":8080"))
}

// very simple hello handler with string
// Try with POST and see that it doesn't work
func HelloHandler(c echo.Context) error {
	return c.String(200, "HELLO")
	// return c.String(http.StatusOK, "HELLO")
}
