package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
	"net/http"
	"os"
)

var APP_STORAGE = "storage"
var APP_NAME = "EchoApp"
var ERROR_LOG = "AppError.log"

func main() {
	//TODO: write a router init
	router := echo.New()

	LogFile, err := os.OpenFile(fmt.Sprintf("%s/%s/%s", APP_STORAGE, APP_NAME, ERROR_LOG), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer LogFile.Close()

	if err != nil {
		panic(err)
	}
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: io.MultiWriter(LogFile),
	}))

	router.GET("/status", func(c echo.Context) error {
		c.JSON(http.StatusOK, map[string]string{"health": "OK"})
		return nil
	})
	router.GET("/redirect", func(c echo.Context) error { return nil }, middleware.HTTPSRedirect())
	//TODO: write a log init
	//TODO: write a db connection init
	//TODO: gracefully close the server with a defer functino
	//TODO: run the server
	router.Start(":8080")
	return
}
