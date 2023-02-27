package main

import (
	"log"
	"net/http"

	"github.com/cnugroho11/blog_api/initializer"
	"github.com/labstack/echo/v4"
)

var (
	server *echo.Echo
)

func init() {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env", err)
	}

	initializer.ConnectDatabase(&config)

	server = echo.New()
}

func main() {
	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Server run properly",
		})
	})

	server.Logger.Fatal(server.Start(":8080"))
}
