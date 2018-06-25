package main

func main() {
	//TODO: write a router init
	router := echo.New()
	router.Use(middleware.Recovery())
	router.Use(middleware.Log())
	router.GET("/status", func(c echo.Context) {
		c.JSON(http.StatusOK, map[string]string{"health": "OK"})
	})
	//TODO: write a log init
	//TODO: write a db connection init
	//TODO: gracefully close the server with a defer functino
	//TODO: run the server
	router.Start(":8080")
}
