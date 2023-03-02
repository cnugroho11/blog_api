package route

import (
	"github.com/cnugroho11/blog_api/controller"
	"github.com/labstack/echo/v4"
)

type BlogRouteController struct {
	blogController controller.BlogController
}

func NewBlogRouteController(blogController controller.BlogController) BlogRouteController {
	return BlogRouteController{blogController}
}

func (bc *BlogRouteController) BlogRoute(rg *echo.Group) {
	router := rg.Group("/blog")

	router.GET("/all", bc.blogController.GetBlog)
	router.GET("/:slug", bc.blogController.GetBySlug)
	router.POST("/create", bc.blogController.CreateBlog)
}
