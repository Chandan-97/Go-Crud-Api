package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", get)
	e.Logger.Fatal(e.Start(":8080"))
}

func get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Server")
}
