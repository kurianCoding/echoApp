package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	//TODO: write a router init
	router := echo.New()
	router.Use(middleware.Logger())

	router.GET("/status", func(c echo.Context) error {
		c.JSON(http.StatusOK, map[string]string{"health": "OK"})
		return nil
	})
	router.GET("/redirect", func(c echo.Context) error { return nil }, middleware.NONHTTPSRedirect())
	//TODO: write a log init
	//TODO: write a db connection init
	//TODO: gracefully close the server with a defer functino
	//TODO: run the server
	router.Start(":8080")
	return
}
