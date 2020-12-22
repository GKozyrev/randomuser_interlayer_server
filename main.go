package main

import (
	"testapi/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// init server
	server := echo.New()
	// server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// routes
	server.GET("/data", handlers.DataGet)
	server.POST("/data", handlers.DataPost)
	server.File("/favicon.ico", "/favicon.ico") // just hate favicon 404 error

	// start server
	server.Logger.Fatal(server.Start(":8080"))
}

// example request:
// http://localhost:8080/data?results=10&from=2012-11-09T07:47:23.904Z&to=2019-11-09T07:47:23.904Z
