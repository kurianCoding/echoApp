package main

import (
	"fmt"
	"github.com/kurianCoding/echoApp/services"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"io"
	"net/http"
	"os"
)

var APP_STORAGE = "temp"
var APP_NAME = "logs"
var ERROR_LOG = "App.log"

func main() {
	//TODO: write a router init
	router := echo.New()

	LogFile, err := os.OpenFile(fmt.Sprintf("%s/%s/%s", APP_STORAGE, APP_NAME, ERROR_LOG), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	redisPool, err := services.NewRedisStore() // redis connection is made here
	defer func() { LogFile.Close() }()

	if err != nil {
		panic(err)
	}
	// declares, a session store, adds the session store to echo context with key `_session_store`
	router.Use(session.MiddlewareWithConfig(session.Config{Store: redisPool}))

	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: io.MultiWriter(LogFile),
	}))

	router.GET("/status", func(c echo.Context) error {
		c.JSON(http.StatusOK, map[string]string{"health": "OK"})
		return nil
	})
	//router.GET("/redirect", func(c echo.Context) error { return nil }, middleware.HTTPSRedirect())
	//TODO: write a log init
	//TODO: write a db connection init
	//TODO: gracefully close the server with a defer functino
	//TODO: run the server
	router.Start(":8080")
	return
}
