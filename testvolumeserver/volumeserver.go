package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		log.Println("test")
		fileData, err := ioutil.ReadFile("/serverdata/test.txt")
		if err != nil {
			log.Fatal("File reading error", err)
		}
		return c.String(http.StatusOK, string(fileData))
	})
	e.Logger.Fatal(e.Start(":8090"))
}
