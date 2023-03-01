package main

import (
	"log"
	"net/http"

	"github.com/cnugroho11/blog_api/controller"
	"github.com/cnugroho11/blog_api/initializer"
	"github.com/cnugroho11/blog_api/route"
	"github.com/labstack/echo/v4"
)

var (
	server *echo.Echo

	BlogController      controller.BlogController
	BlogRouteController route.BlogRouteController
)

func init() {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env", err)
	}

	initializer.ConnectDatabase(&config)

	BlogController = controller.NewBlogController(initializer.DB)
	BlogRouteController = route.NewBlogRouteController(BlogController)

	server = echo.New()
}

func main() {
	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Server run properly",
		})
	})

	router := server.Group("/api/v1")
	BlogRouteController.BlogRoute(router)

	server.Logger.Fatal(server.Start(":8080"))
}
