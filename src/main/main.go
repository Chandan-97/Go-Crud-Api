package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/candidates", get)
	e.GET("/candidates/:id", getCandidate)
	e.POST("/candidates", post)
	e.PUT("/candidates/:id", update)
	e.DELETE("/candidates/:id", delete)

	e.Logger.Fatal(e.Start(":8080"))
}

// func get(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello from Server")
// }

func get(c echo.Context) error {
	return c.String(http.StatusOK, "GET /candidates")
}

func getCandidate(c echo.Context) error {
	return c.String(http.StatusOK, "GET /candidates/:id")
}

func post(c echo.Context) error {
	return c.String(http.StatusOK, "POST /candidates/:id")
}

func update(c echo.Context) error {
	return c.String(http.StatusOK, "PUT /candidates/:id")
}

func delete(c echo.Context) error {
	return c.String(http.StatusOK, "DELETE /candidates/:id")
}
