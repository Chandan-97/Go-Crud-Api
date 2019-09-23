package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

type Candidate struct {
	gorm.Model

	Name         string `gorm:"type:varchar(100)"`
	Phone_number string `gorm:"type:varchar(20)"`
	Status       string `gorm:"type:varchar(10)"`
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password")
	if err != nil {
		fmt.Print(err)
	}

	if !db.HasTable(&Candidate{}) {
		db.CreateTable(&Candidate{})
	}

	db.AutoMigrate(&Candidate{})

	e := echo.New()

	e.GET("/candidates/:id", func(c echo.Context) error {
		id := c.Param("id")
		var candidate Candidate
		db.Table("candidates").Where("id = ?", id).First(&candidate)
		return c.JSON(http.StatusOK, candidate)
	})

	// GET
	e.GET("/candidates", func(c echo.Context) error {
		var candidates []Candidate
		db.Find(&candidates)
		// fmt.Println(candidates)
		return c.JSON(http.StatusOK, candidates)
	})

	// POST
	e.POST("/candidates", func(c echo.Context) error {
		name := c.FormValue("name")
		status := c.FormValue("status")
		phone_number := c.FormValue("phone_number")

		candidate := Candidate{Name: name, Phone_number: phone_number, Status: status}

		db.Create(&candidate)

		return c.JSON(http.StatusOK, candidate)
	})

	// PUT
	e.PUT("/candidates/:id", func(c echo.Context) error {
		id := c.Param("id")
		name := c.FormValue("name")
		status := c.FormValue("status")
		phone_number := c.FormValue("phone_number")

		db.Table("candidates").Where("id = ?", id).Updates(Candidate{Name: name, Status: status, Phone_number: phone_number})

		return c.String(http.StatusOK, "PUT")
	})

	// DELETE
	e.DELETE("/candidates/:id", func(c echo.Context) error {
		id := c.Param("id")

		db.Unscoped().Table("candidates").Where("id = ?", id).Delete(Candidate{})

		return c.String(http.StatusOK, "DELETE")
	})

	e.Logger.Fatal(e.Start(":8080"))
	defer db.Close()
}
