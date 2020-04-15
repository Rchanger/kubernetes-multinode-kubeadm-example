package main

import (
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
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.Logger.Fatal(e.Start(":8090"))
}
